//Small_HG
package services

import (
	"product/datamodels"
	"product/repositories"
)

type IOrderService interface {
	GetOrderByID(int64) (*datamodels.Order,error)
	DeleteOrderByID(int64) bool
	UpdateOrder(*datamodels.Order) error
	InsertOrder(*datamodels.Order) (int64 ,error)
	GetAllOrder()([]*datamodels.Order,error)
	GetAllOrderInfo()(map[int]map[string]string,error)
}

func NewOrderService(repository repositories.IOrderRepository ) IOrderService  {
	return &OrderService{OrderRepository:repository}
}

type OrderService struct {
	OrderRepository repositories.IOrderRepository
}

func (o *OrderService)GetOrderByID(orderID int64) (order *datamodels.Order,err error)  {
	return o.OrderRepository.SelectByKey(orderID)
}

func (o *OrderService)DeleteOrderByID(orderID int64) (isOk bool)  {
	isOk = o.OrderRepository.Delete(orderID)
	return
}

func (o *OrderService)UpdateOrder(order *datamodels.Order) error{
	return o.OrderRepository.Update(order)
}

func (o *OrderService)InsertOrder(order *datamodels.Order) (orderID int64,err error)  {
	return o.OrderRepository.Insert(order)
}

func (o *OrderService) GetAllOrder()([]*datamodels.Order,error) {
	return o.OrderRepository.SelectAll()
}

func (o *OrderService) GetAllOrderInfo()(map[int]map[string]string,error) {
	return o.OrderRepository.SelectAllWithInfo()
}



