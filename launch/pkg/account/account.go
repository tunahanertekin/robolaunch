package account

//The following package resposible create groups and bind them user after creation of namespace & role
//03.01.2022

import (
	"context"
	"fmt"
	"os"

	"github.com/Nerzal/gocloak/v10"
)

var kUrl = os.Getenv("KEYCLOAK_URL")
var kUser = os.Getenv("KEYCLOAK_ADMIN_USERNAME")
var kPass = os.Getenv("KEYCLOAK_ADMIN_PASSWORD")
var kRealm = os.Getenv("KEYCLOAK_REALM")
var kc = gocloak.NewClient(kUrl)

// It responsible for create group with same name with cluster-role
func CreateGroup(namespace string) (string, error) {
	ctx := context.Background()

	token, err := kc.LoginAdmin(ctx, kUser, kPass, "master")
	if err != nil {
		return "", err
	}
	group := gocloak.Group{
		Name: gocloak.StringP(namespace + "_role"),
	}
	result, err := kc.CreateGroup(ctx, token.AccessToken, kRealm, group)
	if err != nil {
		fmt.Printf("Oh no!, failed to create group :( %v\n", err)
		return "", err
	}
	return result, nil

}

//FOR TESTING PURPOSES
func DeleteGroup(id string) error {
	ctx := context.Background()

	token, err := kc.LoginAdmin(ctx, kUser, kPass, "master")
	if err != nil {
		return err
	}
	err = kc.DeleteGroup(ctx, token.AccessToken, kRealm, id)
	if err != nil {
		fmt.Printf("Oh no!, failed to delete group :( %v\n", err)
		return err
	}
	return nil

}

func GetGroup(id string) (*gocloak.Group, error) {
	ctx := context.Background()

	token, err := kc.LoginAdmin(ctx, kUser, kPass, "master")
	if err != nil {
		return nil, err
	}
	group, err := kc.GetGroup(ctx, token.AccessToken, kRealm, id)
	if err != nil {
		fmt.Printf("Oh no!, failed to get groups :( %v\n", err)
		return nil, err
	}

	return group, nil

}
func BindGroup(username string, namespace string) error {
	ctx := context.Background()

	token, err := kc.LoginAdmin(ctx, kUser, kPass, "master")
	if err != nil {
		return err
	}
	user, err := kc.GetUsers(ctx, token.AccessToken, kRealm, gocloak.GetUsersParams{
		Username: gocloak.StringP(username),
	})
	if err != nil {
		return err
	}
	group, err := kc.GetGroups(ctx, token.AccessToken, kRealm, gocloak.GetGroupsParams{
		Search: gocloak.StringP(namespace + "_role"),
	})
	if err != nil {
		return err
	}
	err = kc.AddUserToGroup(ctx, token.AccessToken, kRealm, *user[0].ID, *group[0].ID)
	if err != nil {
		return err
	}
	return nil

}
