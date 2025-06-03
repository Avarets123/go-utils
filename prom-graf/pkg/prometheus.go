package pkg

import "github.com/prometheus/client_golang/prometheus"

func NewCounter(namespace, name, help string) prometheus.Counter {
	return prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func NewGauge(namespace, name, help string) prometheus.Gauge {
	return prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}

func NewHistogram(namespace, name, help string) prometheus.Histogram {
	return prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Name:      name,
		Help:      help,
	})
}
