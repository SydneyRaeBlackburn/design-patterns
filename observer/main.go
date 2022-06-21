package main

import (
	"fmt"
	o "observer/observer"
	s "observer/subject"
)

func main() {
	weatherDataSubject := s.NewWeatherData()
	currentConditionDisplayObserver := o.NewCurrentConditionDisplay(0.0, 0.0, 0.0)
	forecastDisplay := o.NewForecastDisplay(0.0, 0.0)
	statisticsDisplay := o.NewStatisticsDisplay(0.0, 0.0, 0.0, 0)
	heatIndexDisplay := o.NewHeatIndexDisplay(0.0)

	weatherDataSubject.RegisterObserver(currentConditionDisplayObserver)
	weatherDataSubject.RegisterObserver(forecastDisplay)
	weatherDataSubject.RegisterObserver(statisticsDisplay)
	weatherDataSubject.RegisterObserver(heatIndexDisplay)

	fmt.Println()

	weatherDataSubject.SetMeasurements()

	fmt.Println()

	weatherDataSubject.RemoveObserver(heatIndexDisplay)
	weatherDataSubject.SetMeasurements()
}
