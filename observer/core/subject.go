package core

import (
	"fmt"
	"reflect"
)

type Subject interface {
	RegisterObserver(ob Observer)
	RemoveObserver(ob Observer)
	NotifyObservers()
	GetObservers() []Observer
}

func RegisterObserver(obs []Observer, ob Observer) []Observer {
	obs = append(obs, ob)
	fmt.Printf("Observer %s registered.\n", reflect.TypeOf(ob))
	return obs
}

func RemoveObserver(obs []Observer, ob Observer) []Observer {
	for idx, observer := range obs {
		if observer == ob {
			obs[idx] = obs[len(obs)-1] // Copy last element to index idx.
			obs = obs[:len(obs)-1]     // Erase last element (order doesn't matter)
			fmt.Printf("Observer %s removed.\n", reflect.TypeOf(ob))
			break
		}
	}

	return obs
}

func NotifyObservers(obs []Observer) {
	for _, observer := range obs {
		observer.Update()
	}
}
