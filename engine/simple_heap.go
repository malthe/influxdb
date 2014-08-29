package engine

import "github.com/influxdb/influxdb/protocol"

type Value struct {
	streamId int
	s        *protocol.Series
}

type SimpleHeap struct {
	less   func(f, s int64) bool
	values []Value
}

func (sh *SimpleHeap) Size() int {
	return len(sh.values)
}

func (sh *SimpleHeap) Add(streamId int, s *protocol.Series) {
	sh.values = append(sh.values, Value{
		s:        s,
		streamId: streamId,
	})
}

func (sh *SimpleHeap) NextIndex() int {
	minIdx := 0
	minTimestamp := sh.values[0].s.Points[0].GetTimestamp()
	for i, v := range sh.values[1:] {
		t := v.s.Points[0].GetTimestamp()
		// log4go.Debug("comparing %d to %d, result: %v\n", t, minTimestamp, sh.less(t, minTimestamp))
		if sh.less(t, minTimestamp) {
			// log4go.Debug("Updating index to %d", i)
			minIdx = i + 1
			minTimestamp = t
		}
	}
	return minIdx
}

func (sh *SimpleHeap) Next() (int, *protocol.Series) {
	idx := sh.NextIndex()
	s := sh.Size()
	v := sh.values[idx]
	// log4go.Debug("Selected value: %d -- %s", v.streamId, v.s)
	sh.values[idx] = sh.values[s-1]
	sh.values = sh.values[:s-1]
	return v.streamId, v.s
}

func NewSimpleMaxHeap() *SimpleHeap {
	return &SimpleHeap{
		less: func(f, s int64) bool {
			return f > s
		},
	}

}
func NewSimpleMinHeap() *SimpleHeap {
	return &SimpleHeap{
		less: func(f, s int64) bool {
			return f < s
		},
	}
}
