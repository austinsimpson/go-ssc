package ssc

// #cgo CFLAGS: -g -Wall
// #cgo windows LDFLAGS: -L../lib/windows -lssc
// #cgo darwin LDFLAGS: -L../ lib/darwin/ssc.dylib
// #include <stdlib.h>
// #include "../include/sscapi.h"
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

type SscDataType uint8
type SscVariableType uint8

const (
	SSC_Invalid SscDataType = 0
	SSC_String  SscDataType = 1
	SSC_Number  SscDataType = 2
	SSC_Array   SscDataType = 3
	SSC_Matrix  SscDataType = 4
	SSC_Table   SscDataType = 5
)

const (
	Input       SscVariableType = 1
	Output      SscVariableType = 2
	InputOutput SscVariableType = 3
)

type SscData struct {
	data C.ssc_data_t
}

type SscNumber float64

type SscMatrix struct {
	Values  []SscNumber
	Rows    uint32
	Columns uint32
}

func CreateData() *SscData {
	result := new(SscData)
	result.data = C.ssc_data_create()
	return result
}

func (sscData *SscData) FreeData() {
	if sscData.data != nil {
		C.ssc_data_free(sscData.data)
	}
}

func (sscData *SscData) Clear() {
	C.ssc_data_clear(sscData.data)
}

func (sscData *SscData) Unassign(variableName string) {
	C.ssc_data_unassign(sscData.data, C.CString(variableName))
}

func (sscData *SscData) Query(variableName string) SscDataType {
	dataTypeAsInt := C.ssc_data_query(sscData.data, C.CString(variableName))
	return SscDataType(dataTypeAsInt)
}

func (sscData *SscData) Iter() func() *string {
	current := C.ssc_data_first(sscData.data)
	return func() *string {
		if current != nil {
			prev := current
			current = C.ssc_data_next(sscData.data)
			result := C.GoString(prev)
			return &result
		} else {
			return nil
		}
	}
}

func (sscData *SscData) SetString(variableName string, value string) {
	C.ssc_data_set_string(sscData.data, C.CString(variableName), C.CString(value))
}

func (sscData *SscData) GetString(variableName string) string {
	return C.GoString(C.ssc_data_get_string(sscData.data, C.CString(variableName)))
}

func (sscData *SscData) SetNumber(variableName string, value SscNumber) {
	C.ssc_data_set_number(sscData.data, C.CString(variableName), C.double(value))
}

func (sscData *SscData) GetNumber(variableName string) (SscNumber, bool) {
	var result SscNumber
	success := C.ssc_data_get_number(sscData.data, C.CString(variableName), (*C.double)(&result)) == 1
	return result, success
}

func (sscData *SscData) SetArray(variableName string, value []SscNumber) {
	C.ssc_data_set_array(sscData.data, C.CString(variableName), (*C.double)(&value[0]), C.int(len(value)))
}

func (sscData *SscData) GetArray(variableName string) []SscNumber {
	var lengthC C.int
	rawResult := C.ssc_data_get_array(sscData.data, C.CString(variableName), &lengthC)
	if rawResult != nil {
		return convertCArray(unsafe.Slice(rawResult, lengthC))
	} else {
		return []SscNumber{}
	}
}

func convertCArray(rawArray []C.double) []SscNumber {
	result := []SscNumber{}
	for _, rawNumber := range rawArray {
		result = append(result, SscNumber(rawNumber))
	}
	return result[:]
}

func (sscData *SscData) SetMatrix(variableName string, value SscMatrix) error {
	if int(value.Rows*value.Columns) != len(value.Values) {
		return errors.New(fmt.Sprintf("Number of entries in data (%d) does not match shape (%d, %d)", len(value.Values), value.Rows, value.Columns))
	}
	C.ssc_data_set_matrix(sscData.data, C.CString(variableName), (*C.double)(&value.Values[0]), C.int(value.Rows), C.int(value.Columns))
	return nil
}

func (sscData *SscData) SetTable(variableName string, table *SscData) {
	C.ssc_data_set_table(sscData.data, C.CString(variableName), table.data)
}
