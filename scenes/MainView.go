package scenes

import (
	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type FirstView struct {
	window fyne.Window
}

func NewFirstView(window fyne.Window) *FirstView {
	FirstView := &FirstView{
		window: window,
	}
	FirstView.InitApp()
	return FirstView
}

func (m *FirstView) InitApp() {
	m.DrawSceneMenu()
}

func (m *FirstView) DrawSceneMenu() {
	title := canvas.NewText("PARKING CAR SIMULATOR", color.RGBA{R: 200, G: 200, B: 200, A: 200})
	title.TextStyle = fyne.TextStyle{Bold: true, Italic: false}
	title.Resize(fyne.NewSize(60, 60))
	
	titleContainer := container.New(layout.NewVBoxLayout(), layout.NewSpacer(), container.NewCenter(title), layout.NewSpacer())
	exit := widget.NewButton("Exit to the program", m.ExitProgram)
	start := widget.NewButton("Start to emulate", m.StartEmulation)

	buttonsContainer := container.NewHBox(layout.NewSpacer(), exit, layout.NewSpacer(), start, layout.NewSpacer())
	container_center := container.NewVBox(
		layout.NewSpacer(),
		titleContainer,
		layout.NewSpacer(),
		buttonsContainer,
		layout.NewSpacer(),
	)

	m.window.SetContent(container_center)
	m.window.Resize(fyne.NewSize(600, 600))
	m.window.SetFixedSize(true)
}

func (m *FirstView) ExitProgram() {
	m.window.Close()
}

func (m *FirstView) StartEmulation() {
	NewParkingView(m.window)
}
