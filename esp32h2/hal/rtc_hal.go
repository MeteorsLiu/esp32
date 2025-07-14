package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcCntlSleepRetent struct {
	CpuPdMem c.Pointer
}
type RtcCntlSleepRetentT RtcCntlSleepRetent
