package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DacSignalConnT struct {
	DacChannelIoNum [2]c.Uint8T
}
