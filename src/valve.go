package main

import (
	"github.com/stianeikeland/go-rpio/v4"
)

/* Valve Definitions */
type Valve rpio.Pin

const (
	valveBoiler Valve = Valve(pinCh1)
	valveZone1  Valve = Valve(pinCh2)
	valveZone2  Valve = Valve(pinCh3)
	valveZone3  Valve = Valve(pinCh4)
	valveZone4  Valve = Valve(pinCh5)
	valveZone5  Valve = Valve(pinCh6)
	valveZone6  Valve = Valve(pinCh7)
	valveZone7  Valve = Valve(pinCh8)
)

var zoneValves = []Valve{valveZone1, valveZone2, valveZone3, valveZone4, valveZone5, valveZone6, valveZone7}

/* Valve States */
type ValveState bool

const (
	valveStateOpen   ValveState = true
	valveStateClosed ValveState = false
)

var valveStates = map[Valve]ValveState{
	valveZone1: valveStateClosed,
	valveZone2: valveStateClosed,
	valveZone3: valveStateClosed,
	valveZone4: valveStateClosed,
	valveZone5: valveStateClosed,
	valveZone6: valveStateClosed,
	valveZone7: valveStateClosed,
}

/* Open Valve to Pin State Mapping */
var openValveToPinState = map[Valve]rpio.State{
	valveBoiler: rpio.Low,
	valveZone1:  rpio.High,
	valveZone2:  rpio.High,
	valveZone3:  rpio.High,
	valveZone4:  rpio.High,
	valveZone5:  rpio.High,
	valveZone6:  rpio.High,
	valveZone7:  rpio.High,
}

/* Valve Actions */
const (
	valveActionOpenAll  = iota
	valveActionCloseAll = iota
	valveActionOpenOne  = iota
	valveActionCloseOne = iota
)

type valveAction struct {
	action int
	valve  Valve
}

func handleValveActions(valveActionChan chan valveAction) {
	closeAllValves()
	for {
		valveAction := <-valveActionChan
		switch valveAction.action {
		case valveActionOpenAll:
			openAllValves()
		case valveActionCloseAll:
			closeAllValves()
		case valveActionOpenOne:
			openValve(valveAction.valve)
		case valveActionCloseOne:
			closeValve(valveAction.valve)
		}
	}
}

func openValve(valve Valve) {
	valveStates[valve] = valveStateOpen
	commitValveState()
}

func closeValve(valve Valve) {
	valveStates[valve] = valveStateClosed
	commitValveState()
}

func openAllValves() {
	for _, valve := range zoneValves {
		valveStates[valve] = valveStateOpen
	}
	commitValveState()
}

func closeAllValves() {
	for _, valve := range zoneValves {
		valveStates[valve] = valveStateClosed
	}
	commitValveState()
}

func commitValveState() {
	anyZonesOpen := false
	for _, valveState := range valveStates {
		if valveState == valveStateOpen {
			anyZonesOpen = true
			break
		}
	}

	// Close boiler before closing any zones.
	if !anyZonesOpen {
		rpio.Pin(valveBoiler).Write((openValveToPinState[valveBoiler] + 1) % 2)
	}

	// Open/close zones as recorded in state.
	for valve, valveState := range valveStates {
		switch valveState {
		case valveStateOpen:
			rpio.Pin(valve).Write(openValveToPinState[valve])
		case valveStateClosed:
			rpio.Pin(valve).Write((openValveToPinState[valve] + 1) % 2)
		}
	}

	// Open boiler after opening any zones.
	if anyZonesOpen {
		rpio.Pin(valveBoiler).Write(openValveToPinState[valveBoiler])
	}
}
