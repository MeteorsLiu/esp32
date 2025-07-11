package esp_driver_dac

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DacChannelMaskT c.Int

const (
	DAC_CHANNEL_MASK_CH0 DacChannelMaskT = 1
	DAC_CHANNEL_MASK_CH1 DacChannelMaskT = 2
	DAC_CHANNEL_MASK_ALL DacChannelMaskT = 3
)

type DacContinuousS struct {
	Unused [8]uint8
}
type DacContinuousHandleT *DacContinuousS

/**
 * @brief DAC continuous channels' configurations
 *
 */

type DacContinuousConfigT struct {
	ChanMask DacChannelMaskT
	DescNum  c.Uint32T
	BufSize  c.SizeT
	FreqHz   c.Uint32T
	Offset   c.Int8T
	ClkSrc   DacContinuousDigiClkSrcT
	ChanMode DacContinuousChannelModeT
}

/**
 * @brief Event structure used in DAC event queue
 */

type DacEventDataT struct {
	Buf        c.Pointer
	BufSize    c.SizeT
	WriteBytes c.SizeT
}

// llgo:type C
type DacIsrCallbackT func(DacContinuousHandleT, *DacEventDataT, c.Pointer) bool

/**
 * @brief Group of DAC callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_DAC_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in the function should be in the SRAM as well.
 */

type DacEventCallbacksT struct {
	OnConvertDone DacIsrCallbackT
	OnStop        DacIsrCallbackT
}

/**
 * @brief Allocate new DAC channels in continuous mode
 * @note  The DAC channels can't be registered to continuous mode separately
 *
 * @param[in]  cont_cfg    Continuous mode configuration
 * @param[out] ret_handle   The returned continuous mode handle
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The DAC channel has been registered already
 *      - ESP_ERR_NOT_FOUND     Not found the available dma peripheral, might be occupied
 *      - ESP_ERR_NO_MEM        No memory for the DAC continuous mode resources
 *      - ESP_OK                Allocate the new DAC continuous mode success
 */
// llgo:link (*DacContinuousConfigT).DacContinuousNewChannels C.dac_continuous_new_channels
func (recv_ *DacContinuousConfigT) DacContinuousNewChannels(ret_handle *DacContinuousHandleT) EspErrT {
	return 0
}

/**
 * @brief Delete the DAC continuous handle
 *
 * @param[in]  handle       The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The channels have already been deregistered or not disabled
 *      - ESP_OK                Delete the continuous channels success
 */
//go:linkname DacContinuousDelChannels C.dac_continuous_del_channels
func DacContinuousDelChannels(handle DacContinuousHandleT) EspErrT

/**
 * @brief Enabled the DAC continuous mode
 * @note  Must enable the channels before
 *
 * @param[in]  handle       The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The channels have been enabled already
 *      - ESP_OK                Enable the continuous output success
 */
//go:linkname DacContinuousEnable C.dac_continuous_enable
func DacContinuousEnable(handle DacContinuousHandleT) EspErrT

/**
 * @brief Disable the DAC continuous mode
 *
 * @param[in]  handle       The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @return
 *      - ESP_ERR_INVALID_ARG  The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The channels have been enabled already
 *      - ESP_OK                Disable the continuous output success
 */
//go:linkname DacContinuousDisable C.dac_continuous_disable
func DacContinuousDisable(handle DacContinuousHandleT) EspErrT

/**
 * @brief Write DAC data continuously
 * @note  The data in buffer will only be converted one time,
 *        This function will be blocked until all data loaded or timeout
 *        then the DAC output will keep outputting the voltage of the last data in the buffer
 * @note  Specially, on ESP32, the data bit width of DAC continuous data is fixed to 16 bits while only the high 8 bits are available,
 *        The driver will help to expand the inputted buffer automatically by default,
 *        you can also align the data to 16 bits manually by clearing `CONFIG_DAC_DMA_AUTO_16BIT_ALIGN` in menuconfig.
 *
 * @param[in]  handle   The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @param[in]  buf      The digital data buffer to convert
 * @param[in]  buf_size The buffer size of digital data buffer
 * @param[out] bytes_loaded The bytes that has been loaded into DMA buffer, can be NULL if don't need it
 * @param[in]  timeout_ms The timeout time in millisecond, set a minus value means will block forever
 * @return
 *      - ESP_ERR_INVALID_ARG   The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The DAC continuous mode has not been enabled yet
 *      - ESP_ERR_TIMEOUT       Waiting for semaphore or message queue timeout
 *      - ESP_OK                Success to output the acyclic DAC data
 */
//go:linkname DacContinuousWrite C.dac_continuous_write
func DacContinuousWrite(handle DacContinuousHandleT, buf *c.Uint8T, buf_size c.SizeT, bytes_loaded *c.SizeT, timeout_ms c.Int) EspErrT

/**
 * @brief Write DAC continuous data cyclically
 * @note  The data in buffer will be converted cyclically using DMA once this function is called,
 *        This function will return once the data loaded into DMA buffers.
 * @note  The buffer size of cyclically output is limited by the descriptor number and
 *        dma buffer size while initializing the continuous mode.
 *        Concretely, in order to load all the data into descriptors,
 *        the cyclic buffer size is not supposed to be greater than `desc_num * buf_size`
 * @note  Specially, on ESP32, the data bit width of DAC continuous data is fixed to 16 bits while only the high 8 bits are available,
 *        The driver will help to expand the inputted buffer automatically by default,
 *        you can also align the data to 16 bits manually by clearing `CONFIG_DAC_DMA_AUTO_16BIT_ALIGN` in menuconfig.
 *
 * @param[in]  handle   The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @param[in]  buf      The digital data buffer to convert
 * @param[in]  buf_size The buffer size of digital data buffer
 * @param[out] bytes_loaded The bytes that has been loaded into DMA buffer, can be NULL if don't need it
 * @return
 *      - ESP_ERR_INVALID_ARG   The input parameter is invalid
 *      - ESP_ERR_INVALID_STATE The DAC continuous mode has not been enabled yet
 *      - ESP_OK                Success to output the acyclic DAC data
 */
//go:linkname DacContinuousWriteCyclically C.dac_continuous_write_cyclically
func DacContinuousWriteCyclically(handle DacContinuousHandleT, buf *c.Uint8T, buf_size c.SizeT, bytes_loaded *c.SizeT) EspErrT

/**
 * @brief Set event callbacks for DAC continuous mode
 *
 * @note User can deregister a previously registered callback by calling this function and setting the callback member in the `callbacks` structure to NULL.
 * @note When CONFIG_DAC_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 *       The variables used in this function, including the `user_data`, should be in the internal RAM as well.
 *
 * @param[in] handle        The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @param[in] callbacks     Group of callback functions, input NULL to clear the former callbacks
 * @param[in] user_data     User data, which will be passed to callback functions directly
 * @return
 *      - ESP_OK                Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG   Set event callbacks failed because of invalid argument
 */
//go:linkname DacContinuousRegisterEventCallback C.dac_continuous_register_event_callback
func DacContinuousRegisterEventCallback(handle DacContinuousHandleT, callbacks *DacEventCallbacksT, user_data c.Pointer) EspErrT

/**
 * @brief  Start the async writing
 * @note   When the asynchronous writing start, the DAC will keep outputting '0' until the data are loaded into the DMA buffer.
 *         To loaded the data into DMA buffer, 'on_convert_done' callback is required,
 *         which can be registered by 'dac_continuous_register_event_callback' before enabling
 *
 * @param[in] handle        The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @return
 *      - ESP_OK                Start asynchronous writing successfully
 *      - ESP_ERR_INVALID_ARG   The handle is NULL
 *      - ESP_ERR_INVALID_STATE The channel is not enabled or the 'on_convert_done' callback is not registered
 */
//go:linkname DacContinuousStartAsyncWriting C.dac_continuous_start_async_writing
func DacContinuousStartAsyncWriting(handle DacContinuousHandleT) EspErrT

/**
 * @brief Stop the sync writing
 *
 * @param[in] handle        The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @return
 *      - ESP_OK                Stop asynchronous writing successfully
 *      - ESP_ERR_INVALID_ARG   The handle is NULL
 *      - ESP_ERR_INVALID_STATE Asynchronous writing has not started
 */
//go:linkname DacContinuousStopAsyncWriting C.dac_continuous_stop_async_writing
func DacContinuousStopAsyncWriting(handle DacContinuousHandleT) EspErrT

/**
 * @brief Write DAC data asynchronously
 * @note  This function can be called when the asynchronous writing started, and it can be called in the callback directly
 *        but recommend to writing data in a task, referring to :example:`peripherals/dac/dac_continuous/dac_audio`
 *
 * @param[in]  handle        The DAC continuous channel handle that obtained from 'dac_continuous_new_channels'
 * @param[in]  dma_buf       The DMA buffer address, it can be acquired from 'dac_event_data_t' in the 'on_convert_done' callback
 * @param[in]  dma_buf_len   The DMA buffer length, it can be acquired from 'dac_event_data_t' in the 'on_convert_done' callback
 * @param[in]  data          The data that need to be written
 * @param[in]  data_len      The data length the need to be written
 * @param[out] bytes_loaded  The bytes number that has been loaded/written into the DMA buffer
 * @return
 *      - ESP_OK                        Write the data into DMA buffer successfully
 *      - ESP_ERR_INVALID_ARG           NULL pointer
 *      - ESP_ERR_INVALID_STATE         The channels haven't start the asynchronous writing
 *      - ESP_ERR_NOT_FOUND             The param 'dam_buf' not match any existed DMA buffer
 */
//go:linkname DacContinuousWriteAsynchronously C.dac_continuous_write_asynchronously
func DacContinuousWriteAsynchronously(handle DacContinuousHandleT, dma_buf *c.Uint8T, dma_buf_len c.SizeT, data *c.Uint8T, data_len c.SizeT, bytes_loaded *c.SizeT) EspErrT
