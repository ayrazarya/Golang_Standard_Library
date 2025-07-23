package package_strconv

import (
	"fmt"
	"strconv"
)

func DemoFormatBool() {
	fmt.Println("\n=== Demo FormatBool ===")

	fmt.Println("true →", strconv.FormatBool(true))
	fmt.Println("false →", strconv.FormatBool(false))
}
