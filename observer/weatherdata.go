package main

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

func NewWeatherData(temp float32, humidity float32, pressure float32) *WeatherData {
	return &WeatherData{
		Temp:     temp,
		Humidity: humidity,
		Pressure: pressure,
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
	return 73.3
}

func getHumidity() float32 {
	return .61
}

func getPressure() float32 {
	return 29.93
}
