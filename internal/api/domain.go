package api

type Domain struct {
	JokeProvider JokeProvider
}

func New(JokeProvider JokeProvider) *Domain {
	return &Domain{
		JokeProvider: JokeProvider,
	}
}
