package serializer

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/bgildson/enext-challenge/parser"
)

func TestJSONGameSerializer(t *testing.T) {
	tt := []struct {
		description string
		in          *parser.Game
		out         string
	}{
		{
			description: "an empty game",
			in: &parser.Game{
				ID:         "1",
				TotalKills: 0,
				Players:    []string{},
				Kills:      map[string]int{},
			},
			out: `{"id": "1", "total_kills": 0, "players": [], "kills": {}}`,
		},
		{
			description: "a normal game",
			in: &parser.Game{
				ID:         "2",
				TotalKills: 4,
				Players:    []string{"player one", "player two"},
				Kills: map[string]int{
					"player one": 1,
					"player two": 3,
				},
			},
			out: `{
  "id": "2",
  "total_kills": 4,
  "players": ["player one", "player two"],
  "kills": {
    "player one": 1,
    "player two": 3
  }
}`,
		},
	}

	s := NewJSONGameSerializer()

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

func TestJSONGamesSerializer(t *testing.T) {
	tt := []struct {
		description string
		in          []*parser.Game
		out         string
	}{
		{
			description: "an empty games list",
			in:          []*parser.Game{},
			out:         `[]`,
		},
		{
			description: "a populated games list",
			in: []*parser.Game{
				{
					ID:         "1",
					TotalKills: 0,
					Players:    []string{},
					Kills:      map[string]int{},
				},
				{
					ID:         "2",
					TotalKills: 4,
					Players:    []string{"player one", "player two"},
					Kills: map[string]int{
						"player one": 1,
						"player two": 3,
					},
				},
			},
			out: `[
  {
    "id": "1",
    "total_kills": 0,
    "players": [],
    "kills": {}
  },
  {
    "id": "2",
    "total_kills": 4,
    "players": ["player one", "player two"],
    "kills": {
      "player one": 1,
      "player two": 3
    }
  }
]`,
		},
	}

	s := NewJSONGamesSerializer()

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
