package domain

import (
	"errors"
	"time"
)

type OrderItem struct {
	ProductCode string  `json:"product_code"`
	UnitPrice   float32 `json:"unit_price"`
	Quantity    int32   `json:"quantity"`
}

type Order struct {
	ID         int64       `json:"id"`
	CustomerID int64       `json:"customer_id"`
	Status     string      `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt  int64       `json:"created_at"`
}

var ErrOrderItemLimitExceeded = errors.New("order cannot exceed 50 total items")

func NewOrder(customerID int64, orderItems []OrderItem) Order {
	return Order{
		CreatedAt:  time.Now().Unix(),
		Status:     "Pending",
		CustomerID: customerID,
		OrderItems: orderItems,
	}
}

func (o Order) TotalItemQuantity() int32 {
	var total int32
	for _, item := range o.OrderItems {
		total += item.Quantity
	}
	return total
}

func (o *Order) TotalPrice() float32 {
	var totalPrice float32
	for _, orderItem := range o.OrderItems {
		totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
	}
	return totalPrice
}