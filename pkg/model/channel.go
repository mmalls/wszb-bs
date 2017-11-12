package model

import (
	"time"
)

// Channel table
type Channel struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    int       `json:"userId"`
	Name      string    `json:"name"`
	Phone     string    `gorm:"unique" json:"phone"`
	Intro     string    `gorm:"size:512" json:"intro"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName ...
func (Channel) TableName() string {
	return "wszb_channel"
}

func (c *Channel) ListByUserID() (rows []Channel, err error) {
	err = db.Where("user_id = ?", c.UserID).Find(&rows).Error
	return
}

func (c *Channel) Save() error {
	return db.Save(c).Error
}

func (c *Channel) Get() error {
	return db.Find(c).Error
}

func (c *Channel) Delete() error {
	return db.Delete(c).Error
}
