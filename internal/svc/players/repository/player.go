package repository

type (
	//PlayerRepoImpl is an implementation of PlayerRepo
	PlayerRepoImpl struct {
	}
)

//NewPlayer return implementation of PlayerRepo
func NewPlayer() *PlayerRepoImpl {
	return &PlayerRepoImpl{}
}

//GetPlayerScore return player's score
func (r *PlayerRepoImpl) GetPlayerScore(name string) (score int) {
	panic("not implemented") // TODO: Implement
}
