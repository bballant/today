package db

import (
	"math/rand"
	"time"

	"model"
)

type StubPulseDao struct {
	PulseDao
}

func (dao *StubPulseDao) Get(
	email string, year int, month time.Month, day int) *model.Pulse {

	return &model.Pulse{
		Email:     email,
		Year:      year,
		Month:     month,
		Day:       day,
		Weight:    (150 + rand.Intn(50)),
		Happiness: rand.Intn(5),
		Drink:     rand.Intn(5),
		Eat:       rand.Intn(5),
		Exercise:  rand.Intn(5),
	}
}
