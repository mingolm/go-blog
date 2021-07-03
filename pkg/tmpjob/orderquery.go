package tmpjob

import (
	"context"
	"errors"
	"github.com/mingolm/go-recharge/pkg/core"
	"github.com/mingolm/go-recharge/pkg/core/driver"
	"github.com/mingolm/go-recharge/pkg/service"
	"github.com/mingolm/go-recharge/utils/errutil"
	"time"
)

func NewOrderQueryJob() Job {
	return &orderQueryJob{
		Service:      core.Instance(),
		OrderService: service.NewOrderService(),
	}
}

type orderQueryJob struct {
	*core.Service
	OrderService service.Order
}

func (j *orderQueryJob) Run(ctx context.Context) (err error) {
	timeTicker := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-ctx.Done():
			timeTicker.Stop()

			j.Logger.Info("order-query-job shutdown")
			j.Logger.Info("bye~")
			return nil
		case <-timeTicker.C: // 5秒钟执行一次僵尸队列处理，保持绝对一致性
			row, err := j.OrderService.Alignment()
			if err != nil {
				if !errors.Is(err, errutil.ErrNotFound) {
					j.Logger.Errorw("order-query-job order service alignment failed",
						"err", err,
					)
				}
				continue
			}
			j.Logger.Errorw("order-query-job has zombie job",
				"row", row,
			)
		default:
			row, err := j.OrderService.GetNextWaiting()
			if err != nil {
				if !errors.Is(err, errutil.ErrNotFound) {
					j.Logger.Errorw("order-query-job order service get waiting failed",
						"err", err,
					)
				}
				time.Sleep(1 * time.Second)
			}

			j.Logger.Infow("order-query-job running",
				"row", row,
			)

			statusOutput, err := j.ThirdDriver.GetOrderStatus(row.SourceID, row.OrderID)
			if err != nil {
				j.Logger.Errorw("order-query-job third driver get order status failed",
					"err", err,
				)
				if err := j.OrderService.Retry(row, false); err != nil {
					j.Logger.Errorw("order-query-job third order service retry failed",
						"err", err,
					)

					if _, err := j.ThirdDriver.CancelOrder(row.SourceID, row.OrderID); err != nil {
						j.Logger.Errorw("order-query-job third third drive cancel failed",
							"err", err,
						)
					}
				}
				time.Sleep(3 * time.Second)
			}
			result := statusOutput.Result

			switch result.State {
			case driver.OrderStateWaiting:
				fallthrough
			case driver.OrderStateWaitingQRCode:
				fallthrough
			case driver.OrderStateWaitingNotice:
				if err := j.OrderService.Retry(row, true); err != nil {
					j.Logger.Errorw("order-query-job third order service retry failed",
						"err", err,
					)
					if _, err := j.ThirdDriver.CancelOrder(row.SourceID, row.OrderID); err != nil {
						j.Logger.Errorw("order-query-job third third drive cancel failed",
							"err", err,
						)
					}
				}
				time.Sleep(1 * time.Second)
			case driver.OrderStateSuccess:
				j.Logger.Info("order-query-job order task success")

				if err := j.OrderRepo.SetSuccessByOrderID(ctx, row.OrderID); err != nil {
					j.Logger.Errorw("order-query-job order set success failed",
						"err", err,
					)
					time.Sleep(1 * time.Second)
				}
			}
		}
	}
}

func (j *orderQueryJob) String() string {
	return "order-query-job"
}
