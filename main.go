package main

import (
	"fmt"

	s "github.com/solairerove/go-channels-playground/channel-block"
	d "github.com/solairerove/go-channels-playground/direct"
)

func main() {
	fmt.Println("Channels playground")

	s.SimpleChannelBlock()
	d.ChannelDicect()
}
