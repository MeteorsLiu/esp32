package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief I2C hal Context definition
 */

type I2cHalContextT struct {
	Dev *I2cDevT
}

/**
 * @brief I2C hal clock configurations
 */

type I2cHalSclkInfoT struct {
	ClkSel    c.Uint8T
	ClkActive c.Uint8T
	ClkDiv    HalUtilsClkDivT
}

/**
 * @brief Timing configuration structure. Used for I2C reset internally.
 */

type I2cHalTimingConfigT struct {
	HighPeriod     c.Int
	LowPeriod      c.Int
	WaitHighPeriod c.Int
	RstartSetup    c.Int
	StartHold      c.Int
	StopSetup      c.Int
	StopHold       c.Int
	SdaSample      c.Int
	SdaHold        c.Int
	Timeout        c.Int
	ClkCfg         I2cHalSclkInfoT
}

/**
 * @brief  Init the I2C slave.
 *
 * @param  hal Context of the HAL layer
 * @param  i2c_num I2C port number
 *
 * @return None
 */
// llgo:link (*I2cHalContextT).I2cHalSlaveInit C.i2c_hal_slave_init
func (recv_ *I2cHalContextT) I2cHalSlaveInit() {
}

/**
 * @brief  Init the I2C master.
 *
 * @param  hal Context of the HAL layer
 * @param  i2c_num I2C port number
 *
 * @return None
 */
// llgo:link (*I2cHalContextT).I2cHalMasterInit C.i2c_hal_master_init
func (recv_ *I2cHalContextT) I2cHalMasterInit() {
}

/**
 * @brief  I2C hardware FSM reset
 *
 * @param  hal Context of the HAL layer
 *
 * @return None
 */
// llgo:link (*I2cHalContextT).I2cHalMasterFsmRst C.i2c_hal_master_fsm_rst
func (recv_ *I2cHalContextT) I2cHalMasterFsmRst() {
}

/**
 * @brief  I2C master handle tx interrupt event
 *
 * @param  hal Context of the HAL layer
 * @param  event Pointer to accept the interrupt event
 *
 * @return None
 */
// llgo:link (*I2cHalContextT).I2cHalMasterHandleTxEvent C.i2c_hal_master_handle_tx_event
func (recv_ *I2cHalContextT) I2cHalMasterHandleTxEvent(event *I2cIntrEventT) {
}

/**
 * @brief  I2C master handle rx interrupt event
 *
 * @param  hal Context of the HAL layer
 * @param  event Pointer to accept the interrupt event
 *
 * @return None
 */
// llgo:link (*I2cHalContextT).I2cHalMasterHandleRxEvent C.i2c_hal_master_handle_rx_event
func (recv_ *I2cHalContextT) I2cHalMasterHandleRxEvent(event *I2cIntrEventT) {
}

/**
 * @brief Set scl timeout reg value according to given timeout us and source clock frequency
 *
 * @param hal Context of the HAL layer
 * @param timeout_us timeout us
 * @param sclk_clock_hz source clock hz
 */
// llgo:link (*I2cHalContextT).I2cHalMasterSetSclTimeoutVal C.i2c_hal_master_set_scl_timeout_val
func (recv_ *I2cHalContextT) I2cHalMasterSetSclTimeoutVal(timeout_us c.Uint32T, sclk_clock_hz c.Uint32T) {
}

/**
 * @brief Start I2C master transaction
 *
 * @param hal Context of the HAL
 */
// llgo:link (*I2cHalContextT).I2cHalMasterTransStart C.i2c_hal_master_trans_start
func (recv_ *I2cHalContextT) I2cHalMasterTransStart() {
}

/**
 * @brief Get timing configuration
 *
 * @param hal Context of the HAL
 * @param timing_config Pointer to timing config structure.
 */
// llgo:link (*I2cHalContextT).I2cHalGetTimingConfig C.i2c_hal_get_timing_config
func (recv_ *I2cHalContextT) I2cHalGetTimingConfig(timing_config *I2cHalTimingConfigT) {
}

/**
 * @brief Set timing configuration
 *
 * @param hal Context of the HAL
 * @param timing_config Timing config structure.
 */
// llgo:link (*I2cHalContextT).I2cHalSetTimingConfig C.i2c_hal_set_timing_config
func (recv_ *I2cHalContextT) I2cHalSetTimingConfig(timing_config *I2cHalTimingConfigT) {
}
