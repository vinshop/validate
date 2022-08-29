package validate

import "reflect"

type Wrapper struct {
	Data  interface{}
	Type  reflect.Type
	Value reflect.Value
}

func Wrap(data interface{}) *Wrapper {
	wrapper, ok := data.(*Wrapper)
	if ok {
		return wrapper
	}
	t := reflect.TypeOf(data)
	return &Wrapper{
		Data:  data,
		Type:  t,
		Value: reflect.Indirect(reflect.ValueOf(data)),
	}
}
