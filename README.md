![](img/hinge128x128.png)

# Hinge

[![Go Report Card](https://goreportcard.com/badge/github.com/devops-kung-fu/hinge)](https://goreportcard.com/report/github.com/devops-kung-fu/hinge) ![GitHub release (latest by date)](https://img.shields.io/github/v/release/devops-kung-fu/hinge) [![codecov](https://codecov.io/gh/devops-kung-fu/hinge/branch/main/graph/badge.svg?token=BIROb1Npbk)](https://codecov.io/gh/devops-kung-fu/hinge)

Creates and updates your Dependabot configuration file, `dependabot.yml`.

## Overview

**Hinge** automatically creates and updates a repository's `dependabot.yml` file, by recursively walking through the repository, identifying all supported **Dependabot** platform ecosystems, noting their paths relative to the repository root, and finally producing a YAML-compliant configuration in `/.github/dependabot.yml`.

## Dependabot?

[Dependabot](https://docs.github.com/en/code-security/supply-chain-security/keeping-your-dependencies-updated-automatically) is GitHub's flagship product for [Supply Chain Security](https://docs.github.com/en/code-security/supply-chain-security). Dependabot takes the effort out of maintaining your dependencies. You can use it to ensure that your repository automatically keeps up with the latest releases of the packages and applications it depends on.

## Installation

To install Hinge, [download the latest release](https://github.com/devops-kung-fu/hinge/releases) , make it executable, rename it to `hinge` and move it to the `/usr/local/bin` directory for Linux, or on your `PATH` for other operating systems.

### Linux Example

```bash
sudo chmod +x hinge-1.0.0-linux-amd64
sudo mv hinge-1.0.0-linux-amd64 /usr/local/bin/hinge
```

## How To Hinge :dancer: