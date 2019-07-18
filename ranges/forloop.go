package ranges

import (
	"fmt"
)

// ChannelRange tbd
func ChannelRange() {
	fmt.Println("\nChannelRange")

	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}
}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
