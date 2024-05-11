<p align="center">
<img src="img/logo.svg" alt="git-stack logo" height=250>
</p>

# git-stack

[![GitHub release (latest by date)][release-img]][release]
[![GolangCI][golangci-lint-img]][golangci-lint]
[![Go Report Card][report-card-img]][report-card]

Stack based flow for GitHub.

## Installation

### macOS

```sh
brew install fcjr/fcjr/git-stack
```

### Windows

```sh
scoop bucket add fcjr https://github.com/fcjr/scoop-fcjr.git
scoop install git-stack
```

### Other

If you are on another platform such as linux or just rather not use a package manager, you can download a binary from Github releases and use it straight away without having to install any additional dependencies.

1) Find the latest release, download the tar.gz file for your given operating system and extract it.
2) Inside you'll find the `git-stack` executable which you can run directly (Note: you may need to allow execution via `chmod +x <PATH_TO_GIT_STACK_EXE>`).


#### Usage
Basic:

`git stack -h`

### Development

#### Building the current commit

This project uses [goreleaser](https://github.com/goreleaser/goreleaser/).

 1) Install [go](https://golang.org/doc/install).
 2) Install goreleaser via the steps [here](https://goreleaser.com/install/).
 3) Build current commit via `goreleaser release --snapshot --skip-publish --rm-dist`.

[release-img]: https://img.shields.io/github/v/release/fcjr/git-stack
[release]: https://github.com/fcjr/git-stack/releases
[golangci-lint-img]: https://github.com/fcjr/git-stack/workflows/go-lint/badge.svg
[golangci-lint]: https://github.com/fcjr/git-stack/actions?query=workflow%3Ago-lint
[report-card-img]: https://goreportcard.com/badge/github.com/fcjr/git-stack
[report-card]: https://goreportcard.com/report/github.com/fcjr/git-stack