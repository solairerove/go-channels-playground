package selects

import (
	"fmt"
)

// ChannelSelect tbd
func ChannelSelect() {
	fmt.Println("\nChannelSelect")

	odd, even, quit := make(chan int), make(chan int), make(chan int)

	go send(odd, even, quit)

	receive(odd, even, quit)
}

func receive(o, e, q <-chan int) {
	for {
		select {
		case v := <-o:
			fmt.Println("Odd:", v)
		case v := <-e:
			fmt.Println("Eve:", v)
		case v := <-q:
			fmt.Println("Quit:", v)
			return
		}
	}
}

func send(o, e, q chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	q <- 0
}
