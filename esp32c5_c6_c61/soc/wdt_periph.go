package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const TIMG_WDT_WKEY_VALUE = 0x50D83AA1

type TgRegCtxLinkT struct {
	LinkList *RegdmaEntriesConfigT
	LinkNum  c.Uint32T
}
