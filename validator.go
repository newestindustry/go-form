package form

import (
	"errors"
	"fmt"
)

type Validator struct {
	Type string
}

func NewValidator() *Validator {
	validator := &Validator{}

	return validator
}

func (f *Form) Check() {
	var overallSuccess bool = true
	for elkey, elvalue := range f.Elements {
		var elementSuccess bool = true
		for _, valvalue := range elvalue.Validators {
			var success bool
			var err []error
			switch valvalue.Type {
			case "required":
				success, err = validateRequired(elvalue)

				break
			}

			if err != nil {
				f.Elements[elkey].Errors = err
				f.Errors = append(f.Errors, err...)
			}
			if elementSuccess == true && success == false {
				elementSuccess = false
			}
			if overallSuccess == true && success == false {
				overallSuccess = false
			}
		}
		f.Elements[elkey].Valid = elementSuccess
	}

	f.Valid = overallSuccess
}

func validateRequired(el *Element) (bool, []error) {
	var msg string = T("%s is required")

	var err []error = []error{}
	var success bool = false
	// var value interface{}

	if el.Value != "" {
		success = true
	} else {
		err = append(err, errors.New(fmt.Sprintf(msg, el.Label)))
	}

	return success, err
}
