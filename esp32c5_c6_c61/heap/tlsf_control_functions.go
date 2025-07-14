package heap

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TlsfConfig c.Int

const (
	ALIGN_SIZE_LOG2 TlsfConfig = 2
	ALIGN_SIZE      TlsfConfig = 4
)

/* The TLSF control structure. */

type ControlT struct {
	BlockNull        BlockHeaderT
	FlIndexCount     c.Uint
	FlIndexShift     c.Uint
	FlIndexMax       c.Uint
	SlIndexCount     c.Uint
	SlIndexCountLog2 c.Uint
	SmallBlockSize   c.Uint
	Size             c.SizeT
	FlBitmap         c.Uint
	SlBitmap         *c.Uint
	Blocks           **BlockHeaderT
}
