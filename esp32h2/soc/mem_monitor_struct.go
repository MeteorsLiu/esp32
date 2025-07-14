package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configuration registers */
/** Type of log_setting register
 *  log config register
 */

type MemMonitorLogSettingRegT struct {
	Val c.Uint32T
}

/** Type of log_check_data register
 *  check data register
 */

type MemMonitorLogCheckDataRegT struct {
	Val c.Uint32T
}

/** Type of log_data_mask register
 *  check data mask register
 */

type MemMonitorLogDataMaskRegT struct {
	Val c.Uint32T
}

/** Type of log_min register
 *  log boundary register
 */

type MemMonitorLogMinRegT struct {
	Val c.Uint32T
}

/** Type of log_max register
 *  log boundary register
 */

type MemMonitorLogMaxRegT struct {
	Val c.Uint32T
}

/** Type of log_mem_start register
 *  log message store range register
 */

type MemMonitorLogMemStartRegT struct {
	Val c.Uint32T
}

/** Type of log_mem_end register
 *  log message store range register
 */

type MemMonitorLogMemEndRegT struct {
	Val c.Uint32T
}

/** Type of log_mem_current_addr register
 *  current writing address.
 */

type MemMonitorLogMemCurrentAddrRegT struct {
	Val c.Uint32T
}

/** Type of log_mem_addr_update register
 *  writing address update
 */

type MemMonitorLogMemAddrUpdateRegT struct {
	Val c.Uint32T
}

/** Type of log_mem_full_flag register
 *  full flag status register
 */

type MemMonitorLogMemFullFlagRegT struct {
	Val c.Uint32T
}

/** Group: clk register */
/** Type of clock_gate register
 *  clock gate force on register
 */

type MemMonitorClockGateRegT struct {
	Val c.Uint32T
}

/** Group: version register */
/** Type of date register
 *  version register
 */

type MemMonitorDateRegT struct {
	Val c.Uint32T
}

type MemMonitorDevT struct {
	LogSetting        MemMonitorLogSettingRegT
	LogCheckData      MemMonitorLogCheckDataRegT
	LogDataMask       MemMonitorLogDataMaskRegT
	LogMin            MemMonitorLogMinRegT
	LogMax            MemMonitorLogMaxRegT
	LogMemStart       MemMonitorLogMemStartRegT
	LogMemEnd         MemMonitorLogMemEndRegT
	LogMemCurrentAddr MemMonitorLogMemCurrentAddrRegT
	LogMemAddrUpdate  MemMonitorLogMemAddrUpdateRegT
	LogMemFullFlag    MemMonitorLogMemFullFlagRegT
	ClockGate         MemMonitorClockGateRegT
	Reserved02c       [244]c.Uint32T
	Date              MemMonitorDateRegT
}
