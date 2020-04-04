package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// Getting subcommand
	// os.Arg[0] is the current executable
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("Usage flag.go {get | post} <URL> [--d  data] [-X header] [-t thread] [-b bool]")
		os.Exit(1)
	}

	// 2. Subcommands
	getCommand := flag.NewFlagSet("get", flag.ExitOnError)
	postCommand := flag.NewFlagSet("post", flag.ExitOnError)

	// Add the arguments to the subcommand
	getURLPtr := getCommand.String("url", "", "URL of the target (Required)")
	getHeaderPtr := getCommand.String("X", "X-Defalut-Header: roguesecurity", "Header sent along with the request")
	getResponsePtr := getCommand.Bool("b", false, "Return full response.")

	// Add the arguments to the subcommand
	postURLPtr := postCommand.String("url", "", "URL of the target (Required)")
	postThreadPtr := postCommand.Int("t", 5, "No of concurrent connections")
	postDataPtr := postCommand.String("d", "", "Data to be sent")

	// Get the received subcommand
	switch os.Args[1] {
	case "get":
		getCommand.Parse(os.Args[2:])
	case "post":
		postCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check the subcommand received
	if getCommand.Parsed() {

		// Check if required parameter has value
		// --url flag
		if *getURLPtr == "" {
			getCommand.PrintDefaults()
			os.Exit(1)
		} else {
			fmt.Println("The subcommand argument value is :", *getURLPtr)
		}

		fmt.Println("The post no of thread argument value is :", *getHeaderPtr)
		fmt.Println("The post data argument value is :", *postDataPtr)

	}

	// Check the subcommand received
	if postCommand.Parsed() {
		if *postURLPtr == "" {
			postCommand.PrintDefaults()
			os.Exit(1)
		} else {
			fmt.Println("The get header argument value is :", *postURLPtr)
		}
		fmt.Println("The get header argument value is :", *postThreadPtr)
		fmt.Println("The get reponse argument value is :", *getResponsePtr)

	}

}
