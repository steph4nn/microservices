package ports

import (
	"context"

	"github.com/steph4nn/microservices/shipping/internal/application/domain"
)

type APIPort interface {
	CreateShipping(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error)
}
