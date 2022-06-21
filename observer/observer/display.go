package observer

type DisplayElement interface {
	Display()
}

type displayElement struct{}

func NewDisplayElement() DisplayElement {
	return &displayElement{}
}

func (o *displayElement) Display() {}
