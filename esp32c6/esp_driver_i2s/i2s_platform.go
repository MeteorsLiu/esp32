package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2sCtlrT c.Int

const (
	I2S_CTLR_HP I2sCtlrT = 0
	I2S_CTLR_LP I2sCtlrT = 1
)

/**
 * @brief Hold the I2S port occupation
 *
 * @note This private API is used to avoid applications from using the same I2S instance for different purpose.
 * @note This function will help enable the peripheral APB clock as well.
 *
 * @param[in] type I2S controller type
 * @param id I2S port number
 * @param comp_name The name of compnant that occupied this i2s controller
 * @return
 *      - ESP_OK: The specific I2S port is free and register the new device object successfully
 *      - ESP_ERR_INVALID_ARG: Invalid argument, e.g. wrong port_id
 *      - ESP_ERR_NOT_FOUND Specific I2S port is not available
 */
// llgo:link I2sCtlrT.I2sPlatformAcquireOccupation C.i2s_platform_acquire_occupation
func (recv_ I2sCtlrT) I2sPlatformAcquireOccupation(id c.Int, comp_name *c.Char) EspErrT {
	return 0
}

/**
 * @brief Release the I2S port occupation
 *
 * @note This function will help disable the peripheral APB clock as well.
 *
 * @param[in] type I2S controller type
 * @param id I2S port number
 * @return
 *      - ESP_OK: Deregister I2S port successfully (i.e. that I2S port can used used by other users after this function returns)
 *      - ESP_ERR_INVALID_ARG: Invalid argument, e.g. wrong port_id
 *      - ESP_ERR_INVALID_STATE: Specific I2S port is free already
 */
// llgo:link I2sCtlrT.I2sPlatformReleaseOccupation C.i2s_platform_release_occupation
func (recv_ I2sCtlrT) I2sPlatformReleaseOccupation(id c.Int) EspErrT {
	return 0
}

/**
 * @brief This function is only used for getting DMA buffer offset in `test_i2s_iram.c`
 *
 * @return
 *      - The offset of DMA buffers in the `i2s_chan_handle_t` struct (unit: bytes)
 */
//go:linkname I2sPlatformGetDmaBufferOffset C.i2s_platform_get_dma_buffer_offset
func I2sPlatformGetDmaBufferOffset() c.SizeT
