package model

import (
	"time"
)

// Goods table
type Goods struct {
	ID            int `gorm:"primary_key" json:"id"`
	UserID        int `json:"userId"`
	ChannelID     int `json:"channelId"`
	Name          string
	Catalog       string
	Intro         string `gorm:"size:512"`
	SellPrice     float32
	PurchasePrice float32
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// TableName ...
func (Goods) TableName() string {
	return "wszb_goods"
}

func (c *Goods) ListByUserID() (rows []Goods, err error) {
	err = db.Where("user_id = ?", c.UserID).Find(&rows).Error
	return
}

func (c *Goods) Save() error {
	return db.Save(c).Error
}

func (c *Goods) Get() error {
	return db.Find(c).Error
}

func (c *Goods) Delete() error {
	return db.Delete(c).Error
}
