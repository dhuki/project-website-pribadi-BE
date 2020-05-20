package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/website-pribadi/pkg/topic/domain/entity"
	"github.com/website-pribadi/pkg/topic/interface/transport"
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

func (l loggingMiddleware) CreateTopic(ctx context.Context, req transport.Request) (topic entity.Topic, err error) {
	start := time.Now()
	defer func() { // anonymous defer func, there are three different defer func : https://www.geeksforgeeks.org/defer-keyword-in-golang/
		latency := time.Since(start)
		l.histogram.WithLabelValues(fmt.Sprintf("%v", topic.ID)).Observe(latency.Seconds())
		level.Info(l.logger).Log("req : ", req.ID, "result : ", topic.ID)
	}()
	return l.next.CreateTopic(ctx, req)
}

func (l loggingMiddleware) ListTopic(ctx context.Context) ([]entity.Topic, error) {
	defer func() { // anonymous defer func, there are three different defer func : https://www.geeksforgeeks.org/defer-keyword-in-golang/
		level.Info(l.logger).Log("result : ", entity.Topic{}.ID)
	}()
	return l.next.ListTopic(ctx)
}

func (l loggingMiddleware) GetById(ctx context.Context, req transport.Request) (entity.Topic, error) {
	defer func() { // anonymous defer func, there are three different defer func : https://www.geeksforgeeks.org/defer-keyword-in-golang/
		level.Info(l.logger).Log("req : ", req, "result : ", entity.Topic{}.ID)
	}()
	return l.next.GetById(ctx, req)
}
