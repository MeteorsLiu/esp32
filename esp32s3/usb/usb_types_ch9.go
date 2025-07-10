package usb

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_B_DESCRIPTOR_TYPE_DEVICE = 0x01
const USB_B_DESCRIPTOR_TYPE_CONFIGURATION = 0x02
const USB_B_DESCRIPTOR_TYPE_STRING = 0x03
const USB_B_DESCRIPTOR_TYPE_INTERFACE = 0x04
const USB_B_DESCRIPTOR_TYPE_ENDPOINT = 0x05
const USB_B_DESCRIPTOR_TYPE_DEVICE_QUALIFIER = 0x06
const USB_B_DESCRIPTOR_TYPE_OTHER_SPEED_CONFIGURATION = 0x07
const USB_B_DESCRIPTOR_TYPE_INTERFACE_POWER = 0x08
const USB_B_DESCRIPTOR_TYPE_OTG = 0x09
const USB_B_DESCRIPTOR_TYPE_DEBUG = 0x0a
const USB_B_DESCRIPTOR_TYPE_INTERFACE_ASSOCIATION = 0x0b
const USB_B_DESCRIPTOR_TYPE_SECURITY = 0x0c
const USB_B_DESCRIPTOR_TYPE_KEY = 0x0d
const USB_B_DESCRIPTOR_TYPE_ENCRYPTION_TYPE = 0x0e
const USB_B_DESCRIPTOR_TYPE_BOS = 0x0f
const USB_B_DESCRIPTOR_TYPE_DEVICE_CAPABILITY = 0x10
const USB_B_DESCRIPTOR_TYPE_WIRELESS_ENDPOINT_COMP = 0x11
const USB_B_DESCRIPTOR_TYPE_WIRE_ADAPTER = 0x21
const USB_B_DESCRIPTOR_TYPE_RPIPE = 0x22
const USB_B_DESCRIPTOR_TYPE_CS_RADIO_CONTROL = 0x23
const USB_B_DESCRIPTOR_TYPE_PIPE_USAGE = 0x24
const USB_SETUP_PACKET_SIZE = 8
const USB_B_REQUEST_GET_STATUS = 0x00
const USB_B_REQUEST_CLEAR_FEATURE = 0x01
const USB_B_REQUEST_SET_FEATURE = 0x03
const USB_B_REQUEST_SET_ADDRESS = 0x05
const USB_B_REQUEST_GET_DESCRIPTOR = 0x06
const USB_B_REQUEST_SET_DESCRIPTOR = 0x07
const USB_B_REQUEST_GET_CONFIGURATION = 0x08
const USB_B_REQUEST_SET_CONFIGURATION = 0x09
const USB_B_REQUEST_GET_INTERFACE = 0x0A
const USB_B_REQUEST_SET_INTERFACE = 0x0B
const USB_B_REQUEST_SYNCH_FRAME = 0x0C
const USB_W_VALUE_DT_DEVICE = 0x01
const USB_W_VALUE_DT_CONFIG = 0x02
const USB_W_VALUE_DT_STRING = 0x03
const USB_W_VALUE_DT_INTERFACE = 0x04
const USB_W_VALUE_DT_ENDPOINT = 0x05
const USB_W_VALUE_DT_DEVICE_QUALIFIER = 0x06
const USB_W_VALUE_DT_OTHER_SPEED_CONFIG = 0x07
const USB_W_VALUE_DT_INTERFACE_POWER = 0x08
const USB_STANDARD_DESC_SIZE = 2
const USB_DEVICE_DESC_SIZE = 18
const USB_CLASS_PER_INTERFACE = 0x00
const USB_CLASS_AUDIO = 0x01
const USB_CLASS_COMM = 0x02
const USB_CLASS_HID = 0x03
const USB_CLASS_PHYSICAL = 0x05
const USB_CLASS_STILL_IMAGE = 0x06
const USB_CLASS_PRINTER = 0x07
const USB_CLASS_MASS_STORAGE = 0x08
const USB_CLASS_HUB = 0x09
const USB_CLASS_CDC_DATA = 0x0a
const USB_CLASS_CSCID = 0x0b
const USB_CLASS_CONTENT_SEC = 0x0d
const USB_CLASS_VIDEO = 0x0e
const USB_CLASS_WIRELESS_CONTROLLER = 0xe0
const USB_CLASS_PERSONAL_HEALTHCARE = 0x0f
const USB_CLASS_AUDIO_VIDEO = 0x10
const USB_CLASS_BILLBOARD = 0x11
const USB_CLASS_USB_TYPE_C_BRIDGE = 0x12
const USB_CLASS_MISC = 0xef
const USB_CLASS_APP_SPEC = 0xfe
const USB_CLASS_VENDOR_SPEC = 0xff
const USB_SUBCLASS_VENDOR_SPEC = 0xff
const USB_CONFIG_DESC_SIZE = 9
const USB_IAD_DESC_SIZE = 9
const USB_INTF_DESC_SIZE = 9
const USB_EP_DESC_SIZE = 7
const USB_B_ENDPOINT_ADDRESS_EP_NUM_MASK = 0x0f
const USB_B_ENDPOINT_ADDRESS_EP_DIR_MASK = 0x80
const USB_W_MAX_PACKET_SIZE_MPS_MASK = 0x07ff
const USB_W_MAX_PACKET_SIZE_MULT_MASK = 0x1800
const USB_BM_ATTRIBUTES_XFERTYPE_MASK = 0x03
const USB_BM_ATTRIBUTES_SYNCTYPE_MASK = 0x0C
const USB_BM_ATTRIBUTES_USAGETYPE_MASK = 0x30
const USB_STR_DESC_SIZE = 2

type UsbDeviceStateT c.Int

const (
	USB_DEVICE_STATE_NOT_ATTACHED UsbDeviceStateT = 0
	USB_DEVICE_STATE_ATTACHED     UsbDeviceStateT = 1
	USB_DEVICE_STATE_POWERED      UsbDeviceStateT = 2
	USB_DEVICE_STATE_DEFAULT      UsbDeviceStateT = 3
	USB_DEVICE_STATE_ADDRESS      UsbDeviceStateT = 4
	USB_DEVICE_STATE_CONFIGURED   UsbDeviceStateT = 5
	USB_DEVICE_STATE_SUSPENDED    UsbDeviceStateT = 6
)

/**
 * @brief Structure representing a USB control transfer setup packet
 *
 * See Table 9-2 of USB2.0 specification for more details
 */

type UsbSetupPacketT struct {
	Val [8]c.Uint8T
}

/**
 * @brief Structure representing a USB device status
 *
 * See Figures 9-4 Information Returned by a GetStatus() Request to a Device of USB2.0 specification for more details
 */

type UsbDeviceStatusT struct {
	Val c.Uint16T
}

/**
 * @brief USB standard descriptor
 *
 * All USB standard descriptors start with these two bytes. Use this type when traversing over configuration descriptors
 */

type UsbStandardDescT struct {
	Val [2]c.Uint8T
}

/**
 * @brief Structure representing a USB device descriptor
 *
 * See Table 9-8 of USB2.0 specification for more details
 */

type UsbDeviceDescT struct {
	Val [18]c.Uint8T
}

/**
 * @brief Structure representing a short USB configuration descriptor
 *
 * See Table 9-10 of USB2.0 specification for more details
 *
 * @note The full USB configuration includes all the interface and endpoint
 *       descriptors of that configuration.
 */

type UsbConfigDescT struct {
	Val [9]c.Uint8T
}

/**
 * @brief Structure representing a USB interface association descriptor
 */

type UsbIadDescT struct {
	Val [9]c.Uint8T
}

/**
 * @brief Structure representing a USB interface descriptor
 *
 * See Table 9-12 of USB2.0 specification for more details
 */

type UsbIntfDescT struct {
	Val [9]c.Uint8T
}

/**
 * @brief Structure representing a USB endpoint descriptor
 *
 * See Table 9-13 of USB2.0 specification for more details
 */

type UsbEpDescT struct {
	Val [7]c.Uint8T
}

/**
 * @brief Structure representing a USB string descriptor
 */

type UsbStrDescT struct {
	Val [2]c.Uint8T
}
