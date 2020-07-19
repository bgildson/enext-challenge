package handler

import (
	"net/http"

	"github.com/bgildson/enext-challenge/api/serializer"
	"github.com/bgildson/enext-challenge/api/service"
	"github.com/bgildson/enext-challenge/api/util"
	"github.com/go-chi/chi"
)

// GamesHandler indicates how to implements a new GamesHandler
type GamesHandler interface {
	GetOne(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

type gamesHandler struct {
	service service.GamesService
}

// NewGamesHandler creates a new GamesHandler instance
func NewGamesHandler(service service.GamesService) GamesHandler {
	return &gamesHandler{service}
}

func (h *gamesHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	games, err := h.service.List()
	if err != nil {
		handleFailure(w, http.StatusBadGateway, err)
		return
	}

	s := serializer.NewJSONGamesSerializer()

	b, err := s.Serialize(games)
	if err != nil {
		handleFailure(w, http.StatusBadGateway, err)
		return
	}

	handleSuccess(w, http.StatusOK, b)
}

func (h *gamesHandler) GetOne(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	game, err := h.service.Find(id)
	if err != nil {
		handleFailure(w, http.StatusBadGateway, err)
		return
	}

	s := serializer.NewJSONGameSerializer()

	b, err := s.Serialize(game)
	if err != nil {
		handleFailure(w, http.StatusBadGateway, err)
		return
	}

	handleSuccess(w, http.StatusOK, b)
}

func handleSuccess(w http.ResponseWriter, statusCode int, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(body)
}

func handleFailure(w http.ResponseWriter, statusCode int, err error) {
	w.Header().Set("Content-Type", "application/json")

	msg := util.NewMessage(err.Error())

	s := util.NewJSONMessageSerializer()

	b, err := s.Serialize(msg)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(b)
}
