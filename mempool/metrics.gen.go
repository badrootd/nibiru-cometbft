// Code generated by metricsgen. DO NOT EDIT.

package mempool

import (
	"github.com/go-kit/kit/metrics/discard"
	prometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func PrometheusMetrics(namespace string, labelsAndValues ...string) *Metrics {
	labels := []string{}
	for i := 0; i < len(labelsAndValues); i += 2 {
		labels = append(labels, labelsAndValues[i])
	}
	return &Metrics{
		Size: prometheus.NewGaugeFrom(stdprometheus.GaugeOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "size",
			Help:      "Number of uncommitted transactions in the mempool.",
		}, labels).With(labelsAndValues...),
		TxSizeBytes: prometheus.NewHistogramFrom(stdprometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "tx_size_bytes",
			Help:      "Histogram of transaction sizes in bytes.",

			Buckets: stdprometheus.ExponentialBuckets(1, 3, 7),
		}, labels).With(labelsAndValues...),
		FailedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "failed_txs",
			Help:      "Number of failed transactions.",
		}, labels).With(labelsAndValues...),
		RejectedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "rejected_txs",
			Help:      "Number of rejected transactions.",
		}, labels).With(labelsAndValues...),
		RecheckTimes: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "recheck_times",
			Help:      "Number of times transactions are rechecked in the mempool.",
		}, labels).With(labelsAndValues...),
		AlreadyReceivedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "already_received_txs",
			Help:      "Number of duplicate transaction reception.",
		}, labels).With(labelsAndValues...),
		RequestedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "requested_txs",
			Help:      "Number of requested transactions (WantTx messages).",
		}, labels).With(labelsAndValues...),
		RerequestedTxs: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "rerequested_txs",
			Help:      "Number of re-requested transactions.",
		}, labels).With(labelsAndValues...),
		NoPeerForTx: prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: namespace,
			Subsystem: MetricsSubsystem,
			Name:      "no_peer_for_tx",
			Help:      "Number of times we cannot find a peer for a tx.",
		}, labels).With(labelsAndValues...),
	}
}

func NopMetrics() *Metrics {
	return &Metrics{
		Size:               discard.NewGauge(),
		TxSizeBytes:        discard.NewHistogram(),
		FailedTxs:          discard.NewCounter(),
		RejectedTxs:        discard.NewCounter(),
		RecheckTimes:       discard.NewCounter(),
		AlreadyReceivedTxs: discard.NewCounter(),
		RequestedTxs:       discard.NewCounter(),
		RerequestedTxs:     discard.NewCounter(),
		NoPeerForTx:        discard.NewCounter(),
	}
}
