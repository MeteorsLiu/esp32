package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief I2S ETM event configuration
 */

type I2sEtmEventConfigT struct {
	EventType I2sEtmEventTypeT
	Threshold c.Uint32T
}

/**
 * @brief Register the ETM event for I2S channel
 *
 * @note The created ETM event object can be deleted later by calling `esp_etm_del_event`
 *
 * @param[in] handle I2S channel handle, allocated by `i2s_new_channel`
 * @param[in] config I2S ETM event configuration
 * @param[out] out_event Returned ETM event handle
 * @return
 *      - ESP_OK: Get ETM event successfully
 *      - ESP_ERR_INVALID_ARG: Get ETM event failed because of invalid argument
 *      - ESP_ERR_NOT_SUPPORTED: Get ETM event failed because the I2S hardware doesn't support ETM event
 *      - ESP_FAIL: Get ETM event failed because of other error
 */
//go:linkname I2sNewEtmEvent C.i2s_new_etm_event
func I2sNewEtmEvent(handle I2sChanHandleT, config *I2sEtmEventConfigT, out_event *EspEtmEventHandleT) EspErrT

/**
 * @brief I2S ETM task configuration
 */

type I2sEtmTaskConfigT struct {
	TaskType I2sEtmTaskTypeT
}

/**
 * @brief Register the ETM task for I2S channel
 *
 * @note The created ETM task object can be deleted later by calling `esp_etm_del_task`
 *
 * @param[in] handle I2S channel handle, allocated by `i2s_new_channel`
 * @param[in] config I2S ETM task configuration
 * @param[out] out_task Returned ETM task handle
 * @return
 *      - ESP_OK: Get ETM task successfully
 *      - ESP_ERR_INVALID_ARG: Get ETM task failed because of invalid argument
 *      - ESP_ERR_NOT_SUPPORTED: Get ETM task failed because the i2s hardware doesn't support ETM task
 *      - ESP_FAIL: Get ETM task failed because of other error
 */
//go:linkname I2sNewEtmTask C.i2s_new_etm_task
func I2sNewEtmTask(handle I2sChanHandleT, config *I2sEtmTaskConfigT, out_task *EspEtmTaskHandleT) EspErrT
