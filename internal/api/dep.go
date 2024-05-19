package api

import "context"

type JokeProvider interface {
	GetJoke(ctx context.Context, category string) (string, error)
}
