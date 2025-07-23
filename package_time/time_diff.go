package package_time

import (
	"fmt"
	"time"
)

func TimeDifference() {
	start := time.Now()
	time.Sleep(1 * time.Second)
	end := time.Now()

	diff := end.Sub(start)
	fmt.Println("Selisih waktu:", diff)
}
