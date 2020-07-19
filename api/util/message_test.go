package util

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestJSONMessageSerializer(t *testing.T) {
	tt := []struct {
		description string
		in          *Message
		out         string
	}{
		{
			description: "a normal message",
			in:          NewMessage("an error has occurred"),
			out:         `{"message": "an error has occurred"}`,
		},
	}

	s := NewJSONMessageSerializer()

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			var outData interface{}
			err := json.Unmarshal([]byte(tc.out), &outData)
			if err != nil {
				t.Errorf("an unexpected error occurred: %v", err)
			}

			r, err := s.Serialize(tc.in)
			if err != nil {
				t.Errorf("an unexpected error occurred: %v", err)
			}

			var inData interface{}
			_ = json.Unmarshal(r, &inData)

			if !reflect.DeepEqual(inData, outData) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", outData, inData)
			}
		})
	}
}
