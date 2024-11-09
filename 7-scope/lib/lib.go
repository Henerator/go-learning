package lib

import "fmt"

// first upper letter - exportable
const Email = "support@example.com"

// first low case - not exportable
var support string

// first upper letter - exportable
func SetSupport(s string) {
	support = s
}

// first upper letter - exportable
func GetContact() string {
	return fmt.Sprintf("%s <%s>", support, Email)
}
