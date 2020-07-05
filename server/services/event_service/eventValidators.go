package event_service

import (
	"code-test/server/model"
	"code-test/server/repository"
	"net/url"
)

type ScreenResizeEvent struct {
	EventType  string `json:"eventType"`
	WebsiteUrl string `json:"websiteUrl"`
	SessionId  string `json:"sessionId"`
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

//"{\"eventType\":\"timeTaken\",\"websiteUrl\":\"https://ravelin.com\",\"sessionId\":\"123123-123123-123123123\",\"time\":72}"
type TimeTakenEvent struct {
	EventType  string `json:"eventType"`
	WebsiteUrl string `json:"websiteUrl"`
	SessionId  string `json:"sessionId"`
	Time       int    `json:"time"` // seconds
}

type SessionEvent struct {
	WebsiteURL string `json:"websiteUrl"`
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

	isValid = timeEvent.Time > 0

	if err != nil {
		return false, err
	}

	if !isValid {
		return false, nil
	}

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
	_, err := url.ParseRequestURI(sEvent.WebsiteURL)

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
