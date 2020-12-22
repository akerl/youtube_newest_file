package main

import (
	"fmt"
	"os"

	"github.com/akerl/slack_newest_file/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
