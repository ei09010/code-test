package model

type Data struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int             // Seconds
}

type ScreenResizeEvent struct {
	WebsiteUrl string
	SessionId  string
	ResizeFrom Dimension
	ResizeTo   Dimension
}

type CopyPasteEvent struct {
	WebsiteUrl   string
	SessionId    string
	CopyAndPaste map[string]bool // map[fieldId]true
}

type TimeTakenEvent struct {
	WebsiteUrl         string
	SessionId          string
	FormCompletionTime int // Seconds
}

type CompleteData struct {
	WebsiteUrl         string
	SessionId          string
	ResizeFrom         Dimension
	ResizeTo           Dimension
	CopyAndPaste       map[string]bool // map[fieldId]true
	FormCompletionTime int             // Seconds
}

type Dimension struct {
	Width  string
	Height string
}
