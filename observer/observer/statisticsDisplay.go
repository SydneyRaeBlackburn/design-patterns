package observer

import "fmt"

type StatisticsDisplay interface {
	Observer
	DisplayElement
}

type statisticsDisplay struct {
	Observer       Observer
	DisplayElement DisplayElement
	maxTemp        float32
	minTemp        float32
	tempSum        float32
	numReadings    int
}

func NewStatisticsDisplay(maxTemp float32, minTemp float32, tempSum float32, numReadings int) StatisticsDisplay {
	return &statisticsDisplay{
		Observer:       NewObserver(),
		DisplayElement: NewDisplayElement(),
		maxTemp:        maxTemp,
		minTemp:        minTemp,
		tempSum:        tempSum,
		numReadings:    numReadings,
	}
}

func (sd *statisticsDisplay) Update(temperature float32, _ float32, _ float32) {
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
