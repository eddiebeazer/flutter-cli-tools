package fastlane

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetGithubReleases(t *testing.T) {
	// Published release
	release, err := GetLatestGithubRelease("eddiebeazer", "unreal-ci", false)
	if err != nil {
		return
	}

	assert.Equal(t, *release.TagName, "0.1.3")
}

func TestGetLatestDraftRelease(t *testing.T) {
	// Draft Release
	release, err := GetLatestGithubRelease("eddiebeazer", "unreal-ci", true)
	if err != nil {
		return
	}

	assert.Equal(t, *release.TagName, "0.1.4")
}

func TestGetGithubRelease(t *testing.T) {
	// Draft Release
	release, err := GetGithubRelease("eddiebeazer", "unreal-ci", "0.1.3")
	if err != nil {
		return
	}

	assert.NotNil(t, release)
}
