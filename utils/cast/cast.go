package cast

import (
	"reflect"

	"github.com/spf13/cast"
)

// Comment
func Kind(kind reflect.Kind, i interface{}) interface{} {
	switch kind {
	case reflect.Bool:
		return cast.ToBool(i)
	case reflect.Int:
		return cast.ToInt(i)
	case reflect.Int8:
		return int8(cast.ToInt16(i))
	case reflect.Int16:
		return cast.ToInt16(i)
	case reflect.Int32:
		return cast.ToInt32(i)
	case reflect.Int64:
		return cast.ToInt64(i)
	case reflect.Uint:
		return cast.ToUint(i)
	case reflect.Uint8:
		return cast.ToUint8(i)
	case reflect.Uint16:
		return cast.ToUint16(i)
	case reflect.Uint32:
		return cast.ToUint32(i)
	case reflect.Uint64:
		return cast.ToUint64(i)
	case reflect.Uintptr:
		return uintptr(cast.ToUint(i))
	case reflect.Float32:
		return cast.ToFloat32(i)
	case reflect.Float64:
		return cast.ToFloat64(i)
	case reflect.Map:
		return cast.ToStringMap(i)
	case reflect.Slice:
		return cast.ToSlice(i)
	case reflect.String:
		return cast.ToString(i)
	default:
		return i
	}
}
