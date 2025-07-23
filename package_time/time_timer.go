package package_time

import (
	"fmt"
	"time"
)

func TimerExample() {
	timer := time.NewTimer(2 * time.Second)
	fmt.Println("Menunggu timer...")
	<-timer.C
	fmt.Println("Timer selesai:", time.Now())
}
