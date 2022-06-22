package weatherData

import (
	"fmt"
	"observer/core"
	s "observer/subjects"
)

type StatisticsDisplay interface {
	core.Observer
	core.Display
}

type statisticsDisplay struct {
	Subject     s.WeatherData
	maxTemp     float32
	minTemp     float32
	tempSum     float32
	numReadings int
}

func NewStatisticsDisplay(subject s.WeatherData) StatisticsDisplay {
	return &statisticsDisplay{
		Subject:     subject,
		maxTemp:     subject.GetTemperature(),
		minTemp:     subject.GetTemperature(),
		tempSum:     subject.GetTemperature(),
		numReadings: 1,
	}
}

func (sd *statisticsDisplay) Update() {
	temperature := sd.Subject.GetTemperature()
	sd.tempSum += temperature
	sd.numReadings++

	if temperature > sd.maxTemp {
		sd.maxTemp = temperature
	}

	if temperature < sd.minTemp {
		sd.minTemp = temperature
	}

	sd.Display()
}

func (sd *statisticsDisplay) Display() {
	avg := sd.tempSum / float32(sd.numReadings)
	fmt.Printf("Avg/Max/Min temperature = %f/%f/%f\n", avg, sd.maxTemp, sd.minTemp)
}
