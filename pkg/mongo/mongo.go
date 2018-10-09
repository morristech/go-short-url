package mongo

import (
	ctx "context"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/findopt"
)

func Connect(config *Config) (*Client, error) {
	var (
		client *mongo.Client
		err    error
	)

	if client, err = mongo.NewClient(config.uri()); err != nil {
		return nil, err
	}

	if err := client.Connect(ctx.Background()); err != nil {
		return nil, err
	}

	return &Client{
		client,
		client.Database(config.Database),
	}, nil
}

func (c *Client) Disconnect() (bool, error) {
	var err error
	if err = c.session.Disconnect(ctx.Background()); err != nil {
		return false, err
	}

	return true, nil
}

func (c *Client) Create(collection string, data interface{}) (*mongo.InsertOneResult, error) {
	result, err := c.db.Collection(collection).InsertOne(ctx.Background(), data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Client) FindOne(collection string, conditions interface{}) *mongo.DocumentResult {

	sort := findopt.Sort(bson.NewDocument(
		bson.EC.Int32("_id", -1),
	))

	return c.db.Collection(collection).FindOne(ctx.Background(), conditions, sort)
}

func (c *Client) Count(collection string, conditions map[string]string) (int64, error) {
	result, err := c.db.Collection(collection).Count(ctx.Background(), conditions)
	if err != nil {
		return 0, err
	}

	return result, err
}

func (c *Client) FindOneAndUpdate(collection string, filter interface{}, update interface{}) *mongo.DocumentResult {
	return c.db.Collection(collection).FindOneAndUpdate(ctx.Background(), filter, update)
}
