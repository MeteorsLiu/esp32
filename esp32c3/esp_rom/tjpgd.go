package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const JD_SZBUF = 512
const JD_FORMAT = 0
const JD_USE_SCALE = 1
const JD_TBLCLIP = 1

type INT c.Int
type UINT c.Uint
type CHAR c.Char
type UCHAR c.Char
type BYTE c.Char
type SHORT int16
type USHORT uint16
type WORD uint16
type WCHAR uint16
type LONG c.Long
type ULONG c.Ulong
type DWORD c.Ulong
type JRESULT c.Int

const (
	JDR_OK   JRESULT = 0
	JDR_INTR JRESULT = 1
	JDR_INP  JRESULT = 2
	JDR_MEM1 JRESULT = 3
	JDR_MEM2 JRESULT = 4
	JDR_PAR  JRESULT = 5
	JDR_FMT1 JRESULT = 6
	JDR_FMT2 JRESULT = 7
	JDR_FMT3 JRESULT = 8
)

/* Rectangular structure */

type JRECT struct {
	Left   WORD
	Right  WORD
	Top    WORD
	Bottom WORD
}

type JDEC struct {
	Unused [8]uint8
}
