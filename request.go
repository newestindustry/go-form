package form

import (
	//	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func GetResource(req *http.Request) (string, string) {
	_, resource := http.DefaultServeMux.Handler(&http.Request{Method: "GET", URL: req.URL})

	remains := strings.Replace(req.URL.Path, resource, "", -1)

	return resource, remains
}

// ParseUrlValues parses url.Values into the s interface using
// the reflect package. It also checks for the form struct tags
// so they can be used as fieldnames instead of the variable
// names. It returns the error if parsing failed.
func ParseUrlValues(query url.Values, s interface{}) error {
	fieldIndex := []int{0}
	numElements := reflect.TypeOf(s).Elem().NumField()
	for i := 0; i < numElements; i++ {
		fieldIndex[0] = i
		f := reflect.TypeOf(s).Elem().FieldByIndex(fieldIndex)
		v := reflect.ValueOf(s).Elem().FieldByIndex(fieldIndex)

		fieldname := f.Tag.Get("form")
		if fieldname == "" {
			fieldname = strings.ToLower(f.Name)
		}

		if val, ok := query[fieldname]; ok {
			curVal := val[0]
			switch v.Kind() {
			case reflect.Bool:
				castBool, _ := strconv.ParseBool(curVal)
				v.SetBool(castBool)
			case reflect.Float64:
				castFloat, _ := strconv.ParseFloat(curVal, 64)
				v.SetFloat(castFloat)
			case reflect.Int64:
				castInt, _ := strconv.ParseInt(curVal, 0, 64)
				v.SetInt(castInt)
			case reflect.String:
				v.SetString(curVal)
			default:
			}
		}
	}
	return nil
}

// ParseRequest parses a net/http Request into url.Values
// so it can use ParseUrlValues to parse the request
func ParseRequest(req *http.Request, s interface{}) error {
	req.ParseForm()
	req.ParseMultipartForm(10000)

	ParseUrlValues(req.Form, s)

	return nil
}

// ParsePostForm parses a net/http Request into url.Values
// so it can use ParseUrlValues to parse the request
func ParsePostForm(req *http.Request, s interface{}) error {
	req.ParseForm()
	req.ParseMultipartForm(10000)

	ParseUrlValues(req.Form, s)

	return nil
}
