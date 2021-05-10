package lib

// Configuration - represents the content of dependabot.yml
type Configuration struct {
	Version int `yaml:"version"`
	Updates   []Update `yaml:"updates"`
}

// Update - an update set in dependabot.yml
type Update struct {
	PackageEcosystem    string   `yaml:"package-ecosystem"`
	Directory string `yaml:"directory"`
	Schedule struct{
		Interval string `yaml:"interval"`
	} `yaml:"schedule"`
}
