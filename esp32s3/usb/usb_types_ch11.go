package usb

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const USB_CLASS_DESCRIPTOR_TYPE_HUB = 0x29
const USB_PORT_STATUS_SIZE = 4
const USB_HUB_STATUS_SIZE = 4

type UsbHubClassRequestT c.Int

const (
	USB_B_REQUEST_HUB_GET_PORT_STATUS  UsbHubClassRequestT = 0
	USB_B_REQUEST_HUB_CLEAR_FEATURE    UsbHubClassRequestT = 1
	USB_B_REQUEST_HUB_GET_STATE        UsbHubClassRequestT = 2
	USB_B_REQUEST_HUB_SET_PORT_FEATURE UsbHubClassRequestT = 3
	USB_B_REQUEST_HUB_GET_DESCRIPTOR   UsbHubClassRequestT = 6
	USB_B_REQUEST_HUB_SET_DESCRIPTOR   UsbHubClassRequestT = 7
	USB_B_REQUEST_HUB_CLEAR_TT_BUFFER  UsbHubClassRequestT = 8
	USB_B_REQUEST_HUB_RESET_TT         UsbHubClassRequestT = 9
	USB_B_REQUEST_HUB_GET_TT_STATE     UsbHubClassRequestT = 10
	USB_B_REQUEST_HUB_STOP_TT          UsbHubClassRequestT = 11
)

type UsbHubPortStateT c.Int

const (
	USB_PORT_STATE_NOT_CONFIGURED UsbHubPortStateT = 0
	USB_PORT_STATE_POWERED_OFF    UsbHubPortStateT = 1
	USB_PORT_STATE_DISCONNECTED   UsbHubPortStateT = 2
	USB_PORT_STATE_DISABLED       UsbHubPortStateT = 3
	USB_PORT_STATE_RESETTING      UsbHubPortStateT = 4
	USB_PORT_STATE_ENABLED        UsbHubPortStateT = 5
	USB_PORT_STATE_TRANSMIT       UsbHubPortStateT = 6
	USB_PORT_STATE_TRANSMIT_R     UsbHubPortStateT = 7
	USB_PORT_STATE_SUSPENDED      UsbHubPortStateT = 8
	USB_PORT_STATE_RESUMING       UsbHubPortStateT = 9
	USB_PORT_STATE_SEND_EOR       UsbHubPortStateT = 10
	USB_PORT_STATE_RESTART_S      UsbHubPortStateT = 11
	USB_PORT_STATE_RESTART_E      UsbHubPortStateT = 12
	USB_PORT_STATE_TESTING        UsbHubPortStateT = 13
)

type UsbHubPortFeatureT c.Int

const (
	USB_FEATURE_PORT_CONNECTION     UsbHubPortFeatureT = 0
	USB_FEATURE_PORT_ENABLE         UsbHubPortFeatureT = 1
	USB_FEATURE_PORT_SUSPEND        UsbHubPortFeatureT = 2
	USB_FEATURE_PORT_OVER_CURRENT   UsbHubPortFeatureT = 3
	USB_FEATURE_PORT_RESET          UsbHubPortFeatureT = 4
	USB_FEATURE_PORT_POWER          UsbHubPortFeatureT = 8
	USB_FEATURE_PORT_LOWSPEED       UsbHubPortFeatureT = 9
	USB_FEATURE_C_PORT_CONNECTION   UsbHubPortFeatureT = 16
	USB_FEATURE_C_PORT_ENABLE       UsbHubPortFeatureT = 17
	USB_FEATURE_C_PORT_SUSPEND      UsbHubPortFeatureT = 18
	USB_FEATURE_C_PORT_OVER_CURRENT UsbHubPortFeatureT = 19
	USB_FEATURE_C_PORT_RESET        UsbHubPortFeatureT = 20
	USB_FEATURE_PORT_TEST           UsbHubPortFeatureT = 21
	USB_FEATURE_PORT_INDICATOR      UsbHubPortFeatureT = 22
)

/**
 * @brief USB Hub Port Status and Hub Change results
 *
 * See USB 2.0 spec Table 11-19 and Table 11-20
 */

type UsbPortStatusT struct {
	WPortStatus struct {
		Val c.Uint16T
	}
	WPortChange struct {
		Val c.Uint16T
	}
}

/**
 * @brief USB Hub Status
 */

type UsbHubStatusT struct {
	WHubStatus struct {
		Val c.Uint16T
	}
	WHubChange struct {
		Val c.Uint16T
	}
}

/**
 * @brief USB Hub Device descriptor
 */

type UsbHubDescriptorT struct {
	BDescLength         c.Uint8T
	BDescriptorType     c.Uint8T
	BNbrPorts           c.Uint8T
	WHubCharacteristics struct {
		Val c.Uint16T
	}
	BPwrOn2PwrGood   c.Uint8T
	BHubContrCurrent c.Uint8T
}
