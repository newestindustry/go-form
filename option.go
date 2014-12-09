package form

type Option struct {
	Value    string
	Label    string
	Selected bool
}

func NewOption() *Option {
	option := &Option{}

	return option
}
