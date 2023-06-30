package translations

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	validate "github.com/vasapolrittideah/accord/apps/api/internal/validator"
	"reflect"
	"strings"
)

func RegisterTranslations(v *validator.Validate) ut.Translator {
	english := en.New()
	universalTranslator := ut.New(english, english)
	trans, _ := universalTranslator.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v, trans)

	validate.Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return trans
}
