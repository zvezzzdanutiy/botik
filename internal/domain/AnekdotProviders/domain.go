package AnekdotProviders

import "net/http"

type Domain struct {
	client *http.Client
}

func New(client *http.Client) *Domain {
	return &Domain{
		client: client,
	}
}
