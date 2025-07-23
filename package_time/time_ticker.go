package package_time

import (
	"fmt"
	"time"
)

func TickerExample() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	fmt.Println("Ticker berjalan (3x):")
	for i := 0; i < 3; i++ {
		<-ticker.C
		fmt.Println("Tick at", time.Now())
	}
}
