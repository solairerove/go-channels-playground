package direct

import (
	"fmt"
)

// ChannelDicect tbd
func ChannelDicect() {
	fmt.Println("\nChannelDicect")

	c := make(chan int)
	cr := make(<-chan int) // receive
	cs := make(chan<- int) // send

	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", cr)
	fmt.Printf("%T\n", cs)

}
