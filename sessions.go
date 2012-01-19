package webtestutil

import (
	"os"
	"http"

	"gorilla.googlecode.com/hg/gorilla/sessions"
)

// Session data
var sessionData sessions.SessionData

// Transient testing session store with easy inspection and manipulation
type TestingSessionStore struct {
	// List of encoders registered for this store.
	encoders []sessions.SessionEncoder
}

// Register the testing session store and use it as the default store key
func RegisterTestingStore() {
	sessions.DefaultSessionFactory.SetStore("testing", new(TestingSessionStore))
	sessions.DefaultStoreKey = "testing"
}

// Set the session data
func SetSessionData(data sessions.SessionData) {
	sessionData = data
}

func SessionData() sessions.SessionData {
	if sessionData == nil {
		sessionData = make(sessions.SessionData)
	}
	return sessionData
}


// Reset the session data, should be called in each test using the session
func ResetSession() {
	sessionData = nil
}

func (s *TestingSessionStore) Load(r *http.Request, key string, info *sessions.SessionInfo) {
	info.Data = SessionData()
}

func (s *TestingSessionStore) Save(r *http.Request, w http.ResponseWriter, key string, info *sessions.SessionInfo) (bool, os.Error) {
	sessionData = info.Data
	return true, nil
}

func (s *TestingSessionStore) Encoders() []sessions.SessionEncoder {
	return s.encoders
}

func (s *TestingSessionStore) SetEncoders(encoders ...sessions.SessionEncoder) {
	s.encoders = encoders
}