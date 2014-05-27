package form

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func ParseUrlValues(query url.Values, s interface{}) error {
	tempintslice := []int{0}
	ielements := reflect.TypeOf(s).Elem().NumField()
	for i := 0; i < ielements; i++ {
		tempintslice[0] = i
		f := reflect.TypeOf(s).Elem().FieldByIndex(tempintslice)
		v := reflect.ValueOf(s).Elem().FieldByIndex(tempintslice)

		fieldname := f.Tag.Get("form")
		if fieldname == "" {
			fieldname = strings.ToLower(f.Name)
		}

		if val, ok := query[fieldname]; ok {
			curVal := val[0]
			switch v.Kind() {
			case reflect.Bool:
				testBool, _ := strconv.ParseBool(curVal)
				v.SetBool(testBool)
			case reflect.Float64:
				testFloat, _ := strconv.ParseFloat(curVal, 64)
				v.SetFloat(testFloat)
			case reflect.Int64:
				testInt, _ := strconv.ParseInt(curVal, 0, 64)
				v.SetInt(testInt)
			case reflect.String:
				v.SetString(curVal)
			default:
			}
		}
	}
	return nil
}

func ParseRequest(req *http.Request, s interface{}) error {
	req.ParseForm()
	req.ParseMultipartForm(10000)

	ParseUrlValues(req.Form, s)

	return nil
}
