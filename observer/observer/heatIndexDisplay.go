package observer

import "fmt"

type HeatIndexDisplay interface {
	Observer
	DisplayElement
}

type heatIndexDisplay struct {
	Observer       Observer
	DisplayElement DisplayElement
	heatIndex      float32
}

func NewHeatIndexDisplay(heatIndex float32) HeatIndexDisplay {
	return &heatIndexDisplay{
		Observer:       NewObserver(),
		DisplayElement: NewDisplayElement(),
		heatIndex:      heatIndex,
	}
}

func (sd *heatIndexDisplay) Update(temperature float32, humidity float32, _ float32) {
	sd.heatIndex = computeHeatIndex(temperature, humidity)

	sd.Display()
}

func (sd *heatIndexDisplay) Display() {
	fmt.Printf("Heat index is %f\n", sd.heatIndex)
}

func computeHeatIndex(t float32, h float32) float32 {
	index := 16.923 + (0.185212 * t) + (5.37941 * h) - (0.100254 * t * h) + (0.00941695 * (t * t)) + (0.00728898 * (h * h)) + (0.000345372 * (t * t * h)) - (0.000814971 * (t * h * h)) + (0.0000102102 * (t * t * h * h)) - (0.000038646 * (t * t * t)) + (0.0000291583 * (h * h * h)) + (0.00000142721 * (t * t * t * h)) + (0.000000197483 * (t * h * h * h)) - (0.0000000218429 * (t * t * t * h * h)) + 0.000000000843296*(t*t*h*h*h) - (0.0000000000481975 * (t * t * t * h * h * h))
	return index
}
