package main

import (
	"fmt"
	"parking/view"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	w := app.NewWindow("Parking Car Go")
	w.CenterOnScreen()
	view.NewFirstView(w)
	fmt.Println("Program is okay")
	w.ShowAndRun()
}
