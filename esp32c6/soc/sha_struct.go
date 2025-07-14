package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of mode register
 *  Initial configuration register.
 */

type ShaModeRegT struct {
	Val c.Uint32T
}

/** Type of t_string register
 *  SHA 512/t configuration register 0.
 */

type ShaTStringRegT struct {
	Val c.Uint32T
}

/** Type of t_length register
 *  SHA 512/t configuration register 1.
 */

type ShaTLengthRegT struct {
	Val c.Uint32T
}

/** Type of dma_block_num register
 *  DMA configuration register 0.
 */

type ShaDmaBlockNumRegT struct {
	Val c.Uint32T
}

/** Type of start register
 *  Typical SHA configuration register 0.
 */

type ShaStartRegT struct {
	Val c.Uint32T
}

/** Type of continue register
 *  Typical SHA configuration register 1.
 */

type ShaContinueRegT struct {
	Val c.Uint32T
}

/** Type of dma_start register
 *  DMA configuration register 1.
 */

type ShaDmaStartRegT struct {
	Val c.Uint32T
}

/** Type of dma_continue register
 *  DMA configuration register 2.
 */

type ShaDmaContinueRegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of busy register
 *  Busy register.
 */

type ShaBusyRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of clear_irq register
 *  Interrupt clear register.
 */

type ShaClearIrqRegT struct {
	Val c.Uint32T
}

/** Type of irq_ena register
 *  Interrupt enable register.
 */

type ShaIrqEnaRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  Date register.
 */

type ShaDateRegT struct {
	Val c.Uint32T
}

/** Group: memory type */

type ShaDevT struct {
	Mode        ShaModeRegT
	TString     ShaTStringRegT
	TLength     ShaTLengthRegT
	DmaBlockNum ShaDmaBlockNumRegT
	Start       ShaStartRegT
	Conti       ShaContinueRegT
	Busy        ShaBusyRegT
	DmaStart    ShaDmaStartRegT
	DmaContinue ShaDmaContinueRegT
	ClearIrq    ShaClearIrqRegT
	IrqEna      ShaIrqEnaRegT
	Date        ShaDateRegT
	Reserved030 [4]c.Uint32T
	H           [16]c.Uint32T
	M           [16]c.Uint32T
}
