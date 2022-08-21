package main

import (
	"errors"
	"fmt"
	"time"
)

type DateRange struct {
	Start time.Time
	End   time.Time
}

func NewDateRange(start time.Time, end time.Time) (*DateRange, error) {

	if start.IsZero() || end.IsZero() {
		return nil, errors.New("start or end can not be empty !")
	}

	if end.Before(start) {
		return nil, errors.New("end can not be before from start !")
	}

	date := DateRange{
		Start: start,
		End:   end,
	}

	return &date, nil
}

func (d DateRange) Hours() float64 {
	return d.End.Sub(d.Start).Hours()
}

func main() {

	start := time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC)
	end := time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC)

	lifetime, _ := NewDateRange(start, end)

	fmt.Println(lifetime.Hours())

	travelInTime := DateRange{
		Start: time.Date(1852, 11, 27, 0, 0, 0, 0, time.UTC),
		End:   time.Date(1815, 12, 10, 0, 0, 0, 0, time.UTC),
	}

	fmt.Println(travelInTime.Hours())
}
