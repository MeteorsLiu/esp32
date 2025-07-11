package esp_lcd

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ESP_LCD_I80_BUS_WIDTH_MAX = 16

type EspLcdI80BusT struct {
	Unused [8]uint8
}
type EspLcdI80BusHandleT *EspLcdI80BusT

/**
 * @brief LCD Intel 8080 bus configuration structure
 */

type EspLcdI80BusConfigT struct {
	DcGpioNum        c.Int
	WrGpioNum        c.Int
	ClkSrc           LcdClockSourceT
	DataGpioNums     [16]c.Int
	BusWidth         c.SizeT
	MaxTransferBytes c.SizeT
	SramTransAlign   c.SizeT
}

/**
 * @brief Create Intel 8080 bus handle
 *
 * @param[in] bus_config Bus configuration
 * @param[out] ret_bus Returned bus handle
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_ERR_NOT_FOUND     if no free bus is available
 *          - ESP_OK                on success
 */
// llgo:link (*EspLcdI80BusConfigT).EspLcdNewI80Bus C.esp_lcd_new_i80_bus
func (recv_ *EspLcdI80BusConfigT) EspLcdNewI80Bus(ret_bus *EspLcdI80BusHandleT) EspErrT {
	return 0
}

/**
 * @brief Destroy Intel 8080 bus handle
 *
 * @param[in] bus Intel 8080 bus handle, created by `esp_lcd_new_i80_bus()`
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_INVALID_STATE if there still be some device attached to the bus
 *          - ESP_OK                on success
 */
//go:linkname EspLcdDelI80Bus C.esp_lcd_del_i80_bus
func EspLcdDelI80Bus(bus EspLcdI80BusHandleT) EspErrT

/**
 * @brief Panel IO configuration structure, for intel 8080 interface
 */

type EspLcdPanelIoI80ConfigT struct {
	CsGpioNum        c.Int
	PclkHz           c.Uint32T
	TransQueueDepth  c.SizeT
	OnColorTransDone EspLcdPanelIoColorTransDoneCbT
	UserCtx          c.Pointer
	LcdCmdBits       c.Int
	LcdParamBits     c.Int
	DcLevels         struct {
		DcIdleLevel  c.Uint
		DcCmdLevel   c.Uint
		DcDummyLevel c.Uint
		DcDataLevel  c.Uint
	}
	Flags struct {
		CsActiveHigh     c.Uint
		ReverseColorBits c.Uint
		SwapColorBytes   c.Uint
		PclkActiveNeg    c.Uint
		PclkIdleLow      c.Uint
	}
}

/**
 * @brief Create LCD panel IO, for Intel 8080 interface
 *
 * @param[in] bus Intel 8080 bus handle, created by `esp_lcd_new_i80_bus()`
 * @param[in] io_config IO configuration, for i80 interface
 * @param[out] ret_io Returned panel IO handle
 * @return
 *          - ESP_ERR_INVALID_ARG   if parameter is invalid
 *          - ESP_ERR_NOT_SUPPORTED if some configuration can't be satisfied, e.g. pixel clock out of the range
 *          - ESP_ERR_NO_MEM        if out of memory
 *          - ESP_OK                on success
 */
//go:linkname EspLcdNewPanelIoI80 C.esp_lcd_new_panel_io_i80
func EspLcdNewPanelIoI80(bus EspLcdI80BusHandleT, io_config *EspLcdPanelIoI80ConfigT, ret_io *EspLcdPanelIoHandleT) EspErrT

/**
 * @brief Allocate a draw buffer that can be used by I80 interfaced LCD panel
 *
 * @note This function differs from the normal 'heap_caps_*' functions in that it can also automatically handle the alignment required by DMA burst, cache line size, etc.
 *
 * @param[in] io Panel IO handle, created by `esp_lcd_new_panel_io_i80()`
 * @param[in] size Size of memory to be allocated
 * @param[in] caps Bitwise OR of MALLOC_CAP_* flags indicating the type of memory desired for the allocation
 * @return Pointer to a new buffer of size 'size' with capabilities 'caps', or NULL if allocation failed
 */
//go:linkname EspLcdI80AllocDrawBuffer C.esp_lcd_i80_alloc_draw_buffer
func EspLcdI80AllocDrawBuffer(io EspLcdPanelIoHandleT, size c.SizeT, caps c.Uint32T) c.Pointer
