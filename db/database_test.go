package db

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/milosz-1111/dddb.git/config"
)

func Default() *Database {
	return &Database{
		DB: map[string][]byte{
			"Key1": {0, 0, 0},
			"Key2": {1, 1, 1},
			"Key3": {0, 1, 0},
			"Key4": {1, 0, 1},
		},
		Config: *config.Default(),
		Length: 4,
	}
}

func TestRead(t *testing.T) {
	d := Default()

	var tests = []struct {
		key           string
		expectedValue []byte
	}{
		{
			"Key1", []byte{0, 0, 0},
		},
		{
			"Key2", []byte{1, 1, 1},
		},
		{
			"Key3", []byte{0, 1, 0},
		},
		{
			"Key4", []byte{1, 0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Key: %s", tt.key), func(t *testing.T) {
			value, err := d.Read(tt.key)
			if err != nil {
				t.Fatalf("error: %v", err)
			}

			if !reflect.DeepEqual(value, tt.expectedValue) {
				t.Fatalf("got: %v, expected: %v", value, tt.expectedValue)
			}
		})
	}

}

func TestUpdate(t *testing.T) {
	d := Default()

	var tests = []struct {
		key   string
		value []byte
	}{
		{
			"Key4", []byte{1, 1, 0},
		},
		{
			"Key5", []byte{1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Key: %s", tt.key), func(t *testing.T) {
			d.Update(tt.key, tt.value)

			v, _ := d.Read(tt.key)
			if !reflect.DeepEqual(v, tt.value) {
				t.Fatalf("got: %v, expected: %v", v, tt.value)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	d := Default()

	var tests = []struct {
		key    string
		length int
	}{
		{
			"Key4", 3,
		},
		{
			"Key3", 2,
		},
		{
			"Key2", 1,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Key: %s", tt.key), func(t *testing.T) {
			d.Delete(tt.key)

			if d.Length != tt.length {
				t.Fatalf("Current length: %d, expected length: %d", d.Length, tt.length)
			}
		})
	}

}
