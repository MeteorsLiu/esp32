package perfmon

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Performance monitor configuration structure
 *
 * Structure to configure performance counter to measure dedicated function
 */

type XtensaPerfmonConfig struct {
	RepeatCount    c.Int
	MaxDeviation   c.Float
	CallParams     c.Pointer
	CallFunction   c.Pointer
	Callback       c.Pointer
	CallbackParams c.Pointer
	Tracelevel     c.Int
	CountersSize   c.Uint32T
	SelectMask     *c.Uint32T
}
type XtensaPerfmonConfigT XtensaPerfmonConfig

/**
 * @brief      Execute PM
 *
 * Execute performance counter for dedicated function with defined parameters
 *
 * @param[in] config: pointer to the configuration structure
 *
 * @return
 *      - ESP_OK if no errors
 *      - ESP_ERR_INVALID_ARG if one of the required parameters not defined
 *      - ESP_FAIL - counter overflow
 */
// llgo:link (*XtensaPerfmonConfigT).XtensaPerfmonExec C.xtensa_perfmon_exec
func (recv_ *XtensaPerfmonConfigT) XtensaPerfmonExec() EspErrT {
	return 0
}

/**
 * @brief      Dump PM results
 *
 * Callback to dump perfmon result to a FILE* stream specified in
 * perfmon_config_t::callback_params. If callback_params is set to NULL, will print to stdout
 *
 * @param[in] params:   used parameters passed from configuration (callback_params).
 *                      This parameter expected as FILE* hanle, where data will be stored.
 *                      If this parameter NULL, then data will be stored to the stdout.
 * @param[in] select:   select value for current counter
 * @param[in] mask:     mask value for current counter
 * @param[in] value:    counter value for current counter
 *
 */
//go:linkname XtensaPerfmonViewCb C.xtensa_perfmon_view_cb
func XtensaPerfmonViewCb(params c.Pointer, select_ c.Uint32T, mask c.Uint32T, value c.Uint32T)
