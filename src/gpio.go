package main

import (
	"github.com/stianeikeland/go-rpio/v4"
)

const (
	pinCh1 rpio.Pin = 5
	pinCh2 rpio.Pin = 6
	pinCh3 rpio.Pin = 13
	pinCh4 rpio.Pin = 16
	pinCh5 rpio.Pin = 19
	pinCh6 rpio.Pin = 20
	pinCh7 rpio.Pin = 21
	pinCh8 rpio.Pin = 26
)

var chPins = []rpio.Pin{pinCh1, pinCh2, pinCh3, pinCh4, pinCh5, pinCh6, pinCh7, pinCh8}

func initGPIO() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	for _, ch := range chPins {
		rpio.Pin(ch).Output()
	}
}
