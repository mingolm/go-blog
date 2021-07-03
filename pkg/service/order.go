package service

import "encoding/json"

type Order interface {
	// 数据对齐（僵尸队列清理/重试）
	Alignment() (row *OrderTask, err error)
	// 获取待回调任务
	GetNextWaiting() (row *OrderTask, err error)
	// 添加待回调任务
	AddWaiting(row *OrderTask) (err error)
	// ack
	ACK(row *OrderTask) (err error)
	// 任务重试
	Retry(row *OrderTask, force bool) (err error)
}

func NewOrderService() Order {
	return &order{}
}

type OrderTask struct {
	OrderID  string `json:"order_id"`
	SourceID string `json:"source_id"`
}

func (t *OrderTask) Bytes() (bs []byte) {
	bs, _ = json.Marshal(t)
	return bs
}

type order struct {
}

func (s *order) Alignment() (row *OrderTask, err error) {
	return
}

func (s *order) GetNextWaiting() (row *OrderTask, err error) {
	return
}

func (s *order) AddWaiting(row *OrderTask) (err error) {
	return
}

func (s *order) ACK(row *OrderTask) (err error) {
	return
}

func (s *order) Retry(row *OrderTask, force bool) (err error) {
	return
}
