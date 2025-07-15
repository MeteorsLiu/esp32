package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GdmaChannelT struct {
	Unused [8]uint8
}
type GdmaChannelHandleT *GdmaChannelT

/**
 * @brief Collection of configuration items that used for allocating GDMA channel
 *
 */

type GdmaChannelAllocConfigT struct {
	SiblingChan GdmaChannelHandleT
	Direction   GdmaChannelDirectionT
	Flags       struct {
		ReserveSibling c.Int
		IsrCacheSafe   c.Int
	}
}

/**
 * @brief Type of GDMA event data
 */

type GdmaEventDataT struct {
	Flags struct {
		AbnormalEof c.Uint32T
		NormalEof   c.Uint32T
	}
}

// llgo:type C
type GdmaEventCallbackT func(GdmaChannelHandleT, *GdmaEventDataT, c.Pointer) bool

/**
 * @brief Group of supported GDMA TX callbacks
 * @note The callbacks are all running under ISR environment
 */

type GdmaTxEventCallbacksT struct {
	OnTransEof GdmaEventCallbackT
	OnDescrErr GdmaEventCallbackT
}

/**
 * @brief Group of supported GDMA RX callbacks
 * @note The callbacks are all running under ISR environment
 */

type GdmaRxEventCallbacksT struct {
	OnRecvEof  GdmaEventCallbackT
	OnDescrErr GdmaEventCallbackT
	OnRecvDone GdmaEventCallbackT
}

/**
 * @brief Type of GDMA engine trigger
 * @note It's recommended to initialize this structure with `GDMA_MAKE_TRIGGER`.
 *
 */

type GdmaTriggerT struct {
	Periph     GdmaTriggerPeripheralT
	InstanceId c.Int
	BusId      c.Int
}

/**
 * @brief A collection of strategy item that each GDMA channel could apply
 *
 */

type GdmaStrategyConfigT struct {
	OwnerCheck        bool
	AutoUpdateDesc    bool
	EofTillDataPopped bool
}

/**
 * @brief Channel transfer configurations
 */

type GdmaTransferConfigT struct {
	MaxDataBurstSize c.Uint32T
	AccessExtMem     bool
}

/**
 * @brief GDMA transfer ability
 *
 * @note The alignment set in this structure is **not** a guarantee that gdma driver will take care of the nonalignment cases.
 *       Actually the GDMA driver has no knowledge about the DMA buffer (address and size) used by upper layer.
 *       So it's the responsibility of the **upper layer** to take care of the buffer address and size.
 *
 */

type GdmaTransferAbilityT struct {
	SramTransAlign  c.SizeT
	PsramTransAlign c.SizeT
}
