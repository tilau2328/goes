package goes

import (
	"fmt"
	"strings"
)

func TypeOf(item interface{}) string {
	return fmt.Sprintf("%T", item)
}

func MessageType(messageType interface{}) string {
	return strings.TrimPrefix(TypeOf(messageType), "*")
}