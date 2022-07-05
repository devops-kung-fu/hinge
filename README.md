![](img/hinge128x128.png)

# Hinge

![GitHub release (latest by date)](https://img.shields.io/github/v/release/devops-kung-fu/hinge) 
[![Go Report Card](https://goreportcard.com/badge/github.com/devops-kung-fu/hinge)](https://goreportcard.com/report/github.com/devops-kung-fu/hinge) 
[![codecov](https://codecov.io/gh/devops-kung-fu/hinge/branch/main/graph/badge.svg?token=BIROb1Npbk)](https://codecov.io/gh/devops-kung-fu/hinge) 
[![SBOM](https://img.shields.io/badge/CyloneDX-SBoM-informational)](hinge-sbom.json)


Creates and updates your Dependabot configuration file, `dependabot.yml`.

## Overview

**Hinge** automatically creates and updates a repository's `dependabot.yml` file, by recursively walking through the repository, identifying all supported **Dependabot** platform ecosystems, noting their paths relative to the repository root, and finally producing a YAML-compliant configuration in `/.github/dependabot.yml`.

## Dependabot?

[Dependabot](https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically) is GitHub's flagship product for [Supply Chain Security](https://docs.github.com/en/code-security/supply-chain-security). Dependabot takes the effort out of maintaining your dependencies. You can use it to ensure that your repository automatically keeps up with the latest releases of the packages and applications it depends on.

## Installation

To install ```hinge```, [download the latest release](https://github.com/devops-kung-fu/hinge/releases) , make it executable, rename it to `hinge` and move it to the `/usr/local/bin` directory for Linux, or on your `PATH` for other operating systems.

### Linux Example

```bash
sudo chmod +x hinge-1.0.0-linux-amd64
sudo mv hinge-1.0.0-linux-amd64 /usr/local/bin/hinge
```
### With a Go Development Environment

If you have a Go development environment set up, you can also simply do this:

``` bash
go install github.com/devops-kung-fu/hinge@latest
```

## Usage

```
Hinge
DKFM - DevOps Kung Fu Mafia
https://github.com/devops-kung-fu/hinge
Version: 1.0.0

Creates or updates your Dependabot config.

Usage:
  hinge [flags] path/to/repo

Examples:
  hinge path/to/repo

Flags:
  -d, --day string        Specify a day to check for updates when using a weekly interval. (default "monday")
      --debug             Displays debug level log messages.
  -h, --help              help for hinge
  -i, --interval string   How often to check for new versions. (default "daily")
  -t, --time string       Specify a time of day to check for updates using 24 hour format (format: hh:mm). (default "05:00")
  -z, --timezone string   Specify a time zone. Valid timezones are available at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. (default "US/Pacific")
  -v, --verbose           Displays command line output. (default true)
      --version           version for hinge
```
### Flag Notes

| Flag | Notes |
|------|---|
|-d, --day | Must be a valid day of the week. (monday, tuesday, wednesday, thursday, friday, saturday, sunday). Defaults to monday if using a weekly interval.|
|-i, --interval | Must be one of the following: daily, weekly, monthly. Defaults to daily.|
|-z, --timezone | Must be a timezone listed at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones. Defaults to "US/Pacific if not explicitly defined.|



## Development

## Overview

In order to use contribute and participate in the development of ```hinge``` you'll need to have an updated Go environment. Before you start, please view the [Contributing](CONTRIBUTING.md) and [Code of Conduct](CODE_OF_CONDUCT.md) files in this repository.

## Prerequisites

This project makes use of [DKFM](https://github.com/devops-kung-fu) tools such as [Hookz](https://github.com/devops-kung-fu/hookz) and other open source tooling. Install these tools with the following commands:

``` bash

go install github.com/devops-kung-fu/hookz@latest
go install github.com/kisielk/errcheck@latest
go install golang.org/x/lint/golint@latest
go install github.com/fzipp/gocyclo@latest

```

## Software Bill of Materials

```hinge``` uses the CycloneDX to generate a Software Bill of Materials in CycloneDX format (v1.4) every time a developer commits code to this repository (as long as [Hookz](https://github.com/devops-kung-fu/hookz) is being used and has been initialized in the working directory). More information for CycloneDX is available [here](https://cyclonedx.org)

The current SBoM for ```hinge``` is available [here](hinge-sbom.json).

## Credits

A big thank-you to our friends at [Freepik](https://www.freepik.com) for the ```hinge``` logo. 
