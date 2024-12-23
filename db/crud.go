package db

import "fmt"

func NewDatabase(cap int) *Database {
	return &Database{
		// No Lock initialization is needed, as we can
		// just use the default value of RWMutex.
		DB:   make(map[string][]byte),
		Cap:  cap,
		Size: 0,
	}
}

func (d *Database) Read(key string) ([]byte, error) {
	d.Lock.RLock()
	defer d.Lock.RUnlock()

	value, found := d.DB[key]

	if !found {
		return nil, fmt.Errorf("key isn't present in the database, key: %s", key)
	}

	return value, nil
}

// Update is used for both updating old values and adding new ones.
func (d *Database) Update(key string, value []byte) error {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	if d.Size == d.Cap {
		return fmt.Errorf("database has exceeded its maximum capacity")
	}

	d.DB[key] = value
	d.Size++

	return nil
}

func (d *Database) Delete(key string) {
	d.Lock.Lock()
	defer d.Lock.Unlock()

	delete(d.DB, key)
	d.Size--
}
