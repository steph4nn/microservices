package ports

import (
	"context"

	"github.com/steph4nn/microservices/shipping/internal/application/domain"
)

type DBPort interface {
	Get(ctx context.Context, id string) (domain.Shipping, error)
	Save(ctx context.Context, shipping *domain.Shipping) error
}
