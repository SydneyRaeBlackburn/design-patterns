package subject

import (
	"math/rand"
	o "observer/observer"
)

type WeatherData interface {
	Subject
	MeasurementsChanged()
	SetMeasurements()
}

type weatherData struct {
	Subject  *subject
	Temp     float32
	Humidity float32
	Pressure float32
}

func NewWeatherData() WeatherData {
	return &weatherData{
		Subject:  NewSubject(),
		Temp:     getTemperature(),
		Humidity: getHumidity(),
		Pressure: getPressure(),
	}
}

func (wd *weatherData) MeasurementsChanged() {
	wd.NotifyObservers()
}

func (wd *weatherData) SetMeasurements() {
	wd.Temp = getTemperature()
	wd.Humidity = getHumidity()
	wd.Pressure = getPressure()

	wd.MeasurementsChanged()
}

func getTemperature() float32 {
	return float32(rand.Intn(50)+50) + rand.Float32()
}

func getHumidity() float32 {
	return rand.Float32()
}

func getPressure() float32 {
	return float32(rand.Intn(25)+25) + rand.Float32()
}

// adhere to interface but call the root methods

func (wd *weatherData) RegisterObserver(ob o.Observer) {
	wd.Subject.RegisterObserver(ob)
}

func (wd *weatherData) RemoveObserver(ob o.Observer) {
	wd.Subject.RemoveObserver(ob)
}

func (wd *weatherData) NotifyObservers() {
	observers := getObservers(wd.Subject)
	for _, observer := range observers {
		observer.Update(wd.Temp, wd.Humidity, wd.Pressure)
	}
}
