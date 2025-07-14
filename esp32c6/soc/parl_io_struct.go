package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: PARL_IO RX Configuration0 */
/** Type of rx_cfg0 register
 *  Parallel RX module configuration register0.
 */

type ParlIoRxCfg0RegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO RX Configuration1 */
/** Type of rx_cfg1 register
 *  Parallel RX module configuration register1.
 */

type ParlIoRxCfg1RegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO TX Configuration0 */
/** Type of tx_cfg0 register
 *  Parallel TX module configuration register0.
 */

type ParlIoTxCfg0RegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO TX Configuration1 */
/** Type of tx_cfg1 register
 *  Parallel TX module configuration register1.
 */

type ParlIoTxCfg1RegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO TX Status0 */
/** Type of st register
 *  Parallel IO module status register0.
 */

type ParlIoStRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Interrupt Configuration and Status */
/** Type of int_ena register
 *  Parallel IO interrupt enable signal configuration register.
 */

type ParlIoIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_raw register
 *  Parallel IO interrupt raw signal status register.
 */

type ParlIoIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Parallel IO interrupt signal status register.
 */

type ParlIoIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Parallel IO interrupt  clear signal configuration register.
 */

type ParlIoIntClrRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Clock Gating Configuration */
/** Type of clk register
 *  Parallel IO clk configuration register
 */

type ParlIoClkRegT struct {
	Val c.Uint32T
}

/** Group: PARL_IO Version Register */
/** Type of version register
 *  Version register.
 */

type ParlIoVersionRegT struct {
	Val c.Uint32T
}

type ParlIoDevT struct {
	RxCfg0      ParlIoRxCfg0RegT
	RxCfg1      ParlIoRxCfg1RegT
	TxCfg0      ParlIoTxCfg0RegT
	TxCfg1      ParlIoTxCfg1RegT
	St          ParlIoStRegT
	IntEna      ParlIoIntEnaRegT
	IntRaw      ParlIoIntRawRegT
	IntSt       ParlIoIntStRegT
	IntClr      ParlIoIntClrRegT
	Reserved024 [63]c.Uint32T
	Clk         ParlIoClkRegT
	Reserved124 [182]c.Uint32T
	Version     ParlIoVersionRegT
}
