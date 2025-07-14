package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const MAX_BLE_DEVNAME_LEN = 29
const BLE_UUID128_VAL_LENGTH = 16
const MAX_BLE_MANUFACTURER_DATA_LEN = 29
const BLE_ADDR_LEN = 6

type ProtocommTransportBleEventT c.Int

const (
	PROTOCOMM_TRANSPORT_BLE_CONNECTED    ProtocommTransportBleEventT = 0
	PROTOCOMM_TRANSPORT_BLE_DISCONNECTED ProtocommTransportBleEventT = 1
)

/**
 * @brief   This structure maps handler required by protocomm layer to
 *          UUIDs which are used to uniquely identify BLE characteristics
 *          from a smartphone or a similar client device.
 */

type NameUuid struct {
	Name *c.Char
	Uuid c.Uint16T
}
type ProtocommBleNameUuidT NameUuid

/**
 * @brief Structure for BLE events in Protocomm.
 */

type ProtocommBleEventT struct {
	EvtType    c.Uint16T
	ConnHandle c.Uint16T
}

/**
 * @brief   Config parameters for protocomm BLE service
 */

type ProtocommBleConfig struct {
	DeviceName          [30]c.Char
	ServiceUuid         [16]c.Uint8T
	ManufacturerData    *c.Uint8T
	ManufacturerDataLen c.SsizeT
	NuLookupCount       c.SsizeT
	NuLookup            *ProtocommBleNameUuidT
	BleBonding          c.Uint
	BleSmSc             c.Uint
	BleLinkEncryption   c.Uint
	BleAddr             *c.Uint8T
	KeepBleOn           c.Uint
	BleNotify           c.Uint
}
type ProtocommBleConfigT ProtocommBleConfig
