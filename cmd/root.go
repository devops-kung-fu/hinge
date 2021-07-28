// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"os"

	"github.com/devops-kung-fu/hinge/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	version  = "0.1.1"
	verbose  bool
	trace    bool
	debug    bool
	interval string
	day      string
	time     string
	timeZone string
	rootCmd  = &cobra.Command{
		Use:     "hinge [flags] path/to/repo",
		Example: "  hinge path/to/repo",
		Short:   "Creates and updates your Dependabot config.",
		Version: version,
		PreRun: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				color.Style{color.FgRed, color.OpBold}.Println("Please provide the path to the repository.")
				fmt.Println()
				_ = cmd.Usage()
				os.Exit(1)
			} else if len(args) > 1 {
				color.Style{color.FgRed, color.OpBold}.Println("Only one path is allowed.")
				fmt.Println()
				_ = cmd.Usage()
				os.Exit(1)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			repoPath := args[0]
			schedule := lib.Schedule{
				Interval: "daily",
			}
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
			lib.Generator(repoPath, verbose, schedule)

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
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "Displays dependabot.yml configuration in stardard output.")
	rootCmd.PersistentFlags().BoolVar(&trace, "trace", false, "Displays trace level log messages.")
	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Displays debug level log messages.")
	rootCmd.PersistentFlags().StringVarP(&interval, "interval", "i", "daily", "How often to check for new versions.")
	rootCmd.PersistentFlags().StringVarP(&day, "day", "d", "daily", "Specify a day to check for updates.")
	rootCmd.PersistentFlags().StringVarP(&time, "time", "t", "05:00", "Specify a time of day to check for updates (format: hh:mm).")
	rootCmd.PersistentFlags().StringVarP(&timeZone, "timezone", "z", "UTC", "Specify a time zone. The time zone identifier must be from the Time Zone database maintained by IANA.")
	color.Style{color.FgWhite, color.OpBold}.Println("Hinge")
	fmt.Println("https://github.com/devops-kung-fu/hinge")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("")
}
