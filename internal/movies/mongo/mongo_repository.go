package mongo

import (
	"context"
	. "github.com/alperhankendi/devnot-workshop/internal/movies"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const collectionName = "movies"

type repository struct {
	db *mongo.Collection
}

func NewRepository(db *mongo.Database) Repository {

	return &repository{
		db: db.Collection(collectionName),
	}
}
func (receiver *repository) Get(ctx context.Context, id string) (movie Movie, err error) {

	err = receiver.db.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
	return
}

func (receiver *repository) Create(ctx context.Context, item *Movie) (err error) {
	_, err = receiver.db.InsertOne(ctx, item)
	return
}
