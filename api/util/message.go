package util

import (
	"encoding/json"
	"fmt"
)

// Message wraps an api message
type Message struct {
	Message string `json:"message"`
}

// NewMessage creates a new Game instance
func NewMessage(message string) *Message {
	return &Message{message}
}

// MessageSerializer indicates how to implements a new MessageSerializer
type MessageSerializer interface {
	Serialize(message *Message) ([]byte, error)
}

type jsonMessageSerializer struct{}

// NewJSONMessageSerializer creates a new instance of MessageSerializer for json serialization
func NewJSONMessageSerializer() MessageSerializer {
	return &jsonMessageSerializer{}
}

func (s *jsonMessageSerializer) Serialize(message *Message) ([]byte, error) {
	b, err := json.Marshal(message)
	if err != nil {
		return nil, fmt.Errorf("could not serialize message: %v", err)
	}
	return b, nil
}
