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

// AuthUserGroup is an object representing the database table.
type AuthUserGroup struct {
	ID      int `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID  int `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	GroupID int `boil:"group_id" json:"group_id" toml:"group_id" yaml:"group_id"`

	R *authUserGroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authUserGroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthUserGroupColumns = struct {
	ID      string
	UserID  string
	GroupID string
}{
	ID:      "id",
	UserID:  "user_id",
	GroupID: "group_id",
}

var AuthUserGroupTableColumns = struct {
	ID      string
	UserID  string
	GroupID string
}{
	ID:      "auth_user_groups.id",
	UserID:  "auth_user_groups.user_id",
	GroupID: "auth_user_groups.group_id",
}

// Generated where

var AuthUserGroupWhere = struct {
	ID      whereHelperint
	UserID  whereHelperint
	GroupID whereHelperint
}{
	ID:      whereHelperint{field: "\"auth_user_groups\".\"id\""},
	UserID:  whereHelperint{field: "\"auth_user_groups\".\"user_id\""},
	GroupID: whereHelperint{field: "\"auth_user_groups\".\"group_id\""},
}

// AuthUserGroupRels is where relationship names are stored.
var AuthUserGroupRels = struct {
	Group string
	User  string
}{
	Group: "Group",
	User:  "User",
}

// authUserGroupR is where relationships are stored.
type authUserGroupR struct {
	Group *AuthGroup `boil:"Group" json:"Group" toml:"Group" yaml:"Group"`
	User  *AuthUser  `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*authUserGroupR) NewStruct() *authUserGroupR {
	return &authUserGroupR{}
}

func (r *authUserGroupR) GetGroup() *AuthGroup {
	if r == nil {
		return nil
	}
	return r.Group
}

func (r *authUserGroupR) GetUser() *AuthUser {
	if r == nil {
		return nil
	}
	return r.User
}

// authUserGroupL is where Load methods for each relationship are stored.
type authUserGroupL struct{}

var (
	authUserGroupAllColumns            = []string{"id", "user_id", "group_id"}
	authUserGroupColumnsWithoutDefault = []string{"user_id", "group_id"}
	authUserGroupColumnsWithDefault    = []string{"id"}
	authUserGroupPrimaryKeyColumns     = []string{"id"}
	authUserGroupGeneratedColumns      = []string{}
)

type (
	// AuthUserGroupSlice is an alias for a slice of pointers to AuthUserGroup.
	// This should almost always be used instead of []AuthUserGroup.
	AuthUserGroupSlice []*AuthUserGroup
	// AuthUserGroupHook is the signature for custom AuthUserGroup hook methods
	AuthUserGroupHook func(context.Context, boil.ContextExecutor, *AuthUserGroup) error

	authUserGroupQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authUserGroupType                 = reflect.TypeOf(&AuthUserGroup{})
	authUserGroupMapping              = queries.MakeStructMapping(authUserGroupType)
	authUserGroupPrimaryKeyMapping, _ = queries.BindMapping(authUserGroupType, authUserGroupMapping, authUserGroupPrimaryKeyColumns)
	authUserGroupInsertCacheMut       sync.RWMutex
	authUserGroupInsertCache          = make(map[string]insertCache)
	authUserGroupUpdateCacheMut       sync.RWMutex
	authUserGroupUpdateCache          = make(map[string]updateCache)
	authUserGroupUpsertCacheMut       sync.RWMutex
	authUserGroupUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var authUserGroupAfterSelectHooks []AuthUserGroupHook

var authUserGroupBeforeInsertHooks []AuthUserGroupHook
var authUserGroupAfterInsertHooks []AuthUserGroupHook

var authUserGroupBeforeUpdateHooks []AuthUserGroupHook
var authUserGroupAfterUpdateHooks []AuthUserGroupHook

var authUserGroupBeforeDeleteHooks []AuthUserGroupHook
var authUserGroupAfterDeleteHooks []AuthUserGroupHook

var authUserGroupBeforeUpsertHooks []AuthUserGroupHook
var authUserGroupAfterUpsertHooks []AuthUserGroupHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthUserGroup) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthUserGroup) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthUserGroup) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthUserGroup) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthUserGroup) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthUserGroup) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthUserGroup) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthUserGroup) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthUserGroup) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authUserGroupAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthUserGroupHook registers your hook function for all future operations.
func AddAuthUserGroupHook(hookPoint boil.HookPoint, authUserGroupHook AuthUserGroupHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		authUserGroupAfterSelectHooks = append(authUserGroupAfterSelectHooks, authUserGroupHook)
	case boil.BeforeInsertHook:
		authUserGroupBeforeInsertHooks = append(authUserGroupBeforeInsertHooks, authUserGroupHook)
	case boil.AfterInsertHook:
		authUserGroupAfterInsertHooks = append(authUserGroupAfterInsertHooks, authUserGroupHook)
	case boil.BeforeUpdateHook:
		authUserGroupBeforeUpdateHooks = append(authUserGroupBeforeUpdateHooks, authUserGroupHook)
	case boil.AfterUpdateHook:
		authUserGroupAfterUpdateHooks = append(authUserGroupAfterUpdateHooks, authUserGroupHook)
	case boil.BeforeDeleteHook:
		authUserGroupBeforeDeleteHooks = append(authUserGroupBeforeDeleteHooks, authUserGroupHook)
	case boil.AfterDeleteHook:
		authUserGroupAfterDeleteHooks = append(authUserGroupAfterDeleteHooks, authUserGroupHook)
	case boil.BeforeUpsertHook:
		authUserGroupBeforeUpsertHooks = append(authUserGroupBeforeUpsertHooks, authUserGroupHook)
	case boil.AfterUpsertHook:
		authUserGroupAfterUpsertHooks = append(authUserGroupAfterUpsertHooks, authUserGroupHook)
	}
}

// One returns a single authUserGroup record from the query.
func (q authUserGroupQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuthUserGroup, error) {
	o := &AuthUserGroup{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for auth_user_groups")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuthUserGroup records from the query.
func (q authUserGroupQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthUserGroupSlice, error) {
	var o []*AuthUserGroup

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AuthUserGroup slice")
	}

	if len(authUserGroupAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuthUserGroup records in the query.
func (q authUserGroupQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count auth_user_groups rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q authUserGroupQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if auth_user_groups exists")
	}

	return count > 0, nil
}

// Group pointed to by the foreign key.
func (o *AuthUserGroup) Group(mods ...qm.QueryMod) authGroupQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.GroupID),
	}

	queryMods = append(queryMods, mods...)

	return AuthGroups(queryMods...)
}

// User pointed to by the foreign key.
func (o *AuthUserGroup) User(mods ...qm.QueryMod) authUserQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	return AuthUsers(queryMods...)
}

// LoadGroup allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (authUserGroupL) LoadGroup(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuthUserGroup interface{}, mods queries.Applicator) error {
	var slice []*AuthUserGroup
	var object *AuthUserGroup

	if singular {
		var ok bool
		object, ok = maybeAuthUserGroup.(*AuthUserGroup)
		if !ok {
			object = new(AuthUserGroup)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAuthUserGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAuthUserGroup))
			}
		}
	} else {
		s, ok := maybeAuthUserGroup.(*[]*AuthUserGroup)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAuthUserGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAuthUserGroup))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &authUserGroupR{}
		}
		args = append(args, object.GroupID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &authUserGroupR{}
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
		foreign.R.GroupAuthUserGroups = append(foreign.R.GroupAuthUserGroups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.GroupID == foreign.ID {
				local.R.Group = foreign
				if foreign.R == nil {
					foreign.R = &authGroupR{}
				}
				foreign.R.GroupAuthUserGroups = append(foreign.R.GroupAuthUserGroups, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (authUserGroupL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuthUserGroup interface{}, mods queries.Applicator) error {
	var slice []*AuthUserGroup
	var object *AuthUserGroup

	if singular {
		var ok bool
		object, ok = maybeAuthUserGroup.(*AuthUserGroup)
		if !ok {
			object = new(AuthUserGroup)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeAuthUserGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeAuthUserGroup))
			}
		}
	} else {
		s, ok := maybeAuthUserGroup.(*[]*AuthUserGroup)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeAuthUserGroup)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeAuthUserGroup))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &authUserGroupR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &authUserGroupR{}
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
		qm.From(`auth_user`),
		qm.WhereIn(`auth_user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load AuthUser")
	}

	var resultSlice []*AuthUser
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice AuthUser")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for auth_user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for auth_user")
	}

	if len(authUserAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &authUserR{}
		}
		foreign.R.UserAuthUserGroups = append(foreign.R.UserAuthUserGroups, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &authUserR{}
				}
				foreign.R.UserAuthUserGroups = append(foreign.R.UserAuthUserGroups, local)
				break
			}
		}
	}

	return nil
}

// SetGroup of the authUserGroup to the related item.
// Sets o.R.Group to related.
// Adds o to related.R.GroupAuthUserGroups.
func (o *AuthUserGroup) SetGroup(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AuthGroup) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_user_groups\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"group_id"}),
		strmangle.WhereClause("\"", "\"", 2, authUserGroupPrimaryKeyColumns),
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
		o.R = &authUserGroupR{
			Group: related,
		}
	} else {
		o.R.Group = related
	}

	if related.R == nil {
		related.R = &authGroupR{
			GroupAuthUserGroups: AuthUserGroupSlice{o},
		}
	} else {
		related.R.GroupAuthUserGroups = append(related.R.GroupAuthUserGroups, o)
	}

	return nil
}

// SetUser of the authUserGroup to the related item.
// Sets o.R.User to related.
// Adds o to related.R.UserAuthUserGroups.
func (o *AuthUserGroup) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *AuthUser) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"auth_user_groups\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, authUserGroupPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &authUserGroupR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &authUserR{
			UserAuthUserGroups: AuthUserGroupSlice{o},
		}
	} else {
		related.R.UserAuthUserGroups = append(related.R.UserAuthUserGroups, o)
	}

	return nil
}

// AuthUserGroups retrieves all the records using an executor.
func AuthUserGroups(mods ...qm.QueryMod) authUserGroupQuery {
	mods = append(mods, qm.From("\"auth_user_groups\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"auth_user_groups\".*"})
	}

	return authUserGroupQuery{q}
}

// FindAuthUserGroup retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthUserGroup(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*AuthUserGroup, error) {
	authUserGroupObj := &AuthUserGroup{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"auth_user_groups\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, authUserGroupObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from auth_user_groups")
	}

	if err = authUserGroupObj.doAfterSelectHooks(ctx, exec); err != nil {
		return authUserGroupObj, err
	}

	return authUserGroupObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuthUserGroup) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_user_groups provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserGroupColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authUserGroupInsertCacheMut.RLock()
	cache, cached := authUserGroupInsertCache[key]
	authUserGroupInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authUserGroupAllColumns,
			authUserGroupColumnsWithDefault,
			authUserGroupColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(authUserGroupType, authUserGroupMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authUserGroupType, authUserGroupMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"auth_user_groups\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"auth_user_groups\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into auth_user_groups")
	}

	if !cached {
		authUserGroupInsertCacheMut.Lock()
		authUserGroupInsertCache[key] = cache
		authUserGroupInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuthUserGroup.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuthUserGroup) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	authUserGroupUpdateCacheMut.RLock()
	cache, cached := authUserGroupUpdateCache[key]
	authUserGroupUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authUserGroupAllColumns,
			authUserGroupPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update auth_user_groups, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"auth_user_groups\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authUserGroupPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authUserGroupType, authUserGroupMapping, append(wl, authUserGroupPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update auth_user_groups row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for auth_user_groups")
	}

	if !cached {
		authUserGroupUpdateCacheMut.Lock()
		authUserGroupUpdateCache[key] = cache
		authUserGroupUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q authUserGroupQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for auth_user_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for auth_user_groups")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthUserGroupSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"auth_user_groups\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, authUserGroupPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in authUserGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all authUserGroup")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuthUserGroup) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no auth_user_groups provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authUserGroupColumnsWithDefault, o)

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

	authUserGroupUpsertCacheMut.RLock()
	cache, cached := authUserGroupUpsertCache[key]
	authUserGroupUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			authUserGroupAllColumns,
			authUserGroupColumnsWithDefault,
			authUserGroupColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			authUserGroupAllColumns,
			authUserGroupPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert auth_user_groups, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authUserGroupPrimaryKeyColumns))
			copy(conflict, authUserGroupPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"auth_user_groups\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(authUserGroupType, authUserGroupMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authUserGroupType, authUserGroupMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert auth_user_groups")
	}

	if !cached {
		authUserGroupUpsertCacheMut.Lock()
		authUserGroupUpsertCache[key] = cache
		authUserGroupUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AuthUserGroup record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthUserGroup) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AuthUserGroup provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authUserGroupPrimaryKeyMapping)
	sql := "DELETE FROM \"auth_user_groups\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from auth_user_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for auth_user_groups")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q authUserGroupQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no authUserGroupQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from auth_user_groups")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_user_groups")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthUserGroupSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(authUserGroupBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"auth_user_groups\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authUserGroupPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from authUserGroup slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for auth_user_groups")
	}

	if len(authUserGroupAfterDeleteHooks) != 0 {
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
func (o *AuthUserGroup) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthUserGroup(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthUserGroupSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthUserGroupSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authUserGroupPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"auth_user_groups\".* FROM \"auth_user_groups\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authUserGroupPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AuthUserGroupSlice")
	}

	*o = slice

	return nil
}

// AuthUserGroupExists checks if the AuthUserGroup row exists.
func AuthUserGroupExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"auth_user_groups\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if auth_user_groups exists")
	}

	return exists, nil
}

// Exists checks if the AuthUserGroup row exists.
func (o *AuthUserGroup) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AuthUserGroupExists(ctx, exec, o.ID)
}
