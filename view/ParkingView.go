package view

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	_"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	_"fyne.io/fyne/v2/widget"
	"image/color"
	"parking/models"
	"time"
)

var semRenderNewCarWait chan bool
var semQuit chan bool
var parking *models.Parking

type ParkingView struct {
	window               fyne.Window
	waitRectangleStation [models.MaxWait]*canvas.Rectangle
}

func NewParkingView(window fyne.Window) *ParkingView {
	parkingView := &ParkingView{window: window}
	semQuit = make(chan bool)
	semRenderNewCarWait = make(chan bool)
	parking = models.NewParking(semRenderNewCarWait, semQuit)
	parkingView.MakeScene()
	parkingView.StartEmulator()
	return parkingView
}

func (p *ParkingView) MakeScene() {
	cParkView := container.New(layout.NewVBoxLayout())
	cParkOut := container.New(layout.NewHBoxLayout())
	cButttons := container.New(layout.NewHBoxLayout())
	// restart := widget.NewButton("Restart Simulation", func() {
	// 	dialog.ShowConfirm("Restart", "Do you want to restart the app?", func(response bool) {
	// 		if response {
	// 			p.RestarEmulator()
	// 		}
	// 	}, p.window)
	// })

	// exit := widget.NewButton("Menu Exit",
	// 	func() {
	// 		dialog.ShowConfirm("Exit", "Do you want to go out from here?", func(response bool) {
	// 			if response {
	// 				p.BackToMenu()
	// 			}
	// 		}, p.window)
	// 	},
	// )

	// cButttons.Add(restart)
	// cButttons.Add(exit)
	cParkOut.Add(p.MakeWaitStation())
	cParkOut.Add(layout.NewSpacer())
	cParkOut.Add(p.MakeExitStation())
	cParkOut.Add(layout.NewSpacer())
	cParkView.Add(cParkOut)
	cParkView.Add(layout.NewSpacer())
	cParkView.Add(p.MakeParkingLotEntrance())
	cParkView.Add(layout.NewSpacer())
	cParkView.Add(p.MakeEnterAndExitStation())
	cParkView.Add(layout.NewSpacer())
	cParkView.Add(p.MakeParking())
	cParkView.Add(layout.NewSpacer())

	cParkView.Add(container.NewCenter(cButttons))
	p.window.SetContent(cParkView)
	p.window.Resize(fyne.NewSize(600, 600))
	p.window.CenterOnScreen()
}

func (p *ParkingView) MakeParking() *fyne.Container {
	parkingContainer := container.New(layout.NewGridLayout(5))
	parking.MakeParking()

	parkingArray := parking.GetParking()
	for i := 0; i < len(parkingArray); i++ {
		if i == 10 {
			addSpace(parkingContainer)
		}
		parkingContainer.Add(container.NewCenter(parkingArray[i].GetRectangle(), parkingArray[i].GetText()))
	}
	return container.NewCenter(parkingContainer)
}

func (p *ParkingView) MakeWaitStation() *fyne.Container {
	parkingContainer := container.New(layout.NewGridLayout(5))
	for i := len(p.waitRectangleStation) - 1; i >= 0; i-- {
		car := models.NewSpaceCar().GetRectangle()
		p.waitRectangleStation[i] = car
		p.waitRectangleStation[i].Hide()
		parkingContainer.Add(p.waitRectangleStation[i])
	}
	return parkingContainer
}

func (p *ParkingView) MakeExitStation() *fyne.Container {
	out := parking.MakeOutStation()
	return container.NewCenter(out.GetRectangle())
}

func (p *ParkingView) MakeEnterAndExitStation() *fyne.Container {
	parkingContainer := container.New(layout.NewGridLayout(5))
	parkingContainer.Add(layout.NewSpacer())
	entrace := parking.MakeEntraceStation()
	parkingContainer.Add(entrace.GetRectangle())
	parkingContainer.Add(layout.NewSpacer())
	exit := parking.MakeExitStation()
	parkingContainer.Add(exit.GetRectangle())
	parkingContainer.Add(layout.NewSpacer())
	return container.NewCenter(parkingContainer)
}

func (p *ParkingView) MakeParkingLotEntrance() *fyne.Container {
	EntraceContainer := container.New(layout.NewGridLayout(3))
	EntraceContainer.Add(makeBorder())
	EntraceContainer.Add(layout.NewSpacer())
	EntraceContainer.Add(makeBorder())
	return EntraceContainer
}

func (p *ParkingView) RenderNewCarWaitStation() {
	for {
		select {
		case <-semQuit:
			return
		case <-semRenderNewCarWait:
			waitCars := parking.GetWaitCars()
			for i := len(waitCars) - 1; i >= 0; i-- {
				if waitCars[i].ID != -1 {
					p.waitRectangleStation[i].Show()
					p.waitRectangleStation[i].FillColor = waitCars[i].GetRectangle().FillColor
				}
			}
			p.window.Content().Refresh()
		}
	}
}

func (p *ParkingView) RenderUpdate() {
	for {
		select {
		case <-semQuit:
			return
		default:
			p.window.Content().Refresh()
			time.Sleep(1 * time.Second)
		}
	}
}

func (p *ParkingView) StartEmulator() {
	go parking.GenerateCars()
	go parking.GoOutCars()
	go parking.CheckParking()
	go p.RenderNewCarWaitStation()
	go p.RenderUpdate()

}

func (p *ParkingView) BackToMenu() {
	close(semQuit)
	NewFirstView(p.window)
}

func (p *ParkingView) RestarEmulator() {
	close(semQuit)
	NewParkingView(p.window)
}

func addSpace(parkingContainer *fyne.Container) {
	for j := 0; j < 5; j++ {
		parkingContainer.Add(layout.NewSpacer())
	}
}

func makeBorder() *canvas.Rectangle {
	square := canvas.NewRectangle(color.RGBA{R: 255, G: 255, B: 255, A: 0})
	square.SetMinSize(fyne.NewSquareSize(float32(30)))
	square.StrokeColor = color.RGBA{R: 255, G: 255, B: 255, A: 255}
	square.StrokeWidth = float32(1)
	return square
}
