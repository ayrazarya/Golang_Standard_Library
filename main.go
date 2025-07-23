package main

import (
	"GO_Standard_Library/packacge_os"
	"GO_Standard_Library/package_container_list"
	"GO_Standard_Library/package_container_ring"
	"GO_Standard_Library/package_encoding"
	"GO_Standard_Library/package_errors"
	"GO_Standard_Library/package_flag"
	"GO_Standard_Library/package_fmt"
	"GO_Standard_Library/package_io"
	"GO_Standard_Library/package_math"
	"GO_Standard_Library/package_path"
	"GO_Standard_Library/package_reflect"
	"GO_Standard_Library/package_regexp"
	"GO_Standard_Library/package_slices"
	"GO_Standard_Library/package_sort"
	"GO_Standard_Library/package_strconv"
	"GO_Standard_Library/package_strings"
	"GO_Standard_Library/package_time"
)

func main() {
	package_fmt.RunAllDemo()
	package_errors.RunAllErrorDemo()
	package_os.RunAllOSDemo()
	package_flag.RunAllFlagsDemo()
	package_strings.RunAllStringsDemo()
	package_strconv.RunAllStringConvDemo()
	package_math.RunAllMathDemo()
	package_container_list.RunAllListDemo()
	package_container_ring.RunAllRingDemo()
	package_sort.RunAllSortDemo()
	package_time.RunAllTimeDemo()
	package_reflect.RunAllReflectDemo()
	package_regexp.RunAllRegexpDemo()
	package_encoding.RunAllEncodingDemo()
	package_slices.RunAllSlicesDemo()
	package_path.RunAllPathDemo()
	package_io.RunAllIODemo()
}
