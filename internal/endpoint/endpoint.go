package endpoint

import (
	"fmt"
)

type TimeHandler interface {
	HoursLeft() int64
}

type Endpoint struct {
	timehandler TimeHandler
}

func New(th TimeHandler) *Endpoint {
	return &Endpoint{
		timehandler: th,
	}
}

func (e *Endpoint) Status() error {
	d := e.timehandler.HoursLeft()
	fmt.Printf("До 1 апреля осталось %d часов\n", d)
	return nil
}
