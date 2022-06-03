package gossc

// #cgo CFLAGS: -g -Wall
// #cgo LDFLAGS: -L../ -lssc
// #include <stdlib.h>
// #include "../ssc/sscapi.h"
import "C"
import (
	"errors"
	"fmt"
	"time"
)

type SscModuleEntry struct {
	Name        string
	Description string
	Version     int
}

type SscModule struct {
	module       C.ssc_module_t
	VariableInfo []SscVariableInfo
}

type SscLogType int

const (
	Notice  SscLogType = 1
	Warning SscLogType = 2
	Error   SscLogType = 3
)

type SscActionType int

const (
	Log    SscActionType = 0
	Update SscActionType = 1
)

type SscLogMessage struct {
	Type    SscLogType
	Time    time.Time
	Message string
}

type SscVariableInfo struct {
	VariableType SscVariableType
	DataType     SscDataType
	Name         string
	Label        string
	Units        string
	Meta         string
	Group        string
	Required     string
	Constraints  string
}

func CreateModule(moduleName string) (*SscModule, error) {
	result := new(SscModule)
	moduleNameC := C.CString(moduleName)
	result.module = C.ssc_module_create(moduleNameC)
	if result.module == nil {
		return nil, errors.New(fmt.Sprintf("Failed to initialize module: %s", moduleName))
	} else {
		result.VariableInfo = []SscVariableInfo{}
		for variableIndex, rawVariable := 0, C.ssc_module_var_info(result.module, 0); rawVariable != nil; variableIndex, rawVariable = variableIndex+1, C.ssc_module_var_info(result.module, C.int(variableIndex)) {
			variableInfo := SscVariableInfo{}
			variableInfo.VariableType = SscVariableType(C.ssc_info_var_type(rawVariable))
			variableInfo.DataType = SscDataType(C.ssc_info_data_type(rawVariable))
			variableInfo.Name = C.GoString(C.ssc_info_name(rawVariable))
			variableInfo.Label = C.GoString(C.ssc_info_label(rawVariable))
			variableInfo.Units = C.GoString(C.ssc_info_units(rawVariable))
			variableInfo.Meta = C.GoString(C.ssc_info_meta(rawVariable))
			variableInfo.Group = C.GoString(C.ssc_info_group(rawVariable))
			variableInfo.Required = C.GoString(C.ssc_info_required(rawVariable))
			variableInfo.Constraints = C.GoString(C.ssc_info_constraints(rawVariable))
			result.VariableInfo = append(result.VariableInfo, variableInfo)
		}
		return result, nil
	}
}

func (module *SscModule) Free() {
	if module.module != nil {
		C.ssc_module_free(module.module)
	}
}

func (module *SscModule) Execute(data SscData) bool {
	returnCode := int(C.ssc_module_exec(module.module, data.data))
	success := returnCode == 1
	return success
}

func GetModuleEntries() []SscModuleEntry {
	result := []SscModuleEntry{}
	var rawEntry C.ssc_entry_t
	var entryIndex int = 0
	for rawEntry = C.ssc_module_entry(C.int(entryIndex)); rawEntry != nil; rawEntry, entryIndex = C.ssc_module_entry(C.int(entryIndex)), entryIndex+1 {
		entry := SscModuleEntry{}
		entry.Name = C.GoString(C.ssc_entry_name(rawEntry))
		entry.Description = C.GoString(C.ssc_entry_description(rawEntry))
		entry.Version = int(C.ssc_entry_version(rawEntry))
		result = append(result, entry)
	}
	return result
}
