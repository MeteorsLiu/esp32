package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const RMT_DEFAULT_CLK_DIV = 80

/**
 * @brief Definition of RMT item
 */

type RmtItem32T struct {
	Unused [8]uint8
}
type RmtChannelT c.Int

const (
	RMT_CHANNEL_0   RmtChannelT = 0
	RMT_CHANNEL_1   RmtChannelT = 1
	RMT_CHANNEL_2   RmtChannelT = 2
	RMT_CHANNEL_3   RmtChannelT = 3
	RMT_CHANNEL_MAX RmtChannelT = 4
)

type RmtMemOwnerT c.Int

const (
	RMT_MEM_OWNER_TX  RmtMemOwnerT = 0
	RMT_MEM_OWNER_RX  RmtMemOwnerT = 1
	RMT_MEM_OWNER_MAX RmtMemOwnerT = 2
)

type RmtSourceClkT c.Int
type RmtDataModeT c.Int

const (
	RMT_DATA_MODE_FIFO RmtDataModeT = 0
	RMT_DATA_MODE_MEM  RmtDataModeT = 1
	RMT_DATA_MODE_MAX  RmtDataModeT = 2
)

type RmtModeT c.Int

const (
	RMT_MODE_TX  RmtModeT = 0
	RMT_MODE_RX  RmtModeT = 1
	RMT_MODE_MAX RmtModeT = 2
)

type RmtIdleLevelT c.Int

const (
	RMT_IDLE_LEVEL_LOW  RmtIdleLevelT = 0
	RMT_IDLE_LEVEL_HIGH RmtIdleLevelT = 1
	RMT_IDLE_LEVEL_MAX  RmtIdleLevelT = 2
)

type RmtCarrierLevelT c.Int

const (
	RMT_CARRIER_LEVEL_LOW  RmtCarrierLevelT = 0
	RMT_CARRIER_LEVEL_HIGH RmtCarrierLevelT = 1
	RMT_CARRIER_LEVEL_MAX  RmtCarrierLevelT = 2
)

type RmtChannelStatusT c.Int

const (
	RMT_CHANNEL_UNINIT RmtChannelStatusT = 0
	RMT_CHANNEL_IDLE   RmtChannelStatusT = 1
	RMT_CHANNEL_BUSY   RmtChannelStatusT = 2
)

/**
 * @brief Data struct of RMT channel status
 */

type RmtChannelStatusResultT struct {
	Status [4]RmtChannelStatusT
}

/**
 * @brief Data struct of RMT TX configure parameters
 */

type RmtTxConfigT struct {
	CarrierFreqHz      c.Uint32T
	CarrierLevel       RmtCarrierLevelT
	IdleLevel          RmtIdleLevelT
	CarrierDutyPercent c.Uint8T
	LoopCount          c.Uint32T
	CarrierEn          bool
	LoopEn             bool
	IdleOutputEn       bool
}

/**
 * @brief Data struct of RMT RX configure parameters
 */

type RmtRxConfigT struct {
	IdleThreshold     c.Uint16T
	FilterTicksThresh c.Uint8T
	FilterEn          bool
}

/**
 * @brief Data struct of RMT configure parameters
 */

type RmtConfigT struct {
	RmtMode     RmtModeT
	Channel     RmtChannelT
	GpioNum     GpioNumT
	ClkDiv      c.Uint8T
	MemBlockNum c.Uint8T
	Flags       c.Uint32T
}
type RmtIsrHandleT IntrHandleT

// llgo:type C
type RmtTxEndFnT func(RmtChannelT, c.Pointer)

/**
 * @brief Structure encapsulating a RMT TX end callback
 */

type RmtTxEndCallbackT struct {
	Function RmtTxEndFnT
	Arg      c.Pointer
}

// llgo:type C
type SampleToRmtT func(c.Pointer, *RmtItem32T, c.SizeT, c.SizeT, *c.SizeT, *c.SizeT)
