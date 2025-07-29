package main

import (
	_ "github.com/MeteorsLiu/esp32/cpu0/esp32"
	"github.com/goplus/lib/c"
)

func main() {
	c.Printf(c.Str("hello world"))
}
