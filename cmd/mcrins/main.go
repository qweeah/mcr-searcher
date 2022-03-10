package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/qweeah/mcr-searcher/pkg/client"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		if len(args) == 0 {
			cmd.PrintErrln("No argument")
			return
		}
		keyword := args[0]

		repos, err := client.Catalog()
		if err != nil {
			cmd.PrintErrln(err)
			return
		}
		for _, repo := range repos {
			if strings.Contains(repo, keyword) {
				cmd.Println(repo)
			}
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
