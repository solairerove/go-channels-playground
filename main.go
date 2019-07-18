package main

import (
	"fmt"

	c "github.com/solairerove/go-channels-playground/channel-block"
	d "github.com/solairerove/go-channels-playground/direct"
	in "github.com/solairerove/go-channels-playground/fanin"
	out "github.com/solairerove/go-channels-playground/fanout"
	f "github.com/solairerove/go-channels-playground/funcs"
	r "github.com/solairerove/go-channels-playground/ranges"
	s "github.com/solairerove/go-channels-playground/selects"
)

func main() {
	fmt.Println("Channels playground")

	c.SimpleChannelBlock()
	d.ChannelDicect()
	f.ChannelWithFuncs()
	r.ChannelRange()
	s.ChannelSelect()
	in.ChannnelFanIn()
	out.ChannelsFanOut()
}
