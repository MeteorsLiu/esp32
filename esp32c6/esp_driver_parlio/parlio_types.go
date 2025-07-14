package esp_driver_parlio

import _ "unsafe"

type ParlioTxUnitT struct {
	Unused [8]uint8
}
type ParlioTxUnitHandleT *ParlioTxUnitT

type ParlioRxUnitT struct {
	Unused [8]uint8
}
type ParlioRxUnitHandleT *ParlioRxUnitT

type ParlioRxDelimiterT struct {
	Unused [8]uint8
}
type ParlioRxDelimiterHandleT *ParlioRxDelimiterT
