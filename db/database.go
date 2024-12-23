package db

import (
	"sync"

	"github.com/milosz-1111/dddb.git/config"
)

type Database struct {
	DB map[string][]byte

	// Config is used to store settings.
	Config config.Config
	Length int

	// Lock is used to maintain the ability to perform
	// concurrent read operations, while protecting the
	// database from race conditions during writing.
	Lock sync.RWMutex
}
