package esp_driver_gptimer

import _ "unsafe"

/**
 * @brief GPTimer ETM event configuration
 */

type GptimerEtmEventConfigT struct {
	EventType GptimerEtmEventTypeT
}

/**
 * @brief Get the ETM event for GPTimer
 *
 * @note The created ETM event object can be deleted later by calling `esp_etm_del_event`
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[in] config GPTimer ETM event configuration
 * @param[out] out_event Returned ETM event handle
 * @return
 *      - ESP_OK: Get ETM event successfully
 *      - ESP_ERR_INVALID_ARG: Get ETM event failed because of invalid argument
 *      - ESP_FAIL: Get ETM event failed because of other error
 */
//go:linkname GptimerNewEtmEvent C.gptimer_new_etm_event
func GptimerNewEtmEvent(timer GptimerHandleT, config *GptimerEtmEventConfigT, out_event *EspEtmEventHandleT) EspErrT

/**
 * @brief GPTimer ETM task configuration
 */

type GptimerEtmTaskConfigT struct {
	TaskType GptimerEtmTaskTypeT
}

/**
 * @brief Get the ETM task for GPTimer
 *
 * @note The created ETM task object can be deleted later by calling `esp_etm_del_task`
 *
 * @param[in] timer Timer handle created by `gptimer_new_timer`
 * @param[in] config GPTimer ETM task configuration
 * @param[out] out_task Returned ETM task handle
 * @return
 *      - ESP_OK: Get ETM task successfully
 *      - ESP_ERR_INVALID_ARG: Get ETM task failed because of invalid argument
 *      - ESP_FAIL: Get ETM task failed because of other error
 */
//go:linkname GptimerNewEtmTask C.gptimer_new_etm_task
func GptimerNewEtmTask(timer GptimerHandleT, config *GptimerEtmTaskConfigT, out_task *EspEtmTaskHandleT) EspErrT
