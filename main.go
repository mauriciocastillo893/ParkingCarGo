package main

import (
	"fmt"
	"parking/view"
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	w := app.NewWindow("ParkingCarGo")
	w.CenterOnScreen()
	view.NewMainView(w)
	fmt.Println("Program is okay")
	w.ShowAndRun()
}
