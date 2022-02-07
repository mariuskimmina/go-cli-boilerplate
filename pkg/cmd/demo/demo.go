package demo

import (
	"log"

	"github.com/spf13/cobra"
)

type DemoOptions struct {
	cmdBaseName string
}

type DemoFlags struct {
}

func NewDemoFlags() *DemoFlags {
	return &DemoFlags{}
}

func (flags *DemoFlags) AddFlags(cmd *cobra.Command) {

}

func NewCmdDemo(baseName string) *cobra.Command {
	flags := NewDemoFlags()

	cmd := &cobra.Command{
		Use:     "demo",
		Short:   "this is a demo command",
		Example: "demo doesn't do anything",
		Run: func(cmd *cobra.Command, args []string) {
			o, err := flags.ToOptions(cmd, baseName, args)
			if err != nil {
				log.Fatal("Failed to execute demo command")
			}
			err = o.Run()
			if err != nil {
				log.Fatal("Failed to execute demo command")
			}
		},
	}

	flags.AddFlags(cmd)

	return cmd
}

func (flags *DemoFlags) ToOptions(cmd *cobra.Command, baseName string, args []string) (*DemoOptions, error) {
	o := &DemoOptions{
		cmdBaseName: baseName,
	}

	return o, nil
}

func (o DemoOptions) Run() error {
	return nil
}
