package repository

type (
	//PlayerInMemoryRepoImpl is in-memory implementation of PlayerRepo
	PlayerInMemoryRepoImpl struct {
	}
)

//NewPlayerInMemory return implementation of PlayerRepo
func NewPlayerInMemory() *PlayerInMemoryRepoImpl {
	return &PlayerInMemoryRepoImpl{}
}

//GetPlayerScore return player's score
func (r *PlayerInMemoryRepoImpl) GetPlayerScore(name string) (score int) {
	return 0
}
