# unreal-ci
[![License](https://img.shields.io/github/license/eddiebeazer/unreal-ci?color=blue)](https://opensource.org/license/mit/)

[![codecov](https://codecov.io/gh/eddiebeazer/flutter-cli-tools/graph/badge.svg?token=3Hc0j7C3wy)](https://codecov.io/gh/eddiebeazer/flutter-cli-tools)

[![Build status](https://badge.buildkite.com/0dd0a7b5430670feb6ba2b340d430b96c04a007bf5099fb611.svg)](https://buildkite.com/the-digital-sages/flutter-cli-tools)

[![Go Report Card](https://goreportcard.com/badge/github.com/eddiebeazer/flutter-cli-tools)](https://goreportcard.com/report/github.com/eddiebeazer/flutter-cli-tools)

## What Flutter CLI Tools does

This was a simple CLI I made to work in tandem with using fastlane and flutter.  In particular the 2 annoying tasks I felt
like I was dealing with in my deployment pipelines was setting the app version and getting release notes from GitHub (draft).
This takes care of that.  You can get and set flutter app versions as well as get release notes from GitHub

### Installation

```
go install https://github.com/eddiebeazer/flutter-cli-tools
flutter-cli-tools help
```