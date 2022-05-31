package main

import (
	"fmt"
)

func main() {
	data := NewWeatherData()
	fmt.Println("Temp: ", data.Temp)
	fmt.Println("Humidity: ", data.Humidity)
	fmt.Println("Pressure: ", data.Pressure)
	fmt.Println()

	data.MeasurementsChanged()
	fmt.Println("Temp: ", data.Temp)
	fmt.Println("Humidity: ", data.Humidity)
	fmt.Println("Pressure: ", data.Pressure)
	fmt.Println()
}
