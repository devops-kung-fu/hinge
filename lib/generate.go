package lib

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/devops-kung-fu/heybo"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

// Generator - generates the dependabot.yml in the specified repo path.
func Generator(logger *heybo.Logger, repoPath string, verbose bool, schedule Schedule) {
	logger.Trace("Setting up filesystem.")
	fs := afero.NewOsFs()
	logger.Debug("Getting platform ecosystems.")
	bundler := platform(fs, `Gemfile|Gemfile\.lock`, repoPath, "bundler", schedule)
	cargo := platform(fs, `Cargo\.toml`, repoPath, "cargo", schedule)
	composer := platform(fs, `composer\.json`, repoPath, "composer", schedule)
	docker := platform(fs, `Dockerfile(.*)|docker\-compose\.yml`, repoPath, "docker", schedule)
	elm := platform(fs, `elm\-package\.json`, repoPath, "elm", schedule)
	gitsubmodules := platform(fs, `\.gitmodules`, repoPath, "gitsubmodule", schedule)
	github := platform(fs, `(.*)\.(yaml|yml)`, repoPath, "github-actions", schedule)
	var githubActual []Update
	for _, githubPath := range github {
		if strings.Contains(githubPath.Directory, ".github/workflows") {
			githubActual = append(githubActual, githubPath)
		}
	}
	gomod := platform(fs, `go\.mod`, repoPath, "gomod", schedule)
	gradle := platform(fs, `build\.gradle|build\.gradle\.kts`, repoPath, "gradle", schedule)
	hexmix := platform(fs, `mix\.exs|mix\.lock`, repoPath, "mix", schedule)
	maven := platform(fs, `pom\.xml`, repoPath, "maven", schedule)
	npm := platform(fs, `package\.json|package\-lock\.json`, repoPath, "npm", schedule)
	nuget := platform(fs, `\.nuspec`, repoPath, "nuget", schedule)
	pip := platform(fs, `requirements\.txt|requirement\.txt|Pipfile|Pipfile\.lock|setup\.py|requirements\.in|pyproject\.toml`, repoPath, "pip", schedule)
	terraform := platform(fs, `(.*)\.tf`, repoPath, "terraform", schedule)
	logger.Info("Got platform ecosystems.")
	logger.Debug("Begin joining updates.")
	updates := joinUpdates(bundler, cargo, composer, docker, elm, gitsubmodules, githubActual, gomod, gradle, hexmix, maven, npm, nuget, pip, terraform)
	logger.Info("Joined all updates.")
	logger.Debug("Building configuration.")
	config := Configuration{
		Version: 2,
		Updates: updates,
	}
	logger.Info("Configuration complete.")
	if verbose {
		logger.Trace("Output configuration to standard output.")
		outputConfig(config)
	}
	logger.Debug("Writing configuration.")
	writeConfig(fs, repoPath, config)
	logger.Info("Done.")
}

func platform(fs afero.Fs, regex string, repoPath string, ecosystem string, schedule Schedule) []Update {
	dirs := directoryParser(fs, regex, repoPath)
	uniqueDirs := removeDuplicates(dirs)
	updates := updatesBuilder(uniqueDirs, ecosystem, schedule)
	return updates
}

func directoryParser(fs afero.Fs, regex string, repoPath string) []string {
	fileDiscovery, err := FindFiles(fs, repoPath, regex)
	if err != nil {
		fmt.Println(err)
	}
	var cleanDiscovery []string
	for _, file := range fileDiscovery {
		if !strings.HasPrefix(file, ".github") {
			cleanFile := strings.Replace(file, repoPath, "", 1)
			cleanDiscovery = append(cleanDiscovery, cleanFile)
		} else if strings.HasPrefix(file, ".github") {
			cleanDiscovery = append(cleanDiscovery, "/" + file)
		} else {
			cleanDiscovery = append(cleanDiscovery, file)
		}
	}
	var dotDirectory []string
	for _, filePath := range cleanDiscovery {
		cleanDir := path.Dir(filePath)
		dotDirectory = append(dotDirectory, cleanDir)
	}
	var directory []string
	for _, dots := range dotDirectory {
		if len(dots) == 1 {
			slashes := strings.Replace(dots, ".", "/", 1)
			directory = append(directory, slashes)
		} else {
			directory = append(directory, dots)
		}
	}
	return directory
}

func removeDuplicates(stringSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}

	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func updatesBuilder(directories []string, ecosystem string, schedule Schedule) []Update {
	var updates []Update
	for _, dir := range directories {
		update := Update{
			PackageEcosystem: ecosystem,
			Directory:        dir,
			Schedule:         schedule,
		}
		updates = append(updates, update)
	}
	return updates
}

func joinUpdates(updates ...[]Update) []Update {
	var joinedUpdate []Update
	for _, update := range updates {
		joinedUpdate = append(joinedUpdate, update...)
	}
	if len(joinedUpdate) > 200 {
		joinedUpdate = joinedUpdate[:200]
	}
	return joinedUpdate
}

func outputConfig(config Configuration) {
	yamlOutput, err := yaml.Marshal(config)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println(string(yamlOutput))
}

func writeConfig(fs afero.Fs, repoPath string, config Configuration) {
	githubPath := path.Join(repoPath, ".github", "dependabot.yml")
	githubDir := path.Join(repoPath, ".github")
	if _, err := fs.Stat(githubDir); os.IsNotExist(err) {
		_ = fs.Mkdir(githubDir, 0755)
	}
	outFile, err := fs.OpenFile(githubPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
	}
	yamlEncoder := yaml.NewEncoder(outFile)
	encodeErr := yamlEncoder.Encode(config)
	if encodeErr != nil {
		fmt.Println(encodeErr)
	}
	_ = yamlEncoder.Close()
}
