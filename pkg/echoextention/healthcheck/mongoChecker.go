package healthcheck

import (
	"context"
	"github.com/alperhankendi/devnot-workshop/pkg/log"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type pingpong struct {
	db *mongo.Database
}

func (h *pingpong) Check(ctx context.Context) error {

	err := h.db.Client().Ping(context.TODO(), nil)
	if err != nil {
		log.Logger.Errorf("Mongo: Client Ping error. %v", err)
		return errors.New("Mongo Healty Checker: Client Ping error")
	}

	return nil
}
func CheckMongoDb(db *mongo.Database) Checker {
	return &pingpong{
		db: db,
	}
}
