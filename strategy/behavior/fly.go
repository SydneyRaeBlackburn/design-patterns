package behavior

import "fmt"

// concrete strategy

type FlyBehavior interface {
	Fly()
}

// different implementations

type FlyWithWings struct {
	FlyBehavior
}

type FlyNoWay struct {
	FlyBehavior
}

type FlyWithRocket struct {
	FlyBehavior
}

// execution methods

func (fww *FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings!!")
}

func (fnw *FlyNoWay) Fly() {
	fmt.Println("I can't fly")
}

func (fwr *FlyWithRocket) Fly() {
	fmt.Println("I'm flying with a rocket!!")
}
