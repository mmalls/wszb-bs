package model

import (
	"time"
)

// Order table
type Order struct {
	ID            int       `gorm:"primary_key"  json:"id"`
	UserID        int       `json:"userId"`
	CustomID      int       `json:"customId"`
	GoodsID       int       `json:"goodsId"`
	SellPrice     float32   `json:"sellPrice"`
	PurchasePrice float32   `json:"purchasePrice"`
	Quantity      int       `json:"quantity"`
	DiscountType  int       `json:"discountType"` // 0:percent, reduce
	Discount      float32   `json:"discount"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type OrderWithInfo struct {
	Order
	CustomWeixin string `gorm:"column:cwx" json:"customWeixin"`
	CustomPhone  string `gorm:"column:cp" json:"customPhone"`
	GoodsName    string `gorm:"column:gn" json:"goodsName"`
}

// TableName ...
func (Order) TableName() string {
	return "wszb_order"
}

func (c *Order) ListByUserID() (rows []OrderWithInfo, err error) {
	//err = db.Where("user_id = ?", c.UserID).Find(&rows).Error
	err = db.Table("wszb_order").
		Select("wszb_order.*, wszb_custom.weixin as cwx, wszb_custom.phone as cp,  wszb_goods.name as gn").
		Where("wszb_order.user_id = ?", c.UserID).
		Joins("left join wszb_custom on wszb_custom.id = wszb_order.custom_id").
		Joins("left join wszb_goods on wszb_goods.id = wszb_order.goods_id").
		Scan(&rows).Error
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
