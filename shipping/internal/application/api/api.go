package api

import (
	"context"

	"github.com/steph4nn/microservices/shipping/internal/application/domain"
	"github.com/steph4nn/microservices/shipping/internal/ports"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Application struct {
	db ports.DBPort
}

func NewApplication(db ports.DBPort) *Application {
	return &Application{db: db}
}

func (a Application) CreateShipping(ctx context.Context, shipping domain.Shipping) (domain.Shipping, error) {
	if shipping.OrderID <= 0 {
		return domain.Shipping{}, status.Errorf(codes.InvalidArgument, "order_id inválido")
	}
	if len(shipping.Items) == 0 {
		return domain.Shipping{}, status.Errorf(codes.InvalidArgument, "lista de itens vazia")
	}

	shipping.DeliveryDays = shipping.CalculateDeliveryDays()
	if shipping.DeliveryDays < 1 {
		return domain.Shipping{}, status.Errorf(codes.InvalidArgument, "delivery_days inválido")
	}

	if err := a.db.Save(ctx, &shipping); err != nil {
		return domain.Shipping{}, err
	}
	return shipping, nil
}
