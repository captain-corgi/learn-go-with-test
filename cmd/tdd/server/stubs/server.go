package stubs

import (
	"context"
	"testing"
	"time"
)

//SpyStore is a stub store
type SpyStore struct {
	response string
	t        *testing.T
}

//NewSpyStore return a new stub store pointer
func NewSpyStore(res string, t *testing.T) *SpyStore {
	return &SpyStore{res, t}
}

//Fetch return value of store
func (r *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range r.response {
			select {
			case <-ctx.Done():
				r.t.Log("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}
