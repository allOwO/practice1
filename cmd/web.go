package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	clientCmd := &cobra.Command{
		Use:   "dashboard",
		Short: "Messenger web",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	rootCmd.AddCommand(clientCmd)
}