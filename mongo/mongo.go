package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DefaultDatabase = "fandom-coo-main-db-07cc1b8f378"
const BooksCollectionName = "books"
const APITimeout = 1000 * time.Second

type MongoHandler struct {
	client   *mongo.Client
	database string

	Books *mongo.Collection
}

// MongoHandler Constructor
func NewHandler(address string) (*MongoHandler, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()
	cl, err := mongo.Connect(ctx, options.Client().ApplyURI(address))

	if err != nil {
		return nil, err
	}

	mh := &MongoHandler{
		client:   cl,
		database: DefaultDatabase,
	}

	mh.Books = mh.client.Database(mh.database).Collection(BooksCollectionName)

	return mh, nil
}
