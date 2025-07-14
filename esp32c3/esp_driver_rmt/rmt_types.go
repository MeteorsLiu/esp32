package esp_driver_rmt

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RmtChannelT struct {
	Unused [8]uint8
}
type RmtChannelHandleT *RmtChannelT

type RmtSyncManagerT struct {
	Unused [8]uint8
}
type RmtSyncManagerHandleT *RmtSyncManagerT

type RmtEncoderT struct {
	Unused [8]uint8
}
type RmtEncoderHandleT *RmtEncoderT

/**
 * @brief Type of RMT TX done event data
 */

type RmtTxDoneEventDataT struct {
	NumSymbols c.SizeT
}

// llgo:type C
type RmtTxDoneCallbackT func(RmtChannelHandleT, *RmtTxDoneEventDataT, c.Pointer) bool

/**
 * @brief Type of RMT RX done event data
 */

type RmtRxDoneEventDataT struct {
	ReceivedSymbols *RmtSymbolWordT
	NumSymbols      c.SizeT
	Flags           struct {
		IsLast c.Uint32T
	}
}

// llgo:type C
type RmtRxDoneCallbackT func(RmtChannelHandleT, *RmtRxDoneEventDataT, c.Pointer) bool
