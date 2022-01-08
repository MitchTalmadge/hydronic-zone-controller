package main

import (
	"log"
	"net/http"
)

var valveChan = make(chan valveAction)

func main() {
	initGPIO()
	go handleValveActions(valveChan)

	http.HandleFunc("/api/zones/open", zoneOpenHandler)
	http.HandleFunc("/api/zones/openAll", zoneOpenAllHandler)
	http.HandleFunc("/api/zones/close", zoneCloseHandler)
	http.HandleFunc("/api/zones/closeAll", zoneCloseAllHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
