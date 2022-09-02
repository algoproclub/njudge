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
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Judge is an object representing the database table.
type Judge struct {
	ID     int    `boil:"id" json:"id" toml:"id" yaml:"id"`
	State  string `boil:"state" json:"state" toml:"state" yaml:"state"`
	Host   string `boil:"host" json:"host" toml:"host" yaml:"host"`
	Port   string `boil:"port" json:"port" toml:"port" yaml:"port"`
	Ping   int    `boil:"ping" json:"ping" toml:"ping" yaml:"ping"`
	Online bool   `boil:"online" json:"online" toml:"online" yaml:"online"`

	R *judgeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L judgeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var JudgeColumns = struct {
	ID     string
	State  string
	Host   string
	Port   string
	Ping   string
	Online string
}{
	ID:     "id",
	State:  "state",
	Host:   "host",
	Port:   "port",
	Ping:   "ping",
	Online: "online",
}

var JudgeTableColumns = struct {
	ID     string
	State  string
	Host   string
	Port   string
	Ping   string
	Online string
}{
	ID:     "judges.id",
	State:  "judges.state",
	Host:   "judges.host",
	Port:   "judges.port",
	Ping:   "judges.ping",
	Online: "judges.online",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var JudgeWhere = struct {
	ID     whereHelperint
	State  whereHelperstring
	Host   whereHelperstring
	Port   whereHelperstring
	Ping   whereHelperint
	Online whereHelperbool
}{
	ID:     whereHelperint{field: "\"judges\".\"id\""},
	State:  whereHelperstring{field: "\"judges\".\"state\""},
	Host:   whereHelperstring{field: "\"judges\".\"host\""},
	Port:   whereHelperstring{field: "\"judges\".\"port\""},
	Ping:   whereHelperint{field: "\"judges\".\"ping\""},
	Online: whereHelperbool{field: "\"judges\".\"online\""},
}

// JudgeRels is where relationship names are stored.
var JudgeRels = struct {
}{}

// judgeR is where relationships are stored.
type judgeR struct {
}

// NewStruct creates a new relationship struct
func (*judgeR) NewStruct() *judgeR {
	return &judgeR{}
}

// judgeL is where Load methods for each relationship are stored.
type judgeL struct{}

var (
	judgeAllColumns            = []string{"id", "state", "host", "port", "ping", "online"}
	judgeColumnsWithoutDefault = []string{"state", "host", "port"}
	judgeColumnsWithDefault    = []string{"id", "ping", "online"}
	judgePrimaryKeyColumns     = []string{"id"}
	judgeGeneratedColumns      = []string{}
)

type (
	// JudgeSlice is an alias for a slice of pointers to Judge.
	// This should almost always be used instead of []Judge.
	JudgeSlice []*Judge
	// JudgeHook is the signature for custom Judge hook methods
	JudgeHook func(boil.Executor, *Judge) error

	judgeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	judgeType                 = reflect.TypeOf(&Judge{})
	judgeMapping              = queries.MakeStructMapping(judgeType)
	judgePrimaryKeyMapping, _ = queries.BindMapping(judgeType, judgeMapping, judgePrimaryKeyColumns)
	judgeInsertCacheMut       sync.RWMutex
	judgeInsertCache          = make(map[string]insertCache)
	judgeUpdateCacheMut       sync.RWMutex
	judgeUpdateCache          = make(map[string]updateCache)
	judgeUpsertCacheMut       sync.RWMutex
	judgeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var judgeAfterSelectHooks []JudgeHook

var judgeBeforeInsertHooks []JudgeHook
var judgeAfterInsertHooks []JudgeHook

var judgeBeforeUpdateHooks []JudgeHook
var judgeAfterUpdateHooks []JudgeHook

var judgeBeforeDeleteHooks []JudgeHook
var judgeAfterDeleteHooks []JudgeHook

var judgeBeforeUpsertHooks []JudgeHook
var judgeAfterUpsertHooks []JudgeHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Judge) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Judge) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Judge) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Judge) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Judge) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Judge) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Judge) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Judge) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Judge) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range judgeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddJudgeHook registers your hook function for all future operations.
func AddJudgeHook(hookPoint boil.HookPoint, judgeHook JudgeHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		judgeAfterSelectHooks = append(judgeAfterSelectHooks, judgeHook)
	case boil.BeforeInsertHook:
		judgeBeforeInsertHooks = append(judgeBeforeInsertHooks, judgeHook)
	case boil.AfterInsertHook:
		judgeAfterInsertHooks = append(judgeAfterInsertHooks, judgeHook)
	case boil.BeforeUpdateHook:
		judgeBeforeUpdateHooks = append(judgeBeforeUpdateHooks, judgeHook)
	case boil.AfterUpdateHook:
		judgeAfterUpdateHooks = append(judgeAfterUpdateHooks, judgeHook)
	case boil.BeforeDeleteHook:
		judgeBeforeDeleteHooks = append(judgeBeforeDeleteHooks, judgeHook)
	case boil.AfterDeleteHook:
		judgeAfterDeleteHooks = append(judgeAfterDeleteHooks, judgeHook)
	case boil.BeforeUpsertHook:
		judgeBeforeUpsertHooks = append(judgeBeforeUpsertHooks, judgeHook)
	case boil.AfterUpsertHook:
		judgeAfterUpsertHooks = append(judgeAfterUpsertHooks, judgeHook)
	}
}

// OneG returns a single judge record from the query using the global executor.
func (q judgeQuery) OneG() (*Judge, error) {
	return q.One(boil.GetDB())
}

// One returns a single judge record from the query.
func (q judgeQuery) One(exec boil.Executor) (*Judge, error) {
	o := &Judge{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for judges")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Judge records from the query using the global executor.
func (q judgeQuery) AllG() (JudgeSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Judge records from the query.
func (q judgeQuery) All(exec boil.Executor) (JudgeSlice, error) {
	var o []*Judge

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Judge slice")
	}

	if len(judgeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Judge records in the query using the global executor
func (q judgeQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Judge records in the query.
func (q judgeQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count judges rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q judgeQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q judgeQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if judges exists")
	}

	return count > 0, nil
}

// Judges retrieves all the records using an executor.
func Judges(mods ...qm.QueryMod) judgeQuery {
	mods = append(mods, qm.From("\"judges\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"judges\".*"})
	}

	return judgeQuery{q}
}

// FindJudgeG retrieves a single record by ID.
func FindJudgeG(iD int, selectCols ...string) (*Judge, error) {
	return FindJudge(boil.GetDB(), iD, selectCols...)
}

// FindJudge retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindJudge(exec boil.Executor, iD int, selectCols ...string) (*Judge, error) {
	judgeObj := &Judge{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"judges\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, judgeObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from judges")
	}

	if err = judgeObj.doAfterSelectHooks(exec); err != nil {
		return judgeObj, err
	}

	return judgeObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Judge) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Judge) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no judges provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(judgeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	judgeInsertCacheMut.RLock()
	cache, cached := judgeInsertCache[key]
	judgeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			judgeAllColumns,
			judgeColumnsWithDefault,
			judgeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(judgeType, judgeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(judgeType, judgeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"judges\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"judges\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into judges")
	}

	if !cached {
		judgeInsertCacheMut.Lock()
		judgeInsertCache[key] = cache
		judgeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Judge record using the global executor.
// See Update for more documentation.
func (o *Judge) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Judge.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Judge) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	judgeUpdateCacheMut.RLock()
	cache, cached := judgeUpdateCache[key]
	judgeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			judgeAllColumns,
			judgePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update judges, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"judges\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, judgePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(judgeType, judgeMapping, append(wl, judgePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update judges row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for judges")
	}

	if !cached {
		judgeUpdateCacheMut.Lock()
		judgeUpdateCache[key] = cache
		judgeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q judgeQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q judgeQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for judges")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for judges")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o JudgeSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o JudgeSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"judges\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, judgePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in judge slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all judge")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Judge) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Judge) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no judges provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(judgeColumnsWithDefault, o)

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

	judgeUpsertCacheMut.RLock()
	cache, cached := judgeUpsertCache[key]
	judgeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			judgeAllColumns,
			judgeColumnsWithDefault,
			judgeColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			judgeAllColumns,
			judgePrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert judges, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(judgePrimaryKeyColumns))
			copy(conflict, judgePrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"judges\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(judgeType, judgeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(judgeType, judgeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert judges")
	}

	if !cached {
		judgeUpsertCacheMut.Lock()
		judgeUpsertCache[key] = cache
		judgeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single Judge record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Judge) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Judge record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Judge) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Judge provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), judgePrimaryKeyMapping)
	sql := "DELETE FROM \"judges\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from judges")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for judges")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q judgeQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q judgeQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no judgeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from judges")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for judges")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o JudgeSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o JudgeSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(judgeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"judges\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, judgePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from judge slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for judges")
	}

	if len(judgeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Judge) ReloadG() error {
	if o == nil {
		return errors.New("models: no Judge provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Judge) Reload(exec boil.Executor) error {
	ret, err := FindJudge(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JudgeSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty JudgeSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *JudgeSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := JudgeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), judgePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"judges\".* FROM \"judges\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, judgePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in JudgeSlice")
	}

	*o = slice

	return nil
}

// JudgeExistsG checks if the Judge row exists.
func JudgeExistsG(iD int) (bool, error) {
	return JudgeExists(boil.GetDB(), iD)
}

// JudgeExists checks if the Judge row exists.
func JudgeExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"judges\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if judges exists")
	}

	return exists, nil
}
