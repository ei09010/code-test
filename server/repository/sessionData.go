package repository

import (
	"code-test/server/model"
	"errors"
	"log"
	"sync"
)

var SessionsData SessionData

func init() {
	SessionsData = &SessionDataStorage{
		sessionData: make(map[string]*model.Data),
		mu:          sync.Mutex{},
	}
}

type SessionDataStorage struct {
	sessionData map[string]*model.Data
	mu          sync.Mutex
}

// I'm having a persisting issue that prompted me to crate this method: eventhough my test file and this file are in the same package, I'm not being able to access
// SessionDataStorage private properties (lower case). If I have time available, will keep trying to fix
func GetInstance() *SessionDataStorage {
	return &SessionDataStorage{
		sessionData: make(map[string]*model.Data),
		mu:          sync.Mutex{},
	}
}

func (sDataStore *SessionDataStorage) InitUserSession(sessionId string, websiteUrl string) (*model.Data, error) {

	dataToReturn := &model.Data{
		SessionId:    sessionId,
		WebsiteUrl:   websiteUrl,
		CopyAndPaste: make(map[string]bool),
	}

	err := sDataStore.Save(dataToReturn)

	if err != nil {
		log.Println("Error while saving session data", err)
		return nil, err
	}

	return dataToReturn, nil
}

func (sDataStore *SessionDataStorage) Save(receivedSessionData *model.Data) error {

	sDataStore.mu.Lock()
	defer sDataStore.mu.Unlock()

	mapKey := buildKey(receivedSessionData.SessionId, receivedSessionData.WebsiteUrl)

	sDataStore.sessionData[mapKey] = receivedSessionData

	return nil
}

func (sDataStore *SessionDataStorage) Get(sessionId string, websiteUrl string) (*model.Data, error) {

	sDataStore.mu.Lock()
	defer sDataStore.mu.Unlock()

	mapKey := buildKey(sessionId, websiteUrl)

	dataToReturn := sDataStore.sessionData[mapKey]

	return dataToReturn, nil

}

// here we store the Data object, that will always have some properties with the zero-value due to the fragmented nature of the handled events
func (sDataStore *SessionDataStorage) Update(receivedSessionData *model.Data) (*model.Data, error) {

	sDataStore.mu.Lock()
	defer sDataStore.mu.Unlock()

	if receivedSessionData == nil {
		return nil, errors.New("Received nil Data object")
	}

	mapKey := buildKey(receivedSessionData.SessionId, receivedSessionData.WebsiteUrl)

	// store screensize events
	if dataStored, ok := sDataStore.sessionData[mapKey]; ok {

		// store time taken events

		if receivedSessionData.Time > 0 {

			dataStored.Time = receivedSessionData.Time
		}

		// store copy paste events
		// Given that the paste operation will only change from false to true once, I'm only adding to the dictionary

		for k, v := range receivedSessionData.CopyAndPaste {
			if _, ok := dataStored.CopyAndPaste[k]; !ok {

				// delete map default value create when initializing session
				if _, ok := dataStored.CopyAndPaste[""]; ok {
					delete(dataStored.CopyAndPaste, "")
				}

				dataStored.CopyAndPaste[k] = v
			}
		}

		// Since only one re-size happens, I'm assuming that if already stored resize data is empty, we can override with valid (non zero-value) received resize data

		receivedResizeFromValid := receivedSessionData.ResizeFrom.Height != "" && receivedSessionData.ResizeFrom.Width != ""
		dataStoredResizeFromInvalid := dataStored.ResizeFrom.Height == "" && dataStored.ResizeFrom.Width == ""

		if receivedResizeFromValid && dataStoredResizeFromInvalid {
			dataStored.ResizeFrom = receivedSessionData.ResizeFrom
		}

		receivedResizeToValid := receivedSessionData.ResizeTo.Height != "" && receivedSessionData.ResizeTo.Width != ""
		dataStoredResizeToInvalid := dataStored.ResizeTo.Height == "" && dataStored.ResizeTo.Width == ""

		if receivedResizeToValid && dataStoredResizeToInvalid {
			dataStored.ResizeTo = receivedSessionData.ResizeTo
		}

	} else {

		return nil, errors.New("This session is not created. Unable to update any object")
	}

	return sDataStore.sessionData[mapKey], nil

}

func buildKey(sessionId string, websiteUrl string) string {
	return sessionId + "|" + websiteUrl
}
