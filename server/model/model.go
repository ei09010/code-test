package model

// Objects that represent the webServer data model

type Data struct {
	WebsiteUrl   string
	SessionId    string
	ResizeFrom   Dimension
	ResizeTo     Dimension
	CopyAndPaste map[string]bool // map[fieldId]true
	Time         int             // Seconds
}

type Dimension struct {
	Width  string
	Height string
}
