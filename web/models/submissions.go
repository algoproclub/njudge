// Code generated by SQLBoiler 4.13.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Submission is an object representing the database table.
type Submission struct {
	ID         int          `boil:"id" json:"id" toml:"id" yaml:"id"`
	Status     string       `boil:"status" json:"status" toml:"status" yaml:"status"`
	Ontest     null.String  `boil:"ontest" json:"ontest,omitempty" toml:"ontest" yaml:"ontest,omitempty"`
	UserID     int          `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Problemset string       `boil:"problemset" json:"problemset" toml:"problemset" yaml:"problemset"`
	Problem    string       `boil:"problem" json:"problem" toml:"problem" yaml:"problem"`
	Language   string       `boil:"language" json:"language" toml:"language" yaml:"language"`
	Private    bool         `boil:"private" json:"private" toml:"private" yaml:"private"`
	Verdict    int          `boil:"verdict" json:"verdict" toml:"verdict" yaml:"verdict"`
	Source     []byte       `boil:"source" json:"source" toml:"source" yaml:"source"`
	Started    bool         `boil:"started" json:"started" toml:"started" yaml:"started"`
	Submitted  time.Time    `boil:"submitted" json:"submitted" toml:"submitted" yaml:"submitted"`
	Judged     null.Time    `boil:"judged" json:"judged,omitempty" toml:"judged" yaml:"judged,omitempty"`
	Score      null.Float32 `boil:"score" json:"score,omitempty" toml:"score" yaml:"score,omitempty"`

	R *submissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L submissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var SubmissionColumns = struct {
	ID         string
	Status     string
	Ontest     string
	UserID     string
	Problemset string
	Problem    string
	Language   string
	Private    string
	Verdict    string
	Source     string
	Started    string
	Submitted  string
	Judged     string
	Score      string
}{
	ID:         "id",
	Status:     "status",
	Ontest:     "ontest",
	UserID:     "user_id",
	Problemset: "problemset",
	Problem:    "problem",
	Language:   "language",
	Private:    "private",
	Verdict:    "verdict",
	Source:     "source",
	Started:    "started",
	Submitted:  "submitted",
	Judged:     "judged",
	Score:      "score",
}

var SubmissionTableColumns = struct {
	ID         string
	Status     string
	Ontest     string
	UserID     string
	Problemset string
	Problem    string
	Language   string
	Private    string
	Verdict    string
	Source     string
	Started    string
	Submitted  string
	Judged     string
	Score      string
}{
	ID:         "submissions.id",
	Status:     "submissions.status",
	Ontest:     "submissions.ontest",
	UserID:     "submissions.user_id",
	Problemset: "submissions.problemset",
	Problem:    "submissions.problem",
	Language:   "submissions.language",
	Private:    "submissions.private",
	Verdict:    "submissions.verdict",
	Source:     "submissions.source",
	Started:    "submissions.started",
	Submitted:  "submissions.submitted",
	Judged:     "submissions.judged",
	Score:      "submissions.score",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelper__byte struct{ field string }

func (w whereHelper__byte) EQ(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelper__byte) NEQ(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelper__byte) LT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelper__byte) LTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelper__byte) GT(x []byte) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelper__byte) GTE(x []byte) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Float32 struct{ field string }

func (w whereHelpernull_Float32) EQ(x null.Float32) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float32) NEQ(x null.Float32) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float32) LT(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float32) LTE(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float32) GT(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float32) GTE(x null.Float32) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Float32) IN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Float32) NIN(slice []float32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Float32) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float32) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

var SubmissionWhere = struct {
	ID         whereHelperint
	Status     whereHelperstring
	Ontest     whereHelpernull_String
	UserID     whereHelperint
	Problemset whereHelperstring
	Problem    whereHelperstring
	Language   whereHelperstring
	Private    whereHelperbool
	Verdict    whereHelperint
	Source     whereHelper__byte
	Started    whereHelperbool
	Submitted  whereHelpertime_Time
	Judged     whereHelpernull_Time
	Score      whereHelpernull_Float32
}{
	ID:         whereHelperint{field: "\"submissions\".\"id\""},
	Status:     whereHelperstring{field: "\"submissions\".\"status\""},
	Ontest:     whereHelpernull_String{field: "\"submissions\".\"ontest\""},
	UserID:     whereHelperint{field: "\"submissions\".\"user_id\""},
	Problemset: whereHelperstring{field: "\"submissions\".\"problemset\""},
	Problem:    whereHelperstring{field: "\"submissions\".\"problem\""},
	Language:   whereHelperstring{field: "\"submissions\".\"language\""},
	Private:    whereHelperbool{field: "\"submissions\".\"private\""},
	Verdict:    whereHelperint{field: "\"submissions\".\"verdict\""},
	Source:     whereHelper__byte{field: "\"submissions\".\"source\""},
	Started:    whereHelperbool{field: "\"submissions\".\"started\""},
	Submitted:  whereHelpertime_Time{field: "\"submissions\".\"submitted\""},
	Judged:     whereHelpernull_Time{field: "\"submissions\".\"judged\""},
	Score:      whereHelpernull_Float32{field: "\"submissions\".\"score\""},
}

// SubmissionRels is where relationship names are stored.
var SubmissionRels = struct {
	User string
}{
	User: "User",
}

// submissionR is where relationships are stored.
type submissionR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*submissionR) NewStruct() *submissionR {
	return &submissionR{}
}

func (r *submissionR) GetUser() *User {
	if r == nil {
		return nil
	}
	return r.User
}

// submissionL is where Load methods for each relationship are stored.
type submissionL struct{}

var (
	submissionAllColumns            = []string{"id", "status", "ontest", "user_id", "problemset", "problem", "language", "private", "verdict", "source", "started", "submitted", "judged", "score"}
	submissionColumnsWithoutDefault = []string{"status", "user_id", "problemset", "problem", "language", "private", "verdict", "source", "started", "submitted"}
	submissionColumnsWithDefault    = []string{"id", "ontest", "judged", "score"}
	submissionPrimaryKeyColumns     = []string{"id"}
	submissionGeneratedColumns      = []string{}
)

type (
	// SubmissionSlice is an alias for a slice of pointers to Submission.
	// This should almost always be used instead of []Submission.
	SubmissionSlice []*Submission
	// SubmissionHook is the signature for custom Submission hook methods
	SubmissionHook func(boil.Executor, *Submission) error

	submissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	submissionType                 = reflect.TypeOf(&Submission{})
	submissionMapping              = queries.MakeStructMapping(submissionType)
	submissionPrimaryKeyMapping, _ = queries.BindMapping(submissionType, submissionMapping, submissionPrimaryKeyColumns)
	submissionInsertCacheMut       sync.RWMutex
	submissionInsertCache          = make(map[string]insertCache)
	submissionUpdateCacheMut       sync.RWMutex
	submissionUpdateCache          = make(map[string]updateCache)
	submissionUpsertCacheMut       sync.RWMutex
	submissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var submissionAfterSelectHooks []SubmissionHook

var submissionBeforeInsertHooks []SubmissionHook
var submissionAfterInsertHooks []SubmissionHook

var submissionBeforeUpdateHooks []SubmissionHook
var submissionAfterUpdateHooks []SubmissionHook

var submissionBeforeDeleteHooks []SubmissionHook
var submissionAfterDeleteHooks []SubmissionHook

var submissionBeforeUpsertHooks []SubmissionHook
var submissionAfterUpsertHooks []SubmissionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Submission) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Submission) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Submission) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Submission) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Submission) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Submission) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Submission) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Submission) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Submission) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range submissionAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSubmissionHook registers your hook function for all future operations.
func AddSubmissionHook(hookPoint boil.HookPoint, submissionHook SubmissionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		submissionAfterSelectHooks = append(submissionAfterSelectHooks, submissionHook)
	case boil.BeforeInsertHook:
		submissionBeforeInsertHooks = append(submissionBeforeInsertHooks, submissionHook)
	case boil.AfterInsertHook:
		submissionAfterInsertHooks = append(submissionAfterInsertHooks, submissionHook)
	case boil.BeforeUpdateHook:
		submissionBeforeUpdateHooks = append(submissionBeforeUpdateHooks, submissionHook)
	case boil.AfterUpdateHook:
		submissionAfterUpdateHooks = append(submissionAfterUpdateHooks, submissionHook)
	case boil.BeforeDeleteHook:
		submissionBeforeDeleteHooks = append(submissionBeforeDeleteHooks, submissionHook)
	case boil.AfterDeleteHook:
		submissionAfterDeleteHooks = append(submissionAfterDeleteHooks, submissionHook)
	case boil.BeforeUpsertHook:
		submissionBeforeUpsertHooks = append(submissionBeforeUpsertHooks, submissionHook)
	case boil.AfterUpsertHook:
		submissionAfterUpsertHooks = append(submissionAfterUpsertHooks, submissionHook)
	}
}

// OneG returns a single submission record from the query using the global executor.
func (q submissionQuery) OneG() (*Submission, error) {
	return q.One(boil.GetDB())
}

// One returns a single submission record from the query.
func (q submissionQuery) One(exec boil.Executor) (*Submission, error) {
	o := &Submission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for submissions")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Submission records from the query using the global executor.
func (q submissionQuery) AllG() (SubmissionSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Submission records from the query.
func (q submissionQuery) All(exec boil.Executor) (SubmissionSlice, error) {
	var o []*Submission

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Submission slice")
	}

	if len(submissionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Submission records in the query using the global executor
func (q submissionQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Submission records in the query.
func (q submissionQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count submissions rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q submissionQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q submissionQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if submissions exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Submission) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return Users(queryMods...)
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (submissionL) LoadUser(e boil.Executor, singular bool, maybeSubmission interface{}, mods queries.Applicator) error {
	var slice []*Submission
	var object *Submission

	if singular {
		var ok bool
		object, ok = maybeSubmission.(*Submission)
		if !ok {
			object = new(Submission)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeSubmission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeSubmission))
			}
		}
	} else {
		s, ok := maybeSubmission.(*[]*Submission)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeSubmission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeSubmission))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &submissionR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &submissionR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(submissionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Submissions = append(foreign.R.Submissions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Submissions = append(foreign.R.Submissions, local)
				break
			}
		}
	}

	return nil
}

// SetUserG of the submission to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Submissions.
// Uses the global database handle.
func (o *Submission) SetUserG(insert bool, related *User) error {
	return o.SetUser(boil.GetDB(), insert, related)
}

// SetUser of the submission to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Submissions.
func (o *Submission) SetUser(exec boil.Executor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"submissions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, submissionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	if _, err = exec.Exec(updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &submissionR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Submissions: SubmissionSlice{o},
		}
	} else {
		related.R.Submissions = append(related.R.Submissions, o)
	}

	return nil
}

// Submissions retrieves all the records using an executor.
func Submissions(mods ...qm.QueryMod) submissionQuery {
	mods = append(mods, qm.From("\"submissions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"submissions\".*"})
	}

	return submissionQuery{q}
}

// FindSubmissionG retrieves a single record by ID.
func FindSubmissionG(iD int, selectCols ...string) (*Submission, error) {
	return FindSubmission(boil.GetDB(), iD, selectCols...)
}

// FindSubmission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindSubmission(exec boil.Executor, iD int, selectCols ...string) (*Submission, error) {
	submissionObj := &Submission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"submissions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, submissionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from submissions")
	}

	if err = submissionObj.doAfterSelectHooks(exec); err != nil {
		return submissionObj, err
	}

	return submissionObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Submission) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Submission) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no submissions provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(submissionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	submissionInsertCacheMut.RLock()
	cache, cached := submissionInsertCache[key]
	submissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			submissionAllColumns,
			submissionColumnsWithDefault,
			submissionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(submissionType, submissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(submissionType, submissionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"submissions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"submissions\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into submissions")
	}

	if !cached {
		submissionInsertCacheMut.Lock()
		submissionInsertCache[key] = cache
		submissionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Submission record using the global executor.
// See Update for more documentation.
func (o *Submission) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Submission.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Submission) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	submissionUpdateCacheMut.RLock()
	cache, cached := submissionUpdateCache[key]
	submissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			submissionAllColumns,
			submissionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update submissions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"submissions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, submissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(submissionType, submissionMapping, append(wl, submissionPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update submissions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for submissions")
	}

	if !cached {
		submissionUpdateCacheMut.Lock()
		submissionUpdateCache[key] = cache
		submissionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q submissionQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q submissionQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for submissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for submissions")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o SubmissionSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o SubmissionSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), submissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"submissions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, submissionPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in submission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all submission")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Submission) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Submission) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no submissions provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(submissionColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	submissionUpsertCacheMut.RLock()
	cache, cached := submissionUpsertCache[key]
	submissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			submissionAllColumns,
			submissionColumnsWithDefault,
			submissionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			submissionAllColumns,
			submissionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert submissions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(submissionPrimaryKeyColumns))
			copy(conflict, submissionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"submissions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(submissionType, submissionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(submissionType, submissionMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert submissions")
	}

	if !cached {
		submissionUpsertCacheMut.Lock()
		submissionUpsertCache[key] = cache
		submissionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Submission record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Submission) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Submission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Submission) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Submission provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), submissionPrimaryKeyMapping)
	sql := "DELETE FROM \"submissions\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from submissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for submissions")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q submissionQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q submissionQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no submissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from submissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for submissions")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o SubmissionSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o SubmissionSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(submissionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), submissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"submissions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, submissionPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from submission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for submissions")
	}

	if len(submissionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Submission) ReloadG() error {
	if o == nil {
		return errors.New("models: no Submission provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Submission) Reload(exec boil.Executor) error {
	ret, err := FindSubmission(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubmissionSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty SubmissionSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *SubmissionSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := SubmissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), submissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"submissions\".* FROM \"submissions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, submissionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in SubmissionSlice")
	}

	*o = slice

	return nil
}

// SubmissionExistsG checks if the Submission row exists.
func SubmissionExistsG(iD int) (bool, error) {
	return SubmissionExists(boil.GetDB(), iD)
}

// SubmissionExists checks if the Submission row exists.
func SubmissionExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"submissions\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if submissions exists")
	}

	return exists, nil
}
