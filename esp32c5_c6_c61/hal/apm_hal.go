package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Region configuration data.
 */

type ApmCtrlRegionConfigDataT struct {
	RegnNum       c.Uint32T
	RegnStartAddr c.Uint32T
	RegnEndAddr   c.Uint32T
	RegnPms       c.Uint32T
	FilterEnable  bool
	ApmCtrl       ApmLlApmCtrlT
	SecMode       ApmLlSecureModeT
}

/**
 * @brief Secure mode(TEE/REE[0:2] configuration data.
 */

type ApmCtrlSecureModeConfigT struct {
	ApmCtrl   ApmLlApmCtrlT
	SecMode   ApmLlSecureModeT
	ApmMCnt   c.Uint8T
	RegnCount c.Uint32T
	MasterIds c.Uint32T
	PmsData   *ApmCtrlRegionConfigDataT
}

/**
 * @brief Set secure mode
 *
 * @param apm_ctrl APM Ctrl to be configured
 * @param master_id APM master ID
 * @param sec_mode Secure mode
 */
// llgo:link ApmLlApmCtrlT.ApmTeeHalSetMasterSecureMode C.apm_tee_hal_set_master_secure_mode
func (recv_ ApmLlApmCtrlT) ApmTeeHalSetMasterSecureMode(master_id ApmLlMasterIdT, sec_mode ApmLlSecureModeT) {
}

/**
 * @brief TEE controller clock auto gating enable
 *
 * @param enable Flag for HP clock auto gating enable/disable
 */
//go:linkname ApmTeeHalClkGatingEnable C.apm_tee_hal_clk_gating_enable
func ApmTeeHalClkGatingEnable(enable bool)

/**
 * @brief enable/disable APM Ctrl Region access permission filter
 *
 * @param apm_ctrl APM Ctrl to be configured
 * @param regn_num Memory Region number
 * @param enable Flag for Region access filter enable/disable
 */
// llgo:link ApmLlApmCtrlT.ApmHalApmCtrlRegionFilterEnable C.apm_hal_apm_ctrl_region_filter_enable
func (recv_ ApmLlApmCtrlT) ApmHalApmCtrlRegionFilterEnable(regn_num c.Uint32T, enable bool) {
}

/**
 * @brief enable/disable APM Ctrl access path(M[0:n])
 *
 * @param apm_path   APM controller and access path to be configured
 * @param enable     Flag for M path filter enable/disable
 */
// llgo:link (*ApmCtrlPathT).ApmHalApmCtrlFilterEnable C.apm_hal_apm_ctrl_filter_enable
func (recv_ *ApmCtrlPathT) ApmHalApmCtrlFilterEnable(enable bool) {
}

/**
 * @brief enable/disable all available APM Ctrl access path(M[0:n])
 *
 * @param enable Flag for M path filter enable/disable
 */
//go:linkname ApmHalApmCtrlFilterEnableAll C.apm_hal_apm_ctrl_filter_enable_all
func ApmHalApmCtrlFilterEnableAll(enable bool)

/**
 * @brief Region configuration
 *
 * @param pms_data Region configuration data
 */
// llgo:link (*ApmCtrlRegionConfigDataT).ApmHalApmCtrlRegionConfig C.apm_hal_apm_ctrl_region_config
func (recv_ *ApmCtrlRegionConfigDataT) ApmHalApmCtrlRegionConfig() {
}

/**
 * @brief Get APM Ctrl access path(M[0:n]) exception status
 *
 * @param apm_path   APM controller and access path to be configured
 */
// llgo:link (*ApmCtrlPathT).ApmHalApmCtrlExceptionStatus C.apm_hal_apm_ctrl_exception_status
func (recv_ *ApmCtrlPathT) ApmHalApmCtrlExceptionStatus() c.Uint8T {
	return 0
}

/**
 * @brief Clear APM Ctrl access path(M[0:n]) exception
 *
 * @param apm_path   APM controller and access path to be configured
 */
// llgo:link (*ApmCtrlPathT).ApmHalApmCtrlExceptionClear C.apm_hal_apm_ctrl_exception_clear
func (recv_ *ApmCtrlPathT) ApmHalApmCtrlExceptionClear() {
}

/**
 * @brief Get APM Ctrl access path exception information
 *
 * @param excp_info Exception related information like addr,
 * region, amp_ctrl, apm_m_path, sec_mode and master id
 */
// llgo:link (*ApmCtrlExceptionInfoT).ApmHalApmCtrlGetExceptionInfo C.apm_hal_apm_ctrl_get_exception_info
func (recv_ *ApmCtrlExceptionInfoT) ApmHalApmCtrlGetExceptionInfo() {
}

/**
 * @brief APM Ctrl interrupt enable for access path(M[0:n])
 *
 * @param apm_path   APM controller and access path to be configured
 * @param enable     Flag for access path interrupt enable/disable
 */
// llgo:link (*ApmCtrlPathT).ApmHalApmCtrlInterruptEnable C.apm_hal_apm_ctrl_interrupt_enable
func (recv_ *ApmCtrlPathT) ApmHalApmCtrlInterruptEnable(enable bool) {
}

/**
 * @brief APM Ctrl clock auto gating enable
 *
 * @apm_ctrl     APM Ctrl
 * @param enable Flag for HP clock auto gating enable/disable
 */
// llgo:link ApmLlApmCtrlT.ApmHalApmCtrlClkGatingEnable C.apm_hal_apm_ctrl_clk_gating_enable
func (recv_ ApmLlApmCtrlT) ApmHalApmCtrlClkGatingEnable(enable bool) {
}

/**
 * @brief TEE/REE execution environment configuration.
 *
 * This API will be called from TEE mode initialization code which is
 * responsible to setup TEE/REE execution environment.
 * It includes, allocation of all bus masters, memory ranges and other
 * peripherals to the given secure mode.
 * All this information should be passed by the TEE mode initialization code.
 *
 * @sec_mode_data APM Ctl configuration data.
 */
// llgo:link (*ApmCtrlSecureModeConfigT).ApmHalApmCtrlMasterSecModeConfig C.apm_hal_apm_ctrl_master_sec_mode_config
func (recv_ *ApmCtrlSecureModeConfigT) ApmHalApmCtrlMasterSecModeConfig() {
}

/**
 * @brief APM/TEE/HP System Reg reset event bypass enable
 *
 * Disable: tee_reg/apm_reg/hp_system_reg will not only be reset by power-reset,
 * but also some reset events.
 * Enable: tee_reg/apm_reg/hp_system_reg will only be reset by power-reset.
 * Some reset events will be bypassed.
 *
 * @param enable   Flag for event bypass enable/disable
 */
//go:linkname ApmHalApmCtrlResetEventEnable C.apm_hal_apm_ctrl_reset_event_enable
func ApmHalApmCtrlResetEventEnable(enable bool)

/**
 * @brief Fetch the APM Ctrl access path interrupt source number.
 *
 * @param apm_path   APM controller and access path to be configured
 *
 * @return
 *      - valid interrupt source number on success
 *      - -1: invalid interrupt source
 */
// llgo:link (*ApmCtrlPathT).ApmHalApmCtrlGetIntSrcNum C.apm_hal_apm_ctrl_get_int_src_num
func (recv_ *ApmCtrlPathT) ApmHalApmCtrlGetIntSrcNum() c.Int {
	return 0
}
