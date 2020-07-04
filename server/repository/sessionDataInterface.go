package repository

// Session Repository interface is declared for the following reasons:
//   1 - By declaring this methods signature, every session repository implemented behaviour should should follow this methods
//   2 - If in the future I decided to switch to a different storage for my session related data (e. g. persistent data storage), the method signature would stand
// 	everywhere in the app, and all change will be concentrated in this methods implementation
//   3 - Interface declaration would also enable me to use dependency injection

type SessionData interface {
	Save(receivedSessionData interface{}) error                             //set session data object
	Get(sessionId interface{}, websiteUrl interface{}) (interface{}, error) // get session data object
	Update(receivedSessionData interface{}) error                           // Update session data object
}
