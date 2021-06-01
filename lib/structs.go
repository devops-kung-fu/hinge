package lib

// Configuration - represents the content of dependabot.yml
type Configuration struct {
	Version int      `yaml:"version"`
	Updates []Update `yaml:"updates"`
}

// Update - an update set in dependabot.yml
type Update struct {
	PackageEcosystem string `yaml:"package-ecosystem"`
	Directory        string `yaml:"directory"`
	Schedule `yaml:"schedule"`
}

// Schedule - Update check schedule
type Schedule struct {
	Interval string `yaml:"interval"`
	Day string `yaml:"day,omitempty"`
	Time string `yaml:"time,omitempty"`
	TimeZone string `yaml:"timezone,omitempty"`
}