package main

import "time"

type Event struct {
	Email     string
	Date      time.Time
	Weight    int
	Happiness int
	Drink     int
	Eat       int
	Exercize  int
}
