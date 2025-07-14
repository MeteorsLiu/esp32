package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type UsbSerialJtagLlIntrT c.Int

const (
	USB_SERIAL_JTAG_INTR_SOF                 UsbSerialJtagLlIntrT = 2
	USB_SERIAL_JTAG_INTR_SERIAL_OUT_RECV_PKT UsbSerialJtagLlIntrT = 4
	USB_SERIAL_JTAG_INTR_SERIAL_IN_EMPTY     UsbSerialJtagLlIntrT = 8
	USB_SERIAL_JTAG_INTR_TOKEN_REC_IN_EP1    UsbSerialJtagLlIntrT = 256
	USB_SERIAL_JTAG_INTR_BUS_RESET           UsbSerialJtagLlIntrT = 512
	USB_SERIAL_JTAG_INTR_EP1_ZERO_PAYLOAD    UsbSerialJtagLlIntrT = 1024
)
