package gpio

import "github.com/stianeikeland/go-rpio/v4"

const RelayCount = 8

type RelayChannel int

const (
	relayPin1 rpio.Pin = 5
	relayPin2 rpio.Pin = 6
	relayPin3 rpio.Pin = 13
	relayPin4 rpio.Pin = 16
	relayPin5 rpio.Pin = 19
	relayPin6 rpio.Pin = 20
	relayPin7 rpio.Pin = 21
	relayPin8 rpio.Pin = 26
)

var RelayPins = []rpio.Pin{relayPin1, relayPin2, relayPin3, relayPin4, relayPin5, relayPin6, relayPin7, relayPin8}

func initRelays() {
	for _, pin := range RelayPins {
		rpio.Pin(pin).Output()
	}
}

func OpenRelay(channel int) {
	if channel < 1 || channel > RelayCount {
		panic("OpenRelay: Invalid relay channel")
	}
	rpio.Pin(RelayPins[channel-1]).Low()
}

func CloseRelay(channel int) {
	if channel < 1 || channel > RelayCount {
		panic("CloseRelay: Invalid relay channel")
	}
	rpio.Pin(RelayPins[channel-1]).High()
}
