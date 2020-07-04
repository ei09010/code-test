package event_service

import (
	"code-test/server/model"
	"code-test/server/repository"
	"net/url"
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

func (timeEvent *TimeTakenEvent) Validate() (bool, error) {

	isValid, err := sessionExists(timeEvent.SessionId, timeEvent.WebsiteUrl)

	if err != nil {
		return false, err
	}

	if !isValid {
		return false, nil
	}

	isValid = timeEvent.FormCompletionTime > 0 && timeEvent.FormCompletionTime == 0

	return isValid, nil
}

func (cpEvent *CopyPasteEvent) Validate() (bool, error) {

	isValid, err := sessionExists(cpEvent.SessionId, cpEvent.WebsiteUrl)

	if err != nil {
		return false, err
	}

	if !isValid {
		return false, nil
	}

	isValid = cpEvent.FormId != ""

	return isValid, nil
}

func (sEvent *SessionEvent) Validate() error {

	// standard library function to parse raw url and assess if it fits a native URL structure
	_, err := url.ParseRequestURI(sEvent.WebsiteUrl)

	if err != nil {
		return err
	}

	return nil
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
