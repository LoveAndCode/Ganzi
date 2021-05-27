package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"strings"
)

var echoCommand = &cobra.Command{
	Use:   "echo",
	Short: "Echo Argument parameter",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello Ganzi! args: => [" + strings.Join(args, " , ") + "]")
	},
}

func init() {
	rootCmd.AddCommand(echoCommand)
}
