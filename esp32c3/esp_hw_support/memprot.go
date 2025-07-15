package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type MemTypeProtT c.Int

const (
	MEMPROT_NONE            MemTypeProtT = 0
	MEMPROT_IRAM0_SRAM      MemTypeProtT = 1
	MEMPROT_DRAM0_SRAM      MemTypeProtT = 2
	MEMPROT_IRAM0_RTCFAST   MemTypeProtT = 4
	MEMPROT_DRAM0_RTCFAST   MemTypeProtT = 8
	MEMPROT_PERI1_RTCSLOW   MemTypeProtT = 16
	MEMPROT_PERI2_RTCSLOW_0 MemTypeProtT = 32
	MEMPROT_PERI2_RTCSLOW_1 MemTypeProtT = 64
	MEMPROT_ALL             MemTypeProtT = -1
)
