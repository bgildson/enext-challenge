package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/bgildson/enext-challenge/api/service"
	"github.com/bgildson/enext-challenge/api/util"
	"github.com/bgildson/enext-challenge/parser"
)

func TestGamesHandler(t *testing.T) {
	games := []*parser.Game{
		{
			ID:         "1",
			TotalKills: 0,
			Players:    []string{},
			Kills:      map[string]int{},
		},
		{
			ID:         "2",
			TotalKills: 11,
			Players:    []string{"Isgalamido", "Mocinha"},
			Kills: map[string]int{
				"Isgalamido": -7,
				"Mocinha":    0,
			},
		},
	}
	game := &parser.Game{
		ID:         "2",
		TotalKills: 4,
		Players:    []string{"player one", "player two"},
		Kills: map[string]int{
			"player one": 1,
			"player two": 3,
		},
	}
	message := util.NewMessage("an error has occurred")

	serviceSuccess := service.NewMockGamesService(
		func() ([]*parser.Game, error) {
			return games, nil
		},
		func(id string) (*parser.Game, error) {
			return game, nil
		},
	)
	serviceFailure := service.NewMockGamesService(
		func() ([]*parser.Game, error) {
			return nil, fmt.Errorf(message.Message)
		},
		func(id string) (*parser.Game, error) {
			return nil, fmt.Errorf(message.Message)
		},
	)

	handlerSuccess := NewGamesHandler(serviceSuccess)
	handlerFailure := NewGamesHandler(serviceFailure)

	t.Run("GetAll", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/games", nil)
		if err != nil {
			t.Errorf("an unexpected error occurred: %v", err)
		}

		t.Run("success", func(t *testing.T) {
			rec := httptest.NewRecorder()

			handlerSuccess.GetAll(rec, req)

			if rec.Result().StatusCode != http.StatusOK {
				t.Errorf(
					"was expecting %d status code, but returns %d",
					http.StatusOK,
					rec.Result().StatusCode,
				)
			}

			b, err := ioutil.ReadAll(rec.Body)
			if err != nil {
				t.Errorf("could not read response content: %v", err)
			}

			var gs []*parser.Game
			if err := json.Unmarshal(b, &gs); err != nil {
				t.Errorf("could not parse response content: %v", err)
			}

			if !reflect.DeepEqual(gs, games) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", games, gs)
			}
		})

		t.Run("failure", func(t *testing.T) {
			rec := httptest.NewRecorder()

			handlerFailure.GetAll(rec, req)

			if rec.Result().StatusCode != http.StatusBadGateway {
				t.Errorf(
					"was expecting %d status code, but returns %d",
					http.StatusBadGateway,
					rec.Result().StatusCode,
				)
			}

			b, err := ioutil.ReadAll(rec.Body)
			if err != nil {
				t.Errorf("could not read response content: %v", err)
			}

			var msg *util.Message
			if err := json.Unmarshal(b, &msg); err != nil {
				t.Errorf("could not decode api response content: %v", err)
			}

			if !reflect.DeepEqual(msg, message) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", message, msg)
			}
		})
	})

	t.Run("GetOne", func(t *testing.T) {
		req, err := http.NewRequest(http.MethodGet, "/games/2", nil)
		if err != nil {
			t.Errorf("an unexpected error occurred: %v", err)
		}

		t.Run("success", func(t *testing.T) {
			rec := httptest.NewRecorder()

			handlerSuccess.GetOne(rec, req)

			if rec.Result().StatusCode != http.StatusOK {
				t.Errorf(
					"was expecting %d status code, but returns %d",
					http.StatusOK,
					rec.Result().StatusCode,
				)
			}

			b, err := ioutil.ReadAll(rec.Body)
			if err != nil {
				t.Errorf("could not read response content: %v", err)
			}

			var g *parser.Game
			if err := json.Unmarshal(b, &g); err != nil {
				t.Errorf("could not parse response content: %v", err)
			}

			if !reflect.DeepEqual(g, game) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", game, g)
			}
		})

		t.Run("failure", func(t *testing.T) {
			rec := httptest.NewRecorder()

			handlerFailure.GetOne(rec, req)

			if rec.Result().StatusCode == http.StatusOK {
				t.Errorf(
					"was expecting an error status code, but returns %d",
					rec.Result().StatusCode,
				)
			}

			b, err := ioutil.ReadAll(rec.Body)
			if err != nil {
				t.Errorf("could not read response content: %v", err)
			}

			var m *util.Message
			if err := json.Unmarshal(b, &m); err != nil {
				t.Errorf("could not parse response content: %v", err)
			}

			if !reflect.DeepEqual(m, message) {
				t.Errorf("was expecting\n%v\nbut returns\n%v\n", message, m)
			}
		})
	})
}
