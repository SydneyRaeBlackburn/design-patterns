package weatherData

import (
	"fmt"
	"observer/core"
	s "observer/subjects"
)

type ForecastDisplay interface {
	core.Observer
	core.Display
}

type forecastDisplay struct {
	Subject         s.WeatherData
	currentPressure float32
	lastPressure    float32
}

func NewForecastDisplay(subject s.WeatherData) ForecastDisplay {
	return &forecastDisplay{
		Subject:         subject,
		currentPressure: subject.GetPressure(),
		lastPressure:    subject.GetPressure(),
	}
}

func (fd *forecastDisplay) Update() {
	fd.lastPressure = fd.currentPressure
	fd.currentPressure = fd.Subject.GetPressure()

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
