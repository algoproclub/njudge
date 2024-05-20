package problemset

import (
	"fmt"
	"github.com/a-h/templ"
	"github.com/mraron/njudge/internal/web/templates"
	"io"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mraron/njudge/internal/njudge"
	"github.com/mraron/njudge/internal/web/helpers/i18n"
	"github.com/mraron/njudge/pkg/problems"
)

type ProblemListRequest struct {
	Page   int    `query:"page"`
	Order_ string `query:"order"`
	Order  njudge.SortDirection
	By_    string `query:"by"`
	By     njudge.ProblemSortField

	TitleFilter    string `query:"title"`
	CategoryFilter int    `query:"category"`
	TagFilter      string `query:"tags"`

	Problemset string `param:"name"`
}

func NewProblemListRequest(c echo.Context) (*ProblemListRequest, error) {
	data := ProblemListRequest{}
	if err := c.Bind(&data); err != nil {
		return nil, err
	}
	if data.Page <= 0 {
		data.Page = 1
	}

	data.Order, data.By = njudge.SortDESC, njudge.ProblemSortFieldID
	if c.QueryParam("by") == "solver_count" {
		data.By = njudge.ProblemSortFieldSolverCount
	}
	if c.QueryParam("order") == "ASC" {
		data.Order = njudge.SortASC
	}
	return &data, nil
}

func GetProblemList(store problems.Store, ps njudge.Problems, cs njudge.Categories, problemListQuery njudge.ProblemListQuery, pinfo njudge.ProblemInfoQuery, tags njudge.Tags) echo.HandlerFunc {

	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)
		data, err := NewProblemListRequest(c)
		if err != nil {
			return err
		}

		listRequest := njudge.ProblemListRequest{
			Problemset:  data.Problemset,
			Page:        data.Page,
			PerPage:     20,
			SortDir:     data.Order,
			SortField:   data.By,
			TitleFilter: data.TitleFilter,
			User:        c.Get("user").(*njudge.User),
		}

		if data.TagFilter != "" {
			listRequest.TagFilter = strings.Split(data.TagFilter, ",")
		}

		if data.CategoryFilter != 0 {
			if data.CategoryFilter == -1 {
				listRequest.CategoryFilter = njudge.NewCategoryEmptyFilter()
			} else {
				listRequest.CategoryFilter = njudge.NewCategoryIDFilter(data.CategoryFilter)
			}
		}

		problemList, err := problemListQuery.GetProblemList(c.Request().Context(), listRequest)
		if err != nil {
			return err
		}

		u := *c.Request().URL
		links, err := templates.Links(problemList.PaginationData.Page, problemList.PaginationData.PerPage, int64(problemList.PaginationData.Count), u.Query())
		if err != nil {
			return err
		}

		tagsList, err := tags.GetAll(c.Request().Context())
		if err != nil {
			return err
		}
		result := templates.ProblemListViewModel{
			Pages: links,
			Tags:  tagsList,
		}

		categories, err := cs.GetAll(c.Request().Context())
		if err != nil {
			return err
		}

		par := make(map[int]int)
		for ind := range categories {
			if categories[ind].ParentID.Valid {
				par[categories[ind].ID] = categories[ind].ParentID.Int
			}
		}

		categoryNameByID := make(map[int]string)
		for ind := range categories {
			categoryNameByID[categories[ind].ID] = categories[ind].Name
		}

		var getCategoryNameRec func(int) string
		getCategoryNameRec = func(id int) string {
			if _, ok := par[id]; !ok {
				return categoryNameByID[id]
			} else {
				return getCategoryNameRec(par[id]) + " -- " + categoryNameByID[id]
			}
		}

		for ind := range problemList.Problems {
			p, err := ps.Get(c.Request().Context(), problemList.Problems[ind].ID)
			if err != nil {
				return err
			}
			info, err := pinfo.GetProblemData(c.Request().Context(), p.ID, c.Get("userID").(int))
			if err != nil {
				return err
			}
			data, err := p.WithStoredData(store)
			if err != nil {
				return err
			}

			curr := templates.ProblemListProblem{
				Name:        p.Problem,
				Titles:      data.Titles(),
				Visible:     p.Visible,
				UserInfo:    info.UserInfo,
				ShowTags:    true,
				Tags:        nil,
				SolverCount: p.SolverCount,
			}

			for _, tag := range p.Tags {
				curr.Tags = append(curr.Tags, tag.Tag)
			}

			if u := c.Get(templates.UserContextKey).(*njudge.User); u != nil {
				if info.UserInfo.SolvedStatus != njudge.Solved && !u.Settings.ShowUnsolvedTags {
					curr.ShowTags = false
				}
			}

			if p.Category != nil {
				cid := p.Category.ID
				for {
					if _, ok := par[cid]; ok {
						cid = par[cid]
					} else {
						break
					}
				}

				curr.CategoryLink = templates.Link{
					Text: categoryNameByID[cid],
					Href: templ.SafeURL(fmt.Sprintf("/task_archive#category%d", p.Category.ID)),
				}
			}

			result.Problems = append(result.Problems, curr)
		}

		u = *c.Request().URL
		qu := u.Query()
		if data.By == njudge.ProblemSortFieldSolverCount {
			if qu.Get("order") == "DESC" {
				qu.Set("order", "ASC")
			} else {
				qu.Set("order", "")
				qu.Set("by", "")
			}
		} else {
			qu.Set("by", "solver_count")
			qu.Set("order", "DESC")
		}
		result.SolverSorter = templates.SortColumn{
			Order: data.Order,
			Href:  "?" + qu.Encode(),
		}

		result.Filtered = listRequest.IsFiltered()
		result.TitleFilter = data.TitleFilter
		result.TagsFilter = data.TagFilter
		result.CategoryFilterOptions = []templates.CategoryFilterOption{
			{Name: "-"},
		}

		emptySelected := false
		if data.CategoryFilter == -1 {
			emptySelected = true
		}
		result.CategoryFilterOptions = append(result.CategoryFilterOptions, templates.CategoryFilterOption{
			Name:     tr.Translate("No category"),
			Value:    "-1",
			Selected: emptySelected,
		})

		for ind := range categories {
			curr := templates.CategoryFilterOption{
				Name:     getCategoryNameRec(categories[ind].ID),
				Value:    strconv.Itoa(categories[ind].ID),
				Selected: false,
			}

			if strconv.Itoa(categories[ind].ID) == c.QueryParam("category") {
				curr.Selected = true
			}

			result.CategoryFilterOptions = append(result.CategoryFilterOptions, curr)
		}

		sort.Slice(result.CategoryFilterOptions, func(i, j int) bool {
			return result.CategoryFilterOptions[i].Name < result.CategoryFilterOptions[j].Name
		})

		c.Set("title", tr.Translate("Problems"))
		return templates.Render(c, http.StatusOK, templates.ProblemList(result))
	}
}

type GetStatusRequest struct {
	AC         string `query:"ac"`
	UserID     int    `query:"user_id"`
	Problemset string `query:"problem_set"`
	Problem    string `query:"problem"`
	Page       int    `query:"page"`
}

func GetStatus(subList njudge.SubmissionListQuery) echo.HandlerFunc {
	return func(c echo.Context) error {
		tr := c.Get(i18n.TranslatorContextKey).(i18n.Translator)

		data := GetStatusRequest{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		if data.Page <= 0 {
			data.Page = 1
		}

		statusReq := njudge.SubmissionListRequest{
			Page:       data.Page,
			PerPage:    20,
			SortDir:    njudge.SortDESC,
			SortField:  njudge.SubmissionSortFieldID,
			Problemset: data.Problemset,
			Problem:    data.Problem,
			UserID:     data.UserID,
		}

		if data.AC == "1" {
			ac := njudge.VerdictAC
			statusReq.Verdict = &ac
		}

		submissionList, err := subList.GetPagedSubmissionList(c.Request().Context(), statusReq)
		if err != nil {
			return err
		}

		qu := (*c.Request().URL).Query()
		links, err := templates.LinksWithCountLimit(submissionList.PaginationData.Page, submissionList.PaginationData.PerPage, int64(submissionList.PaginationData.Count), qu, 5)
		if err != nil {
			return err
		}

		result := templates.SubmissionsViewModel{
			Submissions: submissionList.Submissions,
			Pages:       links,
		}

		c.Set("title", tr.Translate("Submissions"))
		return templates.Render(c, http.StatusOK, templates.Status(result))
	}
}

func PostSubmit(subService njudge.SubmitService) echo.HandlerFunc {
	type request struct {
		Problemset     string `param:"name"`
		ProblemName    string `form:"problem"`
		LanguageName   string `form:"language"`
		SubmissionCode string `form:"submissionCode"`
	}
	return func(c echo.Context) error {
		u := c.Get("user").(*njudge.User)

		data := request{}
		if err := c.Bind(&data); err != nil {
			return err
		}

		code := data.SubmissionCode
		if len(code) == 0 {
			fileHeader, err := c.FormFile("source")
			if err != nil {
				return err
			}

			f, err := fileHeader.Open()
			if err != nil {
				return err
			}

			contents, err := io.ReadAll(f)
			if err != nil {
				return err
			}

			code = string(contents)
			if err := f.Close(); err != nil {
				return err
			}
		}

		sub, err := subService.Submit(c.Request().Context(), njudge.SubmitRequest{
			UserID:     u.ID,
			Problemset: data.Problemset,
			Problem:    data.ProblemName,
			Language:   data.LanguageName,
			Source:     []byte(code),
		})
		if err != nil {
			return err
		}

		return c.Redirect(http.StatusFound, c.Echo().Reverse("getProblemsetStatus")+"#submission"+strconv.Itoa(sub.ID))
	}
}
