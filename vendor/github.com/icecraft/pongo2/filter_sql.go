package pongo2

import (
	"fmt"
	"reflect"
	"strings"
)

type GError struct {
	Msg string `json:"msg"`
}

func (e *GError) Error() string {
	return e.Msg
}

func commonExpandListInt(in *Value, param *Value, sep string) (*Value, *Error) {

	str := make([]string, 0)

	if in.getResolvedValue().Kind() == reflect.Slice {
		items := in.getResolvedValue()
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			str = append(str, fmt.Sprintf("%v", item))
		}
	} else {
		return nil, &Error{
			Sender:   "filter:sqlExpandIntArray",
			ErrorMsg: "Not list, Pls check it when use expandSql",
		}
	}

	return AsValue(strings.Join(str, sep)), nil
}

func filterExpandListInt(in *Value, param *Value) (*Value, *Error) {
	return commonExpandListInt(in, param, ",")
}

func filterExpandListSep(in *Value, param *Value) (*Value, *Error) {
	return commonExpandListInt(in, param, "|")
}
