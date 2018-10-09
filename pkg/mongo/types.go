package mongo

import (
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Client struct {
	session *mongo.Client
	db      *mongo.Database
}
