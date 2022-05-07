package types

import (
	"errors"
	"fmt"
	"time"
)

type Birthday struct {
	birthdayTime time.Time
}

func NewBirthday(year, month, day int) (Birthday, error) {
	if year < 0 {
		return Birthday{}, errors.New("invalid year")
	}
	t := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
	if t.Year() != year || int(t.Month()) != month || t.Day() != day {
		return Birthday{}, errors.New("invalid birthday")
	}
	fmt.Println(t.Year())
	return Birthday{birthdayTime: t}, nil
}

func (b Birthday) Year() int {
	return b.birthdayTime.Year()
}

func (b Birthday) Month() int {
	return int(b.birthdayTime.Month())
}

func (b Birthday) Day() int {
	return b.birthdayTime.Day()
}
