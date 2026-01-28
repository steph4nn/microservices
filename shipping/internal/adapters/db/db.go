package db

import (
	"context"
	"fmt"

	"github.com/steph4nn/microservices/shipping/internal/application/domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Shipping struct {
	gorm.Model
	OrderID      int64
	Status       string
	DeliveryDays int32
	Items        []ShippingItem
}

type ShippingItem struct {
	gorm.Model
	ProductCode string
	Quantity    int32
	ShippingID  uint
}

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dataSourceUrl string) (*Adapter, error) {
	db, openErr := gorm.Open(mysql.Open(dataSourceUrl), &gorm.Config{})
	if openErr != nil {
		return nil, fmt.Errorf("db connection error: %v", openErr)
	}

	err := db.AutoMigrate(&Shipping{}, &ShippingItem{})
	if err != nil {
		return nil, fmt.Errorf("db migration error: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (a Adapter) Get(ctx context.Context, id string) (domain.Shipping, error) {
	var shippingEntity Shipping

	res := a.db.WithContext(ctx).Preload("Items").First(&shippingEntity, id)

	items := make([]domain.ShippingItem, 0, len(shippingEntity.Items))
	for _, item := range shippingEntity.Items {
		items = append(items, domain.ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	shipping := domain.Shipping{
		ID:           int64(shippingEntity.ID),
		OrderID:      shippingEntity.OrderID,
		Status:       shippingEntity.Status,
		DeliveryDays: shippingEntity.DeliveryDays,
		Items:        items,
		CreatedAt:    shippingEntity.CreatedAt.UnixNano(),
	}

	return shipping, res.Error
}

func (a Adapter) Save(ctx context.Context, shipping *domain.Shipping) error {
	items := make([]ShippingItem, 0, len(shipping.Items))
	for _, item := range shipping.Items {
		items = append(items, ShippingItem{
			ProductCode: item.ProductCode,
			Quantity:    item.Quantity,
		})
	}

	shippingModel := Shipping{
		OrderID:      shipping.OrderID,
		Status:       shipping.Status,
		DeliveryDays: shipping.DeliveryDays,
		Items:        items,
	}

	res := a.db.WithContext(ctx).Create(&shippingModel)
	if res.Error == nil {
		shipping.ID = int64(shippingModel.ID)
		shipping.CreatedAt = shippingModel.CreatedAt.UnixNano()
	}
	return res.Error
}
