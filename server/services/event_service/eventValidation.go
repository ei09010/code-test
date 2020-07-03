package event_service

import (
	"code-test/server/model"
	"code-test/server/repository"
)

type ScreenResizeEvent struct {
	EventType  string
	WebsiteUrl string
	SessionId  string
	ResizeFrom model.Dimension
	ResizeTo   model.Dimension
}

type CopyPasteEvent struct {
	EventType  string
	WebsiteUrl string
	SessionId  string
	Pasted     bool // map[fieldId]true
	FormId     string
}

type TimeTakenEvent struct {
	EventType          string
	WebsiteUrl         string
	SessionId          string
	FormCompletionTime int // Seconds
}

type SessionEvent struct {
	WebsiteUrl string
}

func (scrEvent *ScreenResizeEvent) Validate() (bool, error) {

	isValid, err := sessionExists(scrEvent.SessionId, scrEvent.WebsiteUrl)

	if err != nil {
		return false, err
	}

	if !isValid {
		return false, nil
	}

	validResizeFrom := scrEvent.ResizeFrom.Height != "" && scrEvent.ResizeFrom.Width != ""
	validResizeTo := scrEvent.ResizeTo.Height != "" && scrEvent.ResizeTo.Width != ""

	isValid = validResizeFrom && validResizeTo

	return isValid, nil
}

// check if session exits, otherwise we should stop processing the event
func sessionExists(sessionId string, websiteUrl string) (bool, error) {

	dataFromDb, err := repository.SessionsData.Get(sessionId, websiteUrl)

	if err != nil {
		return false, err
	}

	if dataFromDb == nil {
		return false, nil
	}

	return true, nil
}
