package form

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"strings"
)

func JoinSliceString(class []string) string {
	var s string
	s = strings.Join(class, " ")
	return s
}

func Render(res http.ResponseWriter, filename string, render string, data interface{}) []byte {

	if render != "ni" && render != "bootstrap" {
		render = "bootstrap"
	}

	path := filepath.Join(TemplatePath, render+"-"+filename)
	elementspath := filepath.Join(TemplatePath, render+"-"+"element.tmpl")

	var err error

	tpl := template.New(render + "-" + filename)

	var funcMap = template.FuncMap{
		"join": JoinSliceString,
		"T":    T,
	}

	tpl.Funcs(funcMap)

	_, err = tpl.ParseFiles(path, elementspath)
	if err != nil {
		panic(err)
	}

	b := bytes.NewBuffer(nil)

	err = tpl.Execute(b, data)
	if err != nil {
		panic(err)
	}

	return b.Bytes()
}
