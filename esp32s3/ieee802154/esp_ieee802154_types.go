package ieee802154

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const US_PER_SYMBLE = 16

type EspIeee802154StateT c.Int

const (
	ESP_IEEE802154_RADIO_DISABLE  EspIeee802154StateT = 0
	ESP_IEEE802154_RADIO_IDLE     EspIeee802154StateT = 1
	ESP_IEEE802154_RADIO_SLEEP    EspIeee802154StateT = 2
	ESP_IEEE802154_RADIO_RECEIVE  EspIeee802154StateT = 3
	ESP_IEEE802154_RADIO_TRANSMIT EspIeee802154StateT = 4
)

type EspIeee802154TxErrorT c.Int

const (
	ESP_IEEE802154_TX_ERR_NONE        EspIeee802154TxErrorT = 0
	ESP_IEEE802154_TX_ERR_CCA_BUSY    EspIeee802154TxErrorT = 1
	ESP_IEEE802154_TX_ERR_ABORT       EspIeee802154TxErrorT = 2
	ESP_IEEE802154_TX_ERR_NO_ACK      EspIeee802154TxErrorT = 3
	ESP_IEEE802154_TX_ERR_INVALID_ACK EspIeee802154TxErrorT = 4
	ESP_IEEE802154_TX_ERR_COEXIST     EspIeee802154TxErrorT = 5
	ESP_IEEE802154_TX_ERR_SECURITY    EspIeee802154TxErrorT = 6
)

type EspIeee802154CcaModeT c.Int

const (
	ESP_IEEE802154_CCA_MODE_CARRIER        EspIeee802154CcaModeT = 0
	ESP_IEEE802154_CCA_MODE_ED             EspIeee802154CcaModeT = 1
	ESP_IEEE802154_CCA_MODE_CARRIER_OR_ED  EspIeee802154CcaModeT = 2
	ESP_IEEE802154_CCA_MODE_CARRIER_AND_ED EspIeee802154CcaModeT = 3
)

type EspIeee802154PendingModeT c.Int

const (
	ESP_IEEE802154_AUTO_PENDING_DISABLE  EspIeee802154PendingModeT = 0
	ESP_IEEE802154_AUTO_PENDING_ENABLE   EspIeee802154PendingModeT = 1
	ESP_IEEE802154_AUTO_PENDING_ENHANCED EspIeee802154PendingModeT = 2
	ESP_IEEE802154_AUTO_PENDING_ZIGBEE   EspIeee802154PendingModeT = 3
)

type EspIeee802154MultipanIndexT c.Int

const (
	ESP_IEEE802154_MULTIPAN_0   EspIeee802154MultipanIndexT = 0
	ESP_IEEE802154_MULTIPAN_1   EspIeee802154MultipanIndexT = 1
	ESP_IEEE802154_MULTIPAN_2   EspIeee802154MultipanIndexT = 2
	ESP_IEEE802154_MULTIPAN_3   EspIeee802154MultipanIndexT = 3
	ESP_IEEE802154_MULTIPAN_MAX EspIeee802154MultipanIndexT = 4
)

/**
 * @brief The information of received 15.4 frame.
 *
 */

type EspIeee802154FrameInfoT struct {
	Pending   bool
	Process   bool
	Channel   c.Uint8T
	Rssi      c.Int8T
	Lqi       c.Uint8T
	Timestamp c.Uint64T
}

/**
 * @brief The O-QPSK PHY transmission power table for 2.4 GHz channels as defined by the IEEE 802.15.4 specification.
 *
 */

type EspIeee802154TxpowerTableT struct {
	Channel [16]c.Int8T
}

// llgo:type C
type EspIeee802154ReceiveDoneCbT func(*c.Uint8T, *EspIeee802154FrameInfoT)

// llgo:type C
type EspIeee802154ReceiveSfdDoneCbT func()

// llgo:type C
type EspIeee802154TransmitDoneCbT func(*c.Uint8T, *c.Uint8T, *EspIeee802154FrameInfoT)

// llgo:type C
type EspIeee802154TransmitFailedCbT func(*c.Uint8T, EspIeee802154TxErrorT)

// llgo:type C
type EspIeee802154TransmitSfdDoneCbT func(*c.Uint8T)

// llgo:type C
type EspIeee802154EnergyDetectDoneCbT func(c.Int8T)

// llgo:type C
type EspIeee802154EnhAckGeneratorCbT func(*c.Uint8T, *EspIeee802154FrameInfoT, *c.Uint8T) EspErrT

/**
 * @brief The callback list for event process.
 *
 * @note These callbacks might be called in ISR context.
 *
 */

type EspIeee802154EventCbListT struct {
	RxDoneCb          EspIeee802154ReceiveDoneCbT
	RxSfdDoneCb       EspIeee802154ReceiveSfdDoneCbT
	TxDoneCb          EspIeee802154TransmitDoneCbT
	TxFailedCb        EspIeee802154TransmitFailedCbT
	TxSfdDoneCb       EspIeee802154TransmitSfdDoneCbT
	EdDoneCb          EspIeee802154EnergyDetectDoneCbT
	EnhAckGeneratorCb EspIeee802154EnhAckGeneratorCbT
}
