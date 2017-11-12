package model

import (
	"time"
)

// Order table
type Order struct {
	ID           int       `gorm:"primary_key"  json:"id"`
	UserID       int       `json:"userId"`
	CustomID     int       `json:"customId"`
	GoodsID      int       `json:"goodsId"`
	SellPrice    float32   `json:"sellPrice"`
	Quantity     int       `json:"quantity"`
	DiscountType int       `json:"discountType"` // 0:percent, reduce
	Discount     float32   `json:"discount"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

// TableName ...
func (Order) TableName() string {
	return "wszb_order"
}

func (c *Order) ListByUserID() (rows []Order, err error) {
	err = db.Where("user_id = ?", c.UserID).Find(&rows).Error
	return
}

func (c *Order) Save() error {
	if c.Quantity <= 0 {
		c.Quantity = 1
	}
	return db.Save(c).Error
}

func (c *Order) Get() error {
	return db.Find(c).Error
}

func (c *Order) Delete() error {
	return db.Delete(c).Error
}
