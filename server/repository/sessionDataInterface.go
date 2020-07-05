package repository

import (
	"code-test/server/model"
)

// Session Repository interface is declared for the following reasons:
//   1 - By declaring this methods signature, every session repository implemented behaviour should should follow this methods
//   2 - If in the future I decided to switch to a different storage for my session related data (e. g. persistent data storage), the method signature would stand
// 	everywhere in the app, and all change will be concentrated in this methods implementation
//   3 - Interface declaration enables me to have better testability - I can crate mock implementations of this methods to implement unit tests

type SessionData interface {
	Save(receivedSessionData *model.Data) error                               //set session data object
	Get(sessionId string, websiteUrl string) (*model.Data, error)             // get session data object
	Update(receivedSessionData *model.Data) (*model.Data, error)              // Update session data object
	InitUserSession(sessionId string, websiteUrl string) (*model.Data, error) // create sessionId and Uel in session data object in storage
}
