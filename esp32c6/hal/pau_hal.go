package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type PauHalContextT struct {
	Dev *PauDevT
}

/**
 * @brief Set regdma entry link address
 *
 * @param hal           regdma hal context
 * @param link_addr     entry link address value
 */
// llgo:link (*PauHalContextT).PauHalSetRegdmaEntryLinkAddr C.pau_hal_set_regdma_entry_link_addr
func (recv_ *PauHalContextT) PauHalSetRegdmaEntryLinkAddr(link_addr *PauRegdmaLinkAddrT) {
}

/**
 * @brief Start transmission on regdma modem link
 *
 * @param hal           regdma hal context
 * @param backup_or_restore        false:restore true:backup
 */
// llgo:link (*PauHalContextT).PauHalStartRegdmaModemLink C.pau_hal_start_regdma_modem_link
func (recv_ *PauHalContextT) PauHalStartRegdmaModemLink(backup_or_restore bool) {
}

/**
 * @brief Stop transmission on regdma modem link
 *
 * @param hal           regdma hal context
 */
// llgo:link (*PauHalContextT).PauHalStopRegdmaModemLink C.pau_hal_stop_regdma_modem_link
func (recv_ *PauHalContextT) PauHalStopRegdmaModemLink() {
}

/**
 * @brief Start transmission on regdma extra link
 *
 * @param hal           regdma hal context
 * @param backup_or_restore        false:restore true:backup
 */
// llgo:link (*PauHalContextT).PauHalStartRegdmaExtraLink C.pau_hal_start_regdma_extra_link
func (recv_ *PauHalContextT) PauHalStartRegdmaExtraLink(backup_or_restore bool) {
}

/**
 * @brief Stop transmission on regdma extra link
 *
 * @param hal           regdma hal context
 */
// llgo:link (*PauHalContextT).PauHalStopRegdmaExtraLink C.pau_hal_stop_regdma_extra_link
func (recv_ *PauHalContextT) PauHalStopRegdmaExtraLink() {
}

/**
 * @brief Set PAU module link work timeout threshold
 *
 * @param hal           regdma hal context
 * @param loop_num      the maximum number of regdma link loop num
 * @param count         the maximum number of register access timeout
 */
// llgo:link (*PauHalContextT).PauHalSetRegdmaWorkTimeout C.pau_hal_set_regdma_work_timeout
func (recv_ *PauHalContextT) PauHalSetRegdmaWorkTimeout(loop_num c.Uint32T, count c.Uint32T) {
}

/**
 * @brief Set regdma link wait timeout, include wait retry count and register read interval
 *
 * @param hal           regdma hal context
 * @param count         the maximum number of regdma wait retry count
 * @param interval      the interval of regdma wait link to read register
 */
// llgo:link (*PauHalContextT).PauHalSetRegdmaWaitTimeout C.pau_hal_set_regdma_wait_timeout
func (recv_ *PauHalContextT) PauHalSetRegdmaWaitTimeout(count c.Int, interval c.Int) {
}
