package hal

import _ "unsafe"

/**
 * @brief HAL context type of USB WRAP driver
 */

type UsbWrapHalContextT struct {
	Dev *UsbWrapDevT
}

/**
 * @brief Configure whether USB WRAP is routed to internal/external FSLS PHY
 *
 * @param hal USB WRAP HAL context
 * @param external True if external, False if internal
 */
// llgo:link (*UsbWrapHalContextT).UsbWrapHalPhySetExternal C.usb_wrap_hal_phy_set_external
func (recv_ *UsbWrapHalContextT) UsbWrapHalPhySetExternal(external bool) {
}
