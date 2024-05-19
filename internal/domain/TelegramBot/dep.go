package TelegramBot

import "context"

type AnekdotProvider interface {
	GetJoke(ctx context.Context, category string) (string, error)
}
