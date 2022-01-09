package zone

import (
	"github.com/mitchtalmadge/hydronic-zone-controller/src/gpio"
)

const ZoneCount = gpio.RelayCount - 1 // Reserve 1 relay for boiler
const boilerRelayChannel = 1

var zoneOpen = [ZoneCount]bool{}

func openZone(zone int) {
	zoneOpen[zone-1] = true
	commit()
}

func openAllZones() {
	for i := range zoneOpen {
		zoneOpen[i] = true
	}
	commit()
}

func closeZone(zone int) {
	zoneOpen[zone-1] = false
	commit()
}

func closeAllZones() {
	for i := range zoneOpen {
		zoneOpen[i] = false
	}
	commit()
}

// Updates GPIO to match desired zone state
func commit() {
	anyZonesOpen := false
	for _, open := range zoneOpen {
		if open {
			anyZonesOpen = true
			break
		}
	}

	// Shut off boiler before closing any zones.
	if !anyZonesOpen {
		gpio.CloseRelay(boilerRelayChannel)
	}

	// Set zones as desired.
	for i, open := range zoneOpen {
		if open {
			// The zone motors are wired such that a closed relay = an open zone.
			gpio.CloseRelay(i + 2) // Start at relay channel 2
		} else {
			gpio.OpenRelay(i + 2)
		}
	}

	// Turn on boiler after opening any zones.
	if anyZonesOpen {
		gpio.OpenRelay(boilerRelayChannel)
	}
}
