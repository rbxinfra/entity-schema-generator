package helpers

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

var (
	// NonCamelCaseRegex is a regular expression that matches any string that is not camel case.
	NonCamelCaseRegex = regexp.MustCompile("[^a-zA-Z0-9_ ]+") //nolint:gochecknoglobals

	// AcronymRegex is a regular expression that matches acronyms.
	AcronymRegex = regexp.MustCompile(`[A-Z]{2,}`) //nolint:gochecknoglobals
)

// ToPascalCase converts a string to PascalCase.
// For example, "hello_world" becomes "HelloWorld".
func ToPascalCase(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}

// ToCamelCase converts a string to camelCase.
// For example, "hello_world" becomes "helloWorld".
func ToCamelCase(s string) string {
	s = NonCamelCaseRegex.ReplaceAllString(s, "")
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.AmericanEnglish, cases.NoLower).String(s)
	s = strings.ReplaceAll(s, " ", "")

	if len(s) > 0 {
		s = strings.ToLower(s[:1]) + s[1:]
	}

	return s
}

// NormalizePascalParts normalizes a string by converting all uppercase parts of the string to lowercase.
// Useful for making pascal case acronyms more readable.
// For example, "UserID" becomes "UserId".
func NormalizePascalParts(str string) string {
	matches := AcronymRegex.FindAllString(str, -1)
	for _, match := range matches {
		str = strings.Replace(str, match, match[:1]+strings.ToLower(match[1:]), -1)
	}

	return str
}

// ToString converts an interface to a string while handling fmt.Stringer.
func ToString(o interface{}) string {
	stringer, ok := o.(fmt.Stringer)
	if ok {
		return stringer.String()
	}

	return fmt.Sprintf("%v", o)
}

// ToJson converts an interface to a JSON string.
func ToJson(o interface{}) string {
	data, err := json.Marshal(o)
	if err != nil {
		panic(err)
	}

	return string(data[1 : len(data)-1])
}
