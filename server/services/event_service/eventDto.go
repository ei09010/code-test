package event_service

import (
	"code-test/server/model"
)

// the sole purpose of this data transfer objects is to hold data that will mapped to the repository layer

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

type TimeTakenEvent struct {
	EventType  string `json:"eventType"`
	WebsiteUrl string `json:"websiteUrl"`
	SessionId  string `json:"sessionId"`
	Time       int    `json:"time"` // seconds
}

type SessionEvent struct {
	WebsiteURL string `json:"websiteUrl"`
}
