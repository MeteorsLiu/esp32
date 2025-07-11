package driver

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Set I2S pin number
 *
 * @note
 * The I2S peripheral output signals can be connected to multiple GPIO pads.
 * However, the I2S peripheral input signal can only be connected to one GPIO pad.
 *
 * @param   i2s_num     I2S port number
 *
 * @param   pin         I2S Pin structure, or NULL to set 2-channel 8-bit internal DAC pin configuration (GPIO25 & GPIO26)
 *
 * Inside the pin configuration structure, set I2S_PIN_NO_CHANGE for any pin where
 * the current configuration should not be changed.
 *
 * @note if *pin is set as NULL, this function will initialize both of the built-in DAC channels by default.
 *       if you don't want this to happen and you want to initialize only one of the DAC channels, you can call i2s_set_dac_mode instead.
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_FAIL            IO error
 */
//go:linkname I2sSetPin C.i2s_set_pin
func I2sSetPin(i2s_num I2sPortT, pin *I2sPinConfigT) EspErrT

/**
 * @brief Install and start I2S driver.
 *
 * @param i2s_num         I2S port number
 *
 * @param i2s_config      I2S configurations - see i2s_config_t struct
 *
 * @param queue_size      I2S event queue size/depth.
 *
 * @param i2s_queue       I2S event queue handle, if set NULL, driver will not use an event queue.
 *
 * This function must be called before any I2S driver read/write operations.
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM      Out of memory
 *     - ESP_ERR_INVALID_STATE  Current I2S port is in use
 */
//go:linkname I2sDriverInstall C.i2s_driver_install
func I2sDriverInstall(i2s_num I2sPortT, i2s_config *I2sConfigT, queue_size c.Int, i2s_queue c.Pointer) EspErrT

/**
 * @brief Uninstall I2S driver.
 *
 * @param i2s_num  I2S port number
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_INVALID_STATE I2S port has been uninstalled by others (e.g. LCD i80)
 */
//go:linkname I2sDriverUninstall C.i2s_driver_uninstall
func I2sDriverUninstall(i2s_num I2sPortT) EspErrT

/**
 * @brief Write data to I2S DMA transmit buffer.
 *
 * @param i2s_num             I2S port number
 *
 * @param src                 Source address to write from
 *
 * @param size                Size of data in bytes
 *
 * @param[out] bytes_written  Number of bytes written, if timeout, the result will be less than the size passed in.
 *
 * @param ticks_to_wait       TX buffer wait timeout in RTOS ticks. If this
 * many ticks pass without space becoming available in the DMA
 * transmit buffer, then the function will return (note that if the
 * data is written to the DMA buffer in pieces, the overall operation
 * may still take longer than this timeout.) Pass portMAX_DELAY for no
 * timeout.
 *
 * @return
 *     - ESP_OK               Success
 *     - ESP_ERR_INVALID_ARG  Parameter error
 */
//go:linkname I2sWrite C.i2s_write
func I2sWrite(i2s_num I2sPortT, src c.Pointer, size c.SizeT, bytes_written *c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Write data to I2S DMA transmit buffer while expanding the number of bits per sample. For example, expanding 16-bit PCM to 32-bit PCM.
 *
 * @param i2s_num             I2S port number
 *
 * @param src                 Source address to write from
 *
 * @param size                Size of data in bytes
 *
 * @param src_bits            Source audio bit
 *
 * @param aim_bits            Bit wanted, no more than 32, and must be greater than src_bits
 *
 * @param[out] bytes_written  Number of bytes written, if timeout, the result will be less than the size passed in.
 *
 * @param ticks_to_wait       TX buffer wait timeout in RTOS ticks. If this
 * many ticks pass without space becoming available in the DMA
 * transmit buffer, then the function will return (note that if the
 * data is written to the DMA buffer in pieces, the overall operation
 * may still take longer than this timeout.) Pass portMAX_DELAY for no
 * timeout.
 *
 * Format of the data in source buffer is determined by the I2S
 * configuration (see i2s_config_t).
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2sWriteExpand C.i2s_write_expand
func I2sWriteExpand(i2s_num I2sPortT, src c.Pointer, size c.SizeT, src_bits c.SizeT, aim_bits c.SizeT, bytes_written *c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Read data from I2S DMA receive buffer
 *
 * @param i2s_num         I2S port number
 *
 * @param dest            Destination address to read into
 *
 * @param size            Size of data in bytes
 *
 * @param[out] bytes_read Number of bytes read, if timeout, bytes read will be less than the size passed in.
 *
 * @param ticks_to_wait   RX buffer wait timeout in RTOS ticks. If this many ticks pass without bytes becoming available in the DMA receive buffer, then the function will return (note that if data is read from the DMA buffer in pieces, the overall operation may still take longer than this timeout.) Pass portMAX_DELAY for no timeout.
 *
 * @note If the built-in ADC mode is enabled, we should call i2s_adc_enable and i2s_adc_disable around the whole reading process,
 *       to prevent the data getting corrupted.
 *
 * @return
 *     - ESP_OK               Success
 *     - ESP_ERR_INVALID_ARG  Parameter error
 */
//go:linkname I2sRead C.i2s_read
func I2sRead(i2s_num I2sPortT, dest c.Pointer, size c.SizeT, bytes_read *c.SizeT, ticks_to_wait TickTypeT) EspErrT

/**
 * @brief Set sample rate used for I2S RX and TX.
 *
 * The bit clock rate is determined by the sample rate and i2s_config_t configuration parameters (number of channels, bits_per_sample).
 *
 * `bit_clock = rate * (number of channels) * bits_per_sample`
 *
 * @param i2s_num  I2S port number
 *
 * @param rate I2S sample rate (ex: 8000, 44100...)
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM      Out of memory
 */
//go:linkname I2sSetSampleRates C.i2s_set_sample_rates
func I2sSetSampleRates(i2s_num I2sPortT, rate c.Uint32T) EspErrT

/**
 * @brief Stop I2S driver
 *
 * There is no need to call i2s_stop() before calling i2s_driver_uninstall().
 *
 * Disables I2S TX/RX, until i2s_start() is called.
 *
 * @param i2s_num  I2S port number
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2sStop C.i2s_stop
func I2sStop(i2s_num I2sPortT) EspErrT

/**
 * @brief Start I2S driver
 *
 * It is not necessary to call this function after i2s_driver_install() (it is started automatically), however it is necessary to call it after i2s_stop().
 *
 *
 * @param i2s_num  I2S port number
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2sStart C.i2s_start
func I2sStart(i2s_num I2sPortT) EspErrT

/**
 * @brief Zero the contents of the TX DMA buffer.
 *
 * Pushes zero-byte samples into the TX DMA buffer, until it is full.
 *
 * @param i2s_num  I2S port number
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 */
//go:linkname I2sZeroDmaBuffer C.i2s_zero_dma_buffer
func I2sZeroDmaBuffer(i2s_num I2sPortT) EspErrT

/**
 * @brief Set clock & bit width used for I2S RX and TX.
 *
 * Similar to i2s_set_sample_rates(), but also sets bit width.
 *
 * 1. stop i2s;
 * 2. calculate mclk, bck, bck_factor
 * 3. malloc dma buffer;
 * 4. start i2s
 *
 * @param i2s_num  I2S port number
 *
 * @param rate I2S sample rate (ex: 8000, 44100...)
 *
 * @param bits_cfg I2S bits configuration
 *             the low 16 bits is for data bits per sample in one channel (see 'i2s_bits_per_sample_t')
 *             the high 16 bits is for total bits in one channel (see 'i2s_bits_per_chan_t')
 *             high 16bits =0 means same as the bits per sample.
 *
 * @param ch I2S channel, (I2S_CHANNEL_MONO, I2S_CHANNEL_STEREO or specific channel in TDM mode)
 *
 * @return
 *     - ESP_OK              Success
 *     - ESP_ERR_INVALID_ARG Parameter error
 *     - ESP_ERR_NO_MEM      Out of memory
 */
//go:linkname I2sSetClk C.i2s_set_clk
func I2sSetClk(i2s_num I2sPortT, rate c.Uint32T, bits_cfg c.Uint32T, ch I2sChannelT) EspErrT

/**
 * @brief get clock set on particular port number.
 *
 * @param i2s_num  I2S port number
 *
 * @return
 *     - actual clock set by i2s driver
 */
//go:linkname I2sGetClk C.i2s_get_clk
func I2sGetClk(i2s_num I2sPortT) c.Float
