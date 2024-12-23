package db

import (
	"sync"
)

type Database struct {
	DB map[string][]byte

	// Cap is needed to ensure, that the database won't
	// be flooded with unwated data.
	Cap  int
	Size int

	// Lock is used to maintain the ability to perform
	// concurrent read operations, while protecting the
	// database from race conditions during writing.
	Lock sync.RWMutex
}
