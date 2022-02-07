package main

import (
	"os"

	"github.com/mariuskimmina/go-cli-boilerplate/pkg/base/cli"
	"github.com/mariuskimmina/go-cli-boilerplate/pkg/cmd"
)

func main() {
	command := cmd.NewDefaultDemoCommand()
	code := cli.Run(command)
	os.Exit(code)
}
