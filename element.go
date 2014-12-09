package form

type Element struct {
	ID          string
	Class       []string
	Type        string
	Name        string
	Label       string
	Value       interface{}
	Placeholder string

	Required   bool
	Autofocus  bool
	NoValidate bool
	Multiple   bool

	TabIndex int64

	Valid bool

	Data       map[string]string
	Options    []*Option
	Validators []Validator
	Errors     []error
}

func NewElement() *Element {
	element := &Element{}
	element.Type = "text"
	return element
}
