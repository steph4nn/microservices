package shipping

import (
	"context"
	"log"
	"time"

	shippingpb "github.com/steph4nn/microservices-proto/golang/shipping"
	"github.com/ruandg/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
)

type Adapter struct {
	shipping shippingpb.ShippingClient
}

func NewAdapter(shippingServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		)))
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(shippingServiceUrl, opts...)
	if err != nil {
		log.Fatalf("Failed to create a new shipping adapter. Error: %v", err)
		return nil, err
	}
	client := shippingpb.NewShippingClient(conn)
	return &Adapter{shipping: client}, nil
}

func (a *Adapter) CreateShipping(order domain.Order) (int32, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	items := make([]*shippingpb.ShippingItem, 0, len(order.OrderItems))
	for _, item := range order.OrderItems {
		items = append(items, &shippingpb.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	resp, err := a.shipping.CreateShipping(ctx, &shippingpb.CreateShippingRequest{
		OrderId: order.ID,
		Items:   items,
	})

	if status.Code(err) == codes.DeadlineExceeded {
		log.Printf("Shipping request exceeded deadline for order ID %d", order.ID)
	}

	if err != nil {
		return 0, err
	}

	return resp.GetDeliveryDays(), nil
}
