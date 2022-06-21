package main

import (
	"fmt"
	b "strategy/behavior"
)

// client

func main() {
	// create new generic duck with defaults
	duck := b.NewDuck(nil, nil)

	fmt.Println("Regular Duck:")
	duck.Display()

	fmt.Println()

	// can set the behavior for different kinds of ducks at runtime
	fmt.Println("Rubber Duck (initial):")
	rubberDuck := b.NewDuck(nil, nil)
	rubberDuck.Display()

	fmt.Println()

	var flyBehavior string
	fmt.Println("Set the rubber duck's flying behavior (type none):  ")
	fmt.Scanln(&flyBehavior)
	if flyBehavior == "none" {
		rubberDuck.SetFlyBehavior(&b.FlyNoWay{})
	}
	var quackBehavior string
	fmt.Println("Set the rubber duck's flying behavior (type squeak):  ")
	fmt.Scanln(&quackBehavior)
	if quackBehavior == "squeak" {
		rubberDuck.SetQuackBehvaior(&b.Squeak{})
	}

	fmt.Println()

	fmt.Println("Rubber Duck (updated):")
	rubberDuck.Display()
}
