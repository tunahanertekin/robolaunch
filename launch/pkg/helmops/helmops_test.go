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

// cluster is set
// ns is set
var appRepoResource string = "robot-helm-charts"
var appRepoNamespace string = "default"
var chartName string = "jackal"
var releaseName string = "my-test-jackal"
var version string = "0.1.0"
var values string = ""

func TestCreateRelease(t *testing.T) {
	createRelease, err := CreateRelease(
		os.Getenv("TOKEN"),
		cluster,
		ns,
		CreateReleaseBody{
			AppRepositoryResourceName:      appRepoResource,
			AppRepositoryResourceNamespace: appRepoNamespace,
			ChartName:                      chartName,
			ReleaseName:                    releaseName,
			Version:                        version,
			Values:                         values,
		},
	)

	assert.Equal(t, createRelease.Data.Name, releaseName)
	assert.Equal(t, err, nil)
}

var changedVersion string = "0.1.1"

// for assertion, response must include chart metadata & values
func TestUpdateRelease(t *testing.T) {
	_, err := UpdateRelease(
		os.Getenv("TOKEN"),
		cluster,
		ns,
		releaseName,
		UpdateReleaseBody{
			AppRepositoryResourceName:      appRepoResource,
			AppRepositoryResourceNamespace: appRepoNamespace,
			ChartName:                      chartName,
			ReleaseName:                    releaseName,
			Version:                        changedVersion,
			Values:                         values,
		},
	)

	//assert.Equal(t, updateRelease.Version, changedVersion)
	assert.Equal(t, err, nil)
}

func TestDeleteRelease(t *testing.T) {
	deleteRelease, err := DeleteRelease(
		os.Getenv("TOKEN"),
		cluster,
		ns,
		releaseName,
	)

	assert.True(t, deleteRelease)
	assert.Equal(t, err, nil)
}
