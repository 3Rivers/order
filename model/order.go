package model

import (
	"errors"
)

type Order struct {
	ID int64 `gorm:"id" json:"id"`
	Username string `gorm:"user_name" json:"user_name"`
	Goods string `gorm:"goods" json:"goods"`
}

func (o *Order)GetOrderInfo() (*Order,error) {
	if o.ID == 0 {
		return o,errors.New("缺少订单id参数")
	}
	err := Db.First(&o).Error
	if err != nil {
		return o,err
	}
	return  o,nil
}