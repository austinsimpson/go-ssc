package ssc

// #cgo CFLAGS: -g -Wall
// #cgo windows LDFLAGS: -L../lib/windows -lssc
// #cgo darwin LDFLAGS: -L../ lib/darwin/ssc.dylib
// #include <stdlib.h>
// #include "../include/sscapi.h"
import "C"

type SscLibraryInfo struct {
	Version   int
	BuildInfo string
}

func InitSsc() *SscLibraryInfo {
	result := new(SscLibraryInfo)
	result.Version = int(C.ssc_version())
	result.BuildInfo = C.GoString(C.ssc_build_info())
	return result
}
