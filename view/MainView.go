package view

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type MainView struct {
	window fyne.Window
}

func NewMainView(window fyne.Window) *MainView {
	MainView := &MainView{
		window: window,
	}
	MainView.InitApp()
	return MainView
}

func (m *MainView) InitApp() {
	m.DrawSceneMenu()
}

func (m *MainView) DrawSceneMenu() {
	title := canvas.NewText("PARKING CAR SIMULATOR", color.RGBA{R: 200, G: 200, B: 200, A: 200})
	title.Resize(fyne.NewSize(20, 20))
	titleContainer := container.NewCenter(title)

	start := widget.NewButton("Start Emulate Game", m.StartParkingSimulation)
	exit := widget.NewButton("Exit", m.ExitGame)

	container_center := container.NewVBox(
		titleContainer,
		layout.NewSpacer(),
		start,
		exit,
		layout.NewSpacer(),
	)

	m.window.SetContent(container_center)
	m.window.Resize(fyne.NewSize(600, 600))
	m.window.SetFixedSize(true)
}

func (m *MainView) ExitGame() {
	m.window.Close()
}

func (m *MainView) StartParkingSimulation() {
	NewParkingView(m.window)
}
