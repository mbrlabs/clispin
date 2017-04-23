package main

import (
	"time"

	"github.com/mbrlabs/clispin"
)

func spinner(success bool) {
	spinner := clispin.New(nil)
	spinner.Color(clispin.ColorCyan)

	spinner.Start(func() {
		spinner.Printf("Downloading file %d/2", 1)
		time.Sleep(time.Second)
		spinner.Printf("Downloading file %d/2", 2)
		time.Sleep(time.Second)

		spinner.Print("Processing files")
		time.Sleep(time.Second)

		if success {
			spinner.LastFrame = "✔"
			spinner.Color(clispin.ColorGreen)
			spinner.Print("Success")
		} else {
			spinner.LastFrame = "✖"
			spinner.Color(clispin.ColorRed)
			spinner.Print("Failure")
		}
	})
}

func main() {
	spinner(true)
	spinner(false)
}
