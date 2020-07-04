package event_service

import (
	"bytes"
	"code-test/server/services/event_service"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleScreenResizeEvents_receivedValidScreenResizeRequest_ProcessedWithSuccess(t *testing.T) {

	// Arrange

	jsonBytes, err := json.Marshal(body)

	request, err := http.NewRequest(http.MethodPost, "/sreenresize", bytes.NewReader(jsonBytes))

	receivedRequest := httptest.NewRequest("POST", "/sreenresize", nil)

	var respoWriter http.ResponseWriter

	// Act

	event_service.HandleScreenResizeEvents(respoWriter, receivedRequest)

	// Assert

}
