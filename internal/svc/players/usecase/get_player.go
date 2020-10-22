package usecase

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/captain-corgi/learn-go-with-test/internal/svc/players/repository"
)

type (
	//PlayerServer represent for all get player usecases
	PlayerServer interface {
		ServeHTTP(rw http.ResponseWriter, req *http.Request)
	}
	//PlayerServerImpl is implementation of PlayerServer
	PlayerServerImpl struct {
		playerRepo repository.PlayerRepo
	}
)

//NewPlayerServer return a new player usecase
func NewPlayerServer(playerRepo repository.PlayerRepo) *PlayerServerImpl {
	return &PlayerServerImpl{
		playerRepo: playerRepo,
	}
}

//ServeHTTP is http handler of players
func (r *PlayerServerImpl) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	player := strings.TrimPrefix(req.URL.Path, "/players/")
	playerScore := r.playerRepo.GetPlayerScore(player)
	fmt.Fprint(rw, playerScore)
}
