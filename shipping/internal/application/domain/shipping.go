package domain

import (
	"time"
)

type ShippingItem struct {
	ProductCode string `json:"product_code"`
	Quantity    int32  `json:"quantity"`
}

type Shipping struct {
	ID           int64          `json:"id"`
	OrderID      int64          `json:"order_id"`
	Status       string         `json:"status"`
	DeliveryDays int32          `json:"delivery_days"`
	Items        []ShippingItem `json:"items"`
	CreatedAt    int64          `json:"created_at"`
}

func NewShipping(orderID int64, items []ShippingItem) Shipping {
	return Shipping{
		CreatedAt: time.Now().Unix(),
		Status:    "Pending",
		OrderID:   orderID,
		Items:     items,
	}
}

func (s *Shipping) CalculateDeliveryDays() int32 {
    var totalUnits int32
    for _, item := range s.Items {
        totalUnits += item.Quantity
    }

	if totalUnits <= 0 {
		return 0
	}
	
	deliveryDays := int32(1) + ((totalUnits - 1) / 5)
	return deliveryDays
}