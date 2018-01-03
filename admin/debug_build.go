// +build debug

package admin

import "fmt"

func Log(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
