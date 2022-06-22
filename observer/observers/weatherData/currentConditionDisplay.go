package weatherData

import (
	"fmt"
	"observer/core"
	s "observer/subjects"
)

type CurrentConditionDisplay interface {
	core.Observer
	core.Display
}

type currentConditionDisplay struct {
	Subject     s.WeatherData
	temperature float32
	humidity    float32
	pressure    float32
}

func NewCurrentConditionDisplay(subject s.WeatherData) CurrentConditionDisplay {
	return &currentConditionDisplay{
		Subject:     subject,
		temperature: subject.GetTemperature(),
		humidity:    subject.GetHumidity(),
		pressure:    subject.GetPressure(),
	}
}

func (ccd *currentConditionDisplay) Update() {
	ccd.temperature = ccd.Subject.GetTemperature()
	ccd.humidity = ccd.Subject.GetHumidity()
	ccd.pressure = ccd.Subject.GetPressure()

	ccd.Display()
}

func (ccd *currentConditionDisplay) Display() {
	fmt.Printf("Current condtions: %f degrees fahrenheit and %f%% humidity\n", ccd.temperature, ccd.humidity)
}
