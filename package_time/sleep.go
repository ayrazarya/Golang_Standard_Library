package package_time

import (
	"fmt"
	"time"
)

func SleepExample() {
	fmt.Println("Tidur selama 2 detik...")
	time.Sleep(2 * time.Second)
	fmt.Println("Selesai tidur.")
}
