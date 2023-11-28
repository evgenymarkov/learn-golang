package storage

import "context"

type Store interface {
	Fetch(ctx context.Context) (string, error)
}
