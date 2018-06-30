package web

import (
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/judge"
	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/utils/problems/polygon"
	"github.com/mraron/njudge/web/models"
	_ "github.com/mraron/njudge/web/models"
	"github.com/mraron/njudge/web/roles"
	"html/template"
	"io/ioutil"
	_ "mime"
	"net/http"
	"path/filepath"
	"strconv"
	"time"
)

type Server struct {
	Hostname     string
	Port         string
	ProblemsDir  string
	TemplatesDir string

	MailAccount         string
	MailServerHost      string
	MailServerPort      string
	MailAccountPassword string

	DBAccount  string
	DBPassword string
	DBHost     string
	DBName     string

	GluePort string

	judges   []*models.Judge
	problems map[string]problems.Problem
	db       *sqlx.DB
}

/*
func New(port string, problemsDir string, templatesDir string, mailServerAccount, mailServerHost, mailServerPort, mailAccountPassword string, glueport string) *Server {
	return &Server{port, problemsDir, templatesDir, mailServerAccount, mailServerHost, mailServerPort, mailAccountPassword, glueport, make([]*models.Judge, 0), make(map[string]problems.Problem), nil}
}*/

func (s *Server) updateProblems() {
	if s.problems == nil {
		s.problems = make(map[string]problems.Problem)
	}

	files, err := ioutil.ReadDir(s.ProblemsDir)
	if err != nil {
		panic(err)
	}

	pList := make([]string, 0)

	for _, f := range files {
		if f.IsDir() {
			path := filepath.Join(s.ProblemsDir, f.Name())
			p, err := problems.Parse(path)
			if err == nil {
				s.problems[p.Name()] = p
				pList = append(pList, p.Name())
			} else {
				log.Print(err)
			}
		}
	}
}

func (s *Server) connectToDB() {
	var err error
	s.db, err = sqlx.Open("postgres", "postgres://"+s.DBAccount+":"+s.DBPassword+"@"+s.DBHost+"/"+s.DBName)

	if err != nil {
		panic(err)
	}
}

func (s *Server) loadJudgesFromDB() {
	var err error
	s.judges, err = models.GetJudges(s.db)

	if err != nil {
		panic(err)
	}
}

func (s *Server) syncJudges() {
	for {
		s.loadJudgesFromDB()
		for _, j := range s.judges {
			st, err := judge.NewFromUrl("http://" + j.Host + ":" + j.Port)

			if err != nil {
				log.Print("trying to access judge on ", j.Host, j.Port, " getting error ", err)
				j.Online = false
				j.Ping = -1
				err = j.Update(s.db)
				if err != nil {
					log.Print("also error occured while updating", err)
				}

				continue
			}

			j.Online = true
			j.State = st
			j.Ping = 1

			err = j.Update(s.db)
			if err != nil {
				log.Print("trying to access judge on", j.Host, j.Port, " unsuccesful update in database", err)
				continue
			}
		}
		time.Sleep(10 * time.Second)
	}
}

func (s *Server) internalError(c echo.Context, err error, msg string) error {
	c.Logger().Print("internal error:", err)
	return c.Render(http.StatusInternalServerError, "error.html", msg)
}

func (s *Server) unauthorizedError(c echo.Context) error {
	return c.String(http.StatusUnauthorized, "unauthorized")
}

func (s *Server) runGlue() {
	g := echo.New()
	g.Use(middleware.Logger())

	g.POST("/callback/:id", func(c echo.Context) error {
		id_ := c.Param("id")

		id, err := strconv.Atoi(id_)
		if err != nil {
			return s.internalError(c, err, "err")
		}

		st := judge.Status{}
		if err = c.Bind(&st); err != nil {
			return s.internalError(c, err, "err")
		}

		if st.Done {
			verdict := models.Verdict(st.Status.Verdict())
			if st.Status.Compiled == false {
				verdict = models.VERDICT_CE
			}

			if _, err := s.db.Exec("UPDATE submissions SET verdict=$1, status=$2, ontest=NULL, judged=$3 WHERE id=$4", verdict, st.Status, time.Now(), id); err != nil {
				return s.internalError(c, err, "err")
			}
		} else {
			if _, err := s.db.Exec("UPDATE submissions SET ontest=$1, status=$2, verdict=$3 WHERE id=$4", st.Test, st.Status, models.VERDICT_RU, id); err != nil {
				log.Print("can't realtime update status", err)
			}
		}

		return c.String(http.StatusOK, "ok")
	})

	panic(g.Start(":" + s.GluePort))
}

func (s *Server) judger() {
	for {
		time.Sleep(1 * time.Second)

		ss := []models.Submission{}

		if err := s.db.Select(&ss, "SELECT * FROM submissions WHERE started=false ORDER BY id ASC LIMIT 1"); err != nil {
			log.Print("judger query error", err)
			continue
		}

		if len(ss) == 0 {
			continue
		}

		for _, sub := range ss {
			for _, j := range s.judges {
				if j.State.SupportsProblem(sub.Problem) {
					j.State.Submit(judge.Submission{sub.Problem, sub.Language, sub.Source, "http://" + s.Hostname + ":" + s.GluePort + "/callback/" + strconv.Itoa(int(sub.Id))})
					if _, err := s.db.Exec("UPDATE submissions SET started=true WHERE id=$1", sub.Id); err != nil {
						log.Print("FATAL: ", err)
					}
					break
				}
			}
		}
	}
}

func (s *Server) Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("titkosdolog"))))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := s.currentUser(c)
			c.Set("user", user)

			if err != nil {
				return s.internalError(c, err, "belső hiba")
			}

			return next(c)
		}
	})

	t := &Template{
		templates: template.Must(template.New("templater").Funcs(s.templatefuncs()).ParseGlob(filepath.Join(s.TemplatesDir, "*.html"))),
	}

	e.Renderer = t

	e.GET("/", s.getHome)

	e.Static("/static", "public")
	e.GET("/submission/:id", s.getSubmission)

	ps := e.Group("/problemset", func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("teszt", c.Param("name"))
			c.Set("problemset", c.Param("name"))
			return next(c)
		}
	})

	ps.GET("/:name/", s.getProblemsetMain)
	ps.GET("/:name/:problem/", s.getProblemsetProblem)
	ps.GET("/:name/:problem/pdf/:language/", s.getProblemsetProblemPDFLanguage)
	ps.GET("/:name/:problem/attachment/:attachment/", s.getProblemsetProblemAttachment)
	ps.GET("/:name/:problem/:file", s.getProblemsetProblemFile)
	ps.POST("/:name/submit", s.postProblemsetSubmit)
	ps.GET("/status", s.getProblemsetStatus)

	u := e.Group("/user")

	u.GET("/login", s.getUserLogin)
	u.POST("/login", s.postUserLogin)
	u.GET("/logout", s.getUserLogout)
	u.GET("/register", s.getUserRegister)
	u.POST("/register", s.postUserRegister)
	u.GET("/activate", s.getUserActivate)
	u.GET("/activate/:name/:key", s.getActivateUser)

	v1 := e.Group("/api/v1")

	v1.GET("/problem_rels", s.getAPIProblemRels)
	v1.POST("/problem_rels", s.postAPIProblemRel)
	v1.GET("/problem_rels/:id", s.getAPIProblemRel)
	v1.PUT("/problem_rels/:id", s.putAPIProblemRel)
	v1.DELETE("/problem_rels/:id", s.deleteAPIProblemRel)

	v1.GET("/judges", s.getAPIJudges)
	v1.POST("/judges", s.postAPIJudges)
	v1.GET("/judges/:id", s.getAPIJudge)
	v1.PUT("/judges/:id", s.putAPIJudge)
	v1.DELETE("/judges/:id", s.deleteAPIJudge)

	e.GET("/admin", s.getAdmin)

	s.updateProblems()
	s.connectToDB()
	models.SetDatabase(s.db)
	s.loadJudgesFromDB()
	go s.syncJudges()
	go s.runGlue()
	go s.judger()

	fmt.Println("Ohoho started")

	for idx, judge := range s.judges {
		fmt.Println(idx, judge)
	}

	panic(e.Start(":" + s.Port))
}

func (s *Server) getHome(c echo.Context) error {
	fmt.Println("főoldal")
	return c.Render(http.StatusOK, "home.html", s.problems)
}

func (s *Server) getAdmin(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(u.Role, roles.ActionView, "admin_panel") {
		return c.Render(http.StatusUnauthorized, "error.html", "Engedély megtagadva.")
	}

	return c.Render(http.StatusOK, "admin.html", nil)
}

type paginationData struct {
	_page      int
	_perPage   int
	_sortDir   string
	_sortField string
}

func parsePaginationData(c echo.Context) (*paginationData, error) {
	res := &paginationData{}
	var err error

	_page := c.QueryParam("_page")
	_perPage := c.QueryParam("_perPage")

	res._sortDir = c.QueryParam("_sortDir")
	res._sortField = c.QueryParam("_sortField")

	res._page, err = strconv.Atoi(_page)
	if err != nil {
		return nil, err
	}

	res._perPage, err = strconv.Atoi(_perPage)
	if err != nil {
		return nil, err
	}

	return res, nil
}
