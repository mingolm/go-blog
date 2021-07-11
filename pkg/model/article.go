package model

import "time"

const (
	articleTableName = "articles"
)

type ArticleType int

const (
	ArticleTypeServer = 0
	ArticleTypeGolang = 1
	ArticleTypeLinux  = 2
)

type ArticleStatus int

const (
	ArticleStatusHide   = 0
	ArticleStatusNormal = 1
	ArticleStatusUp     = 2
)

type Article struct {
	ID        uint64        `gorm:"column:id" json:"id"`
	UserID    uint64        `gorm:"column:user_id" json:"user_id"`
	Type      ArticleType   `gorm:"column:type" json:"type"`
	Title     string        `gorm:"column:title" json:"title"`
	Content   string        `gorm:"column:content" json:"content"`
	Status    ArticleStatus `gorm:"column:status" json:"status"`
	IP        IPv4          `gorm:"column:ip" json:"ip"`
	UpdatedAt time.Time     `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt time.Time     `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time    `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Article) TableName() string {
	return articleTableName
}
