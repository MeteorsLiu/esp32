package esp_driver_ana_cmpr

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Analog comparator unit configuration
 */

type AnaCmprConfigT struct {
	Unit         AnaCmprUnitT
	ClkSrc       AnaCmprClkSrcT
	RefSrc       AnaCmprRefSourceT
	CrossType    AnaCmprCrossTypeT
	IntrPriority c.Int
	Flags        struct {
		IoLoopBack c.Uint32T
	}
}

/**
 * @brief Analog comparator internal reference configuration
 */

type AnaCmprInternalRefConfigT struct {
	RefVolt AnaCmprRefVoltageT
}

/**
 * @brief Analog comparator debounce filter configuration
 */

type AnaCmprDebounceConfigT struct {
	WaitUs c.Uint32T
}

/**
 * @brief Group of Analog Comparator callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_ANA_CMPR_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type AnaCmprEventCallbacksT struct {
	OnCross AnaCmprCrossCbT
}

/**
 * @brief Allocating a new analog comparator unit handle
 *
 * @param[in]  config       The config of the analog comparator unit
 * @param[out] ret_cmpr     The returned analog comparator unit handle
 * @return
 *      - ESP_OK                Allocate analog comparator unit handle success
 *      - ESP_ERR_NO_MEM        No memory for the analog comparator structure
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters or wrong unit number
 *      - ESP_ERR_INVALID_STATE The unit has been allocated or the clock source has been occupied
 */
// llgo:link (*AnaCmprConfigT).AnaCmprNewUnit C.ana_cmpr_new_unit
func (recv_ *AnaCmprConfigT) AnaCmprNewUnit(ret_cmpr *AnaCmprHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the analog comparator unit handle
 *
 * @param[in]  cmpr         The handle of analog comparator unit
 * @return
 *      - ESP_OK                Delete analog comparator unit handle success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters or wrong unit number
 *      - ESP_ERR_INVALID_STATE The analog comparator is not disabled yet
 */
//go:linkname AnaCmprDelUnit C.ana_cmpr_del_unit
func AnaCmprDelUnit(cmpr AnaCmprHandleT) EspErrT

/**
 * @brief Set internal reference configuration
 * @note  This function only need to be called when `ana_cmpr_config_t::ref_src`
 *        is ANA_CMPR_REF_SRC_INTERNAL.
 * @note This function is allowed to run within ISR context including intr callbacks
 * @note This function will be placed into IRAM if `CONFIG_ANA_CMPR_CTRL_FUNC_IN_IRAM` is on,
 *       so that it's allowed to be executed when Cache is disabled
 *
 * @param[in]  cmpr         The handle of analog comparator unit
 * @param[in]  ref_cfg      Internal reference configuration
 * @return
 *      - ESP_OK                Set denounce configuration success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters
 *      - ESP_ERR_INVALID_STATE The reference source is not `ANA_CMPR_REF_SRC_INTERNAL`
 */
//go:linkname AnaCmprSetInternalReference C.ana_cmpr_set_internal_reference
func AnaCmprSetInternalReference(cmpr AnaCmprHandleT, ref_cfg *AnaCmprInternalRefConfigT) EspErrT

/**
 * @brief Set debounce configuration to the analog comparator
 * @note This function is allowed to run within ISR context including intr callbacks
 * @note This function will be placed into IRAM if `CONFIG_ANA_CMPR_CTRL_FUNC_IN_IRAM` is on,
 *       so that it's allowed to be executed when Cache is disabled
 *
 * @param[in]  cmpr         The handle of analog comparator unit
 * @param[in]  dbc_cfg      Debounce configuration
 * @return
 *      - ESP_OK                Set denounce configuration success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters
 */
//go:linkname AnaCmprSetDebounce C.ana_cmpr_set_debounce
func AnaCmprSetDebounce(cmpr AnaCmprHandleT, dbc_cfg *AnaCmprDebounceConfigT) EspErrT

/**
 * @brief Set the source signal cross type
 * @note The initial cross type is configured in `ana_cmpr_new_unit`, this function can update the cross type
 * @note This function is allowed to run within ISR context including intr callbacks
 * @note This function will be placed into IRAM if `CONFIG_ANA_CMPR_CTRL_FUNC_IN_IRAM` is on,
 *       so that it's allowed to be executed when Cache is disabled
 *
 * @param[in]  cmpr         The handle of analog comparator unit
 * @param[in]  cross_type   The source signal cross type that can trigger the interrupt
 * @return
 *      - ESP_OK                Set denounce configuration success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters
 */
//go:linkname AnaCmprSetCrossType C.ana_cmpr_set_cross_type
func AnaCmprSetCrossType(cmpr AnaCmprHandleT, cross_type AnaCmprCrossTypeT) EspErrT

/**
 * @brief Register analog comparator interrupt event callbacks
 * @note  This function can only be called before enabling the unit
 *
 * @param[in]  cmpr         The handle of analog comparator unit
 * @param[in]  cbs          Group of callback functions
 * @param[in]  user_data    The user data that will be passed to callback functions directly
 * @return
 *      - ESP_OK                Register callbacks success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters
 *      - ESP_ERR_INVALID_STATE The analog comparator has been enabled
 */
//go:linkname AnaCmprRegisterEventCallbacks C.ana_cmpr_register_event_callbacks
func AnaCmprRegisterEventCallbacks(cmpr AnaCmprHandleT, cbs *AnaCmprEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief Enable the analog comparator unit
 *
 * @param[in]  cmpr         The handle of analog comparator unit
 * @return
 *      - ESP_OK                Enable analog comparator unit success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters
 *      - ESP_ERR_INVALID_STATE The analog comparator has been enabled
 */
//go:linkname AnaCmprEnable C.ana_cmpr_enable
func AnaCmprEnable(cmpr AnaCmprHandleT) EspErrT

/**
 * @brief Disable the analog comparator unit
 *
 * @param[in]  cmpr         The handle of analog comparator unit
 * @return
 *      - ESP_OK                Disable analog comparator unit success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters
 *      - ESP_ERR_INVALID_STATE The analog comparator has disabled already
 */
//go:linkname AnaCmprDisable C.ana_cmpr_disable
func AnaCmprDisable(cmpr AnaCmprHandleT) EspErrT

/**
 * @brief Get the specific GPIO number of the analog comparator unit
 *
 * @param[in]  unit         The handle of analog comparator unit
 * @param[in]  chan_type    The channel type of analog comparator, like source channel or reference channel
 * @param[out] gpio_num     The output GPIO number of this channel
 * @return
 *      - ESP_OK                Get GPIO success
 *      - ESP_ERR_INVALID_ARG   NULL pointer of the parameters or wrong unit number or wrong channel type
 */
// llgo:link AnaCmprUnitT.AnaCmprGetGpio C.ana_cmpr_get_gpio
func (recv_ AnaCmprUnitT) AnaCmprGetGpio(chan_type AnaCmprChannelTypeT, gpio_num *c.Int) EspErrT {
	return 0
}
