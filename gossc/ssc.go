package gossc

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L../ssc ssc.dylib
// #include <stdlib.h>
// #include "../ssc/sscapi.h"
import "C"
import (
	"errors"
	"fmt"
)

type ssc struct {
	Version   int
	BuildInfo string
}

type SscModule struct {
	module C.ssc_module_t
}

type SscData struct {
	data C.ssc_data_t
}

func InitSsc() *ssc {
	result := new(ssc)
	result.Version = int(C.ssc_version())
	result.BuildInfo = C.GoString(C.ssc_build_info())
	return result
}

func CreateModule(moduleName string) (*SscModule, error) {
	result := new(SscModule)
	moduleNameC := C.CString(moduleName)
	result.module = C.ssc_module_create(moduleNameC)
	if result.module == nil {
		return nil, errors.New(fmt.Sprintf("Failed to initialize module: %s", moduleName))
	} else {
		return result, nil
	}
}

func FreeModule(module *SscModule) {
	if module.module != nil {
		C.ssc_module_free(module.module)
	}
}

func CreateData() *SscData {
	result := new(SscData)
	result.data = C.ssc_data_create()
	return result
}

func FreeData(sscData *SscData) {
	if sscData.data != nil {
		C.ssc_data_free(sscData.data)
	}
}
