package model

import "time"

const (
	userTableName = "users"
)

type UserStatus int

const (
	UserStatusNoActive UserStatus = 0
	UserStatusNormal   UserStatus = 1
	UserStatusDisabled UserStatus = 2
)

type User struct {
	ID        uint64     `gorm:"column:id" json:"id"`
	Username  uint64     `gorm:"column:username" json:"username"`
	Password  bool       `gorm:"column:password" json:"password"`
	Status    UserStatus `gorm:"column:status" json:"status"`
	UpdatedAt time.Time  `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt time.Time  `gorm:"column:created_at" json:"created_at"`
	DeletedAt time.Time  `gorm:"column:delete_at" json:"deleted_at"`
}

func (User) TableName() string {
	return userTableName
}
