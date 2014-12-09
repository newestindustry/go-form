package form

import (
	"github.com/nicksnyder/go-i18n/i18n"
)

var T i18n.TranslateFunc

func LoadLocale(locale string) i18n.TranslateFunc {
	if locale != "" {
		T, _ = i18n.Tfunc(locale)
	} else {
		T, _ = i18n.Tfunc("en-US")
	}

	return T
}
