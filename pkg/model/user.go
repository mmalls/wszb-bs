package model

import (
	"time"
)

// User table
type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Name      string    `json:"name"`
	Phone     string    `gorm:"unique" json:"phone"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

// TableName ...
func (User) TableName() string {
	return "wszb_user"
}

func (c *User) Save() error {
	return db.Save(c).Error
}

func (c *User) GetByPhone() error {
	return db.Where("phone = ?", c.Phone).Find(c).Error
}

func (c *User) Get() error {
	return db.Find(c).Error
}

func (c *User) Delete() error {
	return db.Delete(c).Error
}

// Auth ...
type Auth struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type LoginLog struct {
	ID        int       `gorm:"primary_key" json:"id"`
	UserID    int       `json:"userId"`
	IP        string    `json:"ip"`
	Device    string    `json:"device"`
	CreatedAt time.Time `json:"createdAt"`
}

// TableName ...
func (LoginLog) TableName() string {
	return "wszb_user_login"
}

func (c *LoginLog) Save() error {
	return db.Save(c).Error
}

func (c *LoginLog) ListByUserID() (rows []LoginLog, err error) {
	err = db.Where("user_id = ?", c.UserID).Find(&rows).Error
	return
}

func (c *LoginLog) Get() error {
	return db.Find(c).Error
}

func (c *LoginLog) Delete() error {
	return db.Delete(c).Error
}
