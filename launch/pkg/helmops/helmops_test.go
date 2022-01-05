package helmops

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLaunches(t *testing.T) {
	launches, err := GetLaunches()
	jackalVersions := launches["jackal"]

	assert.Greater(t, len(jackalVersions), 0)
	assert.Equal(t, err, nil)
}

func TestGetLaunch(t *testing.T) {
	launch1, err1 := GetLaunch("jackal")
	if err1 != nil {
		panic(err1)
	}

	_, err2 := GetLaunch("wrong-robot")

	assert.Greater(t, len(launch1), 0)
	assert.NotEqual(t, err2, nil)
}

func TestGetLaunchWithVersion(t *testing.T) {
	_, err1 := GetLaunchWithVersion("jackal", "0.1.0")
	_, err2 := GetLaunchWithVersion("wrong-robot", "0.1.0")

	assert.Equal(t, err1, nil)
	assert.NotEqual(t, err2, nil)
}

// Runs consecutively: Register -> Get -> Delete

var cluster string = "default"
var ns string = "default"
var appRepoName string = "test-apprepo"
var appRepoURL string = "https://helm.camunda.cloud" // check availability

func TestRegisterAppRepository(t *testing.T) {

	registerAppRepo, err := RegisterAppRepository(
		os.Getenv("TOKEN"),
		cluster,
		ns,
		RegisterAppRepositoryBody{
			AppRepository: RegisterAppRepositoryBodyDetails{
				Name:        appRepoName,
				Type:        "helm",
				Description: "for testing purposes",
				RepoURL:     appRepoURL,
			},
		},
	)

	assert.Equal(t, err, nil)
	assert.Equal(t, registerAppRepo.Metadata.Name, appRepoName)

}

func TestGetAppRepository(t *testing.T) {

	appRepo, err := GetAppRepository(
		os.Getenv("TOKEN"),
		cluster,
		ns,
		appRepoName,
	)

	assert.Equal(t, err, nil)
	assert.Equal(t, appRepo.Metadata.Name, appRepoName)

}

func TestDeleteAppRepository(t *testing.T) { // also a cleanup for registering

	deleteAppRepo, err1 := DeleteAppRepository(
		os.Getenv("TOKEN"),
		cluster,
		ns,
		appRepoName,
	)

	_, err2 := GetAppRepository(
		os.Getenv("TOKEN"),
		cluster,
		ns,
		appRepoName,
	)

	assert.True(t, deleteAppRepo)
	assert.Equal(t, err1, nil)
	assert.NotEqual(t, err2, nil)

}
