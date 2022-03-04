package cmd

import (
	"fmt"
	"os"
)

var rootCmd = newRootCmd()

func init() {

}

func Execute() {
	if err := rootCmd.getCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
