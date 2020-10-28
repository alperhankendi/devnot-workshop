package movies

import "go.mongodb.org/mongo-driver/mongo"

type (
	repository struct {
		db *mongo.Collection
	}
	Service struct {
		repository
	}
)

func NewService(db *mongo.Collection) *Service {

	return &Service{
		repository{db: db},
	}
}
