package utils

import (
	"fmt"
)

func Test(name string, namespace string, output string) {
	fmt.Println("This is the test function for the implementation")

	fmt.Printf("Pod created with name: %s in namespace: %s with output: %s\n", name, namespace, output)
}
