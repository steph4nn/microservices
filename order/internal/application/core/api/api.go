package api

import (
	"github.com/ruandg/microservices/order/internal/application/core/domain"
	"github.com/ruandg/microservices/order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
	shipping ports.ShippingPort
}

func NewApplication(db ports.DBPort, payment ports.PaymentPort, shipping ports.ShippingPort) *Application {
	return &Application{
		db:      db,
		payment: payment,
		shipping: shipping,
	}
}

func (a Application) PlaceOrder(order domain.Order) (domain.Order, error) {
	if order.TotalItemQuantity() > 50 {
		return domain.Order{}, domain.ErrOrderItemLimitExceeded
	}

	err := a.db.Save(&order)
	if err != nil {
		return domain.Order{}, err
	}

	paymentErr := a.payment.Charge(order)
	if paymentErr != nil {
		return domain.Order{}, paymentErr
	}

	_, shippingErr := a.shipping.CreateShipping(order)
	if shippingErr != nil {
		return domain.Order{}, shippingErr
	}

	return order, nil
}