package serial

import (
	"reflect"
)

// MapData either encodes or decodes the given data structure with the provided
// Coder. Only those data types that can be serialized with the Coder are
// supported.
func MapData(dataStruct interface{}, coder Coder) {
	valueType := reflect.TypeOf(dataStruct).Elem()
	value := reflect.Indirect(reflect.ValueOf(dataStruct))
	fields := valueType.NumField()

	for i := 0; i < fields; i++ {
		structField := valueType.Field(i)
		valueField := value.Field(i)
		fieldKind := structField.Type.Kind()

		if fieldKind == reflect.Uint8 {
			temp := byte(valueField.Uint())
			coder.CodeByte(&temp)
			valueField.SetUint(uint64(temp))
		} else if fieldKind == reflect.Int8 {
			temp := byte(valueField.Int())
			coder.CodeByte(&temp)
			valueField.SetInt(int64(temp))
		} else if fieldKind == reflect.Uint16 {
			temp := uint16(valueField.Uint())
			coder.CodeUint16(&temp)
			valueField.SetUint(uint64(temp))
		} else if fieldKind == reflect.Int16 {
			temp := uint16(valueField.Int())
			coder.CodeUint16(&temp)
			valueField.SetInt(int64(temp))
		} else if fieldKind == reflect.Uint32 {
			temp := uint32(valueField.Uint())
			coder.CodeUint32(&temp)
			valueField.SetUint(uint64(temp))
		} else if fieldKind == reflect.Int32 {
			temp := uint32(valueField.Int())
			coder.CodeUint32(&temp)
			valueField.SetInt(int64(temp))
		} else if fieldKind == reflect.Array {
			temp := valueField.Slice(0, valueField.Len()).Bytes()
			coder.CodeBytes(temp)
		}
	}
}
