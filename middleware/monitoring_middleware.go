package middleware

import (
	"github.com/gin-gonic/gin"
	"go-prometheus/metrics"
	"time"
)

func MonitorMetrics(urlsExcludedFromMonitoring []string) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		startTime := time.Now()
		ctx.Next()

		incomingRequestUrl := ctx.Request.URL.Path
		for _, v := range urlsExcludedFromMonitoring {
			if incomingRequestUrl == v {
				return
			}
		}
		applyRequestCountMetric()
		applyDurationMetric(startTime)
		applyStatusCodeMetric(ctx)
	}
}

func applyRequestCountMetric() {
	metrics.RequestCount.Inc()
}

func applyDurationMetric(startTime time.Time) {
	duration := time.Since(startTime)
	metrics.ResponseTime.Observe(duration.Seconds())
}

func applyStatusCodeMetric(ctx *gin.Context) {
	statusCode := ctx.Writer.Status()
	if statusCode == 200 {
		metrics.StatusOK.Inc()
	} else if statusCode == 500 {
		metrics.StatusInternalServerError.Inc()
	}
}
