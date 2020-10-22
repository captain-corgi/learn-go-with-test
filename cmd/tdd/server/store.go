package server

import "context"

//Store is a simple in-memory storage
type Store interface {
	Fetch(ctx context.Context) (string, error)
}
