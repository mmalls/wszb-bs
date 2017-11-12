package model

import (
	"time"
)

// Custom table
type Custom struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    int       `json:"userId"`
	Weixin    string    `json:"weixin"`
	Phone     string    `gorm:"unique" json:"phone"`
	Address   string    `json:"address"`
	PostCode  string    `json:"postCode"`
	Notes     string    `gorm:"size:512" json:"notes"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
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
