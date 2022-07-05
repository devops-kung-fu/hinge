package lib

import (
	"log"
	"os"
	"path"
	"strings"

	"github.com/devops-kung-fu/common/file"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
)

// Generate generates the dependabot.yml in the specified repo path.
func Generate(afs *afero.Afero, repoPath string, schedule Schedule) (config Configuration, err error) {
	log.Println("Starting file generation")

	bundler := platform(afs, `Gemfile|Gemfile\.lock`, repoPath, "bundler", schedule)
	cargo := platform(afs, `Cargo\.toml`, repoPath, "cargo", schedule)
	composer := platform(afs, `composer\.json`, repoPath, "composer", schedule)
	docker := platform(afs, `Dockerfile(.*)|docker\-compose\.yml`, repoPath, "docker", schedule)
	elm := platform(afs, `elm\-package\.json`, repoPath, "elm", schedule)
	gitsubmodules := platform(afs, `\.gitmodules`, repoPath, "gitsubmodule", schedule)
	github := platform(afs, `(.*)\.(yaml|yml)`, repoPath, "github-actions", schedule)
	var githubActual []Update
	for _, githubPath := range github {
		if strings.Contains(githubPath.Directory, ".github/workflows") {
			githubActual = append(githubActual, githubPath)
		}
	}
	gomod := platform(afs, `go\.mod`, repoPath, "gomod", schedule)
	gradle := platform(afs, `build\.gradle|build\.gradle\.kts`, repoPath, "gradle", schedule)
	hexmix := platform(afs, `mix\.exs|mix\.lock`, repoPath, "mix", schedule)
	maven := platform(afs, `pom\.xml`, repoPath, "maven", schedule)
	npm := platform(afs, `package\.json|package\-lock\.json`, repoPath, "npm", schedule)
	nuget := platform(afs, `\.nuspec`, repoPath, "nuget", schedule)
	pip := platform(afs, `requirements\.txt|requirement\.txt|Pipfile|Pipfile\.lock|setup\.py|requirements\.in|pyproject\.toml`, repoPath, "pip", schedule)
	terraform := platform(afs, `(.*)\.tf`, repoPath, "terraform", schedule)

	log.Println("Got platform ecosystems")

	updates := joinUpdates(bundler, cargo, composer, docker, elm, gitsubmodules, githubActual, gomod, gradle, hexmix, maven, npm, nuget, pip, terraform)

	log.Println("Joined all updates")

	config = Configuration{
		Version: 2,
		Updates: updates,
	}

	log.Println("Configuration complete")

	err = writeConfig(afs, repoPath, config)
	return
}

func platform(afs *afero.Afero, regex string, repoPath string, ecosystem string, schedule Schedule) (updates []Update) {
	uniqueDirs := removeDuplicates(directoryParser(afs, regex, repoPath))
	updates = updatesBuilder(uniqueDirs, ecosystem, schedule)
	return
}

func directoryParser(afs *afero.Afero, regex string, repoPath string) []string {
	fileDiscovery, err := file.FindByRegex(afs, repoPath, regex)
	if err != nil {
		log.Println(err)
	}
	var cleanDiscovery []string
	for _, file := range fileDiscovery {
		if !strings.HasPrefix(file, ".github") {
			cleanFile := strings.Replace(file, repoPath, "", 1)
			cleanDiscovery = append(cleanDiscovery, cleanFile)
		} else if strings.HasPrefix(file, ".github") {
			cleanDiscovery = append(cleanDiscovery, "/"+file)
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

func removeDuplicates(stringSlice []string) (list []string) {
	keys := make(map[string]bool)
	for _, entry := range stringSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return
}

func updatesBuilder(directories []string, ecosystem string, schedule Schedule) (updates []Update) {
	for _, dir := range directories {
		update := Update{
			PackageEcosystem: ecosystem,
			Directory:        dir,
			Schedule:         schedule,
		}
		updates = append(updates, update)
	}
	return
}

func joinUpdates(updates ...[]Update) (joinedUpdate []Update) {
	for _, update := range updates {
		joinedUpdate = append(joinedUpdate, update...)
	}
	if len(joinedUpdate) > 200 {
		joinedUpdate = joinedUpdate[:200]
	}
	return
}

func writeConfig(afs *afero.Afero, repoPath string, config Configuration) (err error) {
	githubPath := path.Join(repoPath, ".github", "dependabot.yml")
	githubDir := path.Join(repoPath, ".github")
	if _, err := afs.Stat(githubDir); os.IsNotExist(err) {
		_ = afs.Mkdir(githubDir, 0755)
	}
	yaml, err := yaml.Marshal(config)
	if err != nil {
		return
	}
	log.Println(config)
	log.Println("Writing configuration")
	err = afs.WriteFile(githubPath, yaml, 0644)
	return
}
