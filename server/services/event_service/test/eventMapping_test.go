package event_service

import (
	"code-test/server/model"
	"code-test/server/repository"
	"code-test/server/services/event_service"
	"testing"
)

// var getMock func(sessionId string, websiteUrl string) (*model.Data, error)
// var initUserMock func(sessionId string, websiteUrl string) (*model.Data, error)
// var updateMock func(receivedSessionData *model.Data) (*model.Data, error)
// var saveMock func(receivedSessionData *model.Data) error

type mapSessionDataStorageMock struct{}

func (sessionStorageMock mapSessionDataStorageMock) Get(sessionId string, websiteUrl string) (*model.Data, error) {
	return getMock(sessionId, websiteUrl)
}

func (sessionStorageMock mapSessionDataStorageMock) InitUserSession(sessionId string, websiteUrl string) (*model.Data, error) {
	return initUserMock(sessionId, websiteUrl)
}

func (sessionStorageMock mapSessionDataStorageMock) Save(*model.Data) error {
	return saveMock(&model.Data{})
}

func (sessionStorageMock mapSessionDataStorageMock) Update(receivedSessionData *model.Data) (*model.Data, error) {
	return updateMock(&model.Data{})
}

func TestMap_receiveScreenResizeEvent_MapToDataModelWithSuccess(t *testing.T) {

	// Arrange
	repository.SessionsData = mapSessionDataStorageMock{}

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{

			SessionId:  "1235",
			WebsiteUrl: "https://ravelin.com",
			ResizeFrom: model.Dimension{
				Height: "",
				Width:  "",
			},
			ResizeTo: model.Dimension{
				Height: "",
				Width:  "",
			},

			Time: 34,
			CopyAndPaste: map[string]bool{
				"formId": true,
			},
		}, nil
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

	expectedResult := &model.Data{
		SessionId:  "1235",
		WebsiteUrl: "https://ravelin.com",
		ResizeFrom: model.Dimension{
			Height: "1",
			Width:  "2",
		},
		ResizeTo: model.Dimension{
			Height: "1",
			Width:  "2",
		},

		Time: 34,
		CopyAndPaste: map[string]bool{
			"formId": true,
		},
	}

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

	res, err := screenResizeEvent.Map()

	// Assert

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if expectedResult.Time != res.Time {
		t.Errorf("Expected %v, got %v", expectedResult.Time, res.Time)
	}

	if expectedResult.CopyAndPaste["formId"] != res.CopyAndPaste["formId"] {
		t.Errorf("Expected %v, got %v", expectedResult.CopyAndPaste, res.CopyAndPaste)
	}

	if expectedResult.ResizeFrom != res.ResizeFrom {
		t.Errorf("Expected %v, got %v", expectedResult.ResizeFrom, res.ResizeFrom)
	}

	if expectedResult.ResizeTo != res.ResizeFrom {
		t.Errorf("Expected %v, got %v", expectedResult.ResizeTo, res.ResizeTo)
	}

	if expectedResult.SessionId != res.SessionId {
		t.Errorf("Expected %v, got %v", expectedResult.SessionId, res.SessionId)
	}

	if expectedResult.WebsiteUrl != res.WebsiteUrl {
		t.Errorf("Expected %v, got %v", expectedResult.WebsiteUrl, res.WebsiteUrl)
	}
}

func TestMap_receiveTimeTakenEvent_MapToDataModelWithSuccess(t *testing.T) {

	// Arrange

	repository.SessionsData = mapSessionDataStorageMock{}

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{

			SessionId:  "1235",
			WebsiteUrl: "https://ravelin.com",
			ResizeFrom: model.Dimension{
				Height: "1",
				Width:  "2",
			},
			ResizeTo: model.Dimension{
				Height: "1",
				Width:  "2",
			},

			Time: 0,
			CopyAndPaste: map[string]bool{
				"formId": true,
			},
		}, nil
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

	expectedResult := &model.Data{
		SessionId:  "1235",
		WebsiteUrl: "https://ravelin.com",
		Time:       12,
		ResizeFrom: model.Dimension{
			Height: "1",
			Width:  "2",
		},
		ResizeTo: model.Dimension{
			Height: "1",
			Width:  "2",
		},
		CopyAndPaste: map[string]bool{
			"formId": true,
		},
	}

	timeTakenEvent := &event_service.TimeTakenEvent{
		EventType:  "timeTaken",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		Time:       12,
	}

	// Act

	res, err := timeTakenEvent.Map()

	// Assert

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}

	if expectedResult.Time != res.Time {
		t.Errorf("Expected %v, got %v", expectedResult.Time, res.Time)
	}

	if expectedResult.CopyAndPaste["formId"] != res.CopyAndPaste["formId"] {
		t.Errorf("Expected %v, got %v", expectedResult.CopyAndPaste, res.CopyAndPaste)
	}

	if expectedResult.ResizeFrom != res.ResizeFrom {
		t.Errorf("Expected %v, got %v", expectedResult.ResizeFrom, res.ResizeFrom)
	}

	if expectedResult.ResizeTo != res.ResizeFrom {
		t.Errorf("Expected %v, got %v", expectedResult.ResizeTo, res.ResizeTo)
	}

	if expectedResult.SessionId != res.SessionId {
		t.Errorf("Expected %v, got %v", expectedResult.SessionId, res.SessionId)
	}

	if expectedResult.WebsiteUrl != res.WebsiteUrl {
		t.Errorf("Expected %v, got %v", expectedResult.WebsiteUrl, res.WebsiteUrl)
	}
}

func TestMap_receiveCopyPasteEvent_MapToDataModelWithSuccess(t *testing.T) {

	// Arrange

	repository.SessionsData = mapSessionDataStorageMock{}

	getMock = func(sessionId string, websiteUrl string) (*model.Data, error) {
		return &model.Data{

			SessionId:  "1235",
			WebsiteUrl: "https://ravelin.com",
			ResizeFrom: model.Dimension{
				Height: "1",
				Width:  "2",
			},
			ResizeTo: model.Dimension{
				Height: "1",
				Width:  "2",
			},

			Time:         34,
			CopyAndPaste: map[string]bool{},
		}, nil
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

	expectedResult := &model.Data{
		SessionId:  "1235",
		WebsiteUrl: "https://ravelin.com",
		CopyAndPaste: map[string]bool{
			"testFormId": true,
		},
		Time: 34,
		ResizeFrom: model.Dimension{
			Height: "1",
			Width:  "2",
		},
		ResizeTo: model.Dimension{
			Height: "1",
			Width:  "2",
		},
	}

	copyPasteEvent := &event_service.CopyPasteEvent{
		EventType:  "timeTaken",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		FormId:     "testFormId",
		Pasted:     true,
	}

	// Act

	res, err := copyPasteEvent.Map()

	// Assert

	if err != nil {
		t.Errorf("Expected %v, got %v", nil, err)
	}
	if expectedResult.Time != res.Time {
		t.Errorf("Expected %v, got %v", expectedResult.Time, res.Time)
	}

	if expectedResult.ResizeFrom != res.ResizeFrom {
		t.Errorf("Expected %v, got %v", expectedResult.ResizeFrom, res.ResizeFrom)
	}

	if expectedResult.ResizeTo != res.ResizeFrom {
		t.Errorf("Expected %v, got %v", expectedResult.ResizeTo, res.ResizeTo)
	}

	if expectedResult.CopyAndPaste["testFormId"] != res.CopyAndPaste["testFormId"] {
		t.Errorf("Expected %v, got %v", expectedResult.WebsiteUrl, res.WebsiteUrl)
	}

	if expectedResult.SessionId != res.SessionId {
		t.Errorf("Expected %v, got %v", expectedResult.WebsiteUrl, res.WebsiteUrl)
	}

	if expectedResult.WebsiteUrl != res.WebsiteUrl {
		t.Errorf("Expected %v, got %v", expectedResult.WebsiteUrl, res.WebsiteUrl)
	}
}
