package movies

import (
	"context"
)

type (
	Service struct {
		repository Repository
	}
)
type Repository interface {
	Get(ctx context.Context, id string) (movie Movie, err error)
	Create(ctx context.Context, item *Movie) (err error)
}

func NewService(repo Repository) *Service {

	return &Service{
		repository: repo,
	}
}
func (receiver *Service) Get(ctx context.Context, id string) (Movie, error) {

	return receiver.repository.Get(ctx, id)
}

func (receiver *Service) Create(ctx context.Context, item *Movie) error {

	return receiver.repository.Create(ctx, item)
}
