package web

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/utils/problems"
	"github.com/mraron/njudge/utils/problems/config/polygon"
	"github.com/mraron/njudge/web/models"
	"github.com/volatiletech/sqlboiler/v4/queries"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func (s *Server) getProblemsetList(c echo.Context) error {
	type problem struct{
		problems.Problem
		SolverCount int
	}

	name := c.Param("name")
	problemLst, err := models.ProblemRels(Where("problemset=?", name)).All(s.db)

	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	if len(problemLst) == 0 {
		return c.Render(http.StatusNotFound, "404.gohtml", "Nem található.")
	}

	lst := make([]problem, len(problemLst))

	for i := 0; i < len(problemLst); i ++ {
		cnt := struct {
			Count int64
		}{0}

		err := queries.Raw("SELECT COUNT(DISTINCT user_id) FROM submissions WHERE problemset=$1 and problem=$2 and verdict=0", name, problemLst[i].Problem).Bind(context.TODO(), s.db, &cnt)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return s.internalError(c, err, "Belső hiba.")
		}

		lst[i] = problem{Problem: s.getProblem(problemLst[i].Problem), SolverCount: int(cnt.Count)}
	}

	return c.Render(http.StatusOK, "problemset_list.gohtml", struct {
		Lst []problem
	}{lst})
}

func (s *Server) getProblemsetProblem(c echo.Context) error {
	name, problem := c.Param("name"), c.Param("problem")

	lst, err := models.ProblemRels(Where("problemset=?", name)).All(s.db)
	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	ok := false
	for _, val := range lst {
		if val.Problem == problem {
			ok = true
		}

		if ok {
			break
		}
	}

	if !ok {
		return c.JSON(http.StatusNotFound, nil)
	}

	lastLanguage := ""
	if u := c.Get("user").(*models.User); u != nil {
		fmt.Println(u)
		sub, err := models.Submissions(Select("language"), Where("user_id = ?", u.ID), OrderBy("id DESC"), Limit(1)).One(s.db)
		if err == nil {
			lastLanguage = sub.Language
		}
	}

	return c.Render(http.StatusOK, "problemset_problem.gohtml",struct{
		Problem problems.Problem
		LastLanguage string
	}{s.getProblem(problem), lastLanguage})
}

func (s *Server) getProblemsetProblemPDFLanguage(c echo.Context) error {
	p, lang := s.getProblem(c.Param("problem")), c.Param("language")

	if p == nil {
		return c.String(http.StatusNotFound, "no such problem")
	}

	if len(p.PDFStatements()) == 0 {
		return c.String(http.StatusNotFound, "no pdf statement")
	}

	return c.Blob(http.StatusOK, "application/pdf", translateContent(lang, p.PDFStatements()).Contents)
}

func (s *Server) getProblemsetProblemFile(c echo.Context) error {
	p := s.getProblem(c.Param("problem"))

	if p == nil {
		return c.String(http.StatusNotFound, "not found")
	}

	fileLoc := ""

	switch p.(type) {
	case polygon.Problem:
		if len(p.HTMLStatements()) == 0 {
			return c.String(http.StatusNotFound, "file not found")
		}

		if strings.HasSuffix(c.Param("file"), ".css") {
			fileLoc = filepath.Join(s.ProblemsDir, p.Name(), "statements", ".html", p.HTMLStatements()[0].Locale, c.Param("file"))
		} else {
			fileLoc = filepath.Join(s.ProblemsDir, p.Name(), "statements", p.HTMLStatements()[0].Locale, c.Param("file"))
		}

	default:
		return c.String(http.StatusNotFound, "not found")
	}

	return c.Attachment(fileLoc, c.Param("file"))
}

func (s *Server) getProblemsetProblemAttachment(c echo.Context) error {
	p, attachment := s.getProblem(c.Param("problem")), c.Param("attachment")
	if p == nil {
		return c.String(http.StatusNotFound, "no such problem")
	}

	for _, val := range p.Attachments() {
		if val.Name == attachment {
			c.Response().Header().Set("Content-Disposition", "attachment; filename="+val.Name)
			c.Response().Header().Set("Content-Type", mime.TypeByExtension(filepath.Ext(val.Name)))
			c.Response().Header().Set("Content-Length", strconv.Itoa(len(val.Contents)))

			io.Copy(c.Response(), bytes.NewReader(val.Contents))

			return c.NoContent(http.StatusOK)
		}
	}

	return c.String(http.StatusNotFound, "no such attachment")
}

type taskArchiveTreeNode struct {
	Id       int
	Type     string
	Name     string
	Link     string
	Children []*taskArchiveTreeNode
}

//@TODO optimize this to use less queries
func (s *Server) getTaskArchive(c echo.Context) error {
	lst, err := models.ProblemCategories(Where("parent_id IS NULL")).All(s.db)
	if err != nil {
		return s.internalError(c, err, "Belső hiba.")
	}

	roots := make([]*taskArchiveTreeNode, 0)

	var dfs func(category *models.ProblemCategory, node *taskArchiveTreeNode) error
	id := 1000
	dfs = func(root *models.ProblemCategory, tree *taskArchiveTreeNode) error {
		problems, err := models.ProblemRels(Where("category_id = ?", root.ID), OrderBy("problem")).All(s.db)
		if err != nil {
			return err
		}

		for _, problem := range problems {
			tree.Children = append(tree.Children, &taskArchiveTreeNode{id, "problem", translateContent("hungarian", s.getProblem(problem.Problem).Titles()).String(), fmt.Sprintf("/problemset/%s/%s/", problem.Problemset, problem.Problem), make([]*taskArchiveTreeNode, 0)})
			id++
		}

		//@TODO make a way to control sorting order from db (add migrations etc.)
		subcats, err := models.ProblemCategories(Where("parent_id = ?", root.ID), OrderBy("name")).All(s.db)
		if err != nil {
			return err
		}

		for _, cat := range subcats {
			akt := &taskArchiveTreeNode{cat.ID, "category", cat.Name, "", make([]*taskArchiveTreeNode, 0)}
			tree.Children = append(tree.Children, akt)
			if err := dfs(cat, akt); err != nil {
				return err
			}
		}

		return nil
	}

	for _, start := range lst {
		roots = append(roots, &taskArchiveTreeNode{start.ID, "category", start.Name, "", make([]*taskArchiveTreeNode, 0)})
		if dfs(start, roots[len(roots)-1]) != nil {
			return s.internalError(c, err, "Belső hiba.")
		}
	}

	return c.Render(http.StatusOK, "task_archive.gohtml", roots)
}


func (s *Server) getProblemsetProblemRanklist(c echo.Context) error {
	problemSet := c.Param("name")
	problem := c.Param("problem")
	prob := s.getProblem(problem)

	sbs := make([]*models.Submission, 0)

	//@TODO
	if err := queries.Raw("SELECT DISTINCT ON (s1.user_id) s1.* FROM (SELECT s1.user_id, MAX(s1.score) as score FROM submissions s1 WHERE problemset=$1 AND problem=$2 GROUP BY s1.user_id) s2 INNER JOIN submissions s1 ON s1.user_id=s2.user_id AND s1.score=s2.score AND s1.problemset=$1 AND s1.problem=$2", problemSet, problem).Bind(context.TODO(), s.db, &sbs); err != nil {
		return s.internalError(c, err, "hiba")
	}

	sort.Slice(sbs, func(i, j int) bool {
		return sbs[i].Score.Float32 > sbs[j].Score.Float32
	})

	return c.Render(http.StatusOK, "problemset_problem_ranklist.gohtml", struct {
		Problem problems.Problem
		Submissions []*models.Submission
	}{prob, sbs})
}