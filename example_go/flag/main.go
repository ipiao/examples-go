package main

import (
	"flag"
	"fmt"
)

var (
	ti = flag.Int("t", 123, "")
)

func Init() {
	flag.Parse()
	fmt.Println("flag parsed in init,", *ti)
}

func main() {
	flag.Parse()
	Init()
	//sync.Once{}
}
