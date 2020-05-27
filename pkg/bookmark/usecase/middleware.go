package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/website-pribadi/pkg/bookmark/domain/entity"
	"github.com/website-pribadi/pkg/bookmark/interface/model"
)

type Middleware func(usecase Usecase) Usecase

type loggingMiddleware struct {
	logger    log.Logger
	histogram *prometheus.HistogramVec
	next      Usecase
}

func NewLoggingInterceptor(logger log.Logger, histogram *prometheus.HistogramVec) Middleware {
	return func(next Usecase) Usecase {
		return &loggingMiddleware{
			logger:    logger,
			histogram: histogram,
			next:      next,
		}
	}
}

func (l loggingMiddleware) CreateReferenceWithTopic(ctx context.Context, req model.ReferenceTopicRequest) (topic model.BaseResponse, err error) {
	start := time.Now()
	defer func() { // anonymous defer func, there are three different defer func : https://www.geeksforgeeks.org/defer-keyword-in-golang/
		latency := time.Since(start)
		l.histogram.WithLabelValues(fmt.Sprintf("%v", topic.Data)).Observe(latency.Seconds())
		// level.Info(l.logger).Log("req : ", req.ID, "result : ", topic.Data)
	}()
	return l.next.CreateReferenceWithTopic(ctx, req)
}

func (l loggingMiddleware) CreateReference(ctx context.Context, req model.ReferenceRequest) (model.BaseResponse, error) {
	defer func() { // anonymous defer func, there are three different defer func : https://www.geeksforgeeks.org/defer-keyword-in-golang/
		level.Info(l.logger).Log("req : ", req, "result : ", entity.Topic{}.ID)
	}()
	return l.next.CreateReference(ctx, req)
}
