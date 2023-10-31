package scenes

import (
	"parking/models"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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

	// ALL THE STATION
	cParkView.Add(p.MakeEnterAndExitStation())

	// SQUARE OF EXIT
	cParkOut.Add(p.MakeExitStation())
	cParkOut.Add(layout.NewSpacer())
	
	// SQUEARE OF WAIT PLACE
	cParkOut.Add(p.MakeWaitStation())

	// VIEW OF THE CARS WHEN ENTRIES AND OUTS + SPACE LINE
	cParkView.Add(cParkOut)
	cParkView.Add(layout.NewSpacer())

	// SPACE BETWEEN PARKING AND ENTRIE-EXIT
	cParkView.Add(layout.NewSpacer())

	cParkView.Add(p.MakeParking())
	// BUTTON SPACE
	cParkView.Add(layout.NewSpacer())

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

func addSpace(parkingContainer *fyne.Container) {
	for j := 0; j < 5; j++ {
		parkingContainer.Add(layout.NewSpacer())
	}
}
