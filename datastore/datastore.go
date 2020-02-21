package datastore

// DataStore is the interfaces for the data repository
type DataStore interface {
	CreateSession(user Session) (int64, error)
	GetSession(username, sessionID string) (*Session, error)
}
