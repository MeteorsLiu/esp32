package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2sPortT c.Int

const (
	I2S_NUM_0    I2sPortT = 0
	I2S_NUM_AUTO I2sPortT = 1
)

type I2sCommModeT c.Int

const (
	I2S_COMM_MODE_STD  I2sCommModeT = 0
	I2S_COMM_MODE_NONE I2sCommModeT = 1
)

type I2sMclkMultipleT c.Int

const (
	I2S_MCLK_MULTIPLE_128  I2sMclkMultipleT = 128
	I2S_MCLK_MULTIPLE_192  I2sMclkMultipleT = 192
	I2S_MCLK_MULTIPLE_256  I2sMclkMultipleT = 256
	I2S_MCLK_MULTIPLE_384  I2sMclkMultipleT = 384
	I2S_MCLK_MULTIPLE_512  I2sMclkMultipleT = 512
	I2S_MCLK_MULTIPLE_576  I2sMclkMultipleT = 576
	I2S_MCLK_MULTIPLE_768  I2sMclkMultipleT = 768
	I2S_MCLK_MULTIPLE_1024 I2sMclkMultipleT = 1024
	I2S_MCLK_MULTIPLE_1152 I2sMclkMultipleT = 1152
)

/**
 * @brief LP I2S transaction type
 */

type LpI2sTransT struct {
	Buffer       c.Pointer
	Buflen       c.SizeT
	ReceivedSize c.SizeT
}

/**
 * @brief Event structure used in I2S event queue
 */

type I2sEventDataT struct {
	Data   c.Pointer
	DmaBuf c.Pointer
	Size   c.SizeT
}

/**
 * @brief Event data structure for LP I2S
 */

type LpI2sEvtDataT struct {
	Trans LpI2sTransT
}

type I2sChannelObjT struct {
	Unused [8]uint8
}
type I2sChanHandleT *I2sChannelObjT

type LpI2sChannelObjT struct {
	Unused [8]uint8
}
type LpI2sChanHandleT *LpI2sChannelObjT

// llgo:type C
type I2sIsrCallbackT func(I2sChanHandleT, *I2sEventDataT, c.Pointer) bool

// llgo:type C
type LpI2sCallbackT func(LpI2sChanHandleT, *LpI2sEvtDataT, c.Pointer) bool
