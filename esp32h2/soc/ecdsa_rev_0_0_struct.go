package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration registers */
/** Type of conf register
 *  ECDSA configure register
 */

type EcdsaRev00ConfRegT struct {
	Val c.Uint32T
}

/** Type of start register
 *  ECDSA start register
 */

type EcdsaRev00StartRegT struct {
	Val c.Uint32T
}

/** Group: Clock and reset registers */
/** Type of clk register
 *  ECDSA clock gate register
 */

type EcdsaRev00ClkRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of int_raw register
 *  ECDSA interrupt raw register, valid in level.
 */

type EcdsaRev00IntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  ECDSA interrupt status register.
 */

type EcdsaRev00IntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  ECDSA interrupt enable register.
 */

type EcdsaRev00IntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  ECDSA interrupt clear register.
 */

type EcdsaRev00IntClrRegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of state register
 *  ECDSA status register
 */

type EcdsaRev00StateRegT struct {
	Val c.Uint32T
}

/** Group: Result registers */
/** Type of result register
 *  ECDSA result register
 */

type EcdsaRev00ResultRegT struct {
	Val c.Uint32T
}

/** Group: SHA register */
/** Type of sha_mode register
 *  ECDSA control SHA register
 */

type EcdsaRev00ShaModeRegT struct {
	Val c.Uint32T
}

/** Type of sha_start register
 *  ECDSA control SHA register
 */

type EcdsaRev00ShaStartRegT struct {
	Val c.Uint32T
}

/** Type of sha_continue register
 *  ECDSA control SHA register
 */

type EcdsaRev00ShaContinueRegT struct {
	Val c.Uint32T
}

/** Type of sha_busy register
 *  ECDSA status register
 */

type EcdsaRev00ShaBusyRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of date register
 *  Version control register
 */

type EcdsaRev00DateRegT struct {
	Val c.Uint32T
}

/**
 *  ECDSA message register
 */

type EcdsaRev00MessageRegT struct {
	Message [8]c.Uint32T
}

/**
 *  ECDSA memory register
 */

type EcdsaRev00MemRegT struct {
	R   [8]c.Uint32T
	S   [8]c.Uint32T
	Z   [8]c.Uint32T
	Qax [8]c.Uint32T
	Qay [8]c.Uint32T
}

type EcdsaDevRev00T struct {
	Reserved000 c.Uint32T
	Conf        EcdsaRev00ConfRegT
	Clk         EcdsaRev00ClkRegT
	IntRaw      EcdsaRev00IntRawRegT
	IntSt       EcdsaRev00IntStRegT
	IntEna      EcdsaRev00IntEnaRegT
	IntClr      EcdsaRev00IntClrRegT
	Start       EcdsaRev00StartRegT
	State       EcdsaRev00StateRegT
	Result      EcdsaRev00ResultRegT
	Reserved028 [53]c.Uint32T
	Date        EcdsaRev00DateRegT
	Reserved100 [64]c.Uint32T
	ShaMode     EcdsaRev00ShaModeRegT
	Reserved204 [3]c.Uint32T
	ShaStart    EcdsaRev00ShaStartRegT
	ShaContinue EcdsaRev00ShaContinueRegT
	ShaBusy     EcdsaRev00ShaBusyRegT
	Reserved21c [25]c.Uint32T
	Message     EcdsaRev00MessageRegT
	Reserved2a0 [472]c.Uint32T
	Mem         EcdsaRev00MemRegT
}
