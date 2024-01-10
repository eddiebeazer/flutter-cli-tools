package flutter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAppVersion(t *testing.T) {
	version, err := GetAppVersion()
	if err != nil {
		return
	}
	assert.Equal(t, version, "0.1.0+0")

	// Manually setting app version
	err = SetAppVersion("2.5.5", "23")
	if err != nil {
		return
	}
	version, err = GetAppVersion()
	if err != nil {
		return
	}
	assert.Equal(t, version, "2.5.5+23")

	// Setting app version from Github
	err = SetAppVersionFromLatestGithubRelease("eddiebeazer", "unreal-ci", false, "532")
	if err != nil {
		return
	}
	version, err = GetAppVersion()
	if err != nil {
		return
	}
	assert.Equal(t, version, "0.1.3+532")
}

func TestSetAppVersion(t *testing.T) {
	// Manually setting app version
	err := SetAppVersion("2.5.5", "23")
	if err != nil {
		return
	}
	version, err := GetAppVersion()
	if err != nil {
		return
	}
	assert.Equal(t, version, "2.5.5+23")
}

func TestSetAppVersionFromGitHub(t *testing.T) {
	// Setting app version from Github
	err := SetAppVersionFromLatestGithubRelease("eddiebeazer", "unreal-ci", false, "532")
	if err != nil {
		return
	}
	version, err := GetAppVersion()
	if err != nil {
		return
	}
	assert.Equal(t, version, "0.1.3+532")
}
