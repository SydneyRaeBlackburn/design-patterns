package subjects

import (
	"math/rand"
	"observer/core"
)

type WeatherData interface {
	core.Subject
	MeasurementsChanged()
	SetMeasurements()
	GetTemperature() float32
	GetHumidity() float32
	GetPressure() float32
}

type weatherData struct {
	Observers []core.Observer
	Temp      float32
	Humidity  float32
	Pressure  float32
}

func NewWeatherData() WeatherData {
	return &weatherData{
		Temp:     setTemperature(),
		Humidity: setHumidity(),
		Pressure: setPressure(),
	}
}

// adhere to Subject interface but call the root methods

func (wd *weatherData) RegisterObserver(ob core.Observer) {
	// use root register from core.Subject interface
	wd.Observers = core.RegisterObserver(wd.Observers, ob)
}

func (wd *weatherData) RemoveObserver(ob core.Observer) {
	// use root remover from core.Subject interface
	wd.Observers = core.RemoveObserver(wd.Observers, ob)
}

func (wd *weatherData) NotifyObservers() {
	// use root notifier from core.Subject interface
	core.NotifyObservers(wd.Observers)
}

func (wd *weatherData) GetObservers() []core.Observer {
	return wd.Observers
}

// weatherData interface

func (wd *weatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

func (wd *weatherData) SetMeasurements() {
	wd.Temp = setTemperature()
	wd.Humidity = setHumidity()
	wd.Pressure = setPressure()

	wd.MeasurementsChanged()
}

func (wd *weatherData) GetTemperature() float32 {
	return wd.Temp
}

func (wd *weatherData) GetHumidity() float32 {
	return wd.Humidity
}

func (wd *weatherData) GetPressure() float32 {
	return wd.Pressure
}

// setters

func setTemperature() float32 {
	return float32(rand.Intn(50)+50) + rand.Float32()
}

func setHumidity() float32 {
	return rand.Float32()
}

func setPressure() float32 {
	return float32(rand.Intn(25)+25) + rand.Float32()
}
