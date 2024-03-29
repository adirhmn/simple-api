// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

func TestUpsert(t *testing.T) {
	t.Run("AuthGroups", testAuthGroupsUpsert)

	t.Run("AuthGroupPermissions", testAuthGroupPermissionsUpsert)

	t.Run("AuthPermissions", testAuthPermissionsUpsert)

	t.Run("AuthUsers", testAuthUsersUpsert)

	t.Run("AuthUserGroups", testAuthUserGroupsUpsert)

	t.Run("AuthUserUserPermissions", testAuthUserUserPermissionsUpsert)

	t.Run("DjangoAdminLogs", testDjangoAdminLogsUpsert)

	t.Run("DjangoContentTypes", testDjangoContentTypesUpsert)

	t.Run("DjangoMigrations", testDjangoMigrationsUpsert)

	t.Run("DjangoSessions", testDjangoSessionsUpsert)

	t.Run("Users", testUsersUpsert)
}
