package cmd

import (
	"strings"

	"github.com/mhaatha/gua-cli/internal/service"
	"github.com/spf13/cobra"
)

var usernameCmd = &cobra.Command{
	Use:   "username [github_username]",
	Short: "Retrieves detailed activity information for a specified GitHub user.",
	Long:  "username command is designed to fetch and display detailed activity data for any specified GitHub user. By entering the username, you can easily retrieve the user's public repositories, recent commits, issues they've opened, and pull requests they've submitted.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
		}

		joinedArgs := strings.Join(args, " ")
		service.GetUsername(joinedArgs)
	},
}

func init() {
	// Disable [flags] after a command name
	usernameCmd.DisableFlagsInUseLine = true

	rootCmd.AddCommand(usernameCmd)
}
