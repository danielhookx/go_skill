package reflect

import (
	"fmt"
	"reflect"
)

func DeepCopy(src interface{}) (interface{}, error) {
	v := reflect.ValueOf(src)
	t := reflect.TypeOf(src)

	if t == nil || !v.IsValid() {
		return src, nil
	}

	switch v.Kind() {
	case reflect.Func, reflect.Chan:
		return nil, fmt.Errorf("unable to copy %v (a %v) as a primitive", v, v.Kind())
	case reflect.Map:
		size := t.Len()
		dst := reflect.MakeMapWithSize(t, size)
		for iter := v.MapRange(); iter.Next(); {
			item, err := DeepCopy(iter.Value().Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to clone map item %v: %v", iter.Key().Interface(), err)
			}
			k, err := DeepCopy(iter.Key())
			if err != nil {
				return nil, fmt.Errorf("failed to clone the map key %v: %v", k, err)
			}
			dst.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(item))
		}
		return dst.Interface(), nil
	case reflect.Array:
		size := t.Len()
		dst := reflect.New(reflect.ArrayOf(size, t.Elem()))
		for i := 0; i < size; i++ {
			item, err := DeepCopy(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to clone array item at index %v: %v", i, err)
			}
			dst.Elem().Index(i).Set(reflect.ValueOf(item))
		}
		return dst.Interface(), nil
	case reflect.Slice:
		size := t.Len()
		dst := reflect.MakeSlice(t, size, size)
		for i := 0; i < size; i++ {
			item, err := DeepCopy(v.Index(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to clone slice item at index %v: %v", i, err)
			}
			if val := reflect.ValueOf(item); val.IsValid() {
				dst.Index(i).Set(val)
			}
		}
		return dst.Interface(), nil
	case reflect.Ptr:
		if v.IsNil() {
			return reflect.Zero(t), nil
		}
		dst := reflect.New(t)
		item, err := DeepCopy(v.Elem().Interface())
		if err != nil {
			return nil, fmt.Errorf("failed to copy the value under the pointer %v: %v", v, err)
		}
		if val := reflect.ValueOf(item); val.IsValid() {
			dst.Elem().Set(val)
		}
		return dst.Interface(), nil
	case reflect.Struct:
		dst := reflect.New(t)
		for i := 0; i < v.NumField(); i++ {
			f := t.Field(i)
			if f.PkgPath != "" {
				continue
			}
			fv, err := DeepCopy(v.Field(i).Interface())
			if err != nil {
				return nil, fmt.Errorf("failed to copy the field %v in the struct %#v: %v", t.Field(i).Name, src, err)
			}
			dst.Elem().Field(i).Set(reflect.ValueOf(fv))
		}
		return dst.Elem().Interface(), nil
	default:
		return v.Interface(), nil
	}
}

func CallMethod(iStruct interface{}, methodName string, params []interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(iStruct)
	if v.Kind() == reflect.Ptr && v.Elem().Kind() != reflect.Struct && v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("type %v not struct", v.Kind())
	}

	method, ok := v.Type().MethodByName(methodName)
	if !ok {
		return nil, fmt.Errorf("method not found")
	}
	ft := method.Type
	params = append([]interface{}{iStruct}, params...)
	if len(params) != ft.NumIn() {
		return nil, fmt.Errorf("method params not match")
	}
	in := make([]reflect.Value, ft.NumIn())
	for i := 0; i < ft.NumIn(); i++ {
		param := reflect.ValueOf(params[i])
		k := ft.In(i).Kind()
		if k != param.Kind() {
			return nil, fmt.Errorf("method param type not match")
		}
		in[i] = param
	}

	response := method.Func.Call(in)
	if len(response) == 0 {
		return nil, nil
	}
	res := make([]interface{}, len(response))
	for i, value := range response {
		res[i] = value.Interface()
	}
	return res, nil
}
