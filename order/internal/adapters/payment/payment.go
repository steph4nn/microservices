package payment

import (
	"context"
	"log"
	"time"

	pb "github.com/ruandg/microservices-proto/golang/payment"
	"github.com/ruandg/microservices/order/internal/application/core/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
)

type Adapter struct {
	payment pb.PaymentClient // comes from generated protobuf code
}

func NewAdapter(paymentServiceUrl string) (*Adapter, error) {
	var opts []grpc.DialOption
	opts = append(opts,
		grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
			grpc_retry.WithCodes(codes.Unavailable, codes.ResourceExhausted),
			grpc_retry.WithMax(5),
			grpc_retry.WithBackoff(grpc_retry.BackoffLinear(time.Second)),
		)))
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
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	
	_, err := a.payment.Create(ctx, &pb.CreatePaymentRequest{
		UserId:     order.CustomerID,
		OrderId:   order.ID,
		TotalPrice: order.TotalPrice(),
	})
	
	if status.Code(err) == codes.DeadlineExceeded {
		log.Printf("Payment request exceeded deadline for order ID %d", order.ID)
	}
	
	return err
}