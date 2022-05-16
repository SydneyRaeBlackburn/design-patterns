package main

import (
	"fmt"
	b "main/behavior"
)

// client

func main() {
	// create new generic duck with defaults
	duck := b.NewDuck(nil, nil)

	fmt.Println("Regular Duck:")
	duck.Display()

	fmt.Println()

	fmt.Println("Rubber Duck:")
	rubberDuck := b.NewDuck(nil, nil)
	rubberDuck.SetFlyBehavior(&b.FlyNoWay{})
	rubberDuck.SetQuackBehvaior(&b.Squeak{})
	rubberDuck.Display()

	fmt.Println()

	fmt.Println("Flying Duck:")

	flyingDuck := b.NewDuck(nil, nil)
	flyingDuck.SetFlyBehavior(&b.FlyWithWings{})
	flyingDuck.Display()

	fmt.Println()

	fmt.Println("Rocket Duck:")
	rocketDuck := b.NewDuck(nil, nil)
	rocketDuck.SetFlyBehavior(&b.FlyWithRocket{})
	rocketDuck.Display()
}
