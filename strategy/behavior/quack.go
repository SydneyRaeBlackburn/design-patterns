package behavior

import "fmt"

// concrete strategy

type QuackBehavior interface {
	Quack()
}

type MuteDuck struct {
	QuackBehavior
}

type Squeak struct {
	QuackBehavior
}

func (md *MuteDuck) Quack() {
	fmt.Println("<< Silence >>")
}

func (s *Squeak) Quack() {
	fmt.Println("Squeak!!")
}
