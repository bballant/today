package db

import (
	"time"

	"model"
)

type PulseDao interface {
	Get(Email string, year int, month time.Month, day int) *model.Pulse
	List(Email string, year int, month time.Month, day int, count int) []model.Pulse
	Delete(Email string, year int, month time.Month, day int)
	Save(pulse *model.Pulse)
	Update(pulse *model.Pulse)
}
