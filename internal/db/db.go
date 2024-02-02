package db

import "time"

// db interface with save and find methods

// I would discourage using this for simple projects.
type DB interface {
    Connect() error
    Disconnect() error
	Save(url string, expiresAt time.Duration) (int, error)
	Find(id int) (string, error)
    Update(id int, options map[string]interface{}) (string, error)
}
