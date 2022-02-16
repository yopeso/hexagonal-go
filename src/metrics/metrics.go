package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	urlCreatedRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "url_created_request",
		Help: "Counter of created URLS",
	})

	urlCreatedSuccessCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "url_created_success",
		Help: "Counter of created URLS",
	})

	urlUsedRequestCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "url_used_request",
		Help: "Counter of used URLS",
	})

	urlUsedSuccessCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "url_used_success",
		Help: "Counter of used URLS",
	})
)

func RecordUrlCreatedRequest() {
	urlCreatedRequestCounter.Inc()
}

func RecordUrlCreatedSuccess() {
	urlCreatedSuccessCounter.Inc()
}

func RecordUrlUsedRequest() {
	urlUsedRequestCounter.Inc()
}

func RecordUrlUseSuccess() {
	urlUsedSuccessCounter.Inc()
}
