package main

import (
	"fmt"

	s "github.com/solairerove/go-channels-playground/channel-block"
	d "github.com/solairerove/go-channels-playground/direct"
	f "github.com/solairerove/go-channels-playground/funcs"
)

func main() {
	fmt.Println("Channels playground")

	s.SimpleChannelBlock()
	d.ChannelDicect()
	f.ChannelWithFuncs()
}
