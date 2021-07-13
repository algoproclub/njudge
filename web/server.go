package web

import (
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/mraron/njudge/utils/problems"
	_ "github.com/mraron/njudge/utils/problems/config/feladat_txt"
	_ "github.com/mraron/njudge/utils/problems/config/polygon"
	_ "github.com/mraron/njudge/utils/problems/config/task_yaml"
	"github.com/mraron/njudge/web/helpers"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/helpers/roles"
	"github.com/mraron/njudge/web/helpers/templates"

	_ "github.com/mraron/njudge/utils/language/cpp11"
	_ "github.com/mraron/njudge/utils/language/cpp14"
	_ "github.com/mraron/njudge/utils/language/golang"
	_ "github.com/mraron/njudge/utils/language/julia"
	_ "github.com/mraron/njudge/utils/language/nim"
	_ "github.com/mraron/njudge/utils/language/octave"
	_ "github.com/mraron/njudge/utils/language/pascal"
	_ "github.com/mraron/njudge/utils/language/python3"
	_ "github.com/mraron/njudge/utils/language/zip"

	"github.com/mraron/njudge/web/models"
	_ "github.com/mraron/njudge/web/models"
	_ "mime"
	"net/http"
)

type Server struct {
	config.Server

	ProblemStore problems.Store

	judges []*models.Judge
	DB     *sqlx.DB
}

func (s *Server) UpdateProblem(pr string) error {
	return s.ProblemStore.UpdateProblem(pr)
}

func (s *Server) GetProblem(pr string) problems.Problem {
	p, _ := s.ProblemStore.Get(pr)
	return p
}

func (s *Server) Run() {
	s.SetupEnvironment()
	s.StartBackgroundJobs()

	e := echo.New()
	store := sessions.NewCookieStore([]byte(s.CookieSecret))

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(store))
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := s.currentUser(c)
			if err != nil {
				return helpers.InternalError(c, err, "belső hiba")
			}
			c.Set("user", user)

			return next(c)
		}
	})

	e.Renderer = templates.New(s.TemplatesDir, s.ProblemStore)

	s.prepareRoutes(e)

	panic(e.Start(":" + s.Port))
}

func (s *Server) getHome(c echo.Context) error {
	return c.Render(http.StatusOK, "home.gohtml", nil)
}

func (s *Server) getAdmin(c echo.Context) error {
	u := c.Get("user").(*models.User)
	if !roles.Can(roles.Role(u.Role), roles.ActionView, "admin_panel") {
		return c.Render(http.StatusUnauthorized, "error.gohtml", "Engedély megtagadva.")
	}

	return c.Render(http.StatusOK, "admin.gohtml", struct {
		Host string
	}{s.Hostname + ":" + s.Port})
}
