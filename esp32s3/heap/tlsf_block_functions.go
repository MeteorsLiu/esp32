package heap

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TlsfptrT PtrdiffT

type BlockHeaderT struct {
	PrevPhysBlock *BlockHeaderT
	Size          c.SizeT
	NextFree      *BlockHeaderT
	PrevFree      *BlockHeaderT
}
