package main

import (
	"github.com/puppetlabs/pe-cli/puppet-access/cmd"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete-token-file",
	Short: "Removes the local token, but does not expire the token on the server",
	Run:   executeDeleteCommand,
}

func init() {
	cmd.RootCmd.AddCommand(deleteCmd)
}

func executeDeleteCommand(cmd *cobra.Command, args []string) {

}
