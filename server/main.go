package main

import (
	"code-test/server/model"
	"code-test/server/repository"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var dbUserData = make(map[string]*model.Data)

func main() {
	http.HandleFunc("/screenresize", handleScreenResizeEvents)
	http.HandleFunc("/timetaken", handleTimeTakenEvents)
	http.HandleFunc("/copypaste", handleCopyPasteEvents)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleScreenResizeEvents(responseWriter http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Unable to read body, returned the following error", err)
		return
	}

	screenResizeReceived := &model.ScreenResizeEvent{}

	if err = json.Unmarshal(body, screenResizeReceived); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	dataToStore := &model.Data{
		WebsiteUrl: screenResizeReceived.WebsiteUrl,
		SessionId:  screenResizeReceived.SessionId,

		ResizeFrom: screenResizeReceived.ResizeFrom,
		ResizeTo:   screenResizeReceived.ResizeTo,
	}

	// validate method POST

	// validate event type -> declare consts with event types expected and check if event type is any of the expected

	// auxiliar method to validate webSiteUrl (regex ? )

	updatedData, err := repository.SessionsData.Update(dataToStore)

	fmt.Printf("Session Data after screenSize update %+v", updatedData)
}

func handleTimeTakenEvents(responseWriter http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Unable to read body, returned the following error", err)
		return
	}

	timeTakenReceived := &model.TimeTakenEvent{}

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

	copyPasteReceived := &model.CopyPasteEvent{}

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

	sessionReceived := &model.SessionEvent{}

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
