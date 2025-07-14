package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of config0 register
 *  need_des
 */

type LpWdtConfig0RegT struct {
	Val c.Uint32T
}

/** Type of config1 register
 *  need_des
 */

type LpWdtConfig1RegT struct {
	Val c.Uint32T
}

/** Type of config2 register
 *  need_des
 */

type LpWdtConfig2RegT struct {
	Val c.Uint32T
}

/** Type of config3 register
 *  need_des
 */

type LpWdtConfig3RegT struct {
	Val c.Uint32T
}

/** Type of config4 register
 *  need_des
 */

type LpWdtConfig4RegT struct {
	Val c.Uint32T
}

/** Type of config5 register
 *  need_des
 */

type LpWdtConfig5RegT struct {
	Val c.Uint32T
}

/** Type of feed register
 *  need_des
 */

type LpWdtFeedRegT struct {
	Val c.Uint32T
}

/** Type of wprotect register
 *  need_des
 */

type LpWdtWprotectRegT struct {
	Val c.Uint32T
}

/** Type of swd_config register
 *  need_des
 */

type LpWdtSwdConfigRegT struct {
	Val c.Uint32T
}

/** Type of swd_wprotect register
 *  need_des
 */

type LpWdtSwdWprotectRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  need_des
 */

type LpWdtIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  need_des
 */

type LpWdtIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  need_des
 */

type LpWdtIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  need_des
 */

type LpWdtIntClrRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  need_des
 */

type LpWdtDateRegT struct {
	Val c.Uint32T
}

type LpWdtDevT struct {
	Config0     LpWdtConfig0RegT
	Config1     LpWdtConfig1RegT
	Config2     LpWdtConfig2RegT
	Config3     LpWdtConfig3RegT
	Config4     LpWdtConfig4RegT
	Config5     LpWdtConfig5RegT
	Feed        LpWdtFeedRegT
	Wprotect    LpWdtWprotectRegT
	SwdConfig   LpWdtSwdConfigRegT
	SwdWprotect LpWdtSwdWprotectRegT
	IntRaw      LpWdtIntRawRegT
	IntSt       LpWdtIntStRegT
	IntEna      LpWdtIntEnaRegT
	IntClr      LpWdtIntClrRegT
	Reserved038 [241]c.Uint32T
	Date        LpWdtDateRegT
}
