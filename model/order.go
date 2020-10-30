package model

import (
	"errors"
)

type Order struct {
	ID int64 `gorm:"id" json:"id"`
	UserName string `gorm:"user_name" json:"user_name"`
	Goods string `gorm:"goods" json:"goods"`
}

func (o *Order)GetOrderInfo() error{
	if o.ID == 0 {
		return errors.New("缺少订单id参数")
	}
	err := Db.Table("order").Where("id = ?", o.ID).First(&o).Error
	if err != nil {
		return err
	}
	return  nil
}