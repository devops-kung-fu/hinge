// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/devops-kung-fu/common/util"
	"github.com/gookit/color"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"

	"github.com/pacoguzman/hinge/lib"
)

var (
	version = "1.0.1"
	//Afs stores a global OS Filesystem that is used throughout hinge
	Afs = &afero.Afero{Fs: afero.NewOsFs()}
	//Verbose determines if the execution of hing should output verbose information
	Verbose        bool
	debug          bool
	interval       string
	day            string
	time           string
	timeZone       string
	rebaseStrategy string
	rootCmd        = &cobra.Command{
		Use:     "hinge [flags] path/to/repo",
		Example: "  hinge path/to/repo",
		Short:   "Creates or updates your Dependabot config.",
		Version: version,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if !debug {
				log.SetOutput(ioutil.Discard)
			}
			util.DoIf(Verbose, func() {
				fmt.Println()
				color.Style{color.FgWhite, color.OpBold}.Println("█▄█ █ █▄ █ ▄▀  ██▀")
				color.Style{color.FgWhite, color.OpBold}.Println("█ █ █ █ ▀█ ▀▄█ █▄▄")
				fmt.Println()
				fmt.Println("DKFM - DevOps Kung Fu Mafia")
				fmt.Println("https://github.com/pacoguzman/hinge")
				fmt.Printf("Version: %s\n", version)
				fmt.Println()
			})
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				util.PrintErr(errors.New("please provide the path to a git repository"))
				_ = cmd.Usage()
				os.Exit(1)
			} else if len(args) > 1 {
				util.PrintErr(errors.New("only one path is allowed"))
				_ = cmd.Usage()
				os.Exit(1)
			}
			path := filepath.Join(args[0], ".git")
			//b, err := Afs.DirExists(".git")
			b, err := Afs.DirExists(path)
			util.IfErrorLog(err)

			if !b {
				e := errors.New("fatal: provided path does not contain a git repository")
				util.PrintErr(e)
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			config, err := lib.Generate(Afs, args[0], buildRebaseStrategy(), buildSchedule())
			if err != nil {
				log.Panic(err)
			}
			util.DoIf(Verbose, func() {
				util.PrintInfof("Found %x ecosystems\n", len(config.Updates))
				for _, update := range config.Updates {
					util.PrintInfof("Processed Ecosystem: %s\n", update.PackageEcosystem)
				}
				util.PrintInfo("Updated .github/dependabot.yml")
				util.PrintSuccess("Done!")
			})
		},
	}
)

// Execute creates the command tree and handles any error condition returned
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", true, "Displays command line output.")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Displays debug level log messages.")
	rootCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "daily", "How often to check for new versions.")
	rootCmd.PersistentFlags().StringVarP(&day, "day", "d", "monday", "Specify a day to check for updates when using a weekly interval.")
	rootCmd.PersistentFlags().StringVarP(&time, "time", "t", "05:00", "Specify a time of day to check for updates using 24 hour format (format: hh:mm).")
	rootCmd.PersistentFlags().StringVarP(&timeZone, "timezone", "z", "US/Pacific", "Specify a time zone. Valid timezones are available at https://en.wikipedia.org/wiki/List_of_tz_database_time_zones.")
	rootCmd.PersistentFlags().StringVarP(&rebaseStrategy, "rebase-strategy", "s", "disable", "Dependabot automatically rebases open pull requests when it detects any changes to the pull request")
}

func buildRebaseStrategy() (rebaseStrategy string) {
	switch {
	case rebaseStrategy == "auto":
		return ""
	case rebaseStrategy == "":
		return ""
	case rebaseStrategy == "disable":
		return "disable"
	}

	return ""
}

func buildSchedule() (schedule lib.Schedule) {
	schedule.Interval = strings.ToLower(schedule.Interval)
	switch {
	case interval == "daily":
		schedule.Interval = interval
		schedule.Time = time
		schedule.TimeZone = timeZone
	case interval == "weekly":
		schedule.Interval = interval
		schedule.Day = day
		schedule.Time = time
		schedule.TimeZone = timeZone
	case interval == "monthly":
		schedule.Interval = interval
	}
	schedule.Day = strings.ToLower(schedule.Day)
	return schedule
}
