package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/mariuskimmina/go-cli-boilerplate/pkg/cmd/demo"
)

type DemoOptions struct {
	Arguments []string
}

func NewDefaultDemoCommand() *cobra.Command {
	return NewDefaultDemoCommandWithArgs(DemoOptions{
		Arguments: os.Args,
	})
}

func NewDefaultDemoCommandWithArgs(o DemoOptions) *cobra.Command {
	cmd := NewDemoCommand(o)

	return cmd

}

func NewDemoCommand(o DemoOptions) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "demo",
		Short: "demo is boilerplate code for building your next go cli",
		Run:   runHelp,
	}
	cmds.AddCommand(demo.NewCmdDemo("demo"))

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	err := cmd.Help()
	if err != nil {
		log.Fatal("Failed to run Help command")
	}
}
