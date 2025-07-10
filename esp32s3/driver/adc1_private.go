package driver

import _ "unsafe"

/**
 * @brief For I2S dma to claim the usage of ADC1.
 *
 * Other tasks will be forbidden to use ADC1 between ``adc1_dma_mode_acquire`` and ``adc1_i2s_release``.
 * The I2S module may have to wait for a short time for the current conversion (if exist) to finish.
 *
 * @return
 *      - ESP_OK success
 *      - ESP_ERR_TIMEOUT reserved for future use. Currently the function will wait until success.
 */
//go:linkname Adc1DmaModeAcquire C.adc1_dma_mode_acquire
func Adc1DmaModeAcquire() EspErrT

/**
 * @brief For ADC1 to claim the usage of ADC1.
 *
 * Other tasks will be forbidden to use ADC1 between ``adc1_rtc_mode_acquire`` and ``adc1_i2s_release``.
 * The ADC1 may have to wait for some time for the I2S read operation to finish.
 *
 * @return
 *      - ESP_OK success
 *      - ESP_ERR_TIMEOUT reserved for future use. Currently the function will wait until success.
 */
//go:linkname Adc1RtcModeAcquire C.adc1_rtc_mode_acquire
func Adc1RtcModeAcquire() EspErrT

/**
 * @brief to let other tasks use the ADC1 when I2S is not work.
 *
 * Other tasks will be forbidden to use ADC1 between ``adc1_adc/i2s_mode_acquire`` and ``adc1_i2s_release``.
 * Call this function to release the occupation of ADC1
 *
 * @return always return ESP_OK.
 */
//go:linkname Adc1LockRelease C.adc1_lock_release
func Adc1LockRelease() EspErrT
