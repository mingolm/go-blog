package main

import (
	"context"
	"github.com/mingolm/go-recharge/pkg/tmpjob"
)

func newOrderQueryJob() tmpjob.Job {
	return tmpjob.NewOrderQueryJob()
}

func runJob(ctx context.Context, job tmpjob.Job) {
	jobDone := make(chan struct{})
	go func() {
		if err := job.Run(ctx); err != nil {
			logger.Fatalw("run job fail",
				"err", err,
			)
		}
		jobDone <- struct{}{}
	}()

	logger.Infow("job running",
		"job", job.String(),
	)
	<-jobDone

	logger.Info("bye!")
}
