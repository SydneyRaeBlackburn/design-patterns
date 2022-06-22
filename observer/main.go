package main

import (
	"fmt"
	"observer/observers/weatherData"
	s "observer/subjects"
)

func main() {
	weatherDataSubject := s.NewWeatherData()
	currentConditionDisplayObserver := weatherData.NewCurrentConditionDisplay(weatherDataSubject)
	forecastDisplay := weatherData.NewForecastDisplay(weatherDataSubject)
	statisticsDisplay := weatherData.NewStatisticsDisplay(weatherDataSubject)
	heatIndexDisplay := weatherData.NewHeatIndexDisplay(weatherDataSubject)

	weatherDataSubject.RegisterObserver(currentConditionDisplayObserver)
	weatherDataSubject.RegisterObserver(forecastDisplay)
	weatherDataSubject.RegisterObserver(statisticsDisplay)
	weatherDataSubject.RegisterObserver(heatIndexDisplay)

	fmt.Println()

	weatherDataSubject.NotifyObservers()

	fmt.Println()

	weatherDataSubject.SetMeasurements()

	fmt.Println()

	weatherDataSubject.RemoveObserver(heatIndexDisplay)
}
