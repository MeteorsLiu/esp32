package vfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspLineEndingsT c.Int

const (
	ESP_LINE_ENDINGS_CRLF EspLineEndingsT = 0
	ESP_LINE_ENDINGS_CR   EspLineEndingsT = 1
	ESP_LINE_ENDINGS_LF   EspLineEndingsT = 2
)
