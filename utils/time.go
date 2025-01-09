package utils

import "time"

var (
	TIME_NS = "ns"
	TIME_UC = "us"
	TIME_MS = "ms"
	TIME_S  = "s"
	TIME_M  = "m"
	TIME_H  = "h"
)
var TIME_SHORTS = []string{TIME_NS, TIME_UC, TIME_MS, TIME_S, TIME_M, TIME_H}
var SHORT_TIME = map[string]int{
	TIME_NS: int(time.Nanosecond),
	TIME_UC: int(time.Microsecond),
	TIME_MS: int(time.Millisecond),
	TIME_S:  int(time.Second),
	TIME_M:  int(time.Minute),
	TIME_H:  int(time.Hour),
}
