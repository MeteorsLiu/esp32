package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration registers */
/** Type of conf register
 *  ECDSA configure register
 */

type EcdsaConfRegT struct {
	Val c.Uint32T
}

/** Type of start register
 *  ECDSA start register
 */

type EcdsaStartRegT struct {
	Val c.Uint32T
}

/** Group: Clock and reset registers */
/** Type of clk register
 *  ECDSA clock gate register
 */

type EcdsaClkRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_raw register
 *  ECDSA interrupt raw register, valid in level.
 */

type EcdsaIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  ECDSA interrupt status register.
 */

type EcdsaIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  ECDSA interrupt enable register.
 */

type EcdsaIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  ECDSA interrupt clear register.
 */

type EcdsaIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of state register
 *  ECDSA status register
 */

type EcdsaStateRegT struct {
	Val c.Uint32T
}

/** Group: Result registers */
/** Type of result register
 *  ECDSA result register
 */

type EcdsaResultRegT struct {
	Val c.Uint32T
}

/** Group: SHA register */
/** Type of sha_mode register
 *  ECDSA control SHA register
 */

type EcdsaShaModeRegT struct {
	Val c.Uint32T
}

/** Type of sha_start register
 *  ECDSA control SHA register
 */

type EcdsaShaStartRegT struct {
	Val c.Uint32T
}

/** Type of sha_continue register
 *  ECDSA control SHA register
 */

type EcdsaShaContinueRegT struct {
	Val c.Uint32T
}

/** Type of sha_busy register
 *  ECDSA status register
 */

type EcdsaShaBusyRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type EcdsaDateRegT struct {
	Val c.Uint32T
}

/**
 *  ECDSA message register
 */

type EcdsaMessageRegT struct {
	Message [8]c.Uint32T
}

/**
 *  ECDSA memory register
 */

type EcdsaMemRegT struct {
	R   [8]c.Uint32T
	S   [8]c.Uint32T
	Z   [8]c.Uint32T
	Qax [8]c.Uint32T
	Qay [8]c.Uint32T
}

type EcdsaDevRev12T struct {
	Reserved000 c.Uint32T
	Conf        EcdsaConfRegT
	Clk         EcdsaClkRegT
	IntRaw      EcdsaIntRawRegT
	IntSt       EcdsaIntStRegT
	IntEna      EcdsaIntEnaRegT
	IntClr      EcdsaIntClrRegT
	Start       EcdsaStartRegT
	State       EcdsaStateRegT
	Result      EcdsaResultRegT
	Reserved028 [53]c.Uint32T
	Date        EcdsaDateRegT
	Reserved100 [64]c.Uint32T
	ShaMode     EcdsaShaModeRegT
	Reserved204 [3]c.Uint32T
	ShaStart    EcdsaShaStartRegT
	ShaContinue EcdsaShaContinueRegT
	ShaBusy     EcdsaShaBusyRegT
	Reserved21c [25]c.Uint32T
	Message     EcdsaMessageRegT
	Reserved2a0 [40]c.Uint32T
	Mem         EcdsaMemRegT
	Reserved300 [432]c.Uint32T
}
