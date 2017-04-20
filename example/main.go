package main

import (
	"time"

	"github.com/mbrlabs/clispin"
)

func main() {
	spinner := clispin.New(nil)
	spinner.Start(func() {
		spinner.Printf("Downloading file %d/2", 1)
		time.Sleep(time.Second)
		spinner.Printf("Downloading file %d/2", 2)
		time.Sleep(time.Second)

		spinner.Printf("Processing files")
		time.Sleep(time.Second)
		spinner.Print("Done")
	})
}
