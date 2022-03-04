package cmd

import "github.com/spf13/cobra"

// 声明基本cmd结构体
type baseCmd struct {
	cmd *cobra.Command
}

// 实现cmder接口，该方法返回basecmd里面的cobra.Command指针
func (c *baseCmd) getCommand() *cobra.Command {
	return c.cmd
}

// 通过 cobra.Command 型指针 新建basecmd变量并返回对应变量的指针
func newBaseCmd(cmd *cobra.Command) *baseCmd {
	return &baseCmd{cmd: cmd}
}

// 向root cmd添加实现了cmder接口的cmd
func addCommands(root *cobra.Command, commands ...cmder) {
	for _, command := range commands {
		cmd := command.getCommand()
		if cmd == nil {
			continue
		}
		root.AddCommand(cmd)
	}
}

// 新建rootcmd
func newRootCmd() *baseCmd {
	return newBaseCmd(
		&cobra.Command{
			Use:   "server",
			Short: "这是c&c的服务端，控制bot",
			Long:  `这是c&c的服务端，控制bot`,
		},
	)
}
