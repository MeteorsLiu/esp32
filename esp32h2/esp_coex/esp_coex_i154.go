package esp_coex

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type Ieee802154CoexEventT c.Int

const (
	IEEE802154_HIGH      Ieee802154CoexEventT = 1
	IEEE802154_MIDDLE    Ieee802154CoexEventT = 2
	IEEE802154_LOW       Ieee802154CoexEventT = 3
	IEEE802154_IDLE      Ieee802154CoexEventT = 4
	IEEE802154_EVENT_MAX Ieee802154CoexEventT = 5
)

type EspIeee802154CoexConfigT struct {
	Idle   Ieee802154CoexEventT
	Txrx   Ieee802154CoexEventT
	TxrxAt Ieee802154CoexEventT
}

// llgo:link Ieee802154CoexEventT.EspCoexIeee802154TxrxPtiSet C.esp_coex_ieee802154_txrx_pti_set
func (recv_ Ieee802154CoexEventT) EspCoexIeee802154TxrxPtiSet() {
}

// llgo:link Ieee802154CoexEventT.EspCoexIeee802154AckPtiSet C.esp_coex_ieee802154_ack_pti_set
func (recv_ Ieee802154CoexEventT) EspCoexIeee802154AckPtiSet() {
}

//go:linkname EspCoexIeee802154CoexBreakNotify C.esp_coex_ieee802154_coex_break_notify
func EspCoexIeee802154CoexBreakNotify()

//go:linkname EspCoexIeee802154ExtcoexTxStage C.esp_coex_ieee802154_extcoex_tx_stage
func EspCoexIeee802154ExtcoexTxStage()

//go:linkname EspCoexIeee802154ExtcoexRxStage C.esp_coex_ieee802154_extcoex_rx_stage
func EspCoexIeee802154ExtcoexRxStage()
