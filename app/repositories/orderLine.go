package repositories

import (
	"errors"

	"github.com/jinzhu/copier"
	"gorm.io/gorm"

	"goshop/app/models"
	"goshop/app/serializers"
	"goshop/dbs"
)

type IOrderLineRepository interface {
	GetOrderLines(query *serializers.OrderLineQueryParam) (*[]models.OrderLine, error)
	GetOrderLineByID(uuid string) (*models.OrderLine, error)
	CreateOrderLine(item *serializers.OrderLineBodyParam) (*models.OrderLine, error)
	UpdateOrderLine(uuid string, item *serializers.OrderLineBodyParam) (*models.OrderLine, error)
}

type OrderLineRepo struct {
	db *gorm.DB
}

func NewOrderLineRepository() *OrderLineRepo {
	return &OrderLineRepo{db: dbs.Database}
}

func (l *OrderLineRepo) GetOrderLines(query *serializers.OrderLineQueryParam) (*[]models.OrderLine, error) {
	var orderLines []models.OrderLine
	if err := l.db.Find(&orderLines, query).Error; err != nil {
		return nil, nil
	}

	return &orderLines, nil
}

func (l *OrderLineRepo) GetOrderLineByID(uuid string) (*models.OrderLine, error) {
	var orderLine models.OrderLine
	if err := l.db.Where("uuid = ?", uuid).First(&orderLine).Error; err != nil {
		return nil, errors.New("not found orderLine")
	}

	return &orderLine, nil
}

func (l *OrderLineRepo) CreateOrderLine(item *serializers.OrderLineBodyParam) (*models.OrderLine, error) {
	var orderLine models.OrderLine
	copier.Copy(&orderLine, &item)

	if err := l.db.Create(&orderLine).Error; err != nil {
		return nil, err
	}

	return &orderLine, nil
}

func (l *OrderLineRepo) UpdateOrderLine(uuid string, item *serializers.OrderLineBodyParam) (*models.OrderLine, error) {
	orderLine, err := l.GetOrderLineByID(uuid)
	if err != nil {
		return nil, err
	}

	copier.Copy(orderLine, &item)
	if err := l.db.Save(&orderLine).Error; err != nil {
		return nil, err
	}

	return orderLine, nil
}
