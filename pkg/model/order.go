package model

import "time"

const (
	orderTableName = "orders"
)

type OrderType int
type BusCode int
type OrderStatus int

const (
	OrderTypeH5     OrderType = 0
	OrderTypeQRCode OrderType = 1

	BusCodeZFB BusCode = 0
	BusCodeWX  BusCode = 1
	BusCodeYSF BusCode = 2

	OrderStatusWaiting = 0
	OrderStatusSuccess = 1
	OrderStatusFailed  = 2
)

type Order struct {
	ID        uint64      `gorm:"column:id" json:"id"`
	Type      OrderType   `gorm:"type" json:"type"`
	SourceID  string      `gorm:"column:source_id" json:"source_id"`
	OrderID   string      `gorm:"column:order_id" json:"order_id"`
	OrderAmt  float64     `gorm:"column:order_amt" json:"order_amt"`
	BusCode   BusCode     `gorm:"column:bus_code" json:"bus_code"`
	Status    OrderStatus `gorm:"column:status" json:"status"`
	IP        IPv4        `gorm:"column:ip" json:"ip"`
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt time.Time   `gorm:"column:created_at" json:"created_at"`
	DeletedAt *time.Time  `gorm:"column:deleted_at" json:"deleted_at"`
}

func (Order) TableName() string {
	return orderTableName
}
