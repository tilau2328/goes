package goes

import (
	"fmt"
	"regexp"
	"strings"
)

func TypeOf(item interface{}) string {
	return fmt.Sprintf("%T", item)
}

func MessageType(messageType interface{}) string {
	return strings.TrimPrefix(TypeOf(messageType), "*")
}

func Regex(regex string, value string) []string {
	r := regexp.MustCompile(regex)
	results := r.FindAll([]byte(value), -1)
	result := make([]string, 0)
	for _, val := range results {
		result = append(result, string(val))
	}
	return result
}
