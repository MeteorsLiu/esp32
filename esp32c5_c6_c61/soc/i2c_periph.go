package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2cSignalConnT struct {
	SdaOutSig c.Uint8T
	SdaInSig  c.Uint8T
	SclOutSig c.Uint8T
	SclInSig  c.Uint8T
	IomuxFunc c.Uint8T
	Irq       c.Uint8T
}

type I2cRegCtxLinkT struct {
	LinkList *RegdmaEntriesConfigT
	LinkNum  c.Uint32T
	ModuleId PeriphRetentionModuleT
}
