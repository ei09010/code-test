package event_service

import (
	"code-test/server/model"
	"code-test/server/services/event_service"
	"testing"
)

func TestMap_receiveScreenResizeEvent_MapToDataModelWithSuccess(t *testing.T) {

	// Arrange

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

	res := screenResizeEvent.Map()

	// Assert

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

	expectedResult := &model.Data{
		SessionId:  "1235",
		WebsiteUrl: "https://ravelin.com",
		Time:       12,
	}

	timeTakenEvent := &event_service.TimeTakenEvent{
		EventType:  "timeTaken",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		Time:       12,
	}

	// Act

	res := timeTakenEvent.Map()

	// Assert

	if expectedResult.Time != res.Time {
		t.Errorf("Expected %v, got %v", expectedResult.Time, res.Time)
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

	expectedResult := &model.Data{
		SessionId:  "1235",
		WebsiteUrl: "https://ravelin.com",
		CopyAndPaste: map[string]bool{
			"testFormId": true,
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

	res := copyPasteEvent.Map()

	// Assert

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
