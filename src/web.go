package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func zoneOpenHandler(w http.ResponseWriter, r *http.Request) {
	zone, err := strconv.Atoi(r.FormValue("zone"))
	if err != nil {
		fmt.Fprintf(w, "Cannot parse zone")
		return
	}
	if zone < 1 || zone > 7 {
		fmt.Fprintf(w, "Invalid zone")
		return
	}

	valveChan <- valveAction{valveActionOpenOne, zoneValves[zone-1]}
}

func zoneOpenAllHandler(w http.ResponseWriter, r *http.Request) {
	valveChan <- valveAction{valveActionOpenAll, 0}
}

func zoneCloseHandler(w http.ResponseWriter, r *http.Request) {
	zone, err := strconv.Atoi(r.FormValue("zone"))
	if err != nil {
		fmt.Fprintf(w, "Cannot parse zone")
		return
	}
	if zone < 1 || zone > 7 {
		fmt.Fprintf(w, "Invalid zone")
		return
	}

	valveChan <- valveAction{valveActionCloseOne, zoneValves[zone-1]}
}

func zoneCloseAllHandler(w http.ResponseWriter, r *http.Request) {
	valveChan <- valveAction{valveActionCloseAll, 0}
}
