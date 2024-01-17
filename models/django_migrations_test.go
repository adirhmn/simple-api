// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testDjangoMigrations(t *testing.T) {
	t.Parallel()

	query := DjangoMigrations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDjangoMigrationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoMigrationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DjangoMigrations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoMigrationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoMigrationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoMigrationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DjangoMigrationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DjangoMigration exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DjangoMigrationExists to return true, but got false.")
	}
}

func testDjangoMigrationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	djangoMigrationFound, err := FindDjangoMigration(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if djangoMigrationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDjangoMigrationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DjangoMigrations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDjangoMigrationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DjangoMigrations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDjangoMigrationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	djangoMigrationOne := &DjangoMigration{}
	djangoMigrationTwo := &DjangoMigration{}
	if err = randomize.Struct(seed, djangoMigrationOne, djangoMigrationDBTypes, false, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoMigrationTwo, djangoMigrationDBTypes, false, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = djangoMigrationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoMigrationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoMigrations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDjangoMigrationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	djangoMigrationOne := &DjangoMigration{}
	djangoMigrationTwo := &DjangoMigration{}
	if err = randomize.Struct(seed, djangoMigrationOne, djangoMigrationDBTypes, false, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoMigrationTwo, djangoMigrationDBTypes, false, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = djangoMigrationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoMigrationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func djangoMigrationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func djangoMigrationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoMigration) error {
	*o = DjangoMigration{}
	return nil
}

func testDjangoMigrationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DjangoMigration{}
	o := &DjangoMigration{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DjangoMigration object: %s", err)
	}

	AddDjangoMigrationHook(boil.BeforeInsertHook, djangoMigrationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	djangoMigrationBeforeInsertHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.AfterInsertHook, djangoMigrationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	djangoMigrationAfterInsertHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.AfterSelectHook, djangoMigrationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	djangoMigrationAfterSelectHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.BeforeUpdateHook, djangoMigrationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	djangoMigrationBeforeUpdateHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.AfterUpdateHook, djangoMigrationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	djangoMigrationAfterUpdateHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.BeforeDeleteHook, djangoMigrationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	djangoMigrationBeforeDeleteHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.AfterDeleteHook, djangoMigrationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	djangoMigrationAfterDeleteHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.BeforeUpsertHook, djangoMigrationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	djangoMigrationBeforeUpsertHooks = []DjangoMigrationHook{}

	AddDjangoMigrationHook(boil.AfterUpsertHook, djangoMigrationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	djangoMigrationAfterUpsertHooks = []DjangoMigrationHook{}
}

func testDjangoMigrationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoMigrationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(djangoMigrationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoMigrationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDjangoMigrationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoMigrationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDjangoMigrationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoMigrations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	djangoMigrationDBTypes = map[string]string{`ID`: `integer`, `App`: `character varying`, `Name`: `character varying`, `Applied`: `timestamp with time zone`}
	_                      = bytes.MinRead
)

func testDjangoMigrationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(djangoMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(djangoMigrationAllColumns) == len(djangoMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDjangoMigrationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(djangoMigrationAllColumns) == len(djangoMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoMigration{}
	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoMigrationDBTypes, true, djangoMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(djangoMigrationAllColumns, djangoMigrationPrimaryKeyColumns) {
		fields = djangoMigrationAllColumns
	} else {
		fields = strmangle.SetComplement(
			djangoMigrationAllColumns,
			djangoMigrationPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := DjangoMigrationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDjangoMigrationsUpsert(t *testing.T) {
	t.Parallel()

	if len(djangoMigrationAllColumns) == len(djangoMigrationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DjangoMigration{}
	if err = randomize.Struct(seed, &o, djangoMigrationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoMigration: %s", err)
	}

	count, err := DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, djangoMigrationDBTypes, false, djangoMigrationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoMigration struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoMigration: %s", err)
	}

	count, err = DjangoMigrations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
