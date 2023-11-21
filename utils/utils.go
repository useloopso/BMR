package utils

import (
	"bytes"
	"reflect"
)

func EncodePacked(input ...[]byte) []byte {
	return bytes.Join(input, nil)
}

func ToReflectValue(args []interface{}) []reflect.Value {
	convertedArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		convertedArgs[i] = reflect.ValueOf(arg)
	}
	return convertedArgs
}
