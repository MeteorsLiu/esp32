package esp_driver_tsens

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type TemperatureSensorObjT struct {
	Unused [8]uint8
}
type TemperatureSensorHandleT *TemperatureSensorObjT

/**
 * @brief Configuration of measurement range for the temperature sensor.
 *
 * @note If you see the log `the boundary you gave cannot meet the range of internal temperature sensor`. You may need to refer to
 *       predefined range listed doc ``api-reference/peripherals/Temperature sensor``.
 */

type TemperatureSensorConfigT struct {
	RangeMin c.Int
	RangeMax c.Int
	ClkSrc   TemperatureSensorClkSrcT
	Flags    struct {
		AllowPd c.Uint32T
	}
}

/**
 * @brief Install temperature sensor driver
 *
 * @param tsens_config Pointer to config structure.
 * @param ret_tsens Return the pointer of temperature sensor handle.
 * @return
 *      - ESP_OK if succeed
 */
// llgo:link (*TemperatureSensorConfigT).TemperatureSensorInstall C.temperature_sensor_install
func (recv_ *TemperatureSensorConfigT) TemperatureSensorInstall(ret_tsens *TemperatureSensorHandleT) EspErrT {
	return 0
}

/**
 * @brief Uninstall the temperature sensor driver
 *
 * @param tsens The handle created by `temperature_sensor_install()`.
 * @return
 *      - ESP_OK if succeed.
 */
//go:linkname TemperatureSensorUninstall C.temperature_sensor_uninstall
func TemperatureSensorUninstall(tsens TemperatureSensorHandleT) EspErrT

/**
 * @brief Enable the temperature sensor
 *
 * @param tsens The handle created by `temperature_sensor_install()`.
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE if temperature sensor is enabled already.
 */
//go:linkname TemperatureSensorEnable C.temperature_sensor_enable
func TemperatureSensorEnable(tsens TemperatureSensorHandleT) EspErrT

/**
 * @brief Disable temperature sensor
 *
 * @param tsens The handle created by `temperature_sensor_install()`.
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE if temperature sensor is not enabled yet.
 */
//go:linkname TemperatureSensorDisable C.temperature_sensor_disable
func TemperatureSensorDisable(tsens TemperatureSensorHandleT) EspErrT

/**
 * @brief Read temperature sensor data that is converted to degrees Celsius.
 * @note  Should not be called from interrupt.
 *
 * @param tsens The handle created by `temperature_sensor_install()`.
 * @param out_celsius The measure output value.
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG invalid arguments
 *     - ESP_ERR_INVALID_STATE Temperature sensor is not enabled yet.
 *     - ESP_FAIL Parse the sensor data into ambient temperature failed (e.g. out of the range).
 */
//go:linkname TemperatureSensorGetCelsius C.temperature_sensor_get_celsius
func TemperatureSensorGetCelsius(tsens TemperatureSensorHandleT, out_celsius *c.Float) EspErrT

type TemperatureValIntrConditionT c.Int

const (
	TEMPERATURE_VAL_HIGHER_THAN_HIGH_THRESHOLD TemperatureValIntrConditionT = 0
	TEMPERATURE_VAL_LOWER_THAN_LOW_THRESHOLD   TemperatureValIntrConditionT = 1
)

/**
 * @brief Temperature sensor event data
 */

type TemperatureSensorThresholdEventDataT struct {
	CelsiusValue  c.Int
	IntrCondition TemperatureValIntrConditionT
}

// llgo:type C
type TemperatureThresCbT func(TemperatureSensorHandleT, *TemperatureSensorThresholdEventDataT, c.Pointer) bool

/**
 * @brief Group of temperature sensor callback functions, all of them will be run in ISR.
 */

type TemperatureSensorEventCallbacksT struct {
	OnThreshold TemperatureThresCbT
}

/**
 * @brief Config options for temperature value absolute interrupt.
 */

type TemperatureSensorAbsThresholdConfigT struct {
	HighThreshold c.Float
	LowThreshold  c.Float
}

/**
 * @brief Set temperature sensor absolute mode automatic monitor.
 *
 * @param tsens The handle created by `temperature_sensor_install()`.
 * @param abs_cfg Configuration of temperature sensor absolute mode interrupt, see `temperature_sensor_abs_threshold_config_t`.
 * @note This function should not be called with `temperature_sensor_set_delta_threshold`.
 *
 * @return
 *      - ESP_OK: Set absolute threshold successfully.
 *      - ESP_ERR_INVALID_STATE: Set absolute threshold failed because of wrong state.
 *      - ESP_ERR_INVALID_ARG: Set absolute threshold failed because of invalid argument.
 */
//go:linkname TemperatureSensorSetAbsoluteThreshold C.temperature_sensor_set_absolute_threshold
func TemperatureSensorSetAbsoluteThreshold(tsens TemperatureSensorHandleT, abs_cfg *TemperatureSensorAbsThresholdConfigT) EspErrT

/**
 * @brief Config options for temperature value delta interrupt.
 */

type TemperatureSensorDeltaThresholdConfigT struct {
	IncreaseDelta c.Float
	DecreaseDelta c.Float
}

/**
 * @brief Set temperature sensor differential mode automatic monitor.
 *
 * @param tsens The handle created by `temperature_sensor_install()`.
 * @param delta_cfg Configuration of temperature sensor delta mode interrupt, see `temperature_sensor_delta_threshold_config_t`.
 * @note This function should not be called with `temperature_sensor_set_absolute_threshold`
 *
 * @return
 *      - ESP_OK: Set differential value threshold successfully.
 *      - ESP_ERR_INVALID_STATE: Set absolute threshold failed because of wrong state.
 *      - ESP_ERR_INVALID_ARG: Set differential value threshold failed because of invalid argument.
 */
//go:linkname TemperatureSensorSetDeltaThreshold C.temperature_sensor_set_delta_threshold
func TemperatureSensorSetDeltaThreshold(tsens TemperatureSensorHandleT, delta_cfg *TemperatureSensorDeltaThresholdConfigT) EspErrT

/**
 * @brief Install temperature sensor interrupt callback. Temperature sensor interrupt will be enabled at same time
 *
 * @param tsens The handle created by `temperature_sensor_install()`.
 * @param cbs Pointer to the group of temperature sensor interrupt callbacks.
 * @param user_arg Callback argument.
 *
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname TemperatureSensorRegisterCallbacks C.temperature_sensor_register_callbacks
func TemperatureSensorRegisterCallbacks(tsens TemperatureSensorHandleT, cbs *TemperatureSensorEventCallbacksT, user_arg c.Pointer) EspErrT
