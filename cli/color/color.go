package main

import (
	"fmt"

	"github.com/fatih/color"
)

// DEBUG
const (
	DEBUG   = "debug"
	SUCCESS = "success"
	ERROR   = "error"
)

func main() {

	color.White("Starting the program...")

	PrintMessage(DEBUG, "Trying to connect...")
	PrintMessage(SUCCESS, "Connection Successful !!!")
	PrintMessage(ERROR, "Error occured")

}

// PrintMessage method prints the colored message
func PrintMessage(infotype, message string) {

	switch infotype {
	case "debug":
		color.Set(color.FgHiYellow)
		defer color.Unset()
		fmt.Println(message)
	case "success":
		color.Set(color.FgHiGreen)
		defer color.Unset()
		fmt.Println(message)
	case "error":
		color.Set(color.FgHiRed)
		defer color.Unset()
		fmt.Println(message)
	default:
		color.Set(color.FgHiWhite)
		defer color.Unset()
		fmt.Println(message)
	}
}
