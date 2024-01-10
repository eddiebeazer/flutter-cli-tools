package fastlane

import (
	"context"
	"errors"
	"github.com/google/go-github/v58/github"
	"os"
)

// GetGithubRelease Get Github release information
func GetGithubRelease(owner string, repo string, version string) (*github.RepositoryRelease, error) {
	client := github.NewClient(nil)
	release, _, err := client.Repositories.GetReleaseByTag(context.Background(), owner, repo, version)
	if err != nil {
		return nil, err
	}

	return release, nil
}

// GetLatestGithubRelease Get the latest draft or released release info from Github
func GetLatestGithubRelease(owner string, repo string, getDraftRelease bool) (*github.RepositoryRelease, error) {
	githubToken := os.Getenv("GITHUB_TOKEN")
	client := github.NewClient(nil).WithAuthToken(githubToken)
	opt := &github.ListOptions{Page: 1, PerPage: 50}
	releases, _, err := client.Repositories.ListReleases(context.Background(), owner, repo, opt)
	if err != nil {
		return nil, err
	}

	var latestRelease *github.RepositoryRelease

	for _, release := range releases {
		if getDraftRelease != *release.Draft {
			continue
		}
		if latestRelease == nil {
			latestRelease = release
			continue
		}
		if release.PublishedAt.After(latestRelease.PublishedAt.Time) {
			latestRelease = release
			continue
		}
	}

	if latestRelease == nil {
		return nil, errors.New("release not found")
	}
	return latestRelease, nil
}
