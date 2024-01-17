// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// DjangoSession is an object representing the database table.
type DjangoSession struct {
	SessionKey  string    `boil:"session_key" json:"session_key" toml:"session_key" yaml:"session_key"`
	SessionData string    `boil:"session_data" json:"session_data" toml:"session_data" yaml:"session_data"`
	ExpireDate  time.Time `boil:"expire_date" json:"expire_date" toml:"expire_date" yaml:"expire_date"`

	R *djangoSessionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L djangoSessionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DjangoSessionColumns = struct {
	SessionKey  string
	SessionData string
	ExpireDate  string
}{
	SessionKey:  "session_key",
	SessionData: "session_data",
	ExpireDate:  "expire_date",
}

var DjangoSessionTableColumns = struct {
	SessionKey  string
	SessionData string
	ExpireDate  string
}{
	SessionKey:  "django_session.session_key",
	SessionData: "django_session.session_data",
	ExpireDate:  "django_session.expire_date",
}

// Generated where

var DjangoSessionWhere = struct {
	SessionKey  whereHelperstring
	SessionData whereHelperstring
	ExpireDate  whereHelpertime_Time
}{
	SessionKey:  whereHelperstring{field: "\"django_session\".\"session_key\""},
	SessionData: whereHelperstring{field: "\"django_session\".\"session_data\""},
	ExpireDate:  whereHelpertime_Time{field: "\"django_session\".\"expire_date\""},
}

// DjangoSessionRels is where relationship names are stored.
var DjangoSessionRels = struct {
}{}

// djangoSessionR is where relationships are stored.
type djangoSessionR struct {
}

// NewStruct creates a new relationship struct
func (*djangoSessionR) NewStruct() *djangoSessionR {
	return &djangoSessionR{}
}

// djangoSessionL is where Load methods for each relationship are stored.
type djangoSessionL struct{}

var (
	djangoSessionAllColumns            = []string{"session_key", "session_data", "expire_date"}
	djangoSessionColumnsWithoutDefault = []string{"session_key", "session_data", "expire_date"}
	djangoSessionColumnsWithDefault    = []string{}
	djangoSessionPrimaryKeyColumns     = []string{"session_key"}
	djangoSessionGeneratedColumns      = []string{}
)

type (
	// DjangoSessionSlice is an alias for a slice of pointers to DjangoSession.
	// This should almost always be used instead of []DjangoSession.
	DjangoSessionSlice []*DjangoSession
	// DjangoSessionHook is the signature for custom DjangoSession hook methods
	DjangoSessionHook func(context.Context, boil.ContextExecutor, *DjangoSession) error

	djangoSessionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	djangoSessionType                 = reflect.TypeOf(&DjangoSession{})
	djangoSessionMapping              = queries.MakeStructMapping(djangoSessionType)
	djangoSessionPrimaryKeyMapping, _ = queries.BindMapping(djangoSessionType, djangoSessionMapping, djangoSessionPrimaryKeyColumns)
	djangoSessionInsertCacheMut       sync.RWMutex
	djangoSessionInsertCache          = make(map[string]insertCache)
	djangoSessionUpdateCacheMut       sync.RWMutex
	djangoSessionUpdateCache          = make(map[string]updateCache)
	djangoSessionUpsertCacheMut       sync.RWMutex
	djangoSessionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var djangoSessionAfterSelectHooks []DjangoSessionHook

var djangoSessionBeforeInsertHooks []DjangoSessionHook
var djangoSessionAfterInsertHooks []DjangoSessionHook

var djangoSessionBeforeUpdateHooks []DjangoSessionHook
var djangoSessionAfterUpdateHooks []DjangoSessionHook

var djangoSessionBeforeDeleteHooks []DjangoSessionHook
var djangoSessionAfterDeleteHooks []DjangoSessionHook

var djangoSessionBeforeUpsertHooks []DjangoSessionHook
var djangoSessionAfterUpsertHooks []DjangoSessionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DjangoSession) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DjangoSession) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DjangoSession) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DjangoSession) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DjangoSession) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DjangoSession) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DjangoSession) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DjangoSession) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DjangoSession) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoSessionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDjangoSessionHook registers your hook function for all future operations.
func AddDjangoSessionHook(hookPoint boil.HookPoint, djangoSessionHook DjangoSessionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		djangoSessionAfterSelectHooks = append(djangoSessionAfterSelectHooks, djangoSessionHook)
	case boil.BeforeInsertHook:
		djangoSessionBeforeInsertHooks = append(djangoSessionBeforeInsertHooks, djangoSessionHook)
	case boil.AfterInsertHook:
		djangoSessionAfterInsertHooks = append(djangoSessionAfterInsertHooks, djangoSessionHook)
	case boil.BeforeUpdateHook:
		djangoSessionBeforeUpdateHooks = append(djangoSessionBeforeUpdateHooks, djangoSessionHook)
	case boil.AfterUpdateHook:
		djangoSessionAfterUpdateHooks = append(djangoSessionAfterUpdateHooks, djangoSessionHook)
	case boil.BeforeDeleteHook:
		djangoSessionBeforeDeleteHooks = append(djangoSessionBeforeDeleteHooks, djangoSessionHook)
	case boil.AfterDeleteHook:
		djangoSessionAfterDeleteHooks = append(djangoSessionAfterDeleteHooks, djangoSessionHook)
	case boil.BeforeUpsertHook:
		djangoSessionBeforeUpsertHooks = append(djangoSessionBeforeUpsertHooks, djangoSessionHook)
	case boil.AfterUpsertHook:
		djangoSessionAfterUpsertHooks = append(djangoSessionAfterUpsertHooks, djangoSessionHook)
	}
}

// One returns a single djangoSession record from the query.
func (q djangoSessionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DjangoSession, error) {
	o := &DjangoSession{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for django_session")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DjangoSession records from the query.
func (q djangoSessionQuery) All(ctx context.Context, exec boil.ContextExecutor) (DjangoSessionSlice, error) {
	var o []*DjangoSession

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DjangoSession slice")
	}

	if len(djangoSessionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DjangoSession records in the query.
func (q djangoSessionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count django_session rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q djangoSessionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if django_session exists")
	}

	return count > 0, nil
}

// DjangoSessions retrieves all the records using an executor.
func DjangoSessions(mods ...qm.QueryMod) djangoSessionQuery {
	mods = append(mods, qm.From("\"django_session\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"django_session\".*"})
	}

	return djangoSessionQuery{q}
}

// FindDjangoSession retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDjangoSession(ctx context.Context, exec boil.ContextExecutor, sessionKey string, selectCols ...string) (*DjangoSession, error) {
	djangoSessionObj := &DjangoSession{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"django_session\" where \"session_key\"=$1", sel,
	)

	q := queries.Raw(query, sessionKey)

	err := q.Bind(ctx, exec, djangoSessionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from django_session")
	}

	if err = djangoSessionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return djangoSessionObj, err
	}

	return djangoSessionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DjangoSession) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no django_session provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(djangoSessionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	djangoSessionInsertCacheMut.RLock()
	cache, cached := djangoSessionInsertCache[key]
	djangoSessionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			djangoSessionAllColumns,
			djangoSessionColumnsWithDefault,
			djangoSessionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(djangoSessionType, djangoSessionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(djangoSessionType, djangoSessionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"django_session\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"django_session\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into django_session")
	}

	if !cached {
		djangoSessionInsertCacheMut.Lock()
		djangoSessionInsertCache[key] = cache
		djangoSessionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DjangoSession.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DjangoSession) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	djangoSessionUpdateCacheMut.RLock()
	cache, cached := djangoSessionUpdateCache[key]
	djangoSessionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			djangoSessionAllColumns,
			djangoSessionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update django_session, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"django_session\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, djangoSessionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(djangoSessionType, djangoSessionMapping, append(wl, djangoSessionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update django_session row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for django_session")
	}

	if !cached {
		djangoSessionUpdateCacheMut.Lock()
		djangoSessionUpdateCache[key] = cache
		djangoSessionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q djangoSessionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for django_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for django_session")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DjangoSessionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), djangoSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"django_session\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, djangoSessionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in djangoSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all djangoSession")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DjangoSession) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no django_session provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(djangoSessionColumnsWithDefault, o)

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

	djangoSessionUpsertCacheMut.RLock()
	cache, cached := djangoSessionUpsertCache[key]
	djangoSessionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			djangoSessionAllColumns,
			djangoSessionColumnsWithDefault,
			djangoSessionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			djangoSessionAllColumns,
			djangoSessionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert django_session, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(djangoSessionPrimaryKeyColumns))
			copy(conflict, djangoSessionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"django_session\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(djangoSessionType, djangoSessionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(djangoSessionType, djangoSessionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert django_session")
	}

	if !cached {
		djangoSessionUpsertCacheMut.Lock()
		djangoSessionUpsertCache[key] = cache
		djangoSessionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DjangoSession record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DjangoSession) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DjangoSession provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), djangoSessionPrimaryKeyMapping)
	sql := "DELETE FROM \"django_session\" WHERE \"session_key\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from django_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for django_session")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q djangoSessionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no djangoSessionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from django_session")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for django_session")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DjangoSessionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(djangoSessionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), djangoSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"django_session\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, djangoSessionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from djangoSession slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for django_session")
	}

	if len(djangoSessionAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DjangoSession) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDjangoSession(ctx, exec, o.SessionKey)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DjangoSessionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DjangoSessionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), djangoSessionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"django_session\".* FROM \"django_session\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, djangoSessionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DjangoSessionSlice")
	}

	*o = slice

	return nil
}

// DjangoSessionExists checks if the DjangoSession row exists.
func DjangoSessionExists(ctx context.Context, exec boil.ContextExecutor, sessionKey string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"django_session\" where \"session_key\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, sessionKey)
	}
	row := exec.QueryRowContext(ctx, sql, sessionKey)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if django_session exists")
	}

	return exists, nil
}

// Exists checks if the DjangoSession row exists.
func (o *DjangoSession) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DjangoSessionExists(ctx, exec, o.SessionKey)
}
