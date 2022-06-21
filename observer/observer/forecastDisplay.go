package observer

import "fmt"

type ForecastDisplay interface {
	Observer
	DisplayElement
}

type forecastDisplay struct {
	Observer        Observer
	DisplayElement  DisplayElement
	currentPressure float32
	lastPressure    float32
}

func NewForecastDisplay(currentPressure float32, lastPressure float32) ForecastDisplay {
	return &forecastDisplay{
		Observer:        NewObserver(),
		DisplayElement:  NewDisplayElement(),
		currentPressure: currentPressure,
		lastPressure:    lastPressure,
	}
}

func (fd *forecastDisplay) Update(_ float32, _ float32, pressure float32) {
	fd.lastPressure = fd.currentPressure
	fd.currentPressure = pressure

	fd.Display()
}

func (fd *forecastDisplay) Display() {
	fmt.Print("Forecast: ")
	if fd.currentPressure > fd.lastPressure {
		fmt.Println("Improving weather on the way!")
	} else if fd.currentPressure == fd.lastPressure {
		fmt.Println("More of the same.")
	} else if fd.currentPressure < fd.lastPressure {
		fmt.Println("Watch out for cooler, rainy weather.")
	}
}
