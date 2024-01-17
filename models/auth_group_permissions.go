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

// AuthGroupPermission is an object representing the database table.
type AuthGroupPermission struct {
	ID           int `boil:"id" json:"id" toml:"id" yaml:"id"`
	GroupID      int `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`
	PermissionID int `boil:"permission_id" json:"permission_id" toml:"permission_id" yaml:"permission_id"`

	R *authGroupPermissionR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authGroupPermissionL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthGroupPermissionColumns = struct {
	ID           string
	GroupID      string
	PermissionID string
}{
	ID:           "id",
	GroupID:      "group_id",
	PermissionID: "permission_id",
}

var AuthGroupPermissionTableColumns = struct {
	ID           string
	GroupID      string
	PermissionID string
}{
	ID:           "auth_group_permissions.id",
	GroupID:      "auth_group_permissions.group_id",
	PermissionID: "auth_group_permissions.permission_id",
}

// Generated where

var AuthGroupPermissionWhere = struct {
	ID           whereHelperint
	GroupID      whereHelperint
	PermissionID whereHelperint
}{
	ID:           whereHelperint{field: "\"auth_group_permissions\".\"id\""},
	GroupID:      whereHelperint{field: "\"auth_group_permissions\".\"group_id\""},
	PermissionID: whereHelperint{field: "\"auth_group_permissions\".\"permission_id\""},
}

// AuthGroupPermissionRels is where relationship names are stored.
var AuthGroupPermissionRels = struct {
	Permission string
	Group      string
}{
	Permission: "Permission",
	Group:      "Group",
}

// authGroupPermissionR is where relationships are stored.
type authGroupPermissionR struct {
	Permission *AuthPermission `boil:"Permission" json:"Permission" toml:"Permission" yaml:"Permission"`
	Group      *AuthGroup      `boil:"Group" json:"Group" toml:"Group" yaml:"Group"`
}

// NewStruct creates a new relationship struct
func (*authGroupPermissionR) NewStruct() *authGroupPermissionR {
	return &authGroupPermissionR{}
}

func (r *authGroupPermissionR) GetPermission() *AuthPermission {
	if r == nil {
		return nil
	}
	return r.Permission
}

func (r *authGroupPermissionR) GetGroup() *AuthGroup {
	if r == nil {
		return nil
	}
	return r.Group
}

// authGroupPermissionL is where Load methods for each relationship are stored.
type authGroupPermissionL struct{}

var (
	authGroupPermissionAllColumns            = []string{"id", "group_id", "permission_id"}
	authGroupPermissionColumnsWithoutDefault = []string{"group_id", "permission_id"}
	authGroupPermissionColumnsWithDefault    = []string{"id"}
	authGroupPermissionPrimaryKeyColumns     = []string{"id"}
	authGroupPermissionGeneratedColumns      = []string{}
)

type (
	// AuthGroupPermissionSlice is an alias for a slice of pointers to AuthGroupPermission.
	// This should almost always be used instead of []AuthGroupPermission.
	AuthGroupPermissionSlice []*AuthGroupPermission
	// AuthGroupPermissionHook is the signature for custom AuthGroupPermission hook methods
	AuthGroupPermissionHook func(context.Context, boil.ContextExecutor, *AuthGroupPermission) error

	authGroupPermissionQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authGroupPermissionType                 = reflect.TypeOf(&AuthGroupPermission{})
	authGroupPermissionMapping              = queries.MakeStructMapping(authGroupPermissionType)
	authGroupPermissionPrimaryKeyMapping, _ = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, authGroupPermissionPrimaryKeyColumns)
	authGroupPermissionInsertCacheMut       sync.RWMutex
	authGroupPermissionInsertCache          = make(map[string]insertCache)
	authGroupPermissionUpdateCacheMut       sync.RWMutex
	authGroupPermissionUpdateCache          = make(map[string]updateCache)
	authGroupPermissionUpsertCacheMut       sync.RWMutex
	authGroupPermissionUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var authGroupPermissionAfterSelectHooks []AuthGroupPermissionHook

var authGroupPermissionBeforeInsertHooks []AuthGroupPermissionHook
var authGroupPermissionAfterInsertHooks []AuthGroupPermissionHook

var authGroupPermissionBeforeUpdateHooks []AuthGroupPermissionHook
var authGroupPermissionAfterUpdateHooks []AuthGroupPermissionHook

var authGroupPermissionBeforeDeleteHooks []AuthGroupPermissionHook
var authGroupPermissionAfterDeleteHooks []AuthGroupPermissionHook

var authGroupPermissionBeforeUpsertHooks []AuthGroupPermissionHook
var authGroupPermissionAfterUpsertHooks []AuthGroupPermissionHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthGroupPermission) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthGroupPermission) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthGroupPermission) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthGroupPermission) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthGroupPermission) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthGroupPermission) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthGroupPermission) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthGroupPermission) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthGroupPermission) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authGroupPermissionAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthGroupPermissionHook registers your hook function for all future operations.
func AddAuthGroupPermissionHook(hookPoint boil.HookPoint, authGroupPermissionHook AuthGroupPermissionHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		authGroupPermissionAfterSelectHooks = append(authGroupPermissionAfterSelectHooks, authGroupPermissionHook)
	case boil.BeforeInsertHook:
		authGroupPermissionBeforeInsertHooks = append(authGroupPermissionBeforeInsertHooks, authGroupPermissionHook)
	case boil.AfterInsertHook:
		authGroupPermissionAfterInsertHooks = append(authGroupPermissionAfterInsertHooks, authGroupPermissionHook)
	case boil.BeforeUpdateHook:
		authGroupPermissionBeforeUpdateHooks = append(authGroupPermissionBeforeUpdateHooks, authGroupPermissionHook)
	case boil.AfterUpdateHook:
		authGroupPermissionAfterUpdateHooks = append(authGroupPermissionAfterUpdateHooks, authGroupPermissionHook)
	case boil.BeforeDeleteHook:
		authGroupPermissionBeforeDeleteHooks = append(authGroupPermissionBeforeDeleteHooks, authGroupPermissionHook)
	case boil.AfterDeleteHook:
		authGroupPermissionAfterDeleteHooks = append(authGroupPermissionAfterDeleteHooks, authGroupPermissionHook)
	case boil.BeforeUpsertHook:
		authGroupPermissionBeforeUpsertHooks = append(authGroupPermissionBeforeUpsertHooks, authGroupPermissionHook)
	case boil.AfterUpsertHook:
		authGroupPermissionAfterUpsertHooks = append(authGroupPermissionAfterUpsertHooks, authGroupPermissionHook)
	}
}

// One returns a single authGroupPermission record from the query.
func (q authGroupPermissionQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuthGroupPermission, error) {
	o := &AuthGroupPermission{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_group_permissions")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuthGroupPermission records from the query.
func (q authGroupPermissionQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthGroupPermissionSlice, error) {
	var o []*AuthGroupPermission

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthGroupPermission slice")
	}

	if len(authGroupPermissionAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuthGroupPermission records in the query.
func (q authGroupPermissionQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_group_permissions rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q authGroupPermissionQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_group_permissions exists")
	}

	return count > 0, nil
}

// Permission pointed to by the foreign key.
func (o *AuthGroupPermission) Permission(mods ...qm.QueryMod) authPermissionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.PermissionID),
	}

	queryMods = append(queryMods, mods...)

	return AuthPermissions(queryMods...)
}

// Group pointed to by the foreign key.
func (o *AuthGroupPermission) Group(mods ...qm.QueryMod) authGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	return AuthGroups(queryMods...)
}

// LoadPermission allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (authGroupPermissionL) LoadPermission(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuthGroupPermission interface{}, mods queries.Applicator) error {
	var slice []*AuthGroupPermission
	var object *AuthGroupPermission

	if singular {
		var ok bool
		object, ok = maybeAuthGroupPermission.(*AuthGroupPermission)
		if !ok {
			object = new(AuthGroupPermission)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAuthGroupPermission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAuthGroupPermission))
			}
		}
	} else {
		s, ok := maybeAuthGroupPermission.(*[]*AuthGroupPermission)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAuthGroupPermission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAuthGroupPermission))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &authGroupPermissionR{}
		}
		args = append(args, object.PermissionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &authGroupPermissionR{}
			}

			for _, a := range args {
				if a == obj.PermissionID {
					continue Outer
				}
			}

			args = append(args, obj.PermissionID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`auth_permission`),
		qm.WhereIn(`auth_permission.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthPermission")
	}

	var resultSlice []*AuthPermission
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthPermission")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for auth_permission")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for auth_permission")
	}

	if len(authPermissionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Permission = foreign
		if foreign.R == nil {
			foreign.R = &authPermissionR{}
		}
		foreign.R.PermissionAuthGroupPermissions = append(foreign.R.PermissionAuthGroupPermissions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PermissionID == foreign.ID {
				local.R.Permission = foreign
				if foreign.R == nil {
					foreign.R = &authPermissionR{}
				}
				foreign.R.PermissionAuthGroupPermissions = append(foreign.R.PermissionAuthGroupPermissions, local)
				break
			}
		}
	}

	return nil
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (authGroupPermissionL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuthGroupPermission interface{}, mods queries.Applicator) error {
	var slice []*AuthGroupPermission
	var object *AuthGroupPermission

	if singular {
		var ok bool
		object, ok = maybeAuthGroupPermission.(*AuthGroupPermission)
		if !ok {
			object = new(AuthGroupPermission)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAuthGroupPermission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAuthGroupPermission))
			}
		}
	} else {
		s, ok := maybeAuthGroupPermission.(*[]*AuthGroupPermission)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAuthGroupPermission)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAuthGroupPermission))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &authGroupPermissionR{}
		}
		args = append(args, object.GroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &authGroupPermissionR{}
			}

			for _, a := range args {
				if a == obj.GroupID {
					continue Outer
				}
			}

			args = append(args, obj.GroupID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`auth_group`),
		qm.WhereIn(`auth_group.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthGroup")
	}

	var resultSlice []*AuthGroup
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthGroup")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for auth_group")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for auth_group")
	}

	if len(authGroupAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Group = foreign
		if foreign.R == nil {
			foreign.R = &authGroupR{}
		}
		foreign.R.GroupAuthGroupPermissions = append(foreign.R.GroupAuthGroupPermissions, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.ID {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &authGroupR{}
				}
				foreign.R.GroupAuthGroupPermissions = append(foreign.R.GroupAuthGroupPermissions, local)
				break
			}
		}
	}

	return nil
}

// SetPermission of the authGroupPermission to the related item.
// Sets o.R.Permission to related.
// Adds o to related.R.PermissionAuthGroupPermissions.
func (o *AuthGroupPermission) SetPermission(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AuthPermission) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_group_permissions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"permission_id"}),
		strmangle.WhereClause("\"", "\"", 2, authGroupPermissionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PermissionID = related.ID
	if o.R == nil {
		o.R = &authGroupPermissionR{
			Permission: related,
		}
	} else {
		o.R.Permission = related
	}

	if related.R == nil {
		related.R = &authPermissionR{
			PermissionAuthGroupPermissions: AuthGroupPermissionSlice{o},
		}
	} else {
		related.R.PermissionAuthGroupPermissions = append(related.R.PermissionAuthGroupPermissions, o)
	}

	return nil
}

// SetGroup of the authGroupPermission to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAuthGroupPermissions.
func (o *AuthGroupPermission) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AuthGroup) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_group_permissions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"group_id"}),
		strmangle.WhereClause("\"", "\"", 2, authGroupPermissionPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.GroupID = related.ID
	if o.R == nil {
		o.R = &authGroupPermissionR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &authGroupR{
			GroupAuthGroupPermissions: AuthGroupPermissionSlice{o},
		}
	} else {
		related.R.GroupAuthGroupPermissions = append(related.R.GroupAuthGroupPermissions, o)
	}

	return nil
}

// AuthGroupPermissions retrieves all the records using an executor.
func AuthGroupPermissions(mods ...qm.QueryMod) authGroupPermissionQuery {
	mods = append(mods, qm.From("\"auth_group_permissions\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"auth_group_permissions\".*"})
	}

	return authGroupPermissionQuery{q}
}

// FindAuthGroupPermission retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthGroupPermission(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AuthGroupPermission, error) {
	authGroupPermissionObj := &AuthGroupPermission{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_group_permissions\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, authGroupPermissionObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_group_permissions")
	}

	if err = authGroupPermissionObj.doAfterSelectHooks(ctx, exec); err != nil {
		return authGroupPermissionObj, err
	}

	return authGroupPermissionObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuthGroupPermission) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_group_permissions provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authGroupPermissionColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authGroupPermissionInsertCacheMut.RLock()
	cache, cached := authGroupPermissionInsertCache[key]
	authGroupPermissionInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authGroupPermissionAllColumns,
			authGroupPermissionColumnsWithDefault,
			authGroupPermissionColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"auth_group_permissions\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"auth_group_permissions\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into auth_group_permissions")
	}

	if !cached {
		authGroupPermissionInsertCacheMut.Lock()
		authGroupPermissionInsertCache[key] = cache
		authGroupPermissionInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuthGroupPermission.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuthGroupPermission) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	authGroupPermissionUpdateCacheMut.RLock()
	cache, cached := authGroupPermissionUpdateCache[key]
	authGroupPermissionUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authGroupPermissionAllColumns,
			authGroupPermissionPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update auth_group_permissions, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_group_permissions\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authGroupPermissionPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, append(wl, authGroupPermissionPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update auth_group_permissions row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for auth_group_permissions")
	}

	if !cached {
		authGroupPermissionUpdateCacheMut.Lock()
		authGroupPermissionUpdateCache[key] = cache
		authGroupPermissionUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q authGroupPermissionQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for auth_group_permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for auth_group_permissions")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthGroupPermissionSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authGroupPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"auth_group_permissions\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, authGroupPermissionPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in authGroupPermission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all authGroupPermission")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuthGroupPermission) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_group_permissions provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authGroupPermissionColumnsWithDefault, o)

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

	authGroupPermissionUpsertCacheMut.RLock()
	cache, cached := authGroupPermissionUpsertCache[key]
	authGroupPermissionUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			authGroupPermissionAllColumns,
			authGroupPermissionColumnsWithDefault,
			authGroupPermissionColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			authGroupPermissionAllColumns,
			authGroupPermissionPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert auth_group_permissions, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authGroupPermissionPrimaryKeyColumns))
			copy(conflict, authGroupPermissionPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"auth_group_permissions\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authGroupPermissionType, authGroupPermissionMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert auth_group_permissions")
	}

	if !cached {
		authGroupPermissionUpsertCacheMut.Lock()
		authGroupPermissionUpsertCache[key] = cache
		authGroupPermissionUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AuthGroupPermission record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthGroupPermission) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuthGroupPermission provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authGroupPermissionPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_group_permissions\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from auth_group_permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for auth_group_permissions")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q authGroupPermissionQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no authGroupPermissionQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auth_group_permissions")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_group_permissions")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthGroupPermissionSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(authGroupPermissionBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authGroupPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"auth_group_permissions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authGroupPermissionPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from authGroupPermission slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_group_permissions")
	}

	if len(authGroupPermissionAfterDeleteHooks) != 0 {
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
func (o *AuthGroupPermission) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthGroupPermission(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthGroupPermissionSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthGroupPermissionSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authGroupPermissionPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"auth_group_permissions\".* FROM \"auth_group_permissions\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authGroupPermissionPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthGroupPermissionSlice")
	}

	*o = slice

	return nil
}

// AuthGroupPermissionExists checks if the AuthGroupPermission row exists.
func AuthGroupPermissionExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"auth_group_permissions\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_group_permissions exists")
	}

	return exists, nil
}

// Exists checks if the AuthGroupPermission row exists.
func (o *AuthGroupPermission) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AuthGroupPermissionExists(ctx, exec, o.ID)
}