package hal

import _ "unsafe"

/**
 * @brief HAL context type of USJ driver
 */

type UsbSerialJtagHalContextT struct {
	Dev *UsbSerialJtagDevT
}

/**
 * @brief Initialize the USJ HAL driver
 *
 * @param hal USJ HAL context
 */
// llgo:link (*UsbSerialJtagHalContextT).UsbSerialJtagHalInit C.usb_serial_jtag_hal_init
func (recv_ *UsbSerialJtagHalContextT) UsbSerialJtagHalInit() {
}
