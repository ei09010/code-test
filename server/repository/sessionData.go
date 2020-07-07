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
	if _, ok := sDataStore.sessionData[mapKey]; ok {

		sDataStore.sessionData[mapKey] = receivedSessionData

	} else {

		return nil, errors.New("This session is not created. Unable to update any object")
	}

	return sDataStore.sessionData[mapKey], nil

}

func buildKey(sessionId string, websiteUrl string) string {
	return sessionId + "|" + websiteUrl
}

// I'm having a persisting issue that prompted me to create this method: eventhough my test file and this file are in the same package, I'm not being able to access
// SessionDataStorage private properties (lower case) to arrange test cases. If I have time available, will keep trying to fix this
func GetInstance() *SessionDataStorage {
	return &SessionDataStorage{
		sessionData: make(map[string]*model.Data),
		mu:          sync.Mutex{},
	}
}
