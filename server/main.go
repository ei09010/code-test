package main

import (
	"code-test/server/model"
	"code-test/server/repository"
	"code-test/server/services/event_service"
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
)

func main() {
	http.HandleFunc("/screenresize", handleScreenResizeEvents)
	http.HandleFunc("/timetaken", handleTimeTakenEvents)
	http.HandleFunc("/copypaste", handleCopyPasteEvents)

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

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Unable to read body, returned the following error", err)
		return
	}

	timeTakenReceived := &event_service.TimeTakenEvent{}

	if err = json.Unmarshal(body, timeTakenReceived); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	dataToStore := &model.Data{
		WebsiteUrl: timeTakenReceived.WebsiteUrl,
		SessionId:  timeTakenReceived.SessionId,

		FormCompletionTime: timeTakenReceived.FormCompletionTime,
	}

	// validate method POST

	// validate event type -> declare consts with event types expected and check if event type is any of the expected

	// auxiliar method to validate webSiteUrl (regex ? )

	updatedData, err := repository.SessionsData.Update(dataToStore)

	fmt.Printf("Session Data after time taken event update %+v", updatedData)
}

func handleCopyPasteEvents(responseWriter http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Unable to read body, returned the following error", err)
		return
	}

	copyPasteReceived := &event_service.CopyPasteEvent{}

	if err = json.Unmarshal(body, copyPasteReceived); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	dataToStore := &model.Data{
		WebsiteUrl: copyPasteReceived.WebsiteUrl,
		SessionId:  copyPasteReceived.SessionId,

		CopyAndPaste: map[string]bool{
			copyPasteReceived.FormId: copyPasteReceived.Pasted,
		},
	}

	// validate method POST

	// validate event type -> declare consts with event types expected and check if event type is any of the expected

	// auxiliar method to validate webSiteUrl (regex ? )

	updatedData, err := repository.SessionsData.Update(dataToStore)

	fmt.Printf("Session Data after copy paste update %+v", updatedData)
}

func handleSessionCreation(responseWriter http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Unable to read body, returned the following error", err)
		return
	}

	sessionReceived := &event_service.SessionEvent{}

	if err = json.Unmarshal(body, sessionReceived); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	// validate method POST

	// auxiliar method to validate webSiteUrl

	// generate session Id
	sessionId := "generated session id" //TODO

	// return to client

	//TODO

	// store in Db

	repository.InitUserSession(sessionId, sessionReceived.WebsiteUrl)

}
