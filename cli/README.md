# Arguments & Subcommands

## Conventions

Arguments/flags can be specified using

* --argument
* -argument

Value of argument can be specified using
* --argument value
* --argument=value
* --argument "value with space"

Help should follow following convention
* Use <> for **requierd** parameter. ping "<hostname>"
* Use [] for **optional** parameter. e.g. mkdir [-p] <dirname>
* Use {} for **choise**. e.g netstat {-t|-u}

## Arguments

Specifying only argument in the script (no subcommand)
```go

// Create the arguments
stringArg := flag.String("arguemnt", "Default Value", "Help Instruction (Required)")

intArg := flag.Int("arguemnt", 5, "Help Instruction (Required)")

boolArg := flag.Bool("arguemnt", false, "Help Instruction (Required)")

// Parse the flag
flag.Parse()

// Access the argument
stringArgValue := *stringArg
intArgValue := *intArg
boolArgValue := *boolArg

// Get help for the flag
flag.PrintDefaults()
```

## Subcommand

Subcommands are the arguments __without__ - or -- after the executable

```go
// Define the list of supported commands
getCommand := flag.NewFlagSet("get", flag.ExitOnError)

postCommand := flag.NewFlagSet("post", flag.ExitOnError)

// Add arguments to the subcommand
getStringArg := getCommand.String("arguemnt", "Default Value", "Help Instruction (Required)")

postIntArg := postCommand.Int("arguemnt", 5, "Help Instruction (Required)")

// Access the subcommand using switch case
switch os.Args[1] {
	case "get":
		getCommand.Parse(os.Args[2:])
	case "post":
		postCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
}
    
// Access the arguments of subcommand
if getCommand.Parsed() {
    // Logic
}

if postCommand.Parsed() {
    //Logic
}
```

## Reference
https://blog.rapid7.com/2016/08/04/build-a-simple-cli-tool-with-golang/
