package package_time

import (
	"fmt"
	"time"
)

func FormattingParsing() {
	t := time.Now()
	fmt.Println("Default format:", t)

	formatted := t.Format("Monday, 02-Jan-2006 15:04:05 MST")
	fmt.Println("Formatted:", formatted)

	parsed, _ := time.Parse("2006-01-02", "2025-12-31")
	fmt.Println("Parsed date:", parsed)
}
