package tba

import "errors"

type Client struct {
	key string
}

func New(key string) (client Client, err error) {
	if key == "" {
		err := errors.New("Invalid arguments.")
		return client, err
	}
	client = Client{
		key: key,
	}
	return client, err
}
