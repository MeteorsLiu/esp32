package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type SigmaDeltaSignalConnT struct {
	Channels [4]struct {
		SdSig c.Int
	}
}
