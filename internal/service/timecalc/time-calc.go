package timecalc

import "time"

// TimeCalc is base structure of package
type TimeCalc struct {
}

// New is constructor of timecalc object
func New() *TimeCalc {
	return &TimeCalc{}
}

// HoursLeft calculates the time in hours before the specified date.
func (s *TimeCalc) HoursLeft() int64 {
	d := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	return int64(dur.Hours())
}
