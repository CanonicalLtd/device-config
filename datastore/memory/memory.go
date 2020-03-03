package memory

import (
	"fmt"
	"github.com/CanonicalLtd/device-config/datastore"
	"sync"
)

// Store implements an in-memory store for sessions
type Store struct {
	lock     sync.RWMutex
	Sessions []datastore.Session
}

// NewStore creates a new memory store
func NewStore() *Store {
	return &Store{
		Sessions: []datastore.Session{},
	}
}

// CreateSession creates a new user session
func (mem *Store) CreateSession(user datastore.Session) (int64, error) {
	mem.lock.Lock()
	defer mem.lock.Unlock()

	user.ID = int64(len(mem.Sessions) + 1)
	mem.Sessions = append(mem.Sessions, user)
	return user.ID, nil
}

// GetSession gets an existing user session
func (mem *Store) GetSession(username, sessionID string) (*datastore.Session, error) {
	mem.lock.RLock()
	defer mem.lock.RUnlock()

	for _, u := range mem.Sessions {
		if u.Username == username && u.SessionID == sessionID {
			return &u, nil
		}
	}

	return nil, fmt.Errorf("cannot find the user session `%s`", username)
}
