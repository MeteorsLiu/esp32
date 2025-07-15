package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const LLDESC_TX_MBLK_SIZE = 268
const LLDESC_RX_SMBLK_SIZE = 64
const LLDESC_RX_MBLK_SIZE = 524
const LLDESC_RX_AMPDU_ENTRY_MBLK_SIZE = 64
const LLDESC_RX_AMPDU_LEN_MBLK_SIZE = 256
const LLDESC_TX_MBLK_NUM = 10
const LLDESC_RX_MBLK_NUM = 10
const LLDESC_RX_AMPDU_ENTRY_MBLK_NUM = 4
const LLDESC_RX_AMPDU_LEN_MLBK_NUM = 8
const LLDESC_OWNER_MASK = 0x80000000
const LLDESC_OWNER_SHIFT = 31
const LLDESC_SW_OWNED = 0
const LLDESC_HW_OWNED = 1
const LLDESC_EOF_MASK = 0x40000000
const LLDESC_EOF_SHIFT = 30
const LLDESC_SOSF_MASK = 0x20000000
const LLDESC_SOSF_SHIFT = 29
const LLDESC_LENGTH_MASK = 0x00fff000
const LLDESC_LENGTH_SHIFT = 12
const LLDESC_SIZE_MASK = 0x00000fff
const LLDESC_SIZE_SHIFT = 0
const LLDESC_ADDR_MASK = 0x000fffff

type TxAmpduEntryS struct {
	SubLen  c.Uint32T
	DiliNum c.Uint32T
	c.Uint32T
	NullByte c.Uint32T
	Data     c.Uint32T
	Enc      c.Uint32T
	Seq      c.Uint32T
}
type TxAmpduEntryT TxAmpduEntryS

type LldescChainS struct {
	Head *LldescT
	Tail *LldescT
}
type LldescChainT LldescChainS
