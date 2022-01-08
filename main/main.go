package main

import "github.com/stianeikeland/go-rpio/v4"

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	pin := rpio.Pin(29)
	pin.Output()
	pin.High()
}
