package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gua",
	Short: "gua-cli (Github User Activity) is a lightweight and fast command-line tool built in Go for fetching and viewing GitHub user activity.",
	Long:  "gua-cli (GitHub User Activity CLI) is a powerful and lightweight command-line tool designed for developers and GitHub enthusiasts. Written in Go, this tool allows users to quickly fetch and view various GitHub user activities, such as repositories, commits, issues, and pull requests, right from the terminal.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
