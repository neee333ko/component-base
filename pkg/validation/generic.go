package validation

import (
	"fmt"
	"net"
	"regexp"
	"strings"
	"unicode"
)

const (
	qnameFmt        = "[0-9a-zA-Z]"
	qnameExtFmt     = "[0-9a-zA-Z._-]"
	qualifiedName   = "(" + qnameFmt + qnameExtFmt + "*)?" + qnameFmt
	qualifiedErrMsg = "must consist of alphanumeric characters, '-', '.' or '_'" +
		"and must start and end with alphanumeric characters."
	nameMaxLen = 63
)

var qnameRegexp *regexp.Regexp = regexp.MustCompile("^" + qualifiedName + "$")

func IsQualifiedName(name string) []string {
	errs := make([]string, 0)
	parts := strings.Split(name, "/")

	switch len(parts) {
	case 2:
		prefix := parts[0]
		name = parts[1]

		prefixErrs := IsValidsubdomainDNS1123(prefix)

		if len(prefixErrs) != 0 {
			errs = append(errs, prefixEach("prefix part: ", prefixErrs...))
		}
	case 1:
		name = parts[0]
	default:
		return append(errs, "a qualified name"+
			RegexpError(qualifiedErrMsg, qnameFmt, "myname", "abc.123")+
			" with an optional DNS subdomain prefix and '/'(e.g. example.com/)")
	}

	if len(name) == 0 {
		errs = append(errs, EmptyError())
	}

	if len(name) > nameMaxLen {
		errs = append(errs, MaxLenError(nameMaxLen))
	}

	if !qnameRegexp.MatchString(name) {
		errs = append(errs, RegexpError(qualifiedErrMsg, qnameFmt, "myName", "my_name", "123.45"))
	}

	return errs
}

const (
	labelFmt    = "(" + qualifiedName + ")" + "?"
	labelErrMsg = "a valid label must be an empty string or consist of" +
		"alphanumeric characters, '-', '_' or '.', and must start and end with an alphanumeric character"
	labelMaxLen = 63
)

var labelRegexp *regexp.Regexp = regexp.MustCompile("^" + labelFmt + "$")

func IsValidLabel(label string) []string {
	errs := make([]string, 0)

	if len(label) > labelMaxLen {
		errs = append(errs, MaxLenError(labelMaxLen))
	}

	if !labelRegexp.MatchString(label) {
		errs = append(errs, RegexpError(labelErrMsg, labelFmt, "MyName", "my_value", "12345"))
	}

	return errs
}

const (
	labelDNS1123Fmt    = "([0-9a-z][0-9a-z-]*)?[0-9a-z]"
	labelDNS1123ErrMsg = "a DNS-1123 label must consist of lower case alphanumeric characters or '-'," +
		"and must start and end with an alphanumeric character"
	labelDNS1123MaxLen = 63
)

var labelDNS1123Regexp *regexp.Regexp = regexp.MustCompile("^" + labelDNS1123Fmt + "$")

func IsValidLabelDNS1123(label string) []string {
	errs := make([]string, 0)

	if len(label) > labelDNS1123MaxLen {
		errs = append(errs, MaxLenError(labelDNS1123MaxLen))
	}

	if !labelDNS1123Regexp.MatchString(label) {
		errs = append(errs, RegexpError(labelDNS1123ErrMsg, labelDNS1123Fmt, "my-name", "123-abc"))
	}

	return errs
}

const (
	subdomainDNS1123Fmt = labelDNS1123Fmt + "(\\." + labelDNS1123Fmt + ")*"
	subdomainDNS1123Msg = "a DNS-1123 subdomain must consist of lower case alphanumeric characters, '-' or '.'," +
		"and must start and end with an alphanumeric character"
	subdomainDNS1123MaxLen = 253
)

var subdomainDNS1123Regexp *regexp.Regexp = regexp.MustCompile("^" + subdomainDNS1123Fmt + "$")

func IsValidsubdomainDNS1123(subdomain string) []string {
	errs := make([]string, 0)

	if len(subdomain) > subdomainDNS1123MaxLen {
		errs = append(errs, MaxLenError(subdomainDNS1123MaxLen))
	}

	if !subdomainDNS1123Regexp.MatchString(subdomain) {
		errs = append(errs, RegexpError(subdomainDNS1123Msg, subdomainDNS1123Fmt, "example.com"))
	}

	return errs
}

func IsValidPort(port int) []string {
	if port >= 1 && port <= 65535 {
		return nil
	}

	return []string{InclusiveRangeError(1, 65535)}
}

func IsInRange(value, lo, hi int) []string {
	if value >= lo && value <= hi {
		return nil
	}

	return []string{InclusiveRangeError(lo, hi)}
}

func IsValidIP(ip string) []string {
	if net.ParseIP(ip) != nil {
		return nil
	}

	return []string{"must be a valid ip address, (e.g. 9.8.7.1)"}
}

func IsValidIPv4Address(ip string) []string {
	if net.ParseIP(ip).To4() != nil {
		return nil
	}

	return []string{"must be a valid ipv4 address"}
}

func IsValidIPv6Address(ip string) []string {
	if net.ParseIP(ip) == nil || net.ParseIP(ip).To4() != nil {
		return []string{"must be a valid ipv6 address"}
	}

	return nil
}

const (
	percentFmt = "(0|[1-9][0-9]*)%"
	percentMsg = "a valid percent string must be a numeric string followed by an ending '%'"
)

var percentRegexp *regexp.Regexp = regexp.MustCompile("^" + percentFmt + "$")

func IsValidPercent(percent string) []string {
	errs := make([]string, 0)

	if !percentRegexp.MatchString(percent) {
		return append(errs, RegexpError(percentMsg, percentFmt, "99%", "0%"))
	}

	return nil
}

func MaxLenError(maxlen int) string {
	return fmt.Sprintf("must be not over %d", maxlen)
}

func RegexpError(msg string, fmt string, examples ...string) string {
	if len(examples) == 0 {
		return msg + "(fmt used for regexp is '" + fmt + "')"
	}

	msg += "(e.g. "

	for i, example := range examples {
		if i > 0 {
			msg += "or "
		}

		msg += example + ", "
	}

	msg += "fmt used for regexp is '" + fmt + "')"

	return msg
}

func EmptyError() string {
	return "must be not empty"
}

func prefixEach(prefix string, msgs ...string) string {
	if msgs == nil {
		return ""
	}

	var ret string

	for i, msg := range msgs {
		if i > 0 {
			ret += ", "
		}

		ret += prefix + msg
	}

	return ret
}

func InclusiveRangeError(lo, hi int) string {
	return fmt.Sprintf("must be between %d and %d, inclusive", lo, hi)
}

const (
	passMaxLen = 16
	passMinLen = 8
)

func IsValidPassword(password string) []string {
	var hasLower bool
	var hasUpper bool
	var hasNum bool
	var hasSpecial bool
	errs := make([]string, 0)

	var len int

	for _, r := range password {
		switch {
		case unicode.IsNumber(r):
			hasNum = true
			len++
		case unicode.IsLower(r):
			hasLower = true
			len++
		case unicode.IsUpper(r):
			hasUpper = true
			len++
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
			len++
		case r == ' ':
			len++
		}
	}

	if len < passMinLen || len > passMaxLen {
		errs = append(errs, InclusiveRangeError(passMinLen, passMaxLen))
	}

	if !hasLower || !hasUpper || !hasNum || !hasSpecial {
		errs = append(errs, "must consist of lower alpha, upper alpha, number and special character")
	}

	return errs
}
