package flutter

import (
	"github.com/eddiebeazer/flutter-cli-tools/pkgs/fastlane"
	"github.com/spf13/viper"
	"log"
	"os"
)

type Pubspec struct {
	Version string `yaml:"version"`
}

func currentDir() (cwd string) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return cwd
}

func ParsePubspecFile() error {
	viper.SetConfigName("pubspec")
	viper.AddConfigPath(currentDir())

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return err
	}

	return nil
}

// GetAppVersion Get App version from pubspec
func GetAppVersion() (string, error) {
	err := ParsePubspecFile()
	if err != nil { // Handle errors reading the config file
		return "", err
	}

	viper.GetViper()

	return viper.GetString("version"), nil
}

// SetAppVersion sets the app version in pubspec to version
func SetAppVersion(version string, buildNumber string) error {
	err := ParsePubspecFile()
	if err != nil { // Handle errors reading the config file
		return err
	}

	viper.Set("version", version+"+"+buildNumber)
	err = viper.WriteConfig()
	if err != nil {
		return err
	}

	return nil
}

// SetAppVersionFromLatestGithubRelease SetAppVersion sets the app version in pubspec to version
func SetAppVersionFromLatestGithubRelease(owner string, repo string, draft bool, buildNumber string) error {
	release, err := fastlane.GetLatestGithubRelease(owner, repo, draft)
	if err != nil {
		return err
	}

	err = SetAppVersion(*release.TagName, buildNumber)
	if err != nil {
		return err
	}

	return nil
}
