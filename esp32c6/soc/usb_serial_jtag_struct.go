package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Registers */
/** Type of ep1 register
 *  FIFO access for the CDC-ACM data IN and OUT endpoints.
 */

type UsbSerialJtagEp1RegT struct {
	Val c.Uint32T
}

/** Type of ep1_conf register
 *  Configuration and control registers for the CDC-ACM FIFOs.
 */

type UsbSerialJtagEp1ConfRegT struct {
	Val c.Uint32T
}

/** Type of conf0 register
 *  PHY hardware configuration.
 */

type UsbSerialJtagConf0RegT struct {
	Val c.Uint32T
}

/** Type of test register
 *  Registers used for debugging the PHY.
 */

type UsbSerialJtagTestRegT struct {
	Val c.Uint32T
}

/** Type of misc_conf register
 *  Clock enable control
 */

type UsbSerialJtagMiscConfRegT struct {
	Val c.Uint32T
}

/** Type of mem_conf register
 *  Memory power control
 */

type UsbSerialJtagMemConfRegT struct {
	Val c.Uint32T
}

/** Type of chip_rst register
 *  CDC-ACM chip reset control.
 */

type UsbSerialJtagChipRstRegT struct {
	Val c.Uint32T
}

/** Type of get_line_code_w0 register
 *  W0 of GET_LINE_CODING command.
 */

type UsbSerialJtagGetLineCodeW0RegT struct {
	Val c.Uint32T
}

/** Type of get_line_code_w1 register
 *  W1 of GET_LINE_CODING command.
 */

type UsbSerialJtagGetLineCodeW1RegT struct {
	Val c.Uint32T
}

/** Type of config_update register
 *  Configuration registers' value update
 */

type UsbSerialJtagConfigUpdateRegT struct {
	Val c.Uint32T
}

/** Type of ser_afifo_config register
 *  Serial AFIFO configure register
 */

type UsbSerialJtagSerAfifoConfigRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Registers */
/** Type of int_raw register
 *  Interrupt raw status register.
 */

type UsbSerialJtagIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Interrupt status register.
 */

type UsbSerialJtagIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable status register.
 */

type UsbSerialJtagIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear status register.
 */

type UsbSerialJtagIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Status Registers */
/** Type of jfifo_st register
 *  JTAG FIFO status and control registers.
 */

type UsbSerialJtagJfifoStRegT struct {
	Val c.Uint32T
}

/** Type of fram_num register
 *  Last received SOF frame index register.
 */

type UsbSerialJtagFramNumRegT struct {
	Val c.Uint32T
}

/** Type of in_ep0_st register
 *  Control IN endpoint status information.
 */

type UsbSerialJtagInEp0StRegT struct {
	Val c.Uint32T
}

/** Type of in_ep1_st register
 *  CDC-ACM IN endpoint status information.
 */

type UsbSerialJtagInEp1StRegT struct {
	Val c.Uint32T
}

/** Type of in_ep2_st register
 *  CDC-ACM interrupt IN endpoint status information.
 */

type UsbSerialJtagInEp2StRegT struct {
	Val c.Uint32T
}

/** Type of in_ep3_st register
 *  JTAG IN endpoint status information.
 */

type UsbSerialJtagInEp3StRegT struct {
	Val c.Uint32T
}

/** Type of out_ep0_st register
 *  Control OUT endpoint status information.
 */

type UsbSerialJtagOutEp0StRegT struct {
	Val c.Uint32T
}

/** Type of out_ep1_st register
 *  CDC-ACM OUT endpoint status information.
 */

type UsbSerialJtagOutEp1StRegT struct {
	Val c.Uint32T
}

/** Type of out_ep2_st register
 *  JTAG OUT endpoint status information.
 */

type UsbSerialJtagOutEp2StRegT struct {
	Val c.Uint32T
}

/** Type of set_line_code_w0 register
 *  W0 of SET_LINE_CODING command.
 */

type UsbSerialJtagSetLineCodeW0RegT struct {
	Val c.Uint32T
}

/** Type of set_line_code_w1 register
 *  W1 of SET_LINE_CODING command.
 */

type UsbSerialJtagSetLineCodeW1RegT struct {
	Val c.Uint32T
}

/** Type of bus_reset_st register
 *  USB Bus reset status register
 */

type UsbSerialJtagBusResetStRegT struct {
	Val c.Uint32T
}

/** Group: Version Registers */
/** Type of date register
 *  Date register
 */

type UsbSerialJtagDateRegT struct {
	Val c.Uint32T
}

type UsbSerialJtagDevS struct {
	Ep1            UsbSerialJtagEp1RegT
	Ep1Conf        UsbSerialJtagEp1ConfRegT
	IntRaw         UsbSerialJtagIntRawRegT
	IntSt          UsbSerialJtagIntStRegT
	IntEna         UsbSerialJtagIntEnaRegT
	IntClr         UsbSerialJtagIntClrRegT
	Conf0          UsbSerialJtagConf0RegT
	Test           UsbSerialJtagTestRegT
	JfifoSt        UsbSerialJtagJfifoStRegT
	FramNum        UsbSerialJtagFramNumRegT
	InEp0St        UsbSerialJtagInEp0StRegT
	InEp1St        UsbSerialJtagInEp1StRegT
	InEp2St        UsbSerialJtagInEp2StRegT
	InEp3St        UsbSerialJtagInEp3StRegT
	OutEp0St       UsbSerialJtagOutEp0StRegT
	OutEp1St       UsbSerialJtagOutEp1StRegT
	OutEp2St       UsbSerialJtagOutEp2StRegT
	MiscConf       UsbSerialJtagMiscConfRegT
	MemConf        UsbSerialJtagMemConfRegT
	ChipRst        UsbSerialJtagChipRstRegT
	SetLineCodeW0  UsbSerialJtagSetLineCodeW0RegT
	SetLineCodeW1  UsbSerialJtagSetLineCodeW1RegT
	GetLineCodeW0  UsbSerialJtagGetLineCodeW0RegT
	GetLineCodeW1  UsbSerialJtagGetLineCodeW1RegT
	ConfigUpdate   UsbSerialJtagConfigUpdateRegT
	SerAfifoConfig UsbSerialJtagSerAfifoConfigRegT
	BusResetSt     UsbSerialJtagBusResetStRegT
	Reserved06c    [5]c.Uint32T
	Date           UsbSerialJtagDateRegT
}
type UsbSerialJtagDevT UsbSerialJtagDevS
