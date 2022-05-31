package main

import "math/rand"

type Weatherdata interface {
	MeasurementsChanged()
	getTemperature()
	getHumidity()
	getPressure()
}

type WeatherData struct {
	Temp     float32
	Humidity float32
	Pressure float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		Temp:     getTemperature(),
		Humidity: getHumidity(),
		Pressure: getPressure(),
	}
}

func (wd *WeatherData) MeasurementsChanged() *WeatherData {
	wd.Temp = getTemperature()
	wd.Humidity = getHumidity()
	wd.Pressure = getPressure()

	// update current conditions display
	// update statistics display
	// update forecast display

	return wd
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
