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

// DjangoMigration is an object representing the database table.
type DjangoMigration struct {
	ID      int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	App     string    `boil:"app" json:"app" toml:"app" yaml:"app"`
	Name    string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Applied time.Time `boil:"applied" json:"applied" toml:"applied" yaml:"applied"`

	R *djangoMigrationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L djangoMigrationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DjangoMigrationColumns = struct {
	ID      string
	App     string
	Name    string
	Applied string
}{
	ID:      "id",
	App:     "app",
	Name:    "name",
	Applied: "applied",
}

var DjangoMigrationTableColumns = struct {
	ID      string
	App     string
	Name    string
	Applied string
}{
	ID:      "django_migrations.id",
	App:     "django_migrations.app",
	Name:    "django_migrations.name",
	Applied: "django_migrations.applied",
}

// Generated where

var DjangoMigrationWhere = struct {
	ID      whereHelperint
	App     whereHelperstring
	Name    whereHelperstring
	Applied whereHelpertime_Time
}{
	ID:      whereHelperint{field: "\"django_migrations\".\"id\""},
	App:     whereHelperstring{field: "\"django_migrations\".\"app\""},
	Name:    whereHelperstring{field: "\"django_migrations\".\"name\""},
	Applied: whereHelpertime_Time{field: "\"django_migrations\".\"applied\""},
}

// DjangoMigrationRels is where relationship names are stored.
var DjangoMigrationRels = struct {
}{}

// djangoMigrationR is where relationships are stored.
type djangoMigrationR struct {
}

// NewStruct creates a new relationship struct
func (*djangoMigrationR) NewStruct() *djangoMigrationR {
	return &djangoMigrationR{}
}

// djangoMigrationL is where Load methods for each relationship are stored.
type djangoMigrationL struct{}

var (
	djangoMigrationAllColumns            = []string{"id", "app", "name", "applied"}
	djangoMigrationColumnsWithoutDefault = []string{"app", "name", "applied"}
	djangoMigrationColumnsWithDefault    = []string{"id"}
	djangoMigrationPrimaryKeyColumns     = []string{"id"}
	djangoMigrationGeneratedColumns      = []string{}
)

type (
	// DjangoMigrationSlice is an alias for a slice of pointers to DjangoMigration.
	// This should almost always be used instead of []DjangoMigration.
	DjangoMigrationSlice []*DjangoMigration
	// DjangoMigrationHook is the signature for custom DjangoMigration hook methods
	DjangoMigrationHook func(context.Context, boil.ContextExecutor, *DjangoMigration) error

	djangoMigrationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	djangoMigrationType                 = reflect.TypeOf(&DjangoMigration{})
	djangoMigrationMapping              = queries.MakeStructMapping(djangoMigrationType)
	djangoMigrationPrimaryKeyMapping, _ = queries.BindMapping(djangoMigrationType, djangoMigrationMapping, djangoMigrationPrimaryKeyColumns)
	djangoMigrationInsertCacheMut       sync.RWMutex
	djangoMigrationInsertCache          = make(map[string]insertCache)
	djangoMigrationUpdateCacheMut       sync.RWMutex
	djangoMigrationUpdateCache          = make(map[string]updateCache)
	djangoMigrationUpsertCacheMut       sync.RWMutex
	djangoMigrationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var djangoMigrationAfterSelectHooks []DjangoMigrationHook

var djangoMigrationBeforeInsertHooks []DjangoMigrationHook
var djangoMigrationAfterInsertHooks []DjangoMigrationHook

var djangoMigrationBeforeUpdateHooks []DjangoMigrationHook
var djangoMigrationAfterUpdateHooks []DjangoMigrationHook

var djangoMigrationBeforeDeleteHooks []DjangoMigrationHook
var djangoMigrationAfterDeleteHooks []DjangoMigrationHook

var djangoMigrationBeforeUpsertHooks []DjangoMigrationHook
var djangoMigrationAfterUpsertHooks []DjangoMigrationHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DjangoMigration) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DjangoMigration) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DjangoMigration) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DjangoMigration) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DjangoMigration) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DjangoMigration) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DjangoMigration) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DjangoMigration) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DjangoMigration) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range djangoMigrationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDjangoMigrationHook registers your hook function for all future operations.
func AddDjangoMigrationHook(hookPoint boil.HookPoint, djangoMigrationHook DjangoMigrationHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		djangoMigrationAfterSelectHooks = append(djangoMigrationAfterSelectHooks, djangoMigrationHook)
	case boil.BeforeInsertHook:
		djangoMigrationBeforeInsertHooks = append(djangoMigrationBeforeInsertHooks, djangoMigrationHook)
	case boil.AfterInsertHook:
		djangoMigrationAfterInsertHooks = append(djangoMigrationAfterInsertHooks, djangoMigrationHook)
	case boil.BeforeUpdateHook:
		djangoMigrationBeforeUpdateHooks = append(djangoMigrationBeforeUpdateHooks, djangoMigrationHook)
	case boil.AfterUpdateHook:
		djangoMigrationAfterUpdateHooks = append(djangoMigrationAfterUpdateHooks, djangoMigrationHook)
	case boil.BeforeDeleteHook:
		djangoMigrationBeforeDeleteHooks = append(djangoMigrationBeforeDeleteHooks, djangoMigrationHook)
	case boil.AfterDeleteHook:
		djangoMigrationAfterDeleteHooks = append(djangoMigrationAfterDeleteHooks, djangoMigrationHook)
	case boil.BeforeUpsertHook:
		djangoMigrationBeforeUpsertHooks = append(djangoMigrationBeforeUpsertHooks, djangoMigrationHook)
	case boil.AfterUpsertHook:
		djangoMigrationAfterUpsertHooks = append(djangoMigrationAfterUpsertHooks, djangoMigrationHook)
	}
}

// One returns a single djangoMigration record from the query.
func (q djangoMigrationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DjangoMigration, error) {
	o := &DjangoMigration{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for django_migrations")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DjangoMigration records from the query.
func (q djangoMigrationQuery) All(ctx context.Context, exec boil.ContextExecutor) (DjangoMigrationSlice, error) {
	var o []*DjangoMigration

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DjangoMigration slice")
	}

	if len(djangoMigrationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DjangoMigration records in the query.
func (q djangoMigrationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count django_migrations rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q djangoMigrationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if django_migrations exists")
	}

	return count > 0, nil
}

// DjangoMigrations retrieves all the records using an executor.
func DjangoMigrations(mods ...qm.QueryMod) djangoMigrationQuery {
	mods = append(mods, qm.From("\"django_migrations\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"django_migrations\".*"})
	}

	return djangoMigrationQuery{q}
}

// FindDjangoMigration retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDjangoMigration(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*DjangoMigration, error) {
	djangoMigrationObj := &DjangoMigration{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"django_migrations\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, djangoMigrationObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from django_migrations")
	}

	if err = djangoMigrationObj.doAfterSelectHooks(ctx, exec); err != nil {
		return djangoMigrationObj, err
	}

	return djangoMigrationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DjangoMigration) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no django_migrations provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(djangoMigrationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	djangoMigrationInsertCacheMut.RLock()
	cache, cached := djangoMigrationInsertCache[key]
	djangoMigrationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			djangoMigrationAllColumns,
			djangoMigrationColumnsWithDefault,
			djangoMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(djangoMigrationType, djangoMigrationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(djangoMigrationType, djangoMigrationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"django_migrations\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"django_migrations\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into django_migrations")
	}

	if !cached {
		djangoMigrationInsertCacheMut.Lock()
		djangoMigrationInsertCache[key] = cache
		djangoMigrationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DjangoMigration.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DjangoMigration) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	djangoMigrationUpdateCacheMut.RLock()
	cache, cached := djangoMigrationUpdateCache[key]
	djangoMigrationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			djangoMigrationAllColumns,
			djangoMigrationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update django_migrations, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"django_migrations\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, djangoMigrationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(djangoMigrationType, djangoMigrationMapping, append(wl, djangoMigrationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update django_migrations row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for django_migrations")
	}

	if !cached {
		djangoMigrationUpdateCacheMut.Lock()
		djangoMigrationUpdateCache[key] = cache
		djangoMigrationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q djangoMigrationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for django_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for django_migrations")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DjangoMigrationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), djangoMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"django_migrations\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, djangoMigrationPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in djangoMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all djangoMigration")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DjangoMigration) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no django_migrations provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(djangoMigrationColumnsWithDefault, o)

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

	djangoMigrationUpsertCacheMut.RLock()
	cache, cached := djangoMigrationUpsertCache[key]
	djangoMigrationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			djangoMigrationAllColumns,
			djangoMigrationColumnsWithDefault,
			djangoMigrationColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			djangoMigrationAllColumns,
			djangoMigrationPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert django_migrations, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(djangoMigrationPrimaryKeyColumns))
			copy(conflict, djangoMigrationPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"django_migrations\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(djangoMigrationType, djangoMigrationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(djangoMigrationType, djangoMigrationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert django_migrations")
	}

	if !cached {
		djangoMigrationUpsertCacheMut.Lock()
		djangoMigrationUpsertCache[key] = cache
		djangoMigrationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DjangoMigration record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DjangoMigration) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DjangoMigration provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), djangoMigrationPrimaryKeyMapping)
	sql := "DELETE FROM \"django_migrations\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from django_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for django_migrations")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q djangoMigrationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no djangoMigrationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from django_migrations")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for django_migrations")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DjangoMigrationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(djangoMigrationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), djangoMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"django_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, djangoMigrationPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from djangoMigration slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for django_migrations")
	}

	if len(djangoMigrationAfterDeleteHooks) != 0 {
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
func (o *DjangoMigration) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDjangoMigration(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DjangoMigrationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DjangoMigrationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), djangoMigrationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"django_migrations\".* FROM \"django_migrations\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, djangoMigrationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DjangoMigrationSlice")
	}

	*o = slice

	return nil
}

// DjangoMigrationExists checks if the DjangoMigration row exists.
func DjangoMigrationExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"django_migrations\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if django_migrations exists")
	}

	return exists, nil
}

// Exists checks if the DjangoMigration row exists.
func (o *DjangoMigration) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DjangoMigrationExists(ctx, exec, o.ID)
}
