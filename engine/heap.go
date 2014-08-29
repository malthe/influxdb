package engine

import "github.com/influxdb/influxdb/protocol"

// A data structure that keeps track of the point with the maximum (or
// minimum) timestamp. We keep the points in Series objects to
// preserve the name and fields. But it's understood that each series
// will have one point only.
type Heap interface {
	Size() int
	Add(int, *protocol.Series)
	Next() (int, *protocol.Series)
}
