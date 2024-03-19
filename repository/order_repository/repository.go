package order_repository

import (
	"assignment2-golang-hacktiv8/entity"
	"assignment2-golang-hacktiv8/pkg/errs"
)

type Repository interface {
	ReadOrderById(orderId int) (*entity.Order, errs.Error)
	CreateOrderWithItems(orderPayload entity.Order, itemPayload []entity.Item) errs.Error
	ReadOrders() ([]OrderItemMapped, errs.Error)
	UpdateOrder(orderPayload entity.Order, itemPayload []entity.Item) errs.Error
	DeleteOrder(orderId int) errs.Error
}
