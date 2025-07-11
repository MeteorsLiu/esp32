package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaHalContextT struct {
	Unused [8]uint8
}

/**
 * @brief GDMA HAL configuration
 */

type GdmaHalConfigT struct {
	GroupId c.Int
}

type GdmaHalCrcConfigT struct {
	InitValue       c.Uint32T
	CrcBitWidth     c.Uint32T
	PolyHex         c.Uint32T
	ReverseDataMask bool
}

/**
 * @brief GDMA HAL private data
 */

type GdmaHalPrivDataT struct {
	M2mFreePeriphMask c.Uint32T
}
