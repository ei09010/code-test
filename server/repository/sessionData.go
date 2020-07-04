package repository

import (
	"code-test/server/model"
	"errors"
	"log"
	"sync"
)

type eventStorage struct {
	sessionData map[string]*model.Data
	mu          sync.Mutex
}

var SessionsData eventStorage

func Init() {
	SessionsData = eventStorage{
		sessionData: make(map[string]*model.Data),
	}
}

func (evStore *eventStorage) InitUserSession(sessionId string, websiteUrl string) (*model.Data, error) {

	dataToReturn := &model.Data{
		SessionId:    sessionId,
		WebsiteUrl:   websiteUrl,
		CopyAndPaste: make(map[string]bool),
	}

	err := SessionsData.Save(dataToReturn)

	if err != nil {
		log.Println("Error while saving session data", err)
		return nil, err
	}

	return dataToReturn, nil
}

func (evStore *eventStorage) Save(receivedSessionData *model.Data) error {

	evStore.mu.Lock()
	defer evStore.mu.Unlock()

	mapKey := buildKey(receivedSessionData.SessionId, receivedSessionData.WebsiteUrl)

	evStore.sessionData[mapKey] = receivedSessionData

	return nil
}

func (evStore *eventStorage) Get(sessionId string, websiteUrl string) (*model.Data, error) {

	evStore.mu.Lock()
	defer evStore.mu.Unlock()

	mapKey := buildKey(sessionId, websiteUrl)

	dataToReturn := evStore.sessionData[mapKey]

	return dataToReturn, nil

}

// here we store the Data object, that will always have some properties with the zero-value due to the fragmented nature of the handled events
func (evStore *eventStorage) Update(receivedSessionData *model.Data) (*model.Data, error) {

	evStore.mu.Lock()
	defer evStore.mu.Unlock()

	if receivedSessionData == nil {
		return nil, errors.New("Received nil Data object")
	}

	mapKey := buildKey(receivedSessionData.SessionId, receivedSessionData.WebsiteUrl)

	// store screensize events
	if dataStored, ok := evStore.sessionData[mapKey]; ok {

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

		// store time taken events

		if receivedSessionData.FormCompletionTime > 0 && dataStored.FormCompletionTime == 0 {

			dataStored.FormCompletionTime = receivedSessionData.FormCompletionTime
		}

		// store copy paste events
		// Given that the paste operation will only change from false to true once, I'm only adding to the dictionary

		for k, v := range receivedSessionData.CopyAndPaste {
			if _, ok := dataStored.CopyAndPaste[k]; !ok {
				dataStored.CopyAndPaste[k] = v
			}
		}

	} else {

		evStore.sessionData[mapKey] = receivedSessionData
	}

	return evStore.sessionData[mapKey], nil

}

func buildKey(sessionId string, websiteUrl string) string {
	return sessionId + "|" + websiteUrl
}
