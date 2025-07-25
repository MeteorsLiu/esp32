package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const PARLIO_LL_RX_MAX_BYTES_PER_FRAME = 0xFFFF
const PARLIO_LL_RX_MAX_CLK_INT_DIV = 0x10000
const PARLIO_LL_RX_MAX_CLK_FRACT_DIV = 0
const PARLIO_LL_RX_MAX_TIMEOUT = 0xFFFF
const PARLIO_LL_TX_MAX_BYTES_PER_FRAME = 0xFFFF
const PARLIO_LL_TX_MAX_CLK_INT_DIV = 0x10000
const PARLIO_LL_TX_MAX_CLK_FRACT_DIV = 0
const PARLIO_LL_TX_DATA_LINE_AS_VALID_SIG = 15

type ParlioLlRxEofCondT c.Int

const (
	PARLIO_LL_RX_EOF_COND_RX_FULL     ParlioLlRxEofCondT = 0
	PARLIO_LL_RX_EOF_COND_EN_INACTIVE ParlioLlRxEofCondT = 1
)
