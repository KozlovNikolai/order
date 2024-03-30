package timecalc

import "time"

type TimeCalc struct {
}

func New() *TimeCalc {
	return &TimeCalc{}
}

func (s *TimeCalc) HoursLeft() int64 {
	d := time.Date(2024, time.April, 1, 0, 0, 0, 0, time.UTC)
	dur := time.Until(d)
	return int64(dur.Hours())
}
