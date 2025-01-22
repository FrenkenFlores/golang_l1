package l1_14

import (
	"reflect"
)

func FindType(a any) string {
	if a == nil {
		return "nil"
	}
	dtype := reflect.TypeOf(a).String()
	return dtype
}
