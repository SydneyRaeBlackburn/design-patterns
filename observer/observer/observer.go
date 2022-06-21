package observer

type Observer interface {
	Update(temp float32, humidity float32, pressure float32)
}

type observer struct{}

func NewObserver() Observer {
	return &observer{}
}

// How to make Update() generic

func (o *observer) Update(_ float32, _ float32, _ float32) {}
