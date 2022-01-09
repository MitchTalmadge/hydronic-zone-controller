package main

import (
	"log"
	"net/http"

	"github.com/mitchtalmadge/hydronic-zone-controller/src/gpio"
	"github.com/mitchtalmadge/hydronic-zone-controller/src/zone"
)

var zoneActionChan = make(chan zone.ZoneAction)

func main() {
	gpio.Init()
	go zone.HandleZoneActions(zoneActionChan)

	http.HandleFunc("/api/zones/open", zoneOpenHandler)
	http.HandleFunc("/api/zones/openAll", zoneOpenAllHandler)
	http.HandleFunc("/api/zones/close", zoneCloseHandler)
	http.HandleFunc("/api/zones/closeAll", zoneCloseAllHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
