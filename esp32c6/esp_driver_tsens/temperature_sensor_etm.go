package esp_driver_tsens

import _ "unsafe"

/**
 * @brief Temperature Sensor ETM event configuration
 */

type TemperatureSensorEtmEventConfigT struct {
	EventType TemperatureSensorEtmEventTypeT
}

/**
 * @brief Get the ETM event for Temperature Sensor
 *
 * @note The created ETM event object can be deleted later by calling `esp_etm_del_event`
 *
 * @param[in] tsens Temperature Sensor handle, allocated by `temperature_sensor_install()`
 * @param[in] config Temperature Sensor ETM event configuration
 * @param[out] out_event Returned ETM event handle
 * @return
 *      - ESP_OK: Get ETM event successfully
 *      - ESP_ERR_INVALID_ARG: Get ETM event failed because of invalid argument
 *      - ESP_FAIL: Get ETM event failed because of other error
 */
//go:linkname TemperatureSensorNewEtmEvent C.temperature_sensor_new_etm_event
func TemperatureSensorNewEtmEvent(tsens TemperatureSensorHandleT, config *TemperatureSensorEtmEventConfigT, out_event *EspEtmEventHandleT) EspErrT

/**
 * @brief Temperature Sensor ETM task configuration
 */

type TemperatureSensorEtmTaskConfigT struct {
	TaskType TemperatureSensorEtmTaskTypeT
}

/**
 * @brief Get the ETM task for Temperature Sensor
 *
 * @note The created ETM task object can be deleted later by calling `esp_etm_del_task`
 *
 * @param[in] tsens Temperature Sensor, allocated by `temperature_sensor_install()`
 * @param[in] config Temperature Sensor ETM task configuration
 * @param[out] out_task Returned ETM task handle
 * @return
 *      - ESP_OK: Get ETM task successfully
 *      - ESP_ERR_INVALID_ARG: Get ETM task failed because of invalid argument
 *      - ESP_FAIL: Get ETM task failed because of other error
 */
//go:linkname TemperatureSensorNewEtmTask C.temperature_sensor_new_etm_task
func TemperatureSensorNewEtmTask(tsens TemperatureSensorHandleT, config *TemperatureSensorEtmTaskConfigT, out_task *EspEtmTaskHandleT) EspErrT
