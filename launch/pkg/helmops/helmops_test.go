package helmops

import (
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
