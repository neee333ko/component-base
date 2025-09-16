package field

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	utilerror "github.com/neee333ko/errors"
)

type Error struct {
	typeError TypeError
	path      *Path
	value     interface{}
	detail    string
}

var _ error = &Error{}

func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.path.String(), e.ErrorBody())
}

func (e *Error) ErrorBody() string {
	var s string
	switch e.typeError {
	case ErrorRequired, ErrorForbidden, ErrorTooLong, ErrorInternal:
		s = e.typeError.String()
	default:
		value := e.value
		valueType := reflect.TypeOf(value)

		if value == nil || valueType == nil {
			value = "null"
		}

		if valueType.Kind() == reflect.Ptr {
			if actualValue := reflect.ValueOf(value); actualValue.IsNil() {
				value = "null"
			} else {
				value = actualValue.Elem().Interface()
			}
		}

		switch v := value.(type) {
		case int, int32, int64, string:
			s = fmt.Sprintf("%s: %v", e.typeError.String(), v)
		case fmt.Stringer:
			s = fmt.Sprintf("%s: %s", e.typeError.String(), v.String())
		default:
			s = fmt.Sprintf("%s: %#v", e.typeError.String(), v)
		}

	}

	if e.detail != "" {
		s += fmt.Sprintf(": %s", e.detail)
	}

	return s
}

type TypeError string

const (
	ErrorNotFound   TypeError = "NotFound"
	ErrorRequired   TypeError = "Required"
	ErrorDuplicate  TypeError = "Duplicate"
	ErrorInvalid    TypeError = "Invalid"
	ErrorNotSupport TypeError = "NotSupport"
	ErrorForbidden  TypeError = "Forbidden"
	ErrorTooLong    TypeError = "TooLong"
	ErrorTooMany    TypeError = "TooMany"
	ErrorInternal   TypeError = "Internal"
)

func (t TypeError) String() string {
	switch t {
	case ErrorNotFound:
		return "ErrorNotFound"
	case ErrorRequired:
		return "ErrorRequired"
	case ErrorDuplicate:
		return "ErrorDuplicate"
	case ErrorInvalid:
		return "ErrorInvalid"
	case ErrorNotSupport:
		return "ErrorNotSupport"
	case ErrorForbidden:
		return "ErrorForbidden"
	case ErrorTooLong:
		return "ErrorTooLong"
	case ErrorTooMany:
		return "ErrorTooMany"
	case ErrorInternal:
		return "ErrorInternal"
	default:
		return "Unrecognized TypeError"
	}
}

func NotFound(path *Path, value interface{}) *Error {
	return &Error{
		typeError: ErrorNotFound,
		path:      path,
		value:     value,
		detail:    "",
	}
}

func Required(path *Path) *Error {
	return &Error{
		typeError: ErrorRequired,
		path:      path,
		value:     nil,
		detail:    "",
	}
}

func Duplicate(path *Path, value interface{}) *Error {
	return &Error{
		typeError: ErrorDuplicate,
		path:      path,
		value:     value,
		detail:    "",
	}
}

func Invalid(path *Path, value interface{}, detail string) *Error {
	return &Error{
		typeError: ErrorInvalid,
		path:      path,
		value:     value,
		detail:    detail,
	}
}

func NotSupport(path *Path, value interface{}, validItems []string) *Error {
	detail := "supported values: "
	supportValues := make([]string, len(validItems))

	for i, item := range validItems {
		supportValues[i] = strconv.Quote(item)
	}

	detail += strings.Join(supportValues, ", ")

	return &Error{
		typeError: ErrorNotSupport,
		path:      path,
		value:     value,
		detail:    detail,
	}
}

func Forbidden(path *Path, detail string) *Error {
	return &Error{
		typeError: ErrorForbidden,
		path:      path,
		value:     nil,
		detail:    detail,
	}
}

func TooLong(path *Path, value interface{}, maxLength int) *Error {
	return &Error{
		typeError: ErrorTooLong,
		path:      path,
		value:     value,
		detail:    fmt.Sprintf("must have at most %v bytes", maxLength),
	}
}

func TooMany(path *Path, actualQuantity int, maxQuantity int) *Error {
	return &Error{
		typeError: ErrorTooMany,
		path:      path,
		value:     actualQuantity,
		detail:    fmt.Sprintf("must have at most %v", maxQuantity),
	}
}

func Internal(path *Path, err error) *Error {
	return &Error{
		typeError: ErrorInternal,
		path:      path,
		value:     nil,
		detail:    err.Error(),
	}
}

type ErrorList []*Error

func NewTypeErrorMatcher(t TypeError) utilerror.Matcher {
	return func(err error) bool {
		if e, ok := err.(*Error); ok {
			return e.typeError == t
		}

		return false
	}
}

func (errlist ErrorList) ToAggregate() utilerror.Aggregate {
	sets := utilerror.NewString()
	errs := make([]error, 0, len(errlist))

	for _, err := range errlist {
		s := err.Error()

		if sets.Has(s) {
			continue
		}

		sets.Insert(s)
		errs = append(errs, err)
	}

	return utilerror.NewAggregate(errs)
}

func fromAggregate(err error) ErrorList {
	errlist := make(ErrorList, 0)

	for _, e := range err.(utilerror.Aggregate).Errors() {
		errlist = append(errlist, e.(*Error))
	}

	return errlist
}

func (errlist ErrorList) Filter(fns ...utilerror.Matcher) ErrorList {
	err := utilerror.FilterOut(errlist.ToAggregate(), fns...)

	if err == nil {
		return nil
	}

	return fromAggregate(err)
}
