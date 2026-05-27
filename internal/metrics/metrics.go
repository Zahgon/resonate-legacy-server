package metrics

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	AioTotal             *prometheus.CounterVec
	AioInFlight          *prometheus.GaugeVec
	AioDuration          *prometheus.HistogramVec
	AioWorker            *prometheus.GaugeVec
	AioWorkerInFlight    *prometheus.GaugeVec
	AioPluginConnections *prometheus.GaugeVec
	ApiTotal             *prometheus.CounterVec
	ApiInFlight          *prometheus.GaugeVec
	ApiDuration          *prometheus.HistogramVec
	CoroutinesTotal      *prometheus.CounterVec
	CoroutinesInFlight   *prometheus.GaugeVec
	CoroutinesDuration   *prometheus.HistogramVec
	HttpRequestsTotal    *prometheus.CounterVec
	HttpRequestsDuration *prometheus.HistogramVec
	PromisesTotal        *prometheus.CounterVec
	SchedulesTotal       *prometheus.CounterVec
	TasksTotal           *prometheus.CounterVec
}

func New(reg prometheus.Registerer) *Metrics { _ = "STUB: not implemented"; return nil }

func (m *Metrics) Enable(reg prometheus.Registerer) { _ = "STUB: not implemented"; return }

func (m *Metrics) Disable(reg prometheus.Registerer) { _ = "STUB: not implemented"; return }
