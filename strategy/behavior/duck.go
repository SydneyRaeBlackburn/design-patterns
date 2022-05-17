package behavior

import "fmt"

// DuckBehavior interface defines what all Ducks have in common (strategy)
type DuckBehavior interface {
	Swim()
	Display()
	FlyBehavior
	QuackBehavior
}

// Define duck (context/class)
type Duck struct {
	DuckBehavior  DuckBehavior
	FlyBehavior   FlyBehavior
	QuackBehavior QuackBehavior
}

// Declare a new duck
func NewDuck(fb FlyBehavior, qb QuackBehavior) *Duck {
	return &Duck{
		FlyBehavior:   fb,
		QuackBehavior: qb,
	}
}

// define top level behavior

func (d *Duck) Display() {
	d.swim()

	if d.FlyBehavior != nil {
		d.FlyBehavior.Fly()
	} else {
		d.fly()
	}

	if d.QuackBehavior != nil {
		d.QuackBehavior.Quack()
	} else {
		d.quack()
	}
}

func (d *Duck) swim() {
	fmt.Println("I'm swimming!!")
}

func (d *Duck) fly() {
	fmt.Println("I'm flying!!")
}

func (d *Duck) quack() {
	fmt.Println("Quack!!")
}

// setter methods

func (d *Duck) SetFlyBehavior(fb FlyBehavior) *Duck {
	d.FlyBehavior = fb

	return d
}

func (d *Duck) SetQuackBehvaior(qb QuackBehavior) *Duck {
	d.QuackBehavior = qb

	return d
}
