package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Order table
type Order struct {
	ID             int          `gorm:"primary_key"  json:"id"`
	UserID         int          `json:"userId"`
	CustomID       int          `json:"customId"`
	TotalSellPrice float32      `json:"totalSellPrice"`
	Goods          []OrderGoods `json:"goods"`
	CreatedAt      time.Time    `json:"createdAt"`
	UpdatedAt      time.Time    `json:"updatedAt"`
}

// OrderGoods table
type OrderGoods struct {
	ID            int     `gorm:"primary_key"  json:"id"`
	OrderID       int     `json:"orderId"`
	GoodsID       int     `json:"goodsId"`
	GoodsName     string  `json:"goodsName"`
	Unit          string  `json:"unit"`
	SellPrice     float32 `json:"sellPrice"`
	PurchasePrice float32 `json:"purchasePrice"`
	Quantity      int     `json:"quantity"`
}

type OrderWithCustom struct {
	Order
	Custom Custom `json:"custom"`
}

type OrderStats struct {
	UserID        int          `json:"userId"`
	TotalOrder    int          `json:"totalOrder"`
	TotalIncoming float32      `json:"totalIncoming"`
	TotalGoods    int          `json:"totalGoods"`
	TotalQuantity int          `json:"totalQuantity"`
	TotalCustom   int          `json:"totalCustom"`
	Items         []*OrderItem `json:"items"`
}

type OrderItem struct {
	Key      string  `json:"key"`
	Incoming float32 `json:"incoming"`
}

// TableName ...
func (Order) TableName() string {
	return "wszb_order"
}

// TableName ...
func (OrderGoods) TableName() string {
	return "wszb_order_goods"
}

func (c *Order) ListByUserID(offset, limit int) (rows []OrderWithCustom, err error) {
	err = db.Where("user_id = ?", c.UserID).
		Offset(offset).Limit(limit).
		Order("id DESC").Find(&rows).Error
	if err == nil {
		for idx := range rows {
			row := &rows[idx]
			err = db.Where("order_id = ? ", row.ID).Find(&row.Goods).Error
			err = db.Where("id = ? ", row.CustomID).Find(&row.Custom).Error
		}
	}
	return
}

func (c *Order) Save() error {
	return c.save(db)
}

func (c *Order) save(tx *gorm.DB) error {
	for _, g := range c.Goods {
		if g.Quantity <= 0 {
			g.Quantity = 1
		}
	}

	return tx.Save(c).Error
}

func (c *Order) Update() error {
	tx := db.Begin()
	err := tx.Where("order_id = ?", c.ID).Delete(&OrderGoods{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = c.save(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *OrderWithCustom) Get() error {
	err := db.Model(&Order{}).Find(c).Error
	if err == nil {
		err = db.Where("order_id = ?", c.ID).Find(&c.Goods).Error
		err = db.Where("id = ? ", c.CustomID).Find(&c.Custom).Error
	}

	return err
}

func (c *Order) Delete() error {
	tx := db.Begin()
	err := tx.Where("order_id = ?", c.ID).Delete(&OrderGoods{}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = db.Delete(c).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (c *Order) Query(begin, end time.Time) (rows []OrderWithCustom, err error) {
	err = db.Where("user_id = ?", c.UserID).
		Where("created_at >= ?", begin).
		Where("created_at <= ?", end).
		Order("id DESC").
		Find(&rows).Error
	if err == nil {
		for idx := range rows {
			row := &rows[idx]
			err = db.Where("order_id = ? ", row.ID).Find(&row.Goods).Error
		}
	}
	return
}
