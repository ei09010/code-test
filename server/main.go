package main

import (
	"code-test/server/repository"
	"code-test/server/services/event_service"
	"code-test/server/services/hash_service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	invalidMethodReceived = "Received http method is not the expected one"
	unableToReadBody      = "Unable to read body"
	unableToUnmarshall    = "Unable to unMarshall request body"
	errorUpdatingData     = "Error updating data in the repository"
	invalidObject         = "The received object is invalid"
	errorValidatingObject = "Error validating object"
	errorSessionId        = "Error in session id generation"

	// this value corresponds to 30 minutes in seconds. It is for demo purposes, and is based in a quick google search: "average session duration"
	sessionLength = 1800
)

func main() {
	http.HandleFunc("/screenresize", handleScreenResizeEvents)
	http.HandleFunc("/timetaken", handleTimeTakenEvents)
	http.HandleFunc("/copypaste", handleCopyPasteEvents)
	http.HandleFunc("/session", handleSessionCreation)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleScreenResizeEvents(responseWriter http.ResponseWriter, request *http.Request) {

	// validate request method and body
	if request.Method != http.MethodPost {

		log.Println(invalidMethodReceived)
		http.Error(responseWriter, invalidMethodReceived, http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {

		log.Println(unableToReadBody, "with error", err)
		http.Error(responseWriter, unableToReadBody, http.StatusBadRequest)
		return
	}

	screenResizeReceived := &event_service.ScreenResizeEvent{}

	if err = json.Unmarshal(body, screenResizeReceived); err != nil {

		log.Println(unableToUnmarshall, "with error", err)
		http.Error(responseWriter, unableToUnmarshall, http.StatusBadRequest)
		return
	}

	// validate payload content
	isValid, err := screenResizeReceived.Validate()

	if err != nil {
		log.Println(errorValidatingObject, "with error", err)
		http.Error(responseWriter, errorValidatingObject, http.StatusInternalServerError)
		return
	}

	if !isValid {
		log.Println(invalidObject, "with error", err)
		http.Error(responseWriter, invalidObject, http.StatusBadRequest)
		return
	}

	// process payload content
	dataToStore := screenResizeReceived.Map()

	updatedData, err := repository.SessionsData.Update(dataToStore)

	if err != nil {

		log.Println(errorUpdatingData, "with error", err)
		http.Error(responseWriter, errorUpdatingData, http.StatusInternalServerError)
		return
	}

	fmt.Printf("Session Data after screenSize update:\n %+v", updatedData)
}

func handleTimeTakenEvents(responseWriter http.ResponseWriter, request *http.Request) {

	// validate request method and body
	if request.Method != http.MethodPost {

		log.Println(invalidMethodReceived)
		http.Error(responseWriter, invalidMethodReceived, http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {

		log.Println(unableToReadBody, "with error", err)
		http.Error(responseWriter, unableToReadBody, http.StatusBadRequest)
		return
	}

	timeTakenReceived := &event_service.TimeTakenEvent{}

	if err = json.Unmarshal(body, timeTakenReceived); err != nil {

		log.Println(unableToUnmarshall, "with error", err)
		http.Error(responseWriter, unableToUnmarshall, http.StatusBadRequest)
		return
	}

	// validate payload content
	isValid, err := timeTakenReceived.Validate()

	if err != nil {
		log.Println(errorValidatingObject, "with error", err)
		http.Error(responseWriter, errorValidatingObject, http.StatusInternalServerError)
		return
	}

	if !isValid {
		log.Println(invalidObject, "with error", err)
		http.Error(responseWriter, invalidObject, http.StatusBadRequest)
		return
	}

	// process payload content
	dataToStore := timeTakenReceived.Map()

	updatedData, err := repository.SessionsData.Update(dataToStore)

	if err != nil {

		log.Println(errorUpdatingData, "with error", err)
		http.Error(responseWriter, errorUpdatingData, http.StatusInternalServerError)
		return
	}

	fmt.Printf("Session Data after time taken event update:\n %+v", updatedData)
}

func handleCopyPasteEvents(responseWriter http.ResponseWriter, request *http.Request) {

	// validate request method and body
	if request.Method != http.MethodPost {

		log.Println(invalidMethodReceived)
		http.Error(responseWriter, invalidMethodReceived, http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {

		log.Println(unableToReadBody, "with error", err)
		http.Error(responseWriter, unableToReadBody, http.StatusBadRequest)
		return
	}

	copyPasteReceived := &event_service.CopyPasteEvent{}

	if err = json.Unmarshal(body, copyPasteReceived); err != nil {

		log.Println(unableToUnmarshall, "with error", err)
		http.Error(responseWriter, unableToUnmarshall, http.StatusBadRequest)
		return
	}

	// validate payload content
	isValid, err := copyPasteReceived.Validate()

	if err != nil {
		log.Println(errorValidatingObject, "with error", err)
		http.Error(responseWriter, errorValidatingObject, http.StatusInternalServerError)
		return
	}

	if !isValid {
		log.Println(invalidObject, "with error", err)
		http.Error(responseWriter, invalidObject, http.StatusBadRequest)
		return
	}

	// process payload content
	dataToStore := copyPasteReceived.Map()

	updatedData, err := repository.SessionsData.Update(dataToStore)

	if err != nil {
		log.Println(errorUpdatingData, "with error", err)
		http.Error(responseWriter, errorUpdatingData, http.StatusInternalServerError)
		return
	}

	fmt.Printf("Session Data after copy paste update:\n %+v", updatedData)
}

func handleSessionCreation(responseWriter http.ResponseWriter, request *http.Request) {

	// validate request method and body
	if request.Method != http.MethodPost {

		log.Println(invalidMethodReceived)
		http.Error(responseWriter, invalidMethodReceived, http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {

		log.Println(unableToReadBody, "with error", err)
		http.Error(responseWriter, unableToReadBody, http.StatusBadRequest)
		return
	}

	sessionReceived := &event_service.SessionEvent{}

	if err = json.Unmarshal(body, sessionReceived); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	// validate payload content
	err = sessionReceived.Validate()

	if err != nil {
		log.Println(invalidObject, "with error", err)
		http.Error(responseWriter, invalidObject, http.StatusBadRequest)
		return
	}

	// process payload content

	sessionId, err := generateSessionId()

	if err != nil {
		log.Println(errorSessionId, "with error", err)
		http.Error(responseWriter, errorSessionId, http.StatusInternalServerError)
		return
	}

	updatedData, err := repository.SessionsData.InitUserSession(sessionId, sessionReceived.WebsiteUrl)

	if err != nil {
		log.Println(errorSessionId, "with error", err)
		http.Error(responseWriter, errorSessionId, http.StatusInternalServerError)
		return
	}

	cookieObject := &http.Cookie{
		Name:   "session",
		Value:  sessionId,
		MaxAge: sessionLength,
	}

	http.SetCookie(responseWriter, cookieObject)

	fmt.Printf("Session Data after sessionId creation :\n %+v", updatedData)

	fmt.Printf("Hashed webSiteUrl: %s", hash_service.Generate(sessionReceived.WebsiteUrl))

}

func generateSessionId() (string, error) {

	// DEFINE SESSION ID GENERATION
	return "", nil
}
