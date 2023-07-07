package main

import (
	"os"

	"github.com/pacoguzman/hinge/cmd"
)

func main() {
	defer os.Exit(0)
	cmd.Execute()
}
