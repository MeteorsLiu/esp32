package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configure_register */
/** Type of store register
 *  need_des
 */

type LpAonStoreRegT struct {
	Val c.Uint32T
}

/** Type of gpio_mux register
 *  need_des
 */

type LpAonGpioMuxRegT struct {
	Val c.Uint32T
}

/** Type of gpio_hold0 register
 *  need_des
 */

type LpAonGpioHold0RegT struct {
	Val c.Uint32T
}

/** Type of gpio_hold1 register
 *  need_des
 */

type LpAonGpioHold1RegT struct {
	Val c.Uint32T
}

/** Type of sys_cfg register
 *  need_des
 */

type LpAonSysCfgRegT struct {
	Val c.Uint32T
}

/** Type of cpucore0_cfg register
 *  need_des
 */

type LpAonCpucore0CfgRegT struct {
	Val c.Uint32T
}

/** Type of io_mux register
 *  need_des
 */

type LpAonIoMuxRegT struct {
	Val c.Uint32T
}

/** Type of ext_wakeup_cntl register
 *  need_des
 */

type LpAonExtWakeupCntlRegT struct {
	Val c.Uint32T
}

/** Type of usb register
 *  need_des
 */

type LpAonUsbRegT struct {
	Val c.Uint32T
}

/** Type of lpbus register
 *  need_des
 */

type LpAonLpbusRegT struct {
	Val c.Uint32T
}

/** Type of sdio_active register
 *  need_des
 */

type LpAonSdioActiveRegT struct {
	Val c.Uint32T
}

/** Type of lpcore register
 *  need_des
 */

type LpAonLpcoreRegT struct {
	Val c.Uint32T
}

/** Type of sar_cct register
 *  need_des
 */

type LpAonSarCctRegT struct {
	Val c.Uint32T
}

/** Type of jtag_sel register
 *  need_des
 */

type LpAonJtagSelRegT struct {
	Val c.Uint32T
}

/** Type of date register
 *  need_des
 */

type LpAonDateRegT struct {
	Val c.Uint32T
}

type LpAonDevT struct {
	Store         [10]LpAonStoreRegT
	GpioMux       LpAonGpioMuxRegT
	GpioHold0     LpAonGpioHold0RegT
	GpioHold1     LpAonGpioHold1RegT
	SysCfg        LpAonSysCfgRegT
	Cpucore0Cfg   LpAonCpucore0CfgRegT
	IoMux         LpAonIoMuxRegT
	ExtWakeupCntl LpAonExtWakeupCntlRegT
	Usb           LpAonUsbRegT
	Lpbus         LpAonLpbusRegT
	SdioActive    LpAonSdioActiveRegT
	Lpcore        LpAonLpcoreRegT
	SarCct        LpAonSarCctRegT
	JtagSel       LpAonJtagSelRegT
	Reserved05c   [232]c.Uint32T
	Date          LpAonDateRegT
}
