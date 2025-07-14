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
