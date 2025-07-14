package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: FIFO Configuration */
/** Type of fifo register
 *  FIFO data register
 */

type LpUartFifoRegT struct {
	Val c.Uint32T
}

/** Type of mem_conf register
 *  UART memory power configuration
 */

type LpUartMemConfRegT struct {
	Val c.Uint32T
}

/** Type of tout_conf_sync register
 *  UART threshold and allocation configuration
 */

type LpUartToutConfSyncRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  Raw interrupt status
 */

type LpUartIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Masked interrupt status
 */

type LpUartIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable bits
 */

type LpUartIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type LpUartIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Configuration Register */
/** Type of clkdiv_sync register
 *  Clock divider configuration
 */

type LpUartClkdivSyncRegT struct {
	Val c.Uint32T
}

/** Type of rx_filt register
 *  Rx Filter configuration
 */

type LpUartRxFiltRegT struct {
	Val c.Uint32T
}

/** Type of conf0_sync register
 *  Configuration register 0
 */

type LpUartConf0SyncRegT struct {
	Val c.Uint32T
}

/** Type of conf1 register
 *  Configuration register 1
 */

type LpUartConf1RegT struct {
	Val c.Uint32T
}

/** Type of hwfc_conf_sync register
 *  Hardware flow-control configuration
 */

type LpUartHwfcConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of sleep_conf0 register
 *  UART sleep configure register 0
 */

type LpUartSleepConf0RegT struct {
	Val c.Uint32T
}

/** Type of sleep_conf1 register
 *  UART sleep configure register 1
 */

type LpUartSleepConf1RegT struct {
	Val c.Uint32T
}

/** Type of sleep_conf2 register
 *  UART sleep configure register 2
 */

type LpUartSleepConf2RegT struct {
	Val c.Uint32T
}

/** Type of swfc_conf0_sync register
 *  Software flow-control character configuration
 */

type LpUartSwfcConf0SyncRegT struct {
	Val c.Uint32T
}

/** Type of swfc_conf1 register
 *  Software flow-control character configuration
 */

type LpUartSwfcConf1RegT struct {
	Val c.Uint32T
}

/** Type of txbrk_conf_sync register
 *  Tx Break character configuration
 */

type LpUartTxbrkConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of idle_conf_sync register
 *  Frame-end idle configuration
 */

type LpUartIdleConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of rs485_conf_sync register
 *  RS485 mode configuration
 */

type LpUartRs485ConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of clk_conf register
 *  UART core clock configuration
 */

type LpUartClkConfRegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of status register
 *  UART status register
 */

type LpUartStatusRegT struct {
	Val c.Uint32T
}

/** Type of mem_tx_status register
 *  Tx-SRAM write and read offset address.
 */

type LpUartMemTxStatusRegT struct {
	Val c.Uint32T
}

/** Type of mem_rx_status register
 *  Rx-SRAM write and read offset address.
 */

type LpUartMemRxStatusRegT struct {
	Val c.Uint32T
}

/** Type of fsm_status register
 *  UART transmit and receive status.
 */

type LpUartFsmStatusRegT struct {
	Val c.Uint32T
}

/** Type of afifo_status register
 *  UART AFIFO Status
 */

type LpUartAfifoStatusRegT struct {
	Val c.Uint32T
}

/** Group: AT Escape Sequence Selection Configuration */
/** Type of at_cmd_precnt_sync register
 *  Pre-sequence timing configuration
 */

type LpUartAtCmdPrecntSyncRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_postcnt_sync register
 *  Post-sequence timing configuration
 */

type LpUartAtCmdPostcntSyncRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_gaptout_sync register
 *  Timeout configuration
 */

type LpUartAtCmdGaptoutSyncRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_char_sync register
 *  AT escape sequence detection configuration
 */

type LpUartAtCmdCharSyncRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  UART Version register
 */

type LpUartDateRegT struct {
	Val c.Uint32T
}

/** Type of reg_update register
 *  UART Registers Configuration Update register
 */

type LpUartRegUpdateRegT struct {
	Val c.Uint32T
}

/** Type of id register
 *  UART ID register
 */

type LpUartIdRegT struct {
	Val c.Uint32T
}

type LpUartDevT struct {
	Fifo             LpUartFifoRegT
	IntRaw           LpUartIntRawRegT
	IntSt            LpUartIntStRegT
	IntEna           LpUartIntEnaRegT
	IntClr           LpUartIntClrRegT
	ClkdivSync       LpUartClkdivSyncRegT
	RxFilt           LpUartRxFiltRegT
	Status           LpUartStatusRegT
	Conf0Sync        LpUartConf0SyncRegT
	Conf1            LpUartConf1RegT
	Reserved028      c.Uint32T
	HwfcConfSync     LpUartHwfcConfSyncRegT
	SleepConf0       LpUartSleepConf0RegT
	SleepConf1       LpUartSleepConf1RegT
	SleepConf2       LpUartSleepConf2RegT
	SwfcConf0Sync    LpUartSwfcConf0SyncRegT
	SwfcConf1        LpUartSwfcConf1RegT
	TxbrkConfSync    LpUartTxbrkConfSyncRegT
	IdleConfSync     LpUartIdleConfSyncRegT
	Rs485ConfSync    LpUartRs485ConfSyncRegT
	AtCmdPrecntSync  LpUartAtCmdPrecntSyncRegT
	AtCmdPostcntSync LpUartAtCmdPostcntSyncRegT
	AtCmdGaptoutSync LpUartAtCmdGaptoutSyncRegT
	AtCmdCharSync    LpUartAtCmdCharSyncRegT
	MemConf          LpUartMemConfRegT
	ToutConfSync     LpUartToutConfSyncRegT
	MemTxStatus      LpUartMemTxStatusRegT
	MemRxStatus      LpUartMemRxStatusRegT
	FsmStatus        LpUartFsmStatusRegT
	Reserved074      [5]c.Uint32T
	ClkConf          LpUartClkConfRegT
	Date             LpUartDateRegT
	AfifoStatus      LpUartAfifoStatusRegT
	Reserved094      c.Uint32T
	RegUpdate        LpUartRegUpdateRegT
	Id               LpUartIdRegT
}
