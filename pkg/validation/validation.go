package validation

import (
	"errors"
	"fmt"
	"os"

	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/en"
	"github.com/neee333ko/component-base/pkg/validation/field"
)

const (
	DescriptionMaxLen = 256
)

type Validator struct {
	val   *validator.Validate
	data  interface{}
	trans ut.Translator
}

func NewValidator(data interface{}) *Validator {
	val := validator.New()

	val.RegisterValidation("dir", ValidateDir, true)
	val.RegisterValidation("file", ValidateFile, true)
	val.RegisterValidation("description", ValidateDescription, true)
	val.RegisterValidation("name", ValidateName, true)

	e := english.New()

	utInstance := ut.New(e, e)

	t, _ := utInstance.GetTranslator("en")

	err := en.RegisterDefaultTranslations(val, t)
	if err != nil {
		panic(err)
	}

	regs := []struct {
		tag         string
		translation string
	}{
		{
			tag:         "dir",
			translation: "{0} must point to an existing dir, but found {1}",
		},
		{
			tag:         "file",
			translation: "{0} must point to an existing file, but found {1}",
		},
		{
			tag:         "description",
			translation: fmt.Sprintf("must be not more than %d", DescriptionMaxLen),
		},
		{
			tag:         "name",
			translation: "is not a valid name",
		},
	}

	for _, r := range regs {
		err := val.RegisterTranslation(r.tag, t, registerFn(r.tag, r.translation), translationFn)
		if err != nil {
			panic(err)
		}
	}

	validator := &Validator{
		val:   val,
		data:  data,
		trans: t,
	}

	return validator
}

func translationFn(ut ut.Translator, fe validator.FieldError) string {
	res, err := ut.T(fe.Tag(), fe.Field(), fe.Value().(string))

	if err != nil {
		return fe.Error()
	}

	return res
}

func registerFn(tag, translation string) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		err := ut.Add(tag, translation, false)

		return err
	}
}

func (v *Validator) Validate() field.ErrorList {
	err := v.val.Struct(v.data)
	if err == nil {
		return nil
	}

	var invalidValidationErrs *validator.InvalidValidationError
	if errors.As(err, &invalidValidationErrs) {
		return field.ErrorList{field.Invalid(field.NewPath(""), invalidValidationErrs.Error(), "")}
	}

	errlist := field.ErrorList{}

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, e := range validationErrs {
			errlist = append(errlist, field.Invalid(field.NewPath(e.Namespace()), e.Translate(v.trans), ""))
		}
	}

	return errlist
}

func ValidateDir(fl validator.FieldLevel) bool {
	dir := fl.Field().String()

	if info, err := os.Stat(dir); err == nil && info.IsDir() {
		return true
	}

	return false
}

func ValidateFile(fl validator.FieldLevel) bool {
	file := fl.Field().String()

	if info, err := os.Stat(file); err == nil && !info.IsDir() {
		return true
	}

	return false
}

func ValidateDescription(fl validator.FieldLevel) bool {
	description := fl.Field().String()

	return len(description) <= DescriptionMaxLen
}

func ValidateName(fl validator.FieldLevel) bool {
	name := fl.Field().String()

	errs := IsQualifiedName(name)

	return len(errs) == 0
}
