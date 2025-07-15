package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DwGdmaChannelT struct {
	Unused [8]uint8
}
type DwGdmaChannelHandleT *DwGdmaChannelT

type DwGdmaLinkListT struct {
	Unused [8]uint8
}
type DwGdmaLinkListHandleT *DwGdmaLinkListT

type DwGdmaLinkListItemT struct {
	Unused [8]uint8
}
type DwGdmaLliHandleT *DwGdmaLinkListItemT

/**
 * @brief A group of channel's static configurations
 *
 * @note By static, we mean these channel end configurations shouldn't be changed after the DMA channel is created.
 */

type DwGdmaChannelStaticConfigT struct {
	BlockTransferType      DwGdmaBlockTransferTypeT
	Role                   DwGdmaRoleT
	HandshakeType          DwGdmaHandshakeTypeT
	NumOutstandingRequests c.Uint8T
	StatusFetchAddr        c.Uint32T
}

/**
 * @brief Configurations for allocating a DMA channel
 */

type DwGdmaChannelAllocConfigT struct {
	Src            DwGdmaChannelStaticConfigT
	Dst            DwGdmaChannelStaticConfigT
	FlowController DwGdmaFlowControllerT
	ChanPriority   c.Int
	IntrPriority   c.Int
}

/**
 * @brief A group of channel's dynamic configurations
 *
 * @note By dynamic, we mean these channel end configurations can be changed in each transfer.
 */

type DwGdmaChannelDynamicConfigT struct {
	Addr       c.Uint32T
	Width      DwGdmaTransferWidthT
	BurstMode  DwGdmaBurstModeT
	BurstItems DwGdmaBurstItemsT
	BurstLen   c.Uint8T
	Flags      struct {
		EnStatusWriteBack c.Uint32T
	}
}

/**
 * @brief Channel block transfer configurations
 */

type DwGdmaBlockTransferConfigT struct {
	Src  DwGdmaChannelDynamicConfigT
	Dst  DwGdmaChannelDynamicConfigT
	Size c.SizeT
}

/**
 * @brief Type of DW_GDMA trans done event data
 */

type DwGdmaTransDoneEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type DwGdmaTransDoneEventCallbackT func(DwGdmaChannelHandleT, *DwGdmaTransDoneEventDataT, c.Pointer) bool

/**
 * @brief Type of DW_GDMA break event data
 */

type DwGdmaBreakEventDataT struct {
	InvalidLli DwGdmaLliHandleT
}

// llgo:type C
type DwGdmaBreakEventCallbackT func(DwGdmaChannelHandleT, *DwGdmaBreakEventDataT, c.Pointer) bool

/**
 * @brief Group of supported DW_GDMA callbacks
 * @note The callbacks are all running under ISR environment
 */

type DwGdmaEventCallbacksT struct {
	OnBlockTransDone DwGdmaTransDoneEventCallbackT
	OnFullTransDone  DwGdmaTransDoneEventCallbackT
	OnInvalidBlock   DwGdmaBreakEventCallbackT
}
type DwGdmaLinkListTypeT c.Int

const (
	DW_GDMA_LINKED_LIST_TYPE_SINGLY   DwGdmaLinkListTypeT = 0
	DW_GDMA_LINKED_LIST_TYPE_CIRCULAR DwGdmaLinkListTypeT = 1
)

/**
 * @brief DMA link list configurations
 */

type DwGdmaLinkListConfigT struct {
	NumItems c.Uint32T
	LinkType DwGdmaLinkListTypeT
}

/**
 * @brief Markers of a DW_GDMA block
 */

type DwGdmaBlockMarkersT struct {
	IsLast          c.Uint32T
	IsValid         c.Uint32T
	EnTransDoneIntr c.Uint32T
}
