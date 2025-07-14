package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of out_data register
 *  need des
 */

type LpIoOutDataRegT struct {
	Val c.Uint32T
}

/** Type of out_data_w1ts register
 *  need des
 */

type LpIoOutDataW1tsRegT struct {
	Val c.Uint32T
}

/** Type of out_data_w1tc register
 *  need des
 */

type LpIoOutDataW1tcRegT struct {
	Val c.Uint32T
}

/** Type of out_enable register
 *  need des
 */

type LpIoOutEnableRegT struct {
	Val c.Uint32T
}

/** Type of out_enable_w1ts register
 *  need des
 */

type LpIoOutEnableW1tsRegT struct {
	Val c.Uint32T
}

/** Type of out_enable_w1tc register
 *  need des
 */

type LpIoOutEnableW1tcRegT struct {
	Val c.Uint32T
}

/** Type of status register
 *  need des
 */

type LpIoStatusRegT struct {
	Val c.Uint32T
}

/** Type of status_w1ts register
 *  need des
 */

type LpIoStatusW1tsRegT struct {
	Val c.Uint32T
}

/** Type of status_w1tc register
 *  need des
 */

type LpIoStatusW1tcRegT struct {
	Val c.Uint32T
}

/** Type of in register
 *  need des
 */

type LpIoInRegT struct {
	Val c.Uint32T
}

/** Type of pin register
 *  need des
 */

type LpIoPinRegT struct {
	Val c.Uint32T
}

/** Type of gpio register
 *  need des
 */

type LpIoGpioRegT struct {
	Val c.Uint32T
}

/** Type of status_interrupt register
 *  need des
 */

type LpIoStatusInterruptRegT struct {
	Val c.Uint32T
}

/** Type of debug_sel0 register
 *  need des
 */

type LpIoDebugSel0RegT struct {
	Val c.Uint32T
}

/** Type of debug_sel1 register
 *  need des
 */

type LpIoDebugSel1RegT struct {
	Val c.Uint32T
}

/** Type of lpi2c register
 *  need des
 */

type LpIoLpi2cRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  need des
 */

type LpIoDateRegT struct {
	Val c.Uint32T
}

type LpIoDevT struct {
	OutData         LpIoOutDataRegT
	OutDataW1ts     LpIoOutDataW1tsRegT
	OutDataW1tc     LpIoOutDataW1tcRegT
	OutEnable       LpIoOutEnableRegT
	OutEnableW1ts   LpIoOutEnableW1tsRegT
	OutEnableW1tc   LpIoOutEnableW1tcRegT
	Status          LpIoStatusRegT
	StatusW1ts      LpIoStatusW1tsRegT
	StatusW1tc      LpIoStatusW1tcRegT
	In              LpIoInRegT
	Pin             [8]LpIoPinRegT
	Gpio            [8]LpIoGpioRegT
	StatusInterrupt LpIoStatusInterruptRegT
	DebugSel0       LpIoDebugSel0RegT
	DebugSel1       LpIoDebugSel1RegT
	Lpi2c           LpIoLpi2cRegT
	Reserved078     [225]c.Uint32T
	Date            LpIoDateRegT
}
