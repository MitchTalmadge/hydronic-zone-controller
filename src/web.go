package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/mitchtalmadge/hydronic-zone-controller/src/zone"
)

func zoneOpenHandler(w http.ResponseWriter, r *http.Request) {
	zoneNum, err := strconv.Atoi(r.FormValue("zone"))
	if err != nil {
		fmt.Fprintf(w, "Cannot parse zone")
		return
	}
	if zoneNum < 1 || zoneNum > zone.ZoneCount {
		fmt.Fprintf(w, "Invalid zone")
		return
	}

	zoneActionChan <- zone.ZoneAction{ActionType: zone.ZoneActionOpenOne, Zone: zoneNum}
}

func zoneOpenAllHandler(w http.ResponseWriter, r *http.Request) {
	zoneActionChan <- zone.ZoneAction{ActionType: zone.ZoneActionOpenAll, Zone: 0}
}

func zoneCloseHandler(w http.ResponseWriter, r *http.Request) {
	zoneNum, err := strconv.Atoi(r.FormValue("zone"))
	if err != nil {
		fmt.Fprintf(w, "Cannot parse zone")
		return
	}
	if zoneNum < 1 || zoneNum > zone.ZoneCount {
		fmt.Fprintf(w, "Invalid zone")
		return
	}

	zoneActionChan <- zone.ZoneAction{ActionType: zone.ZoneActionCloseOne, Zone: zoneNum}
}

func zoneCloseAllHandler(w http.ResponseWriter, r *http.Request) {
	zoneActionChan <- zone.ZoneAction{ActionType: zone.ZoneActionCloseAll, Zone: 0}
}
