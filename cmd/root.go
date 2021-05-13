// Package cmd contains all of the commands that may be executed in the cli
package cmd

import (
	"fmt"
	"os"

	"github.com/devops-kung-fu/heybo"
	"github.com/devops-kung-fu/hinge/lib"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

var (
	version = "0.0.9"
	verbose bool
	trace   bool
	debug   bool
	rootCmd = &cobra.Command{
		Use:     "hinge [flags] path/to/repo",
		Example: "  hinge path/to/repo",
		Short:   "Creates and updates your Dependabot config.",
		Version: version,
		Run: func(cmd *cobra.Command, args []string) {
			heyBo := heybo.HeyBo(heybo.DEBUG)
			switch {
			case trace:
				heyBo.ChangeGlobalLevel(heybo.ALL)
			case debug:
				heyBo.ChangeGlobalLevel(heybo.TRACE)
			}
			heyBo.ChangeTagText(heybo.INFO, "pass")
			if len(args) == 0 {
				color.Style{color.FgRed, color.OpBold}.Println("Please provide the path to the repository.")
				fmt.Println()
				_ = cmd.Usage()
			} else if len(args) > 1 {
				color.Style{color.FgRed, color.OpBold}.Println("Only one path is allowed.")
				fmt.Println()
				_ = cmd.Usage()
			} else {
				repoPath := args[0]
				lib.Generator(heyBo, repoPath, verbose)
			}
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
	rootCmd.PersistentFlags().BoolVarP(&trace, "trace", "t", false, "Displays trace level log messages.")
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Displays debug level log messages.")
	color.Style{color.FgWhite, color.OpBold}.Println("Hinge")
	fmt.Println("https://github.com/devops-kung-fu/hinge")
	fmt.Printf("Version: %s\n", version)
	fmt.Println("")
}
