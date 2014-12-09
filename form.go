package form

import (
	"encoding/json"
	"fmt"
	// "net/url"
)

var Locale string = "en-US"

type Form struct {
	ID       string
	Class    []string
	Method   string
	Action   string
	Data     map[string]string
	Valid    bool
	Validate bool
	Locale   string
	Render   string

	Errors   []error
	Success  []string
	Messages []string

	Elements []*Element
}

// NewForm creates a default form implementation and fills it
// with default values where necessery
func NewForm() *Form {
	form := &Form{}

	form.Class = []string{"fluxum"}
	form.Method = "post"
	form.Data = make(map[string]string)
	form.Locale = "en-US"
	form.Render = "bootstrap" // bootstrap, ni

	LoadLocale(form.Locale)

	return form
}

// CreateFormByJson creates a NewForm() and fills it using the form json
// provided
func CreateFormByJson(data []byte) *Form {
	frm := NewForm()

	err := json.Unmarshal(data, frm)
	if err != nil {
		fmt.Println(err)
	}

	hasFluxum := false
	for _, val := range frm.Class {
		if val == "fluxum" {
			hasFluxum = true
		}
	}
	if !hasFluxum {
		frm.Class = append(frm.Class, "fluxum")
	}

	if frm.Locale == "" {
		frm.Locale = "en-US"
	}
	Locale = frm.Locale
	LoadLocale(Locale)

	return frm
}
