package resize

import (
	"fmt"
	"github.com/cloudinary/cloudinary-go/logger"
	"reflect"
)

type Dimensions struct {
	width       interface{} `cld:"w"`
	height      interface{} `cld:"h"`
	aspectRatio interface{} `cld:"ar"`
}

func (d Dimensions) String() string {
	v := reflect.ValueOf(d)
	typeOfS := v.Type()

	res := ""
	for i := 0; i < v.NumField(); i++ {
		if v.Field(i).IsNil() {
			continue
		}
		res += fmt.Sprintf("%s_%v,", typeOfS.Field(i).Tag.Get("cld"), v.Field(i).Elem())
	}

	return res
}

type Position struct {
	X float64 `cld:"x"`
	Y float64 `cld:"y"`
}

type Gravity struct {
	Gravity interface{}
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
