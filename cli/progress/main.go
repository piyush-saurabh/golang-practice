// go get github.com/cheggaaa/pb/v3

// Start this progress bar as a goroutine and pass the "channel" to track progress
package main

import (
	"time"

	"github.com/cheggaaa/pb"
)

func main() {

	// Total unit of work to be done
	counter := 1000

	// Initialize the progress bar
	progressBar := pb.StartNew(counter)

	// Increase the bar based on the amount of work done
	// This should loop till the finish of the counter
	for i := 0; i < counter; i++ {
		progressBar.Increment()
		time.Sleep(time.Millisecond)
	}

	// Finish the progress
	progressBar.Finish()
}
