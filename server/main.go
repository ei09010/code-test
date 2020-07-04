package main

import (
	"code-test/server/services/event_service"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/screenresize", event_service.HandleScreenResizeEvents)
	http.HandleFunc("/timetaken", event_service.HandleTimeTakenEvents)
	http.HandleFunc("/copypaste", event_service.HandleCopyPasteEvents)
	http.HandleFunc("/session", event_service.HandleSessionCreation)

	serverPort := ":8080"

	log.Fatal(http.ListenAndServe(serverPort, nil))
	log.Println("Server startd on: ", serverPort)
}
