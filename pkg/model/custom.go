package model

import (
	"time"
)

// Custom table
type Custom struct {
	ID        int `gorm:"primary_key" json:"id"`
	UserID    int `json:"userId"`
	Weixin    string
	Phone     string `gorm:"unique"`
	Address   string
	PostCode  string
	Notes     string `gorm:"size:512"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// TableName ...
func (Custom) TableName() string {
	return "wszb_custom"
}

func (c *Custom) ListByUserID() (rows []Custom, err error) {
	err = db.Where("user_id = ?", c.UserID).Find(&rows).Error
	return
}

func (c *Custom) Save() error {
	return db.Save(c).Error
}

func (c *Custom) Get() error {
	return db.Find(c).Error
}

func (c *Custom) Delete() error {
	return db.Delete(c).Error
}
