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

func testDjangoContentTypes(t *testing.T) {
	t.Parallel()

	query := DjangoContentTypes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDjangoContentTypesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
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

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoContentTypesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := DjangoContentTypes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoContentTypesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoContentTypeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDjangoContentTypesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DjangoContentTypeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if DjangoContentType exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DjangoContentTypeExists to return true, but got false.")
	}
}

func testDjangoContentTypesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	djangoContentTypeFound, err := FindDjangoContentType(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if djangoContentTypeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDjangoContentTypesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = DjangoContentTypes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDjangoContentTypesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := DjangoContentTypes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDjangoContentTypesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	djangoContentTypeOne := &DjangoContentType{}
	djangoContentTypeTwo := &DjangoContentType{}
	if err = randomize.Struct(seed, djangoContentTypeOne, djangoContentTypeDBTypes, false, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoContentTypeTwo, djangoContentTypeDBTypes, false, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = djangoContentTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoContentTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoContentTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDjangoContentTypesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	djangoContentTypeOne := &DjangoContentType{}
	djangoContentTypeTwo := &DjangoContentType{}
	if err = randomize.Struct(seed, djangoContentTypeOne, djangoContentTypeDBTypes, false, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}
	if err = randomize.Struct(seed, djangoContentTypeTwo, djangoContentTypeDBTypes, false, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = djangoContentTypeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = djangoContentTypeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func djangoContentTypeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func djangoContentTypeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *DjangoContentType) error {
	*o = DjangoContentType{}
	return nil
}

func testDjangoContentTypesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &DjangoContentType{}
	o := &DjangoContentType{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize DjangoContentType object: %s", err)
	}

	AddDjangoContentTypeHook(boil.BeforeInsertHook, djangoContentTypeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeBeforeInsertHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.AfterInsertHook, djangoContentTypeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeAfterInsertHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.AfterSelectHook, djangoContentTypeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeAfterSelectHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.BeforeUpdateHook, djangoContentTypeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeBeforeUpdateHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.AfterUpdateHook, djangoContentTypeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeAfterUpdateHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.BeforeDeleteHook, djangoContentTypeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeBeforeDeleteHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.AfterDeleteHook, djangoContentTypeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeAfterDeleteHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.BeforeUpsertHook, djangoContentTypeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeBeforeUpsertHooks = []DjangoContentTypeHook{}

	AddDjangoContentTypeHook(boil.AfterUpsertHook, djangoContentTypeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	djangoContentTypeAfterUpsertHooks = []DjangoContentTypeHook{}
}

func testDjangoContentTypesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoContentTypesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(djangoContentTypeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDjangoContentTypeToManyContentTypeAuthPermissions(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoContentType
	var b, c AuthPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, authPermissionDBTypes, false, authPermissionColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.ContentTypeID = a.ID
	c.ContentTypeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ContentTypeAuthPermissions().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.ContentTypeID == b.ContentTypeID {
			bFound = true
		}
		if v.ContentTypeID == c.ContentTypeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DjangoContentTypeSlice{&a}
	if err = a.L.LoadContentTypeAuthPermissions(ctx, tx, false, (*[]*DjangoContentType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentTypeAuthPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ContentTypeAuthPermissions = nil
	if err = a.L.LoadContentTypeAuthPermissions(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentTypeAuthPermissions); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDjangoContentTypeToManyContentTypeDjangoAdminLogs(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoContentType
	var b, c DjangoAdminLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, djangoAdminLogDBTypes, false, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, djangoAdminLogDBTypes, false, djangoAdminLogColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.ContentTypeID, a.ID)
	queries.Assign(&c.ContentTypeID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.ContentTypeDjangoAdminLogs().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if queries.Equal(v.ContentTypeID, b.ContentTypeID) {
			bFound = true
		}
		if queries.Equal(v.ContentTypeID, c.ContentTypeID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DjangoContentTypeSlice{&a}
	if err = a.L.LoadContentTypeDjangoAdminLogs(ctx, tx, false, (*[]*DjangoContentType)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentTypeDjangoAdminLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.ContentTypeDjangoAdminLogs = nil
	if err = a.L.LoadContentTypeDjangoAdminLogs(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.ContentTypeDjangoAdminLogs); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDjangoContentTypeToManyAddOpContentTypeAuthPermissions(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoContentType
	var b, c, d, e AuthPermission

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*AuthPermission{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, authPermissionDBTypes, false, strmangle.SetComplement(authPermissionPrimaryKeyColumns, authPermissionColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*AuthPermission{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddContentTypeAuthPermissions(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.ContentTypeID {
			t.Error("foreign key was wrong value", a.ID, first.ContentTypeID)
		}
		if a.ID != second.ContentTypeID {
			t.Error("foreign key was wrong value", a.ID, second.ContentTypeID)
		}

		if first.R.ContentType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ContentType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ContentTypeAuthPermissions[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ContentTypeAuthPermissions[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ContentTypeAuthPermissions().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDjangoContentTypeToManyAddOpContentTypeDjangoAdminLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoContentType
	var b, c, d, e DjangoAdminLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DjangoAdminLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, djangoAdminLogDBTypes, false, strmangle.SetComplement(djangoAdminLogPrimaryKeyColumns, djangoAdminLogColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*DjangoAdminLog{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddContentTypeDjangoAdminLogs(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.ContentTypeID) {
			t.Error("foreign key was wrong value", a.ID, first.ContentTypeID)
		}
		if !queries.Equal(a.ID, second.ContentTypeID) {
			t.Error("foreign key was wrong value", a.ID, second.ContentTypeID)
		}

		if first.R.ContentType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.ContentType != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.ContentTypeDjangoAdminLogs[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.ContentTypeDjangoAdminLogs[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.ContentTypeDjangoAdminLogs().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDjangoContentTypeToManySetOpContentTypeDjangoAdminLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoContentType
	var b, c, d, e DjangoAdminLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DjangoAdminLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, djangoAdminLogDBTypes, false, strmangle.SetComplement(djangoAdminLogPrimaryKeyColumns, djangoAdminLogColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetContentTypeDjangoAdminLogs(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ContentTypeDjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetContentTypeDjangoAdminLogs(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ContentTypeDjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ContentTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ContentTypeID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.ContentTypeID) {
		t.Error("foreign key was wrong value", a.ID, d.ContentTypeID)
	}
	if !queries.Equal(a.ID, e.ContentTypeID) {
		t.Error("foreign key was wrong value", a.ID, e.ContentTypeID)
	}

	if b.R.ContentType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.ContentType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.ContentType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.ContentType != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.ContentTypeDjangoAdminLogs[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.ContentTypeDjangoAdminLogs[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testDjangoContentTypeToManyRemoveOpContentTypeDjangoAdminLogs(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a DjangoContentType
	var b, c, d, e DjangoAdminLog

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, djangoContentTypeDBTypes, false, strmangle.SetComplement(djangoContentTypePrimaryKeyColumns, djangoContentTypeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DjangoAdminLog{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, djangoAdminLogDBTypes, false, strmangle.SetComplement(djangoAdminLogPrimaryKeyColumns, djangoAdminLogColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddContentTypeDjangoAdminLogs(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.ContentTypeDjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveContentTypeDjangoAdminLogs(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.ContentTypeDjangoAdminLogs().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.ContentTypeID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.ContentTypeID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.ContentType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.ContentType != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.ContentType != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.ContentType != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.ContentTypeDjangoAdminLogs) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.ContentTypeDjangoAdminLogs[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.ContentTypeDjangoAdminLogs[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testDjangoContentTypesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
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

func testDjangoContentTypesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DjangoContentTypeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDjangoContentTypesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := DjangoContentTypes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	djangoContentTypeDBTypes = map[string]string{`ID`: `integer`, `AppLabel`: `character varying`, `Model`: `character varying`}
	_                        = bytes.MinRead
)

func testDjangoContentTypesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(djangoContentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(djangoContentTypeAllColumns) == len(djangoContentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDjangoContentTypesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(djangoContentTypeAllColumns) == len(djangoContentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &DjangoContentType{}
	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, djangoContentTypeDBTypes, true, djangoContentTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(djangoContentTypeAllColumns, djangoContentTypePrimaryKeyColumns) {
		fields = djangoContentTypeAllColumns
	} else {
		fields = strmangle.SetComplement(
			djangoContentTypeAllColumns,
			djangoContentTypePrimaryKeyColumns,
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

	slice := DjangoContentTypeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDjangoContentTypesUpsert(t *testing.T) {
	t.Parallel()

	if len(djangoContentTypeAllColumns) == len(djangoContentTypePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := DjangoContentType{}
	if err = randomize.Struct(seed, &o, djangoContentTypeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoContentType: %s", err)
	}

	count, err := DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, djangoContentTypeDBTypes, false, djangoContentTypePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize DjangoContentType struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert DjangoContentType: %s", err)
	}

	count, err = DjangoContentTypes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
