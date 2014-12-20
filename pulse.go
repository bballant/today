package main

import (
	"net/url"
	"strconv"
	"time"
)

type Pulse struct {
	Email     string
	Year      int
	Month     time.Month
	Day       int
	Weight    int
	Happiness int
	Drink     int
	Eat       int
	Exercise  int
}

func NewPulse(userEmail string, values url.Values) *Pulse {

	year, month, day := time.Now().Date()

	weight, err := strconv.Atoi(values.Get("weight"))
	if err != nil {
		weight = -1
	}

	happiness, err := strconv.Atoi(values.Get("happiness"))
	if err != nil {
		happiness = -1
	}

	drink, err := strconv.Atoi(values.Get("drink"))
	if err != nil {
		drink = -1
	}

	eat, err := strconv.Atoi(values.Get("eat"))
	if err != nil {
		eat = -1
	}

	exercise, err := strconv.Atoi(values.Get("exercise"))
	if err != nil {
		exercise = -1
	}

	return &Pulse{
		Email:     userEmail,
		Year:      year,
		Month:     month,
		Day:       day,
		Weight:    weight,
		Happiness: happiness,
		Drink:     drink,
		Eat:       eat,
		Exercise:  exercise,
	}
}
