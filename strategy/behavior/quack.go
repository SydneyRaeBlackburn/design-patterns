package behavior

import "fmt"

// concrete strategy

type QuackBehavior interface {
	Quack()
}

// different implementations

type MuteDuck struct {
	QuackBehavior
}

type Squeak struct {
	QuackBehavior
}

// execution methods

func (md *MuteDuck) Quack() {
	fmt.Println("<< Silence >>")
}

func (s *Squeak) Quack() {
	fmt.Println("Squeak!!")
}
