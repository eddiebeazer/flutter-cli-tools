package main

import (
	"fmt"
	"github.com/eddiebeazer/flutter-cli-tools/pkgs/fastlane"
	"github.com/eddiebeazer/flutter-cli-tools/pkgs/flutter"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "flutter-ci",
		Usage:     "Tools for Flutter and Fastlane CI",
		Version:   "0.1.0",
		UsageText: "Tools for Flutter and Fastlane CI.  Please run in your projects root directory",
		Suggest:   true,
		Authors: []*cli.Author{
			{
				Name:  "Edward Beazer",
				Email: "eddiebeazer@gmail.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "release",
				Usage: "Commands that help parse release information from a file or Github",
				Subcommands: []*cli.Command{
					{
						Name:  "getGithubRelease",
						Usage: "Gets a specific release from Github",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "owner",
								Aliases:  []string{"o"},
								Value:    "",
								Usage:    "Owner/Org of the repo",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "repo",
								Aliases:  []string{"r"},
								Value:    "",
								Usage:    "Name of the repo",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "version",
								Aliases:  []string{"v"},
								Value:    "",
								Usage:    "Version to fetch from github",
								Required: true,
							},
						},
						Action: func(cCtx *cli.Context) error {
							owner := cCtx.String("owner")
							repo := cCtx.String("repo")
							version := cCtx.String("version")
							release, err := fastlane.GetGithubRelease(owner, repo, version)
							if err != nil {
								return err
							}
							fmt.Printf(*release.TagName)
							return nil
						},
					},
					{
						Name:  "getLatestGithubRelease",
						Usage: "Gets the latest Github Release.  Can also get drafts with -d",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "owner",
								Aliases:  []string{"o"},
								Value:    "",
								Usage:    "Owner/Org of the repo",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "repo",
								Aliases:  []string{"r"},
								Value:    "",
								Usage:    "Name of the repo",
								Required: true,
							},
							&cli.BoolFlag{
								Name:     "draft",
								Aliases:  []string{"d"},
								Value:    false,
								Usage:    "If true, will fetch the latest draft release",
								Required: false,
							},
						},
						Action: func(cCtx *cli.Context) error {
							owner := cCtx.String("owner")
							repo := cCtx.String("repo")
							draft := cCtx.Bool("draft")
							release, err := fastlane.GetLatestGithubRelease(owner, repo, draft)
							if err != nil {
								return err
							}
							fmt.Printf(*release.TagName)
							return nil
						},
					},
				},
			},
			{
				Name:  "appVersion",
				Usage: "Get and set your projects app version",
				Subcommands: []*cli.Command{
					{
						Name:  "get",
						Usage: "Gets the app version from the pubspec file",
						Action: func(cCtx *cli.Context) error {
							version, err := flutter.GetAppVersion()
							if err != nil {
								return err
							}
							fmt.Printf(version)
							return err
						},
					},
					{
						Name:  "set",
						Usage: "Sets the app version",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "version",
								Aliases:  []string{"v"},
								Value:    "",
								Usage:    "Version of the app",
								Required: true,
							},
							&cli.StringFlag{
								Name:    "build",
								Aliases: []string{"b"},
								Value:   "0",
								Usage:   "Build number of the version.  Defaults to 0",
							},
						},
						Action: func(cCtx *cli.Context) error {
							version := cCtx.String("version")
							build := cCtx.String("build")
							return flutter.SetAppVersion(version, build)
						},
					},
					{
						Name:  "setFromGithub",
						Usage: "Sets the app version from the latest Github release",
						Flags: []cli.Flag{
							&cli.StringFlag{
								Name:     "owner",
								Aliases:  []string{"o"},
								Value:    "",
								Usage:    "Owner/Org of the repo",
								Required: true,
							},
							&cli.StringFlag{
								Name:     "repo",
								Aliases:  []string{"r"},
								Value:    "",
								Usage:    "Name of the repo",
								Required: true,
							},
							&cli.BoolFlag{
								Name:     "draft",
								Aliases:  []string{"d"},
								Value:    false,
								Usage:    "If true, will fetch the latest draft release",
								Required: false,
							},
							&cli.StringFlag{
								Name:    "build",
								Aliases: []string{"b"},
								Usage:   "Build number of the version",
							},
						},
						Action: func(cCtx *cli.Context) error {
							owner := cCtx.String("owner")
							repo := cCtx.String("repo")
							draft := cCtx.Bool("draft")
							build := cCtx.String("build")
							err := flutter.SetAppVersionFromLatestGithubRelease(owner, repo, draft, build)
							return err
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
