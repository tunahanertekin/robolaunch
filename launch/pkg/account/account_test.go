package account

import (
	"context"
	"fmt"
	"testing"

	"github.com/Nerzal/gocloak/v10"
	"github.com/stretchr/testify/assert"
)

var groupName string = "test-group"
var testUser string = "tester-user"

func TestGroupCRUD(t *testing.T) {
	resp, err := CreateGroup(groupName)
	if err != nil {
		panic(err)
	}
	group, err := GetGroup(resp)
	if err != nil {
		panic(err)

	}

	test := group.Name
	fmt.Println(*test)

	assert.Equal(t, groupName+"_role", *test)

}

func TestBindGroup(t *testing.T) {
	// Create user & role & bind & check
	ctx := context.Background()
	var kc = gocloak.NewClient(kUrl)

	token, err := kc.LoginAdmin(ctx, kUser, kPass, "master")
	if err != nil {
		panic(err)
	}
	user := gocloak.User{
		FirstName: gocloak.StringP(""),
		LastName:  gocloak.StringP(""),
		Email:     gocloak.StringP(""),
		Enabled:   gocloak.BoolP(true),
		Username:  gocloak.StringP(testUser),
	}

	userId, err := kc.CreateUser(ctx, token.AccessToken, kRealm, user)
	if err != nil {
		panic(err)
	}

	err = BindGroup(*user.Username, groupName)
	if err != nil {
		panic(err)
	}

	// To check new user
	newGroup, err := kc.GetUserGroups(ctx, token.AccessToken, kRealm, userId, gocloak.GetGroupsParams{
		Search: gocloak.StringP(groupName + "_role"),
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(newGroup)
	assert.Equal(t, groupName+"_role", *newGroup[0].Name)

}

func TestCleanUp(t *testing.T) {
	// clear group
	ctx := context.Background()
	token, err := kc.LoginAdmin(ctx, kUser, kPass, "master")
	if err != nil {
		panic(err)
	}

	// get group id first
	group, err := kc.GetGroups(ctx, token.AccessToken, kRealm, gocloak.GetGroupsParams{
		Search: gocloak.StringP(groupName + "_role"),
	})
	if err != nil {
		panic(err)
	}
	err = DeleteGroup(*group[0].ID)
	if err != nil {
		panic(err)
	}
	// delete user
	user, err := kc.GetUsers(ctx, token.AccessToken, kRealm, gocloak.GetUsersParams{
		Username: gocloak.StringP(testUser),
	})
	if err != nil {
		panic(err)
	}
	err = kc.DeleteUser(ctx, token.AccessToken, kRealm, *user[0].ID)
	if err != nil {
		panic(err)
	}
}
