package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SPI_LL_MOSI_FREE_LEVEL = 0

type SpiLlClockValT c.Uint32T
type SpiDmaDevT SpiDevT
