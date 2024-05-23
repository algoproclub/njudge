// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
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

// Partial is an object representing the database table.
type Partial struct {
	Name string `boil:"name" json:"name" toml:"name" yaml:"name"`
	HTML string `boil:"html" json:"html" toml:"html" yaml:"html"`

	R *partialR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L partialL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PartialColumns = struct {
	Name string
	HTML string
}{
	Name: "name",
	HTML: "html",
}

var PartialTableColumns = struct {
	Name string
	HTML string
}{
	Name: "partials.name",
	HTML: "partials.html",
}

// Generated where

var PartialWhere = struct {
	Name whereHelperstring
	HTML whereHelperstring
}{
	Name: whereHelperstring{field: "\"partials\".\"name\""},
	HTML: whereHelperstring{field: "\"partials\".\"html\""},
}

// PartialRels is where relationship names are stored.
var PartialRels = struct {
}{}

// partialR is where relationships are stored.
type partialR struct {
}

// NewStruct creates a new relationship struct
func (*partialR) NewStruct() *partialR {
	return &partialR{}
}

// partialL is where Load methods for each relationship are stored.
type partialL struct{}

var (
	partialAllColumns            = []string{"name", "html"}
	partialColumnsWithoutDefault = []string{"name", "html"}
	partialColumnsWithDefault    = []string{}
	partialPrimaryKeyColumns     = []string{"name"}
	partialGeneratedColumns      = []string{}
)

type (
	// PartialSlice is an alias for a slice of pointers to Partial.
	// This should almost always be used instead of []Partial.
	PartialSlice []*Partial
	// PartialHook is the signature for custom Partial hook methods
	PartialHook func(context.Context, boil.ContextExecutor, *Partial) error

	partialQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	partialType                 = reflect.TypeOf(&Partial{})
	partialMapping              = queries.MakeStructMapping(partialType)
	partialPrimaryKeyMapping, _ = queries.BindMapping(partialType, partialMapping, partialPrimaryKeyColumns)
	partialInsertCacheMut       sync.RWMutex
	partialInsertCache          = make(map[string]insertCache)
	partialUpdateCacheMut       sync.RWMutex
	partialUpdateCache          = make(map[string]updateCache)
	partialUpsertCacheMut       sync.RWMutex
	partialUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var partialAfterSelectMu sync.Mutex
var partialAfterSelectHooks []PartialHook

var partialBeforeInsertMu sync.Mutex
var partialBeforeInsertHooks []PartialHook
var partialAfterInsertMu sync.Mutex
var partialAfterInsertHooks []PartialHook

var partialBeforeUpdateMu sync.Mutex
var partialBeforeUpdateHooks []PartialHook
var partialAfterUpdateMu sync.Mutex
var partialAfterUpdateHooks []PartialHook

var partialBeforeDeleteMu sync.Mutex
var partialBeforeDeleteHooks []PartialHook
var partialAfterDeleteMu sync.Mutex
var partialAfterDeleteHooks []PartialHook

var partialBeforeUpsertMu sync.Mutex
var partialBeforeUpsertHooks []PartialHook
var partialAfterUpsertMu sync.Mutex
var partialAfterUpsertHooks []PartialHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Partial) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Partial) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Partial) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Partial) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Partial) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Partial) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Partial) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Partial) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Partial) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range partialAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPartialHook registers your hook function for all future operations.
func AddPartialHook(hookPoint boil.HookPoint, partialHook PartialHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		partialAfterSelectMu.Lock()
		partialAfterSelectHooks = append(partialAfterSelectHooks, partialHook)
		partialAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		partialBeforeInsertMu.Lock()
		partialBeforeInsertHooks = append(partialBeforeInsertHooks, partialHook)
		partialBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		partialAfterInsertMu.Lock()
		partialAfterInsertHooks = append(partialAfterInsertHooks, partialHook)
		partialAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		partialBeforeUpdateMu.Lock()
		partialBeforeUpdateHooks = append(partialBeforeUpdateHooks, partialHook)
		partialBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		partialAfterUpdateMu.Lock()
		partialAfterUpdateHooks = append(partialAfterUpdateHooks, partialHook)
		partialAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		partialBeforeDeleteMu.Lock()
		partialBeforeDeleteHooks = append(partialBeforeDeleteHooks, partialHook)
		partialBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		partialAfterDeleteMu.Lock()
		partialAfterDeleteHooks = append(partialAfterDeleteHooks, partialHook)
		partialAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		partialBeforeUpsertMu.Lock()
		partialBeforeUpsertHooks = append(partialBeforeUpsertHooks, partialHook)
		partialBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		partialAfterUpsertMu.Lock()
		partialAfterUpsertHooks = append(partialAfterUpsertHooks, partialHook)
		partialAfterUpsertMu.Unlock()
	}
}

// OneG returns a single partial record from the query using the global executor.
func (q partialQuery) OneG(ctx context.Context) (*Partial, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single partial record from the query.
func (q partialQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Partial, error) {
	o := &Partial{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for partials")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all Partial records from the query using the global executor.
func (q partialQuery) AllG(ctx context.Context) (PartialSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Partial records from the query.
func (q partialQuery) All(ctx context.Context, exec boil.ContextExecutor) (PartialSlice, error) {
	var o []*Partial

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Partial slice")
	}

	if len(partialAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all Partial records in the query using the global executor
func (q partialQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Partial records in the query.
func (q partialQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count partials rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table using the global executor.
func (q partialQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q partialQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if partials exists")
	}

	return count > 0, nil
}

// Partials retrieves all the records using an executor.
func Partials(mods ...qm.QueryMod) partialQuery {
	mods = append(mods, qm.From("\"partials\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"partials\".*"})
	}

	return partialQuery{q}
}

// FindPartialG retrieves a single record by ID.
func FindPartialG(ctx context.Context, name string, selectCols ...string) (*Partial, error) {
	return FindPartial(ctx, boil.GetContextDB(), name, selectCols...)
}

// FindPartial retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPartial(ctx context.Context, exec boil.ContextExecutor, name string, selectCols ...string) (*Partial, error) {
	partialObj := &Partial{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"partials\" where \"name\"=$1", sel,
	)

	q := queries.Raw(query, name)

	err := q.Bind(ctx, exec, partialObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from partials")
	}

	if err = partialObj.doAfterSelectHooks(ctx, exec); err != nil {
		return partialObj, err
	}

	return partialObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Partial) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Partial) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no partials provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(partialColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	partialInsertCacheMut.RLock()
	cache, cached := partialInsertCache[key]
	partialInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			partialAllColumns,
			partialColumnsWithDefault,
			partialColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(partialType, partialMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(partialType, partialMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"partials\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"partials\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into partials")
	}

	if !cached {
		partialInsertCacheMut.Lock()
		partialInsertCache[key] = cache
		partialInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// UpdateG a single Partial record using the global executor.
// See Update for more documentation.
func (o *Partial) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Partial.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Partial) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	partialUpdateCacheMut.RLock()
	cache, cached := partialUpdateCache[key]
	partialUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			partialAllColumns,
			partialPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update partials, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"partials\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, partialPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(partialType, partialMapping, append(wl, partialPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update partials row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for partials")
	}

	if !cached {
		partialUpdateCacheMut.Lock()
		partialUpdateCache[key] = cache
		partialUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q partialQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q partialQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for partials")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for partials")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PartialSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PartialSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partialPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"partials\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, partialPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in partial slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all partial")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Partial) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns, opts...)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Partial) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns, opts ...UpsertOptionFunc) error {
	if o == nil {
		return errors.New("models: no partials provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(partialColumnsWithDefault, o)

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

	partialUpsertCacheMut.RLock()
	cache, cached := partialUpsertCache[key]
	partialUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			partialAllColumns,
			partialColumnsWithDefault,
			partialColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			partialAllColumns,
			partialPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert partials, could not build update column list")
		}

		ret := strmangle.SetComplement(partialAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 && updateOnConflict && len(update) != 0 {
			if len(partialPrimaryKeyColumns) == 0 {
				return errors.New("models: unable to upsert partials, could not build conflict column list")
			}

			conflict = make([]string, len(partialPrimaryKeyColumns))
			copy(conflict, partialPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"partials\"", updateOnConflict, ret, update, conflict, insert, opts...)

		cache.valueMapping, err = queries.BindMapping(partialType, partialMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(partialType, partialMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert partials")
	}

	if !cached {
		partialUpsertCacheMut.Lock()
		partialUpsertCache[key] = cache
		partialUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// DeleteG deletes a single Partial record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Partial) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Partial record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Partial) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Partial provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), partialPrimaryKeyMapping)
	sql := "DELETE FROM \"partials\" WHERE \"name\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from partials")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for partials")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

func (q partialQuery) DeleteAllG(ctx context.Context) (int64, error) {
	return q.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all matching rows.
func (q partialQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no partialQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from partials")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for partials")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o PartialSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PartialSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(partialBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partialPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"partials\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, partialPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from partial slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for partials")
	}

	if len(partialAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Partial) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Partial provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Partial) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPartial(ctx, exec, o.Name)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PartialSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty PartialSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PartialSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PartialSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), partialPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"partials\".* FROM \"partials\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, partialPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PartialSlice")
	}

	*o = slice

	return nil
}

// PartialExistsG checks if the Partial row exists.
func PartialExistsG(ctx context.Context, name string) (bool, error) {
	return PartialExists(ctx, boil.GetContextDB(), name)
}

// PartialExists checks if the Partial row exists.
func PartialExists(ctx context.Context, exec boil.ContextExecutor, name string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"partials\" where \"name\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, name)
	}
	row := exec.QueryRowContext(ctx, sql, name)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if partials exists")
	}

	return exists, nil
}

// Exists checks if the Partial row exists.
func (o *Partial) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return PartialExists(ctx, exec, o.Name)
}
