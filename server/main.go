package main

import (
	"code-test/server/model"
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

	screenResizeObject := &model.ScreenResizeEvent{}

	if err = json.Unmarshal(body, screenResizeObject); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	// session will need to be thread safe

	if dataStored, ok := dbUserData[screenResizeObject.SessionId]; ok {

		dataStored.ResizeFrom = screenResizeObject.ResizeFrom
		dataStored.ResizeTo = screenResizeObject.ResizeTo

	} else {

		// auxiliar method to validate webSiteUrl

		// should session Id be generated here? It would definetly enable backend to have more control regarding expiration for example
		newData := &model.Data{
			SessionId:  screenResizeObject.SessionId,
			WebsiteUrl: screenResizeObject.WebsiteUrl,

			ResizeTo:   screenResizeObject.ResizeTo,
			ResizeFrom: screenResizeObject.ResizeFrom,
		}

		dbUserData[screenResizeObject.SessionId] = newData
	}

	fmt.Printf("%+v", screenResizeObject)
}

func handleTimeTakenEvents(responseWriter http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Unable to read body, returned the following error", err)
		return
	}

	timeTakenObject := &model.TimeTakenEvent{}

	if err = json.Unmarshal(body, timeTakenObject); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	fmt.Printf("%+v", timeTakenObject)
}

func handleCopyPasteEvents(responseWriter http.ResponseWriter, request *http.Request) {

	body, err := ioutil.ReadAll(request.Body)

	if err != nil {
		log.Println("Unable to read body, returned the following error", err)
		return
	}

	copyPasteObject := &model.CopyPasteEvent{}

	if err = json.Unmarshal(body, copyPasteObject); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error", err)
		return
	}

	fmt.Printf("%+v", copyPasteObject)
}

func handleSessionCreation(responseWriter http.ResponseWriter, request *http.Request) {

}
