package observer

import "fmt"

type CurrentConditionDisplay interface {
	Observer
	DisplayElement
}

type currentConditionDisplay struct {
	Observer       Observer
	DisplayElement DisplayElement
	temperature    float32
	humidity       float32
	pressure       float32
}

func NewCurrentConditionDisplay(temperature float32, humidity float32, pressure float32) CurrentConditionDisplay {
	return &currentConditionDisplay{
		Observer:       NewObserver(),
		DisplayElement: NewDisplayElement(),
		temperature:    temperature,
		humidity:       humidity,
		pressure:       pressure,
	}
}

func (ccd *currentConditionDisplay) Update(temperature float32, humidity float32, pressure float32) {
	ccd.temperature = temperature
	ccd.humidity = humidity
	ccd.pressure = pressure

	ccd.Display()
}

func (ccd *currentConditionDisplay) Display() {
	fmt.Printf("Current condtions: %f degrees fahrenheit and %f%% humidity\n", ccd.temperature, ccd.humidity)
}
