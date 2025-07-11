package usb

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_TRANSFER_FLAG_ZERO_PACK = 0x01

type UsbSpeedT c.Int

const (
	USB_SPEED_LOW  UsbSpeedT = 0
	USB_SPEED_FULL UsbSpeedT = 1
	USB_SPEED_HIGH UsbSpeedT = 2
)

type UsbTransferTypeT c.Int

const (
	USB_TRANSFER_TYPE_CTRL        UsbTransferTypeT = 0
	USB_TRANSFER_TYPE_ISOCHRONOUS UsbTransferTypeT = 1
	USB_TRANSFER_TYPE_BULK        UsbTransferTypeT = 2
	USB_TRANSFER_TYPE_INTR        UsbTransferTypeT = 3
)

type UsbDeviceHandleS struct {
	Unused [8]uint8
}
type UsbDeviceHandleT *UsbDeviceHandleS

// llgo:type C
type UsbHostEnumFilterCbT func(*UsbDeviceDescT, *c.Uint8T) bool

/**
 * @brief Parent device information
 */
type UsbParentDevInfoT struct {
	DevHdl  UsbDeviceHandleT
	PortNum c.Uint8T
}

/**
 * @brief Basic information of an enumerated device
 */

type UsbDeviceInfoT struct {
	Parent              UsbParentDevInfoT
	Speed               UsbSpeedT
	DevAddr             c.Uint8T
	BMaxPacketSize0     c.Uint8T
	BConfigurationValue c.Uint8T
	StrDescManufacturer *UsbStrDescT
	StrDescProduct      *UsbStrDescT
	StrDescSerialNum    *UsbStrDescT
}
type UsbTransferStatusT c.Int

const (
	USB_TRANSFER_STATUS_COMPLETED UsbTransferStatusT = 0
	USB_TRANSFER_STATUS_ERROR     UsbTransferStatusT = 1
	USB_TRANSFER_STATUS_TIMED_OUT UsbTransferStatusT = 2
	USB_TRANSFER_STATUS_CANCELED  UsbTransferStatusT = 3
	USB_TRANSFER_STATUS_STALL     UsbTransferStatusT = 4
	USB_TRANSFER_STATUS_OVERFLOW  UsbTransferStatusT = 5
	USB_TRANSFER_STATUS_SKIPPED   UsbTransferStatusT = 6
	USB_TRANSFER_STATUS_NO_DEVICE UsbTransferStatusT = 7
)

/**
 * @brief Isochronous packet descriptor
 *
 * If the number of bytes in an Isochronous transfer is larger than the MPS of the endpoint, the transfer is split
 * into multiple packets transmitted at the endpoint's specified interval. An array of Isochronous packet descriptors
 * describes how an Isochronous transfer should be split into multiple packets.
 */

type UsbIsocPacketDescT struct {
	NumBytes       c.Int
	ActualNumBytes c.Int
	Status         UsbTransferStatusT
}

type UsbTransferS struct {
	Unused [8]uint8
}
type UsbTransferT UsbTransferS

// llgo:type C
type UsbTransferCbT func(*UsbTransferT)
