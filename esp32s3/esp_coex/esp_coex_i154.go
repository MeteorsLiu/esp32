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
