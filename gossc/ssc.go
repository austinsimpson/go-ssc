package gossc

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L../ssc -lssc
// #include <stdlib.h>
// #include "../ssc/sscapi.h"
import "C"

type ssc struct {
	Version   int
	BuildInfo string
}

func InitSsc() *ssc {
	result := new(ssc)
	result.Version = int(C.ssc_version())
	result.BuildInfo = C.GoString(C.ssc_build_info())
	return result
}
