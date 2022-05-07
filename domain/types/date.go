package types

import (
	"errors"
	"fmt"
	"time"
)

type Date struct {
	birthdayTime time.Time
}

func NewDate(year, month, day int) (Date, error) {
	if year < 0 {
		return Date{}, errors.New("invalid year")
	}
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if t.Year() != year || int(t.Month()) != month || t.Day() != day {
		return Date{}, errors.New("invalid birthday")
	}
	fmt.Println(t.Year())
	return Date{birthdayTime: t}, nil
}

func (b Date) Year() int {
	return b.birthdayTime.Year()
}

func (b Date) Month() int {
	return int(b.birthdayTime.Month())
}

func (b Date) Day() int {
	return b.birthdayTime.Day()
}
