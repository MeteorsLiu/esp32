package esp_driver_tsens

import _ "unsafe"

/**
 * @brief Temperature Sensor ETM event configuration
 */

type TemperatureSensorEtmEventConfigT struct {
	EventType TemperatureSensorEtmEventTypeT
}

/**
 * @brief Temperature Sensor ETM task configuration
 */

type TemperatureSensorEtmTaskConfigT struct {
	TaskType TemperatureSensorEtmTaskTypeT
}
