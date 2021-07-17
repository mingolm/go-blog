package model

import "time"

const (
	leaveTableName = "leave"
)

type LeaveStatus int

const (
	LeaveStatusNormal LeaveStatus = 1
	LeaveStatusUp     LeaveStatus = 2
	LeaveStatusHide   LeaveStatus = 3
)

type Leave struct {
	ID          uint64      `gorm:"column:id" json:"id"`
	UserID      uint64      `gorm:"column:user_id" json:"user_id"`
	SenderEmail string      `gorm:"column:sender_email" json:"sender_email"`
	Content     string      `gorm:"column:content" json:"content"`
	Status      LeaveStatus `gorm:"column:status" json:"status"`
	IP          IPv4        `gorm:"column:ip" json:"ip"`
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt   time.Time   `gorm:"column:created_at" json:"created_at"`
	DeletedAt   *time.Time  `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Leave) TableName() string {
	return leaveTableName
}
