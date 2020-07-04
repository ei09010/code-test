package event_service

import (
	"code-test/server/model"
	"code-test/server/repository"
	"code-test/server/services/event_service"
	"testing"
)

// This helps in assigning mock at the runtime instead of compile time
var getMock func(sessionId string, websiteUrl string) (*model.Data, error)
var initUserMock func(sessionId string, websiteUrl string) (*model.Data, error)
var updateMock func(receivedSessionData *model.Data) (*model.Data, error)
var saveMock func(receivedSessionData *model.Data) error

type sessionDataStorageMock struct{}

func (sessionStorageMock sessionDataStorageMock) Get(sessionId string, websiteUrl string) (*model.Data, error) {
	return getMock(sessionId, websiteUrl)
}

func (sessionStorageMock sessionDataStorageMock) InitUserSession(sessionId string, websiteUrl string) (*model.Data, error) {
	return initUserMock(sessionId, websiteUrl)
}

func (sessionStorageMock sessionDataStorageMock) Save(*model.Data) error {
	return saveMock(&model.Data{})
}

func (sessionStorageMock sessionDataStorageMock) Update(receivedSessionData *model.Data) (*model.Data, error) {
	return updateMock(&model.Data{})
}

func TestValidate_receiveValidScreenResizeEvent_returnsTrueAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := true

	screenResizeEvent := &event_service.ScreenResizeEvent{
		EventType:  "screenResize",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		ResizeFrom: model.Dimension{
			Height: "1",
			Width:  "2",
		},
		ResizeTo: model.Dimension{
			Height: "1",
			Width:  "2",
		},
	}

	// Act

	res, err := screenResizeEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestValidate_receiveInvalidScreenResizeEvent_returnsFalseAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return nil, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := false

	screenResizeEvent := &event_service.ScreenResizeEvent{
		EventType:  "screenResize",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		ResizeFrom: model.Dimension{
			Height: "1",
			Width:  "2",
		},
		ResizeTo: model.Dimension{
			Height: "1",
			Width:  "2",
		},
	}

	// Act

	res, err := screenResizeEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestValidate_receiveValidTimeTakenEvent_returnsTrueAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := true

	timeTakenEvent := &event_service.TimeTakenEvent{
		EventType:          "screenResize",
		WebsiteUrl:         "https://ravelin.com",
		SessionId:          "1235",
		FormCompletionTime: 30,
	}

	// Act

	res, err := timeTakenEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestValidate_receiveInvalidBecauseSessionIdTimeTakenEvent_returnsFalseAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return nil, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := false

	timeTakenEvent := &event_service.TimeTakenEvent{
		EventType:          "screenResize",
		WebsiteUrl:         "https://ravelin.com",
		SessionId:          "1235",
		FormCompletionTime: 30,
	}

	// Act

	res, err := timeTakenEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestValidate_receiveInvalidBecauseTimeTakenZeroTimeTakenEvent_returnsFalseAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := false

	timeTakenEvent := &event_service.TimeTakenEvent{
		EventType:          "screenResize",
		WebsiteUrl:         "https://ravelin.com",
		SessionId:          "1235",
		FormCompletionTime: 0,
	}

	// Act

	res, err := timeTakenEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestValidate_receiveValidCopyPasteEvent_returnsTrueAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := true

	copyPasteEvent := &event_service.CopyPasteEvent{
		EventType:  "screenResize",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		FormId:     "testFromId",
		Pasted:     true,
	}

	// Act

	res, err := copyPasteEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestValidate_receiveInvalidBecauseSessionIdCopyPasteEvent_returnsFalseAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return nil, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := false

	copyPasteEvent := &event_service.CopyPasteEvent{
		EventType:  "screenResize",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		FormId:     "testFromId",
		Pasted:     true,
	}

	// Act

	res, err := copyPasteEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}

func TestValidate_receiveInvalidBecauseFormIdCopyPasteEvent_returnsFalseAndNil(t *testing.T) {

	// Arrange

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	initUserMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{}, nil
	}

	saveMock = func(receivedSessionData *model.Data) error {
		return nil
	}

	updateMock = func(receivedSessionData *model.Data) (*model.Data, error) {
		return &model.Data{}, nil
	}

	repository.SessionsData = sessionDataStorageMock{}
	expectedResult := false

	copyPasteEvent := &event_service.CopyPasteEvent{
		EventType:  "screenResize",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		FormId:     "",
		Pasted:     true,
	}

	// Act

	res, err := copyPasteEvent.Validate()

	// Assert

	// to validate booleans usually an expression such as Assert.IsTrue() is easier to read, but that doesn't seem to be available in the std library
	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

}
