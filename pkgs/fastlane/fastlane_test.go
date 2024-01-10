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

	// Draft Release
	release, err = GetLatestGithubRelease("eddiebeazer", "unreal-ci", true)
	if err != nil {
		return
	}

	assert.Equal(t, *release.TagName, "0.1.4")
}
