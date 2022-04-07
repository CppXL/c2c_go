package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.getCommand().AddCommand(newConfigCmd().cmd)
}

func newConfigCmd() *baseCmd {
	return newBaseCmdPoint(
		&cobra.Command{
			Use:   "config",
			Short: "使用特定的配置文件",
			Long:  `让C&C服务端加载特定的配置文件`,
			Run: func(cmd *cobra.Command, args []string) {
				fmt.Println("Hugo Static Site Generator v0.9 -- HEAD")
			},
		},
	)
}
