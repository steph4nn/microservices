package grpc

import (
	"context"
	"fmt"

	"github.com/steph4nn/microservices/shipping/internal/application/domain"
	pb "github.com/steph4nn/microservices-proto/golang/shipping"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a Adapter) CreateShipping(
	ctx context.Context,
	request *pb.CreateShippingRequest,
) (*pb.CreateShippingResponse, error) {
	log.WithContext(ctx).Info("Creating shipping...")

	items := make([]domain.ShippingItem, 0, len(request.GetItems()))
	for _, it := range request.GetItems() {
		items = append(items, domain.ShippingItem{
			ProductCode: it.GetProductCode(),
			Quantity:    it.GetQuantity(),
		})
	}

	newShipping := domain.NewShipping(request.GetOrderId(), items)
	result, err := a.api.CreateShipping(ctx, newShipping)
	code := status.Code(err)
	if code == codes.InvalidArgument {
		return nil, err
	} else if err != nil {
		return nil, status.New(codes.Internal, fmt.Sprintf("failed to create shipping. %v", err)).Err()
	}

	return &pb.CreateShippingResponse{DeliveryDays: result.DeliveryDays, ShippingId: result.ID}, nil
}
