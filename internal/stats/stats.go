package stats

import (
	"reflect"
	"runtime"
	"strconv"
)

type Stat struct {
	ID    string
	Type  string
	Value string
}

type RuntimeMetrics struct{}

var memStatsKeys = [...]string{
	"Alloc",
	"TotalAlloc",
	"Sys",
	"Mallocs",
	"Frees",
	"HeapAlloc",
	"HeapSys",
	"HeapIdle",
	"GCSys",
	"OtherSys",
	"NextGC",
	"LastGC",
	"PauseTotalNs",
	"PauseNs",
	"PauseEnd",
	"NumGC",
	"NumForcedGC",
	"GCCPUFractionHeapInuse",
	"HeapReleased",
	"HeapObjects",
	"StackInuse",
	"StackSys",
	"MSpanInuse",
	"MSpanSys",
	"MCacheInuse",
	"MCacheSys",
	"BuckHashSys",
	"GCSys",
	"OtherSys",
	"NextGC",
	"LastGC",
	"PauseTotalNs",
	"PauseNs",
	"PauseEnd",
	"NumGC",
	"NumForcedGC",
	"GCCPUFraction",
}

func (rm *RuntimeMetrics) Read() []Stat {
	memstats := &runtime.MemStats{}

	runtime.ReadMemStats(memstats)

	var stats []Stat

	for _, memStatsKey := range memStatsKeys {
		r := reflect.ValueOf(memStatsKey)
		f := reflect.Indirect(r).FieldByName(memStatsKey)

		value := f.Uint()

		stats = append(stats, Stat{
			ID:    memStatsKey,
			Value: strconv.FormatUint(value, 10),
			Type:  "counter",
		})
	}

	return stats
}
