package tmpjob

import "context"

type Job interface {
	Run(ctx context.Context) (err error)
	String() string
}
