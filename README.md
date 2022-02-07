# Go CLI Boilerplate

This project can be used a baseline to build your next go cli project. The code structure is based on `kubectl` by Kubernetes.


## Usage

The easiest way to use this code is to clone it and remove the git directory

```
git clone github.com/mariuskimmina/go-cli-boilerplate {yourProjectName}
cd {yourProjectName}
rm -r .git
```

You can then replace the occurences of "Demo" inside the codebase with the name of your own project.

## What's in it?

Running the base command with `go run cmd/demo/main.go`

```
demo is boilerplate code for building your next go cli

Usage:
  demo [flags]
  demo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  demo        this is a demo command
  help        Help about any command

Flags:
  -h, --help   help for demo

Use "demo [command] --help" for more information about a command.
```

An exmaple command also called `demo`, that you can run with `go run cmd/demo/main.go demo -h`

```
this is a demo command

Usage:
  demo demo [flags]

Examples:
demo doesn't do anything

Flags:
  -h, --help   help for demo
```
