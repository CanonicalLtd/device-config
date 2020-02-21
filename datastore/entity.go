package datastore

// Session details for a specific user session
type Session struct {
	ID        int64
	Username  string
	SessionID string
}
