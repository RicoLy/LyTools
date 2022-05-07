package main


import (
	"fmt"
	"strings"
	"time"

	ext "github.com/reugn/go-streams/extension"
	"github.com/reugn/go-streams/flow"
)

func main() {
	var s = []string{}
	println(s)
	var a = 11
	println(a)
	source := ext.NewFileSource("./GoStreams/in.txt")
	flow := flow.NewMap(reverse, 1)
	sink := ext.NewFileSink("out.txt")

	source.Via(flow).To(sink)

	time.Sleep(time.Second)
}

var reverse = func(in interface{}) interface{} {
	msg := in.(string)
	fmt.Printf("Got: %s", msg)
	return strings.ToUpper(msg)
}