package mongo

import (
	"context"
	m "github.com/mongodb/mongo-go-driver/mongo"
)

type Client struct {
	session *m.Client
	db      *m.Database
}

func Connect(config *Config) (*Client, error) {
	var (
		client *m.Client
		err    error
	)

	if client, err = m.NewClient(config.uri()); err != nil {
		return nil, err
	}

	if err := client.Connect(context.Background()); err != nil {
		return nil, err
	}

	return &Client{
		client,
		client.Database(config.Database),
	}, nil
}

func (c *Client) Disconnect() (bool, error) {
	var err error
	if err = c.session.Disconnect(context.Background()); err != nil {
		return false, err
	}

	return true, nil
}
