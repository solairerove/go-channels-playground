package block

import (
	"fmt"
	"runtime"
)

// SimpleChannelBlock tbd
func SimpleChannelBlock() {
	fmt.Println("SimpleChannelBlock")
	fmt.Println("Goroutines", runtime.NumGoroutine())

	c := make(chan int)

	go func() {
		c <- 42
	}()

	fmt.Println(<-c)
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Goroutines", runtime.NumGoroutine())
}
