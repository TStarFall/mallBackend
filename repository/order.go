package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mall/model"
	"mall/query"
	"mall/utils"
)

type OrderRepository struct {
	DB *gorm.DB
}

type OrderRepoInterface interface {
	List(req *query.ListQuery) (orders []*model.Order, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(order model.Order) (*model.Order,error)
	Exist(order model.Order) *model.Order
	ExistByOrderID(id string) *model.Order
	Add(order model.Order) (*model.Order, error)
	Edit(order model.Order) (bool, error)
	Delete(order model.Order) (bool, error)
}

func (repo *OrderRepository) List(req *query.ListQuery) (orders []*model.Order, err error) {
	db := repo.DB
	limit, offset := utils.Page(req.PageSize, req.Page)
	if err = db.Limit(limit).Offset(offset).Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil
}

func (repo *OrderRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var orders []model.Order
	db := repo.DB
	if err = db.Find(&orders).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *OrderRepository) Get(order model.Order) (*model.Order, error) {
	if err := repo.DB.Where(&order).Find(&order).Error; err != nil {
		return nil, err
	}
	return &order, nil
}

func (repo *OrderRepository) Exist(order model.Order) *model.Order {
	if order.OrderId != "" {
		repo.DB.Model(&order).Where("order_id = ?", order.OrderId)
		return &order
	}
	return nil
}

func (repo *OrderRepository) ExistByOrderID(id string) *model.Order {
	var order model.Order
	repo.DB.Where("order_id = ?", id).First(&order)
	return &order
}

func (repo *OrderRepository) Add(order model.Order) (*model.Order, error) {
	err := repo.DB.Create(order).Error
	if err != nil {
		return &order, fmt.Errorf("订单添加失败")
	}
	return &order, nil
}

func (repo *OrderRepository) Edit(order model.Order) (bool, error) {
	if order.OrderId == "" {
		return false, fmt.Errorf("请输入更新ID")
	}
	o := &model.Order{}
	err := repo.DB.Model(o).Where("order_id = ?", order.OrderId).Updates(map[string]interface{}{
		"nick_name" : order.NickName,
		"mobile" : order.Mobile,
		"pay_status" : order.PayStatus,
		"order_status" : order.OrderStatus,
		"extra_info" : order.ExtraInfo,
		"user_address" : order.UserAddress,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *OrderRepository) Delete(order model.Order) (bool, error) {
	err := repo.DB.Model(&order).Where("order_id = ?", order.OrderId).Update("isDeleted", order.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

















