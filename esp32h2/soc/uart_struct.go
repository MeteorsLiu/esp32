package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: FIFO Configuration */
/** Type of fifo register
 *  FIFO data register
 */

type UartFifoRegT struct {
	Val c.Uint32T
}

/** Type of mem_conf register
 *  UART memory power configuration
 */

type UartMemConfRegT struct {
	Val c.Uint32T
}

/** Type of tout_conf_sync register
 *  UART threshold and allocation configuration
 */

type UartToutConfSyncRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  Raw interrupt status
 */

type UartIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  Masked interrupt status
 */

type UartIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  Interrupt enable bits
 */

type UartIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  Interrupt clear bits
 */

type UartIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Configuration Register */
/** Type of clkdiv_sync register
 *  Clock divider configuration
 */

type UartClkdivSyncRegT struct {
	Val c.Uint32T
}

/** Type of rx_filt register
 *  Rx Filter configuration
 */

type UartRxFiltRegT struct {
	Val c.Uint32T
}

/** Type of conf0_sync register
 *  a
 */

type UartConf0SyncRegT struct {
	Val c.Uint32T
}

/** Type of conf1 register
 *  Configuration register 1
 */

type UartConf1RegT struct {
	Val c.Uint32T
}

/** Type of hwfc_conf_sync register
 *  Hardware flow-control configuration
 */

type UartHwfcConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of sleep_conf0 register
 *  UART sleep configure register 0
 */

type UartSleepConf0RegT struct {
	Val c.Uint32T
}

/** Type of sleep_conf1 register
 *  UART sleep configure register 1
 */

type UartSleepConf1RegT struct {
	Val c.Uint32T
}

/** Type of sleep_conf2 register
 *  UART sleep configure register 2
 */

type UartSleepConf2RegT struct {
	Val c.Uint32T
}

/** Type of swfc_conf0_sync register
 *  Software flow-control character configuration
 */

type UartSwfcConf0SyncRegT struct {
	Val c.Uint32T
}

/** Type of swfc_conf1 register
 *  Software flow-control character configuration
 */

type UartSwfcConf1RegT struct {
	Val c.Uint32T
}

/** Type of txbrk_conf_sync register
 *  Tx Break character configuration
 */

type UartTxbrkConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of idle_conf_sync register
 *  Frame-end idle configuration
 */

type UartIdleConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of rs485_conf_sync register
 *  RS485 mode configuration
 */

type UartRs485ConfSyncRegT struct {
	Val c.Uint32T
}

/** Type of clk_conf register
 *  UART core clock configuration
 */

type UartClkConfRegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of status register
 *  UART status register
 */

type UartStatusRegT struct {
	Val c.Uint32T
}

/** Type of mem_tx_status register
 *  Tx-SRAM write and read offset address.
 */

type UartMemTxStatusRegT struct {
	Val c.Uint32T
}

/** Type of mem_rx_status register
 *  Rx-SRAM write and read offset address.
 */

type UartMemRxStatusRegT struct {
	Val c.Uint32T
}

/** Type of fsm_status register
 *  UART transmit and receive status.
 */

type UartFsmStatusRegT struct {
	Val c.Uint32T
}

/** Type of afifo_status register
 *  UART AFIFO Status
 */

type UartAfifoStatusRegT struct {
	Val c.Uint32T
}

/** Group: AT Escape Sequence Selection Configuration */
/** Type of at_cmd_precnt_sync register
 *  Pre-sequence timing configuration
 */

type UartAtCmdPrecntSyncRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_postcnt_sync register
 *  Post-sequence timing configuration
 */

type UartAtCmdPostcntSyncRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_gaptout_sync register
 *  Timeout configuration
 */

type UartAtCmdGaptoutSyncRegT struct {
	Val c.Uint32T
}

/** Type of at_cmd_char_sync register
 *  AT escape sequence detection configuration
 */

type UartAtCmdCharSyncRegT struct {
	Val c.Uint32T
}

/** Group: Autobaud Register */
/** Type of pospulse register
 *  Autobaud high pulse register
 */

type UartPospulseRegT struct {
	Val c.Uint32T
}

/** Type of negpulse register
 *  Autobaud low pulse register
 */

type UartNegpulseRegT struct {
	Val c.Uint32T
}

/** Type of lowpulse register
 *  Autobaud minimum low pulse duration register
 */

type UartLowpulseRegT struct {
	Val c.Uint32T
}

/** Type of highpulse register
 *  Autobaud minimum high pulse duration register
 */

type UartHighpulseRegT struct {
	Val c.Uint32T
}

/** Type of rxd_cnt register
 *  Autobaud edge change count register
 */

type UartRxdCntRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  UART Version register
 */

type UartDateRegT struct {
	Val c.Uint32T
}

/** Type of reg_update register
 *  UART Registers Configuration Update register
 */

type UartRegUpdateRegT struct {
	Val c.Uint32T
}

/** Type of id register
 *  UART ID register
 */

type UartIdRegT struct {
	Val c.Uint32T
}

type UartDevS struct {
	Fifo             UartFifoRegT
	IntRaw           UartIntRawRegT
	IntSt            UartIntStRegT
	IntEna           UartIntEnaRegT
	IntClr           UartIntClrRegT
	ClkdivSync       UartClkdivSyncRegT
	RxFilt           UartRxFiltRegT
	Status           UartStatusRegT
	Conf0Sync        UartConf0SyncRegT
	Conf1            UartConf1RegT
	Reserved028      c.Uint32T
	HwfcConfSync     UartHwfcConfSyncRegT
	SleepConf0       UartSleepConf0RegT
	SleepConf1       UartSleepConf1RegT
	SleepConf2       UartSleepConf2RegT
	SwfcConf0Sync    UartSwfcConf0SyncRegT
	SwfcConf1        UartSwfcConf1RegT
	TxbrkConfSync    UartTxbrkConfSyncRegT
	IdleConfSync     UartIdleConfSyncRegT
	Rs485ConfSync    UartRs485ConfSyncRegT
	AtCmdPrecntSync  UartAtCmdPrecntSyncRegT
	AtCmdPostcntSync UartAtCmdPostcntSyncRegT
	AtCmdGaptoutSync UartAtCmdGaptoutSyncRegT
	AtCmdCharSync    UartAtCmdCharSyncRegT
	MemConf          UartMemConfRegT
	ToutConfSync     UartToutConfSyncRegT
	MemTxStatus      UartMemTxStatusRegT
	MemRxStatus      UartMemRxStatusRegT
	FsmStatus        UartFsmStatusRegT
	Pospulse         UartPospulseRegT
	Negpulse         UartNegpulseRegT
	Lowpulse         UartLowpulseRegT
	Highpulse        UartHighpulseRegT
	RxdCnt           UartRxdCntRegT
	ClkConf          UartClkConfRegT
	Date             UartDateRegT
	AfifoStatus      UartAfifoStatusRegT
	Reserved094      c.Uint32T
	RegUpdate        UartRegUpdateRegT
	Id               UartIdRegT
}
type UartDevT UartDevS
