package main

import (
	"fmt"
	"parking/view"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("ParkingCarGo")
	window.CenterOnScreen()
	view.NewMainView(window)
	fmt.Println("Program is okay")
	window.ShowAndRun()
}
