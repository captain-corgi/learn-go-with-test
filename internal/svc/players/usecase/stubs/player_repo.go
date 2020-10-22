package stubs

type StubPlayerStore struct {
	scores map[string]int
}

func NewStubPlayerStore(scores map[string]int) *StubPlayerStore {
	return &StubPlayerStore{scores: scores}
}

func (r *StubPlayerStore) GetPlayerScore(name string) (score int) {
	score = r.scores[name]
	return
}
