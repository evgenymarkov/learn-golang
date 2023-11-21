package walk

import "reflect"

func walk(data any, callback func(string)) {
	value := getValue(data)

	switch value.Kind() {
	case reflect.String:
		callback(value.String())
	case reflect.Struct:
		for i := 0; i < value.NumField(); i++ {
			walk(value.Field(i).Interface(), callback)
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < value.Len(); i++ {
			walk(value.Index(i).Interface(), callback)
		}
	case reflect.Map:
		for _, key := range value.MapKeys() {
			walk(value.MapIndex(key).Interface(), callback)
		}
	case reflect.Chan:
		for v, ok := value.Recv(); ok; v, ok = value.Recv() {
			walk(v.Interface(), callback)
		}
	case reflect.Func:
		functionResult := value.Call(nil)

		for _, result := range functionResult {
			walk(result.Interface(), callback)
		}
	}
}

func getValue(data any) reflect.Value {
	value := reflect.ValueOf(data)

	if value.Kind() == reflect.Pointer {
		value = value.Elem()
	}

	return value
}
