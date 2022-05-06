package model

import (
	"errors"
	"fmt"
	"time"
)

type UserName struct {
	first string
	last  string
}

func NewName(first, last string) (UserName, error) {
	if len(first) == 0 || len(last) == 0 {
		return UserName{}, errors.New("invalid name")
	}
	return UserName{first, last}, nil
}

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
