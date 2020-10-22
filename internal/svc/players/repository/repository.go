package repository

type (
	//PlayerRepo is an interface for player repo
	PlayerRepo interface {
		//GetPlayerScore return player's score
		GetPlayerScore(name string) (score int)
	}
)
