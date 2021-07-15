package tpoller_server

import (
	tstorage_pb "github.com/bartmika/tstorage-server/proto"
)

// Function will convert the data structure from `TelemetryLabel` to `Label`
// so that `tstorage-server` can accept the values.
func ToLabels(polledLabels []*TelemetryLabel) []*tstorage_pb.Label {
	labels := make([]*tstorage_pb.Label, len(polledLabels))
	for _, v := range polledLabels {
		label := &tstorage_pb.Label{
			Name:  v.Name,
			Value: v.Value,
		}
		labels = append(labels, label)
	}
	return labels
}

// Function will convert the data structure so that `tstorage-server` can
// accept the time-series data.
func ToTimeSeriesDatum(d *TelemetryDatum) *tstorage_pb.TimeSeriesDatum {
	return &tstorage_pb.TimeSeriesDatum{
		Labels:    ToLabels(d.Labels),
		Metric:    d.Metric,
		Value:     d.Value,
		Timestamp: d.Timestamp,
	}
}
