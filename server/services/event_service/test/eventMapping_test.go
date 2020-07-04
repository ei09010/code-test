package event_service

import (
	"code-test/server/model"
	"code-test/server/services/event_service"
	"testing"
)
func TestMap_receiveScreenResizeEvent_MapToDataModelWithSuccess(t *testing.T){
	
	// Arrange

	expectedResult := &model.Data{
		SessionId: "1235",
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

	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}
}

func TestMap_receiveTimeTakenEvent_MapToDataModelWithSuccess(t *testing.T){
	
	// Arrange

	expectedResult := &model.Data{
		SessionId: "1235",
		WebsiteUrl: "https://ravelin.com",
		FormCompletionTime: 12,
	}

	timeTakenEvent := &event_service.TimeTakenEvent{
		EventType:  "timeTaken",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		FormCompletionTime: 12,
	}

	// Act

	res := screenResizeEvent.Map()


	// Assert

	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}
}

func TestMap_receiveCopyPasteEvent_MapToDataModelWithSuccess(t *testing.T){
	
	// Arrange

	expectedResult := &model.Data{
		SessionId: "1235",
		WebsiteUrl: "https://ravelin.com",
		CopyAndPaste: map[string]bool{
			"testFormId": true
		},
	}

	copyPasteEvent := &event_service.CopyPasteEvent{
		EventType:  "timeTaken",
		WebsiteUrl: "https://ravelin.com",
		SessionId:  "1235",
		FormId: "testFormId",
		Pasted: true,
	}

	// Act

	res := copyPasteEvent.Map()


	// Assert

	if expectedResult != res {
		t.Errorf("Expected %v, got %v", expectedResult, res)
	}
}