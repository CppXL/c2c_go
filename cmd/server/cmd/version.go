package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.getCommand().AddCommand(newVersionCmd())
}

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "v0.0.1",
		Long:  `v0.0.1`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("c2c_go version v0.0.1")
		},
	}
}
