package database

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

// Database indicates how to implements a Database
type Database interface {
	Get() ([]map[string]interface{}, error)
	GetByID(id string) (map[string]interface{}, error)
}

// Create generic errors for futures comparations
var (
	ErrDatabaseFileNotFound               = errors.New("could not find the database file")
	ErrCouldNotDeserializeDatabaseContent = errors.New("could not deserialize database file")
	ErrGameNotFound                       = errors.New("could not found game")
)

type jsonDatabase struct {
	data map[string]map[string]interface{}
}

// NewJSONDatabase creates a new Database implementation for JSON source
// func NewJSONDatabase(gamesJSONPath string) (*Database, error) {
func NewJSONDatabase(gamesJSONPath string) (Database, error) {
	b, err := ioutil.ReadFile(gamesJSONPath)
	if err != nil {
		fmt.Println(err)
		return nil, ErrDatabaseFileNotFound
	}

	var data map[string]map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil, ErrCouldNotDeserializeDatabaseContent
	}

	return &jsonDatabase{data}, nil
}

func (d *jsonDatabase) Get() ([]map[string]interface{}, error) {
	var r []map[string]interface{}
	for _, v := range d.data {
		r = append(r, v)
	}
	return r, nil
}

func (d *jsonDatabase) GetByID(id string) (map[string]interface{}, error) {
	if g, ok := d.data[id]; ok {
		return g, nil
	}
	return nil, ErrGameNotFound
}
