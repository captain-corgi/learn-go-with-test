package players

import (
	"log"
	"net/http"

	"github.com/captain-corgi/learn-go-with-test/internal/svc/players/repository"
	"github.com/captain-corgi/learn-go-with-test/internal/svc/players/usecase"
)

const (
	port = "8080"
)

//RegisterPlayerSvc register all player routers to main function
func RegisterPlayerSvc() {
	playerRepo := repository.NewPlayerInMemory()

	getPlayerUseCase := usecase.NewPlayerServer(
		playerRepo,
	)

	handler := http.HandlerFunc(getPlayerUseCase.ServeHTTP)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("failed to listening on port %s: %v", port, err)
	}
}
