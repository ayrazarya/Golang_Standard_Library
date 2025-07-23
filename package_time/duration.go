package package_time

import (
	"fmt"
	"time"
)

func DurationExample() {
	d := 2*time.Hour + 30*time.Minute
	fmt.Println("Durasi:", d)
	fmt.Println("Dalam menit:", d.Minutes())
}
