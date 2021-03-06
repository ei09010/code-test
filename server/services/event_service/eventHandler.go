package event_service

import (
	"code-test/server/repository"
	"code-test/server/services/hash_service"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	invalidMethodReceived      = "Received http method is not the expected one"
	unableToReadBody           = "Unable to read body"
	unableToUnmarshall         = "Unable to unMarshall request body"
	errorUpdatingData          = "Error updating data in the repository"
	invalidObject              = "The received object is invalid"
	errorValidatingObject      = "Error validating object"
	errorSessionId             = "Error in session id generation"
	errorReturningSessionId    = "Error returning session id"
	errorRetrievingObjectToMap = "Error retrieving object to map"

	// this value corresponds to 30 minutes in seconds. It is for demo purposes, and is based in a quick google search: "average session duration"
	sessionLength = 1800
)

func HandleScreenResizeEvents(responseWriter http.ResponseWriter, request *http.Request) {

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

	var screenResizeReceived ScreenResizeEvent

	if err = json.Unmarshal(body, &screenResizeReceived); err != nil {

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
		log.Println(invalidObject)
		return
	}

	// process payload content
	dataToStore, err := screenResizeReceived.Map()

	if err != nil {
		http.Error(responseWriter, errorRetrievingObjectToMap, http.StatusInternalServerError)
		return
	}

	updatedData, err := repository.SessionsData.Update(dataToStore)

	if err != nil {

		log.Println(errorUpdatingData, "with error", err)
		http.Error(responseWriter, errorUpdatingData, http.StatusInternalServerError)
		return
	}

	log.Printf("Session Data after screenSize update:\n %+v \n\n", updatedData)
}

func HandleTimeTakenEvents(responseWriter http.ResponseWriter, request *http.Request) {

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

	defer request.Body.Close()

	var timeTakenReceived TimeTakenEvent

	if err = json.Unmarshal(body, &timeTakenReceived); err != nil {

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
		log.Println(invalidObject)
		http.Error(responseWriter, invalidObject, http.StatusBadRequest)
		return
	}

	// process and store payload content
	dataToStore, err := timeTakenReceived.Map()

	if err != nil {
		http.Error(responseWriter, errorRetrievingObjectToMap, http.StatusInternalServerError)
		return
	}

	updatedData, err := repository.SessionsData.Update(dataToStore)

	if err != nil {

		log.Println(errorUpdatingData, "with error", err)
		http.Error(responseWriter, errorUpdatingData, http.StatusInternalServerError)
		return
	}

	log.Printf("Session Data after time taken event update:\n %+v \n\n", updatedData)
}

func HandleCopyPasteEvents(responseWriter http.ResponseWriter, request *http.Request) {

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

	var copyPasteReceived CopyPasteEvent

	if err = json.Unmarshal(body, &copyPasteReceived); err != nil {

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
		log.Println(invalidObject)
		http.Error(responseWriter, invalidObject, http.StatusBadRequest)
		return
	}

	// process and store payload content
	dataToStore, err := copyPasteReceived.Map()

	if err != nil {
		http.Error(responseWriter, errorRetrievingObjectToMap, http.StatusInternalServerError)
		return
	}

	updatedData, err := repository.SessionsData.Update(dataToStore)

	if err != nil {
		log.Println(errorUpdatingData, "with error", err)
		http.Error(responseWriter, errorUpdatingData, http.StatusInternalServerError)
		return
	}

	log.Printf("Session Data after copy paste update:\n %+v \n\n", updatedData)
}

func HandleSessionCreation(responseWriter http.ResponseWriter, request *http.Request) {

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

	var sessionReceived SessionEvent

	if err = json.Unmarshal([]byte(body), &sessionReceived); err != nil {
		log.Println("Unable to unMarshall request body, returned the following error: ", err)
		return
	}

	// validate payload content
	err = sessionReceived.Validate()

	if err != nil {
		log.Println(invalidObject)
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

	log.Printf("Initiating user session with sessionId: %v and url: %v.....\n\n", sessionId, sessionReceived.WebsiteURL)

	updatedData, err := repository.SessionsData.InitUserSession(sessionId, sessionReceived.WebsiteURL)

	if err != nil {
		log.Println(errorSessionId, "with error", err)
		http.Error(responseWriter, errorSessionId, http.StatusInternalServerError)
		return
	}

	// not being able to access cookies on client side to obtain a sessionId stored there, will have to return through response writer
	sessionCreated := SessionCreatedEvent{
		SessionId: updatedData.SessionId,
	}

	sessionResponseBytes, err := json.Marshal(sessionCreated)

	if err != nil {
		log.Println(errorSessionId, "with error", err)
		http.Error(responseWriter, errorSessionId, http.StatusInternalServerError)
		return
	}

	responseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Add("Access-Control-Allow-Credentials", "omit")
	responseWriter.Header().Add("Access-Control-Allow-Methods", "Allow")
	responseWriter.Header().Add("Access-Control-Allow-Methods", "OPTIONS, POST")
	responseWriter.Header().Add("Content-Type", "application/json")

	_, err = responseWriter.Write(sessionResponseBytes)

	if err != nil {
		log.Println(errorReturningSessionId, "with error", err)
		http.Error(responseWriter, errorReturningSessionId, http.StatusInternalServerError)
		return
	}

	log.Printf("Session Data after sessionId creation :\n %+#v \n\n", updatedData)

	log.Printf("Hashed websiteUrl: %s \n", hash_service.Generate(sessionReceived.WebsiteURL))
}

// Returns a 36-character string in the form XXXXXXXX-XXXX-XXXX-XXXX-XXXXXXXXXXXX
// where "X" is an "upper-case" hexadecimal digit [0-9A-F].
func generateSessionId() (string, error) {

	// 1- Generate 16 random bytes (=128 bits) in thread safe way
	randBytes := make([]byte, 16)

	rand.Read(randBytes)

	//  3 -the adjusted bytes as 32 hexadecimal digits
	hexEncoded := hex.EncodeToString(randBytes)

	//  4 - four hyphen "-" characters to obtain blocks of 8, 4, 4, 4 and 12 hex digits
	sessionIdToReturn := hexEncoded[:8] + "-" + hexEncoded[8:12] + "-" + hexEncoded[12:16] + "-" + hexEncoded[16:20] + "-" + hexEncoded[20:]

	return sessionIdToReturn, nil
}
