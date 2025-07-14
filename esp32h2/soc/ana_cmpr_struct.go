package soc

import _ "unsafe"

/**
 * @brief The Analog Comparator Device struct
 * @note  The field in it are register pointers, which point to the physical address
 *        of the corresponding configuration register
 * @note  see 'ana_cmpr_periph.c' for the device instance
 */

type AnalogCmprDevT struct {
	PadCompConfig *GpioPadCompConfigRegT
	PadCompFilter *GpioPadCompFilterRegT
	IntSt         *GpioExtIntStRegT
	IntEna        *GpioExtIntEnaRegT
	IntClr        *GpioExtIntClrRegT
}
