package model

import (
	"time"
)

// Goods table
type Goods struct {
	ID            int       `gorm:"primary_key" json:"id"`
	UserID        int       `json:"userId"`
	ChannelID     int       `json:"channelId"`
	Name          string    `json:"name"`
	Catalog       string    `json:"catalog"`
	Intro         string    `gorm:"size:512" json:"intro"`
	SellPrice     float32   `json:"sellPrice"`
	PurchasePrice float32   `json:"purchasePrice"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type GoodsWitchChl struct {
	Goods
	ChannelName string `gorm:"column:cn" json:"channelName"`
}

// TableName ...
func (Goods) TableName() string {
	return "wszb_goods"
}

func (c *Goods) ListByUserID() (rows []GoodsWitchChl, err error) {
	//err = db.Joins("user_id = ?", c.UserID).Find(&rows).Error
	err = db.Table("wszb_goods").
		Select("wszb_goods.*, wszb_channel.name as cn").
		Where("wszb_goods.user_id = ?", c.UserID).
		Joins("left join wszb_channel on wszb_channel.id = wszb_goods.channel_id").
		Scan(&rows).Error
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
