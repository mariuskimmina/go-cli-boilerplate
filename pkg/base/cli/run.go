package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command) int {
	if err := cmd.Execute(); err != nil {
		fmt.Println("Failed to execute command")
		return 1
	}
	return 0
}
