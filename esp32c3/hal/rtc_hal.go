package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcCntlSleepRetent struct {
	CpuPdMem c.Pointer
}
type RtcCntlSleepRetentT RtcCntlSleepRetent

//go:linkname RtcCntlHalDmaLinkInit C.rtc_cntl_hal_dma_link_init
func RtcCntlHalDmaLinkInit(elem c.Pointer, buff c.Pointer, size c.Int, next c.Pointer) c.Pointer

//go:linkname RtcCntlHalEnableCpuRetention C.rtc_cntl_hal_enable_cpu_retention
func RtcCntlHalEnableCpuRetention(addr c.Pointer)

//go:linkname RtcCntlHalDisableCpuRetention C.rtc_cntl_hal_disable_cpu_retention
func RtcCntlHalDisableCpuRetention(addr c.Pointer)
