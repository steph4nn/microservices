package payment

import (
	"context"
	"log"

	pb "github.com/ruandg/microservices-proto/golang/payment"
	"github.com/ruandg/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Adapter struct {
	payment pb.PaymentClient // comes from generated protobuf code
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial(paymentServiceUrl, opts...)
	if err != nil {
		log.Fatalf("Failed to create a new adapter adapter.Error: %v", err)
		return nil, err
	}
	client := pb.NewPaymentClient(conn) // initialize the stub
	return &Adapter{payment: client}, nil
}

func (a *Adapter) Charge(order domain.Order) error {
	_, err := a.payment.Create(context.Background(), &pb.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:   order.ID,
		TotalPrice: order.TotalPrice(),
	})
	return err
}