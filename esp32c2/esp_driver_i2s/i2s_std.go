package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief I2S slot configuration for standard mode
 */

type I2sStdSlotConfigT struct {
	DataBitWidth I2sDataBitWidthT
	SlotBitWidth I2sSlotBitWidthT
	SlotMode     I2sSlotModeT
	SlotMask     I2sStdSlotMaskT
	WsWidth      c.Uint32T
	WsPol        bool
	BitShift     bool
	LeftAlign    bool
	BigEndian    bool
	BitOrderLsb  bool
}

/**
 * @brief I2S clock configuration for standard mode
 */

type I2sStdClkConfigT struct {
	SampleRateHz c.Uint32T
	ClkSrc       I2sClockSrcT
	MclkMultiple I2sMclkMultipleT
}

/**
 * @brief I2S standard mode GPIO pins configuration
 */

type I2sStdGpioConfigT struct {
	Mclk        GpioNumT
	Bclk        GpioNumT
	Ws          GpioNumT
	Dout        GpioNumT
	Din         GpioNumT
	InvertFlags struct {
		MclkInv c.Uint32T
		BclkInv c.Uint32T
		WsInv   c.Uint32T
	}
}

/**
 * @brief I2S standard mode major configuration that including clock/slot/GPIO configuration
 */

type I2sStdConfigT struct {
	ClkCfg  I2sStdClkConfigT
	SlotCfg I2sStdSlotConfigT
	GpioCfg I2sStdGpioConfigT
}

/**
 * @brief Initialize I2S channel to standard mode
 * @note  Only allowed to be called when the channel state is REGISTERED, (i.e., channel has been allocated, but not initialized)
 *        and the state will be updated to READY if initialization success, otherwise the state will return to REGISTERED.
 *
 * @param[in]   handle      I2S channel handler
 * @param[in]   std_cfg     Configurations for standard mode, including clock, slot and GPIO
 *                          The clock configuration can be generated by the helper macro `I2S_STD_CLK_DEFAULT_CONFIG`
 *                          The slot configuration can be generated by the helper macro `I2S_STD_PHILIPS_SLOT_DEFAULT_CONFIG`,
 *                          `I2S_STD_PCM_SLOT_DEFAULT_CONFIG` or `I2S_STD_MSB_SLOT_DEFAULT_CONFIG`
 *
 * @return
 *      - ESP_OK    Initialize successfully
 *      - ESP_ERR_NO_MEM        No memory for storing the channel information
 *      - ESP_ERR_INVALID_ARG   NULL pointer or invalid configuration
 *      - ESP_ERR_INVALID_STATE This channel is not registered
 */
//go:linkname I2sChannelInitStdMode C.i2s_channel_init_std_mode
func I2sChannelInitStdMode(handle I2sChanHandleT, std_cfg *I2sStdConfigT) EspErrT

/**
 * @brief Reconfigure the I2S clock for standard mode
 * @note  Only allowed to be called when the channel state is READY, i.e., channel has been initialized, but not started
 *        this function won't change the state. `i2s_channel_disable` should be called before calling this function if I2S has started.
 * @note  The input channel handle has to be initialized to standard mode, i.e., `i2s_channel_init_std_mode` has been called before reconfiguring
 *
 * @param[in]   handle      I2S channel handler
 * @param[in]   clk_cfg     Standard mode clock configuration, can be generated by `I2S_STD_CLK_DEFAULT_CONFIG`
 * @return
 *      - ESP_OK    Set clock successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer, invalid configuration or not standard mode
 *      - ESP_ERR_INVALID_STATE This channel is not initialized or not stopped
 */
//go:linkname I2sChannelReconfigStdClock C.i2s_channel_reconfig_std_clock
func I2sChannelReconfigStdClock(handle I2sChanHandleT, clk_cfg *I2sStdClkConfigT) EspErrT

/**
 * @brief Reconfigure the I2S slot for standard mode
 * @note  Only allowed to be called when the channel state is READY, i.e., channel has been initialized, but not started
 *        this function won't change the state. `i2s_channel_disable` should be called before calling this function if I2S has started.
 * @note  The input channel handle has to be initialized to standard mode, i.e., `i2s_channel_init_std_mode` has been called before reconfiguring
 *
 * @param[in]   handle      I2S channel handler
 * @param[in]   slot_cfg    Standard mode slot configuration, can be generated by `I2S_STD_PHILIPS_SLOT_DEFAULT_CONFIG`,
 *                          `I2S_STD_PCM_SLOT_DEFAULT_CONFIG` and `I2S_STD_MSB_SLOT_DEFAULT_CONFIG`.
 * @return
 *      - ESP_OK    Set clock successfully
 *      - ESP_ERR_NO_MEM        No memory for DMA buffer
 *      - ESP_ERR_INVALID_ARG   NULL pointer, invalid configuration  or not standard mode
 *      - ESP_ERR_INVALID_STATE This channel is not initialized or not stopped
 */
//go:linkname I2sChannelReconfigStdSlot C.i2s_channel_reconfig_std_slot
func I2sChannelReconfigStdSlot(handle I2sChanHandleT, slot_cfg *I2sStdSlotConfigT) EspErrT

/**
 * @brief Reconfigure the I2S GPIO for standard mode
 * @note  Only allowed to be called when the channel state is READY, i.e., channel has been initialized, but not started
 *        this function won't change the state. `i2s_channel_disable` should be called before calling this function if I2S has started.
 * @note  The input channel handle has to be initialized to standard mode, i.e., `i2s_channel_init_std_mode` has been called before reconfiguring
 *
 * @param[in]   handle      I2S channel handler
 * @param[in]   gpio_cfg    Standard mode GPIO configuration, specified by user
 * @return
 *      - ESP_OK    Set clock successfully
 *      - ESP_ERR_INVALID_ARG   NULL pointer, invalid configuration  or not standard mode
 *      - ESP_ERR_INVALID_STATE This channel is not initialized or not stopped
 */
//go:linkname I2sChannelReconfigStdGpio C.i2s_channel_reconfig_std_gpio
func I2sChannelReconfigStdGpio(handle I2sChanHandleT, gpio_cfg *I2sStdGpioConfigT) EspErrT
