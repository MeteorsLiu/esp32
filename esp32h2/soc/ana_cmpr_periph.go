package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type AnaCmprPeriphT struct {
	SrcGpio    c.Int
	ExtRefGpio c.Int
	IntrSrc    c.Int
}
