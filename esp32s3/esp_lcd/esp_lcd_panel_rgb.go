package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief LCD RGB timing structure
 * @verbatim
 *                                                 Total Width
 *                             <--------------------------------------------------->
 *                       HSYNC width HBP             Active Width                HFP
 *                             <---><--><--------------------------------------><--->
 *                         ____    ____|_______________________________________|____|
 *                             |___|   |                                       |    |
 *                                     |                                       |    |
 *                         __|         |                                       |    |
 *            /|\    /|\  |            |                                       |    |
 *             | VSYNC|   |            |                                       |    |
 *             |Width\|/  |__          |                                       |    |
 *             |     /|\     |         |                                       |    |
 *             |  VBP |      |         |                                       |    |
 *             |     \|/_____|_________|_______________________________________|    |
 *             |     /|\     |         | / / / / / / / / / / / / / / / / / / / |    |
 *             |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *    Total    |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *    Height   |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *             |Active|      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *             |Height|      |         |/ / / / / / Active Display Area / / / /|    |
 *             |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *             |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *             |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *             |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *             |      |      |         |/ / / / / / / / / / / / / / / / / / / /|    |
 *             |     \|/_____|_________|_______________________________________|    |
 *             |     /|\     |                                                      |
 *             |  VFP |      |                                                      |
 *            \|/    \|/_____|______________________________________________________|
 * @endverbatim
 */

type EspLcdRgbTimingT struct {
	PclkHz          c.Uint32T
	HRes            c.Uint32T
	VRes            c.Uint32T
	HsyncPulseWidth c.Uint32T
	HsyncBackPorch  c.Uint32T
	HsyncFrontPorch c.Uint32T
	VsyncPulseWidth c.Uint32T
	VsyncBackPorch  c.Uint32T
	VsyncFrontPorch c.Uint32T
	Flags           struct {
		HsyncIdleLow  c.Uint32T
		VsyncIdleLow  c.Uint32T
		DeIdleHigh    c.Uint32T
		PclkActiveNeg c.Uint32T
		PclkIdleHigh  c.Uint32T
	}
}

/**
 * @brief Type of RGB LCD panel event data
 */

type EspLcdRgbPanelEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type EspLcdRgbPanelGeneralCbT func(EspLcdPanelHandleT, *EspLcdRgbPanelEventDataT, c.Pointer) bool
type EspLcdRgbPanelDrawBufCompleteCbT EspLcdRgbPanelGeneralCbT
type EspLcdRgbPanelFrameBufCompleteCbT EspLcdRgbPanelGeneralCbT
type EspLcdRgbPanelVsyncCbT EspLcdRgbPanelGeneralCbT

// llgo:type C
type EspLcdRgbPanelBounceBufFillCbT func(EspLcdPanelHandleT, c.Pointer, c.Int, c.Int, c.Pointer) bool
type EspLcdRgbPanelBounceBufFinishCbT EspLcdRgbPanelFrameBufCompleteCbT

/**
 * @brief Group of supported RGB LCD panel callbacks
 * @note The callbacks are all running under ISR environment
 * @note When CONFIG_LCD_RGB_ISR_IRAM_SAFE is enabled, the callback itself and functions called by it should be placed in IRAM.
 */

type EspLcdRgbPanelEventCallbacksT struct {
	OnColorTransDone EspLcdRgbPanelDrawBufCompleteCbT
	OnVsync          EspLcdRgbPanelVsyncCbT
	OnBounceEmpty    EspLcdRgbPanelBounceBufFillCbT
}

/**
 * @brief LCD RGB panel configuration structure
 */

type EspLcdRgbPanelConfigT struct {
	ClkSrc             LcdClockSourceT
	Timings            EspLcdRgbTimingT
	DataWidth          c.SizeT
	BitsPerPixel       c.SizeT
	NumFbs             c.SizeT
	BounceBufferSizePx c.SizeT
	SramTransAlign     c.SizeT
	HsyncGpioNum       c.Int
	VsyncGpioNum       c.Int
	DeGpioNum          c.Int
	PclkGpioNum        c.Int
	DispGpioNum        c.Int
	DataGpioNums       [16]c.Int
	Flags              struct {
		DispActiveLow     c.Uint32T
		RefreshOnDemand   c.Uint32T
		FbInPsram         c.Uint32T
		DoubleFb          c.Uint32T
		NoFb              c.Uint32T
		BbInvalidateCache c.Uint32T
	}
}

/**
 * @brief Create RGB LCD panel
 *
 * @param[in] rgb_panel_config RGB panel configuration
 * @param[out] ret_panel Returned LCD panel handle
 * @return
 *      - ESP_ERR_INVALID_ARG: Create RGB LCD panel failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create RGB LCD panel failed because of out of memory
 *      - ESP_ERR_NOT_FOUND: Create RGB LCD panel failed because some mandatory hardware resources are not found
 *      - ESP_OK: Create RGB LCD panel successfully
 */
// llgo:link (*EspLcdRgbPanelConfigT).EspLcdNewRgbPanel C.esp_lcd_new_rgb_panel
func (recv_ *EspLcdRgbPanelConfigT) EspLcdNewRgbPanel(ret_panel *EspLcdPanelHandleT) EspErrT {
	return 0
}

/**
 * @brief Register LCD RGB panel event callbacks
 *
 * @param[in] panel LCD panel handle, returned from `esp_lcd_new_rgb_panel`
 * @param[in] callbacks Group of callback functions
 * @param[in] user_ctx User data, which will be passed to the callback functions directly
 * @return
 *      - ESP_OK: Set event callbacks successfully
 *      - ESP_ERR_INVALID_ARG: Set event callbacks failed because of invalid argument
 *      - ESP_FAIL: Set event callbacks failed because of other error
 */
//go:linkname EspLcdRgbPanelRegisterEventCallbacks C.esp_lcd_rgb_panel_register_event_callbacks
func EspLcdRgbPanelRegisterEventCallbacks(panel EspLcdPanelHandleT, callbacks *EspLcdRgbPanelEventCallbacksT, user_ctx c.Pointer) EspErrT

/**
 * @brief Set frequency of PCLK for RGB LCD panel
 *
 * @note The PCLK frequency is set in the `esp_lcd_rgb_timing_t` and gets configured during LCD panel initialization.
 *       Usually you don't need to call this function to set the PCLK again, but in some cases, you might want to change the PCLK frequency.
 *       e.g. slow down the PCLK frequency to reduce power consumption or to reduce the memory throughput during OTA.
 * @note This function doesn't cause the hardware to update the PCLK immediately but to record the new frequency and set a flag internally.
 *       Only in the next VSYNC event handler, will the driver attempt to update the PCLK frequency.
 *
 * @param[in] panel LCD panel handle, returned from `esp_lcd_new_rgb_panel`
 * @param[in] freq_hz Frequency of pixel clock, in Hz
 * @return
 *      - ESP_ERR_INVALID_ARG: Set PCLK frequency failed because of invalid argument
 *      - ESP_OK: Set PCLK frequency successfully
 */
//go:linkname EspLcdRgbPanelSetPclk C.esp_lcd_rgb_panel_set_pclk
func EspLcdRgbPanelSetPclk(panel EspLcdPanelHandleT, freq_hz c.Uint32T) EspErrT

/**
 * @brief Restart the LCD transmission
 *
 * @note This function can be useful when the LCD controller is out of sync with the DMA because of insufficient bandwidth.
 *       To save the screen from a permanent shift, you can call this function to restart the LCD DMA.
 * @note This function doesn't restart the DMA immediately but to set a flag internally.
 *       Only in the next VSYNC event handler, will the driver attempt to do the restart job.
 * @note If CONFIG_LCD_RGB_RESTART_IN_VSYNC is enabled, you don't need to call this function manually,
 *       because the restart job will be done automatically in the VSYNC event handler.
 *
 * @param[in] panel panel LCD panel handle, returned from `esp_lcd_new_rgb_panel`
 * @return
 *      - ESP_ERR_INVALID_ARG: Restart the LCD failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Restart the LCD failed because the LCD diver is working in refresh-on-demand mode
 *      - ESP_OK: Restart the LCD successfully
 */
//go:linkname EspLcdRgbPanelRestart C.esp_lcd_rgb_panel_restart
func EspLcdRgbPanelRestart(panel EspLcdPanelHandleT) EspErrT

/**
 * @brief Get the address of the frame buffer(s) that allocated by the driver
 *
 * @param[in] panel LCD panel handle, returned from `esp_lcd_new_rgb_panel`
 * @param[in] fb_num Number of frame buffer(s) to get. This value must be the same as the number of the following parameters.
 * @param[out] fb0 Returned address of the frame buffer 0
 * @param[out] ... List of other frame buffer addresses
 * @return
 *      - ESP_ERR_INVALID_ARG: Get frame buffer address failed because of invalid argument
 *      - ESP_OK: Get frame buffer address successfully
 */
//go:linkname EspLcdRgbPanelGetFrameBuffer C.esp_lcd_rgb_panel_get_frame_buffer
func EspLcdRgbPanelGetFrameBuffer(panel EspLcdPanelHandleT, fb_num c.Uint32T, fb0 *c.Pointer, __llgo_va_list ...interface{}) EspErrT

/**
 * @brief Manually trigger once transmission of the frame buffer to the LCD panel
 *
 * @note This function should only be called when the RGB panel is working under the `refresh_on_demand` mode.
 *
 * @param[in] panel LCD panel handle, returned from `esp_lcd_new_rgb_panel`
 * @return
 *      - ESP_ERR_INVALID_ARG: Start a refresh failed because of invalid argument
 *      - ESP_ERR_INVALID_STATE: Start a refresh failed because the LCD panel is not created with the `refresh_on_demand` flag enabled.
 *      - ESP_OK: Start a refresh successfully
 */
//go:linkname EspLcdRgbPanelRefresh C.esp_lcd_rgb_panel_refresh
func EspLcdRgbPanelRefresh(panel EspLcdPanelHandleT) EspErrT

/**
 * @brief LCD color conversion profile
 */

type EspLcdColorConvProfileT struct {
	ColorSpace LcdColorSpaceT
	ColorRange LcdColorRangeT
	YuvSample  LcdYuvSampleT
}

/**
 * @brief Configuration of YUG-RGB conversion
 */

type EspLcdYuvConvConfigT struct {
	Std LcdYuvConvStdT
	Src EspLcdColorConvProfileT
	Dst EspLcdColorConvProfileT
}

/**
 * @brief Configure how to convert the color format between RGB and YUV
 *
 * @note Pass in `config` as NULL will disable the RGB-YUV converter.
 * @note The hardware converter can only parse a "packed" storage format, while "planar" and "semi-planar" format is not supported.
 *
 * @param[in] panel LCD panel handle, returned from `esp_lcd_new_rgb_panel`
 * @param[in] config Configuration of RGB-YUV conversion
 * @return
 *      - ESP_ERR_INVALID_ARG: Configure RGB-YUV conversion failed because of invalid argument
 *      - ESP_ERR_NOT_SUPPORTED: Configure RGB-YUV conversion failed because the conversion mode is not supported by the hardware
 *      - ESP_OK: Configure RGB-YUV conversion successfully
 */
//go:linkname EspLcdRgbPanelSetYuvConversion C.esp_lcd_rgb_panel_set_yuv_conversion
func EspLcdRgbPanelSetYuvConversion(panel EspLcdPanelHandleT, config *EspLcdYuvConvConfigT) EspErrT
