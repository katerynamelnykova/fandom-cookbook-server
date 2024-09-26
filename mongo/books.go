package mongo

import (
	"context"

	"github.com/katerynamelnykova/fandom-cookbook-server/models"
	"go.mongodb.org/mongo-driver/mongo"
)

func (mh *MongoHandler) GetShortFandomsInfo(filter interface{}) ([]*models.ShortFandomInfo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	cur, err := mh.Books.Find(ctx, filter)

	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var result []*models.ShortFandomInfo
	for cur.Next(ctx) {
		book := &models.ShortFandomInfo{}
		er := cur.Decode(book)
		if er != nil {
			return nil, er
		}
		result = append(result, book)
	}
	return result, nil
}

func (mh *MongoHandler) GetOneBook(c *models.Book, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()
	return mh.Books.FindOne(ctx, filter).Decode(c)
}

func (mh *MongoHandler) GetOneFullBook(c *models.FullBook, filter interface{}) error {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()
	return mh.Books.FindOne(ctx, filter).Decode(c)
}

func (mh *MongoHandler) AddBooks(c []interface{}) (*mongo.InsertManyResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), APITimeout)
	defer cancel()

	return mh.Books.InsertMany(ctx, c)
}
