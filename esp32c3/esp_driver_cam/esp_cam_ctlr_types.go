package esp_driver_cam

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspCamCtlrT struct {
	Unused [8]uint8
}
type EspCamCtlrHandleT *EspCamCtlrT

/**
 * @brief ESP CAM controller transaction type
 */

type EspCamCtlrTransT struct {
	Buffer       c.Pointer
	Buflen       c.SizeT
	ReceivedSize c.SizeT
}

/**
 * @brief ESP CAM controller event callbacks
 */

type EspCamCtlrEvtCbsT struct {
	OnGetNewTrans   c.Pointer
	OnTransFinished c.Pointer
}
