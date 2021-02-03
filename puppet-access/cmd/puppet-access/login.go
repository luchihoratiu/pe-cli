package main

import (
	"github.com/puppetlabs/pe-cli/puppet-access/cmd"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to generate a token in RBAC",
	Run:   executeLoginCommand,
}

func init() {
	cmd.RootCmd.AddCommand(loginCmd)
}

func executeLoginCommand(cmd *cobra.Command, args []string) {

}
