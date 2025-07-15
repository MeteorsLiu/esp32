package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type Dma2dGroupT struct {
	Unused [8]uint8
}
type Dma2dPoolHandleT *Dma2dGroupT

type Dma2dTransS struct {
	Unused [8]uint8
}
type Dma2dTransT Dma2dTransS

/**
 * @brief A collection of configuration items that used for allocating a 2D-DMA pool
 */

type Dma2dPoolConfigT struct {
	PoolId       c.Uint32T
	IntrPriority c.Uint32T
}

type Dma2dChannelT struct {
	Unused [8]uint8
}
type Dma2dChannelHandleT *Dma2dChannelT

/**
 * @brief Struct to save the necessary information of a 2D-DMA channel for upper drivers to configure the channels
 */

type Dma2dTransChannelInfoT struct {
	Dir  Dma2dChannelDirectionT
	Chan Dma2dChannelHandleT
}

// llgo:type C
type Dma2dTransOnPickedCallbackT func(c.Uint32T, *Dma2dTransChannelInfoT, c.Pointer) bool

/**
 * @brief A collection of configuration items for a 2D-DMA transaction
 */

type Dma2dTransConfigT struct {
	TxChannelNum           c.Uint32T
	RxChannelNum           c.Uint32T
	ChannelFlags           c.Uint32T
	SpecifiedTxChannelMask c.Uint32T
	SpecifiedRxChannelMask c.Uint32T
	OnJobPicked            Dma2dTransOnPickedCallbackT
	UserConfig             c.Pointer
}

/**
 * @brief Type of 2D-DMA engine trigger
 */

type Dma2dTriggerT struct {
	Periph      Dma2dTriggerPeripheralT
	PeriphSelId c.Int
}

/**
 * @brief A collection of strategy items that each 2D-DMA channel could apply
 */

type Dma2dStrategyConfigT struct {
	OwnerCheck        bool
	AutoUpdateDesc    bool
	EofTillDataPopped bool
}

/**
 * @brief A collection of transfer ability items that each 2D-DMA channel could apply to improve transfer efficiency
 *
 * @note The 2D-DMA driver has no knowledge about the DMA buffer (address and size) used by upper layer.
 *       So it's the responsibility of the **upper layer** to take care of the buffer address and size.
 *       Usually RX buffer at least requires 4-byte alignment to avoid overwriting other data by DMA write PSRAM process
 *       or its data being overwritten.
 */

type Dma2dTransferAbilityT struct {
	DescBurstEn     bool
	DataBurstLength Dma2dDataBurstLengthT
	MbSize          Dma2dMacroBlockSizeT
}

/**
 * @brief A collection of color space conversion (CSC) items that each 2D-DMA channel could apply
 */

type Dma2dCscConfigT struct {
	PreScramble  Dma2dScrambleOrderT
	PostScramble Dma2dScrambleOrderT
}

/**
 * @brief A collection of configurations apply to 2D-DMA channel DSCR-PORT mode
 */

type Dma2dDscrPortModeConfigT struct {
	BlockH c.Uint32T
	BlockV c.Uint32T
}

/**
 * @brief Type of 2D-DMA event data
 */

type Dma2dEventDataT struct {
	Transaction *Dma2dTransT
}

// llgo:type C
type Dma2dEventCallbackT func(Dma2dChannelHandleT, *Dma2dEventDataT, c.Pointer) bool

/**
 * @brief Group of supported 2D-DMA TX callbacks
 * @note The callbacks are all running under ISR environment
 */

type Dma2dTxEventCallbacksT struct {
	OnDescDone Dma2dEventCallbackT
}

/**
 * @brief Group of supported 2D-DMA RX callbacks
 * @note The callbacks are all running under ISR environment
 *
 * Users should be clear on the unique responsibility of each callback when writing the callback functions, such as
 * where to free the transaction memory.
 */

type Dma2dRxEventCallbacksT struct {
	OnRecvEof   Dma2dEventCallbackT
	OnDescDone  Dma2dEventCallbackT
	OnDescEmpty Dma2dEventCallbackT
}
