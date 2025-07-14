package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Interrupt registers */
/** Type of int_raw register
 *  ECC interrupt raw register, valid in level.
 */

type EccMultIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  ECC interrupt status register.
 */

type EccMultIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  ECC interrupt enable register.
 */

type EccMultIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  ECC interrupt clear register.
 */

type EccMultIntClrRegT struct {
	Val c.Uint32T
}

/** Group: RX Control and configuration registers */
/** Type of conf register
 *  ECC configure register
 */

type EccMultConfRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type EccMultDateRegT struct {
	Val c.Uint32T
}

type EccMultDevT struct {
	Reserved000 [3]c.Uint32T
	IntRaw      EccMultIntRawRegT
	IntSt       EccMultIntStRegT
	IntEna      EccMultIntEnaRegT
	IntClr      EccMultIntClrRegT
	Conf        EccMultConfRegT
	Reserved020 [55]c.Uint32T
	Date        EccMultDateRegT
	K           [8]c.Uint32T
	Px          [8]c.Uint32T
	Py          [8]c.Uint32T
}
