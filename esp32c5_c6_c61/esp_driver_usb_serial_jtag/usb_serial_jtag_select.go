package esp_driver_usb_serial_jtag

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UsjSelectNotifT c.Int

const (
	USJ_SELECT_READ_NOTIF  UsjSelectNotifT = 0
	USJ_SELECT_WRITE_NOTIF UsjSelectNotifT = 1
	USJ_SELECT_ERROR_NOTIF UsjSelectNotifT = 2
)

// llgo:type C
type UsjSelectNotifCallbackT func(UsjSelectNotifT, *BaseTypeT)

/**
 * @brief Set notification callback function for select() events
 * @param usb_serial_jtag_select_notif_callback callback function
 */
//go:linkname UsbSerialJtagSetSelectNotifCallback C.usb_serial_jtag_set_select_notif_callback
func UsbSerialJtagSetSelectNotifCallback(usb_serial_jtag_select_notif_callback UsjSelectNotifCallbackT)

/**
 * @brief Return the readiness status of the driver for read operation
 *
 * @return true if driver read ready, false if not
 */
//go:linkname UsbSerialJtagReadReady C.usb_serial_jtag_read_ready
func UsbSerialJtagReadReady() bool

/**
 * @brief Return the readiness status of the driver for write operation
 *
 * @return true if driver is write ready, false if not
 */
//go:linkname UsbSerialJtagWriteReady C.usb_serial_jtag_write_ready
func UsbSerialJtagWriteReady() bool
