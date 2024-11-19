package database_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
	"opencsg.com/csghub-server/builder/store/database"
	"opencsg.com/csghub-server/common/tests"
)

func TestUserStore_Roles(t *testing.T) {
	type fields struct {
		RoleMask string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
		{
			name: "test no role",
			fields: fields{
				RoleMask: "",
			},
			want: []string{},
		},
		{
			name: "test one role",
			fields: fields{
				RoleMask: "admin",
			},
			want: []string{"admin"},
		},
		{
			name: "test two roles",
			fields: fields{
				RoleMask: "admin,super_user",
			},
			want: []string{"admin", "super_user"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &database.User{
				RoleMask: tt.fields.RoleMask,
			}
			if got := u.Roles(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Roles() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserStore_SetRoles(t *testing.T) {
	type fields struct {
		RoleMask string
	}
	type args struct {
		roles []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "test no role",
			fields: fields{
				RoleMask: "",
			},
			args: args{
				roles: []string{""},
			},
		},
		{
			name: "test one role",
			fields: fields{
				RoleMask: "admin",
			},
			args: args{
				roles: []string{"admin"},
			},
		},
		{
			name: "test two roles",
			fields: fields{
				RoleMask: "admin,super_user",
			},
			args: args{
				roles: []string{"admin", "super_user"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &database.User{}
			u.SetRoles(tt.args.roles)
			if u.RoleMask != tt.fields.RoleMask {
				t.Errorf("User.SetRoles() = %v, want %v", u.RoleMask, tt.fields.RoleMask)
			}
		})
	}
}

func TestUserStore_IndexWithSearch(t *testing.T) {
	db := tests.InitTestDB()
	defer db.Close()
	ctx := context.TODO()

	userStore := database.NewUserStoreWithDB(db)
	err := userStore.Create(ctx, &database.User{
		GitID:    3321,
		Username: "u-foo",
	}, &database.Namespace{Path: "1"})
	require.Nil(t, err)

	err = userStore.Create(ctx, &database.User{
		GitID:    3322,
		Username: "u-bar",
		Email:    "efoo@z.com",
	}, &database.Namespace{Path: "2"})
	require.Nil(t, err)

	err = userStore.Create(ctx, &database.User{
		GitID:    3323,
		Username: "u-barz",
		Email:    "ebar@z.com",
	}, &database.Namespace{Path: "3"})
	require.Nil(t, err)

	cases := []struct {
		per      int
		page     int
		total    int
		expected []int64
	}{
		{10, 1, 2, []int64{3321, 3322}},
		{1, 1, 2, []int64{3321}},
		{1, 2, 2, []int64{3322}},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("page %d, per %d", c.page, c.per), func(t *testing.T) {

			users, count, err := userStore.IndexWithSearch(ctx, "foo", c.per, c.page)
			require.Nil(t, err)
			require.Equal(t, c.total, count)

			gids := []int64{}
			for _, u := range users {
				gids = append(gids, u.GitID)
			}
			require.Equal(t, c.expected, gids)
		})
	}

}

func TestUserStore_CreateUser(t *testing.T) {
	db := tests.InitTestDB()
	defer db.Close()
	ctx := context.TODO()

	us := database.NewUserStoreWithDB(db)
	err := us.Create(ctx, &database.User{
		GitID:    3321,
		Username: "u-foo",
	}, &database.Namespace{Path: "u-foo"})
	require.Nil(t, err)

	user, err := us.FindByUsername(ctx, "u-foo")
	require.Nil(t, err)
	require.Equal(t, 3321, int(user.GitID))
	require.Equal(t, "u-foo", user.Username)
}

func TestUserStore_ChangeUserName(t *testing.T) {
	db := tests.InitTestDB()
	defer db.Close()
	ctx := context.TODO()

	us := database.NewUserStoreWithDB(db)
	err := us.Create(ctx, &database.User{
		GitID:    3321,
		Username: "u-foo",
	}, &database.Namespace{Path: "u-foo"})
	require.Nil(t, err)

	err = us.ChangeUserName(ctx, "u-foo", "u-bar")
	require.Nil(t, err)

	user, err := us.FindByUsername(ctx, "u-bar")
	require.Nil(t, err)
	require.Equal(t, "u-bar", user.Username)
}
