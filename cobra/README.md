# CLI Boilerplate

## Project Setup
```bash
# Create a go module for the cli
go mod init testctl

# Install cobra (This will create the cobra executable in $GOPATH/bin directory)
env GO111MODULE=on go get github.com/spf13/cobra/cobra

# Initialize cobra library
cobra init --pkg-name testctl 
```

## Create new commands
```bash
cobra add create

cobra add pod
```

## Run and Install CLI
```bash
# Run 
go run main.go create pod
# Install
go install testctl
```

### Instructions

* main.go calls `Execute()` function inside cmd/root.go
* root.go is the top level command
* All the commands by default are created as top level command. To a commond as subcommand, modify the `AddCommand` inside `init()`
* Flags for a command/subcomand are defined inside `init()`
* Command line arguments are read inside `Run()`