// Copyright 2016 Michal Witkowski. All Rights Reserved.
// See LICENSE for licensing terms.

package grpc_prometheus

import (
	"time"

	"google.golang.org/grpc/codes"
)

type serverReporter struct {
	metrics     *ServerMetrics
	rpcType     grpcType
	serviceName string
	methodName  string
	startTime   time.Time
}

func newServerReporter(m *ServerMetrics, rpcType grpcType, fullMethod string) *serverReporter {
	r := &serverReporter{
		metrics: m,
		rpcType: rpcType,
	}
	if r.metrics.reqDurationHistogramEnabled {
		r.startTime = time.Now()
	}
	//r.serviceName, r.methodName = splitMethodName(fullMethod)
	//r.metrics.serverStartedCounter.WithLabelValues(string(r.rpcType), r.serviceName, r.methodName).Inc()
	return r
}

/*func (r *serverReporter) ReceivedMessage() {
	r.metrics.respSizeCounter.WithLabelValues(string(r.rpcType), r.serviceName, r.methodName).Inc()
}

func (r *serverReporter) SentMessage() {
	r.metrics.serverStreamMsgSent.WithLabelValues(string(r.rpcType), r.serviceName, r.methodName).Inc()
}*/

func (r *serverReporter) Handled(code codes.Code, size float64) {
	r.metrics.respSizeCounter.WithLabelValues(string(r.rpcType), code.String(), "METHOD", r.methodName, "false", "").Add(size)
	if r.metrics.reqDurationHistogramEnabled {
		r.metrics.reqDurationHistogram.WithLabelValues(string(r.rpcType), code.String(), "METHOD", r.methodName, "false", "").Observe(time.Since(r.startTime).Seconds())
	}
}
