package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/* this bitfield is start from the LSB!!! */

type SdioSlaveLlDescS struct {
	Size   c.Uint32T
	Length c.Uint32T
	Offset c.Uint32T
	Sosf   c.Uint32T
	Eof    c.Uint32T
	Owner  c.Uint32T
	Buf    *c.Uint8T
}
type SdioSlaveLlDescT SdioSlaveLlDescS
type SdioSlaveLlSlvintT c.Int

const (
	SDIO_SLAVE_LL_SLVINT_0 SdioSlaveLlSlvintT = 1
	SDIO_SLAVE_LL_SLVINT_1 SdioSlaveLlSlvintT = 2
	SDIO_SLAVE_LL_SLVINT_2 SdioSlaveLlSlvintT = 4
	SDIO_SLAVE_LL_SLVINT_3 SdioSlaveLlSlvintT = 8
	SDIO_SLAVE_LL_SLVINT_4 SdioSlaveLlSlvintT = 16
	SDIO_SLAVE_LL_SLVINT_5 SdioSlaveLlSlvintT = 32
	SDIO_SLAVE_LL_SLVINT_6 SdioSlaveLlSlvintT = 64
	SDIO_SLAVE_LL_SLVINT_7 SdioSlaveLlSlvintT = 128
)
