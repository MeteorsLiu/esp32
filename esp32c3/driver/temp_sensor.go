package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Set parameter of temperature sensor.
 * @param tsens
 * @return
 *     - ESP_OK Success
 */
// llgo:link TempSensorConfigT.TempSensorSetConfig C.temp_sensor_set_config
func (recv_ TempSensorConfigT) TempSensorSetConfig() EspErrT {
	return 0
}

/**
 * @brief Get parameter of temperature sensor.
 * @param tsens
 * @return
 *     - ESP_OK Success
 */
// llgo:link (*TempSensorConfigT).TempSensorGetConfig C.temp_sensor_get_config
func (recv_ *TempSensorConfigT) TempSensorGetConfig() EspErrT {
	return 0
}

/**
 * @brief Start temperature sensor measure.
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE if temperature sensor is started already.
 */
//go:linkname TempSensorStart C.temp_sensor_start
func TempSensorStart() EspErrT

/**
 * @brief Stop temperature sensor measure.
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_STATE if temperature sensor is stopped already.
 */
//go:linkname TempSensorStop C.temp_sensor_stop
func TempSensorStop() EspErrT

/**
 * @brief Read temperature sensor raw data.
 * @param tsens_out Pointer to raw data, Range: 0 ~ 255
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG `tsens_out` is NULL
 *     - ESP_ERR_INVALID_STATE temperature sensor dont start
 */
//go:linkname TempSensorReadRaw C.temp_sensor_read_raw
func TempSensorReadRaw(tsens_out *c.Uint32T) EspErrT

/**
 * @brief Read temperature sensor data that is converted to degrees Celsius.
 * @note  Should not be called from interrupt.
 * @param celsius The measure output value.
 * @return
 *     - ESP_OK Success
 *     - ESP_ERR_INVALID_ARG ARG is NULL.
 *     - ESP_ERR_INVALID_STATE The ambient temperature is out of range.
 */
//go:linkname TempSensorReadCelsius C.temp_sensor_read_celsius
func TempSensorReadCelsius(celsius *c.Float) EspErrT
