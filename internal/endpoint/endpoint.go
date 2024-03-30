package endpoint

import (
	"fmt"
)

// TimeHandler is interface to timecalc package
type TimeHandler interface {
	HoursLeft() int64
}

// Endpoint is base structure of package
type Endpoint struct {
	timehandler TimeHandler
}

// New is constructor of endpoint object
func New(th TimeHandler) *Endpoint {
	return &Endpoint{
		timehandler: th,
	}
}

// Status show the time in hours before the specified date by interface TimeHandler
func (e *Endpoint) Status() error {
	d := e.timehandler.HoursLeft()
	fmt.Printf("До 1 апреля осталось %d часов\n", d)
	return nil
}
