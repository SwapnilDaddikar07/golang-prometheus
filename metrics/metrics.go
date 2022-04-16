package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var RequestCount = promauto.NewCounter(prometheus.CounterOpts{
	Name:        "total_requests",
	Help:        "gives total number of requests received by this service",
	ConstLabels: nil,
})

var ResponseTime = promauto.NewHistogram(prometheus.HistogramOpts{
	Name:        "response_time",
	Help:        "gives response time for requests",
	ConstLabels: nil,
	Buckets:     []float64{0.001, 0.025, 0.1, 0.75, 1, 2.5, 5, 10},
})

var StatusOK = promauto.NewCounter(prometheus.CounterOpts{
	Name:        "total_requests_with_status_200",
	Help:        "gives total requests which have status 200",
	ConstLabels: nil,
})

var StatusInternalServerError = promauto.NewCounter(prometheus.CounterOpts{
	Name:        "total_requests_with_internal_server_error",
	Help:        "gives total number of requests which failed with status 500",
	ConstLabels: nil,
})
