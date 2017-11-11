package model

import (
	"time"
)

// Order table
type Order struct {
	ID        int `gorm:"primary_key"`
	UserID    int
	CustomID  int
	GoodsID   int
	SellPrice float32
	CreatedAt time.Time
	UpdatedAt time.Time
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
	return db.Save(c).Error
}

func (c *Order) Get() error {
	return db.Find(c).Error
}

func (c *Order) Delete() error {
	return db.Delete(c).Error
}
