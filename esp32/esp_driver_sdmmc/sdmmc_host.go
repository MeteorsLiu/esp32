package esp_driver_sdmmc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Extra configuration for SDMMC peripheral slot
 */

type SdmmcSlotConfigT struct {
	Clk   GpioNumT
	Cmd   GpioNumT
	D0    GpioNumT
	D1    GpioNumT
	D2    GpioNumT
	D3    GpioNumT
	D4    GpioNumT
	D5    GpioNumT
	D6    GpioNumT
	D7    GpioNumT
	Width c.Uint8T
	Flags c.Uint32T
}

/**
 * SD/MMC host state structure
 */

type SdmmcHostStateT struct {
	HostInitialized bool
	NumOfInitSlots  c.Int
}

/**
 * @brief Initialize SDMMC host peripheral
 *
 * @note This function is not thread safe
 *
 * @return
 *      - ESP_OK on success or if sdmmc_host_init was already initialized with this function
 *      - ESP_ERR_NO_MEM if memory can not be allocated
 */
//go:linkname SdmmcHostInit C.sdmmc_host_init
func SdmmcHostInit() EspErrT

/**
 * @brief Initialize given slot of SDMMC peripheral
 *
 * On the ESP32, SDMMC peripheral has two slots:
 *  - Slot 0: 8-bit wide, maps to HS1_* signals in PIN MUX
 *  - Slot 1: 4-bit wide, maps to HS2_* signals in PIN MUX
 *
 * Card detect and write protect signals can be routed to
 * arbitrary GPIOs using GPIO matrix.
 *
 * @note This function is not thread safe
 *
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param slot_config  additional configuration for the slot
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if host has not been initialized using sdmmc_host_init
 *      - ESP_ERR_INVALID_ARG if GPIO pins from slot_config are not valid
 */
//go:linkname SdmmcHostInitSlot C.sdmmc_host_init_slot
func SdmmcHostInitSlot(slot c.Int, slot_config *SdmmcSlotConfigT) EspErrT

/**
 * @brief Select bus width to be used for data transfer
 *
 * SD/MMC card must be initialized prior to this command, and a command to set
 * bus width has to be sent to the card (e.g. SD_APP_SET_BUS_WIDTH)
 *
 * @note This function is not thread safe
 *
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param width  bus width (1, 4, or 8 for slot 0; 1 or 4 for slot 1)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if slot number or width is not valid
 */
//go:linkname SdmmcHostSetBusWidth C.sdmmc_host_set_bus_width
func SdmmcHostSetBusWidth(slot c.Int, width c.SizeT) EspErrT

/**
 * @brief Get bus width configured in ``sdmmc_host_init_slot`` to be used for data transfer
 *
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @return configured bus width of the specified slot.
 */
//go:linkname SdmmcHostGetSlotWidth C.sdmmc_host_get_slot_width
func SdmmcHostGetSlotWidth(slot c.Int) c.SizeT

/**
 * @brief Set card clock frequency
 *
 * Currently only integer fractions of 40MHz clock can be used.
 * For High Speed cards, 40MHz can be used.
 * For Default Speed cards, 20MHz can be used.
 *
 * @note This function is not thread safe
 *
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param freq_khz  card clock frequency, in kHz
 * @return
 *      - ESP_OK on success
 *      - other error codes may be returned in the future
 */
//go:linkname SdmmcHostSetCardClk C.sdmmc_host_set_card_clk
func SdmmcHostSetCardClk(slot c.Int, freq_khz c.Uint32T) EspErrT

/**
 * @brief Enable or disable DDR mode of SD interface
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param ddr_enabled  enable or disable DDR mode
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_NOT_SUPPORTED if DDR mode is not supported on this slot
 */
//go:linkname SdmmcHostSetBusDdrMode C.sdmmc_host_set_bus_ddr_mode
func SdmmcHostSetBusDdrMode(slot c.Int, ddr_enabled bool) EspErrT

/**
 * @brief Enable or disable always-on card clock
 * When cclk_always_on is false, the host controller is allowed to shut down
 * the card clock between the commands. When cclk_always_on is true, the clock
 * is generated even if no command is in progress.
 * @param slot  slot number
 * @param cclk_always_on  enable or disable always-on clock
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG if the slot number is invalid
 */
//go:linkname SdmmcHostSetCclkAlwaysOn C.sdmmc_host_set_cclk_always_on
func SdmmcHostSetCclkAlwaysOn(slot c.Int, cclk_always_on bool) EspErrT

/**
 * @brief Send command to the card and get response
 *
 * This function returns when command is sent and response is received,
 * or data is transferred, or timeout occurs.
 *
 * @note This function is not thread safe w.r.t. init/deinit functions,
 *       and bus width/clock speed configuration functions. Multiple tasks
 *       can call sdmmc_host_do_transaction as long as other sdmmc_host_*
 *       functions are not called.
 *
 * @attention Data buffer passed in cmdinfo->data must be in DMA capable memory and aligned to 4 byte boundary. If it's
 *            behind the cache, both cmdinfo->data and cmdinfo->buflen need to be aligned to cache line boundary.
 *
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param cmdinfo   pointer to structure describing command and data to transfer
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_TIMEOUT if response or data transfer has timed out
 *      - ESP_ERR_INVALID_CRC if response or data transfer CRC check has failed
 *      - ESP_ERR_INVALID_RESPONSE if the card has sent an invalid response
 *      - ESP_ERR_INVALID_SIZE if the size of data transfer is not valid in SD protocol
 *      - ESP_ERR_INVALID_ARG if the data buffer is not in DMA capable memory
 */
//go:linkname SdmmcHostDoTransaction C.sdmmc_host_do_transaction
func SdmmcHostDoTransaction(slot c.Int, cmdinfo *SdmmcCommandT) EspErrT

/**
 * @brief Enable IO interrupts
 *
 * This function configures the host to accept SDIO interrupts.
 *
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @return returns ESP_OK, other errors possible in the future
 */
//go:linkname SdmmcHostIoIntEnable C.sdmmc_host_io_int_enable
func SdmmcHostIoIntEnable(slot c.Int) EspErrT

/**
 * @brief Block until an SDIO interrupt is received, or timeout occurs
 * @param slot  slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param timeout_ticks  number of RTOS ticks to wait for the interrupt
 * @return
 *  - ESP_OK on success (interrupt received)
 *  - ESP_ERR_TIMEOUT if the interrupt did not occur within timeout_ticks
 */
//go:linkname SdmmcHostIoIntWait C.sdmmc_host_io_int_wait
func SdmmcHostIoIntWait(slot c.Int, timeout_ticks TickTypeT) EspErrT

/**
 * @brief Disable SDMMC host and release allocated resources gracefully
 *
 * @note If there are more than 1 active slots, this function will just decrease the reference count
 *       and won't actually disable the host until the last slot is disabled
 *
 * @note This function is not thread safe
 *
 * @param slot slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if SDMMC host has not been initialized
 *      - ESP_ERR_INVALID_ARG if invalid slot number is used
 */
//go:linkname SdmmcHostDeinitSlot C.sdmmc_host_deinit_slot
func SdmmcHostDeinitSlot(slot c.Int) EspErrT

/**
 * @brief Disable SDMMC host and release allocated resources forcefully
 *
 * @note This function will deinitialize the host immediately, regardless of the number of active slots
 *
 * @note This function is not thread safe
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_STATE if SDMMC host has not been initialized
 */
//go:linkname SdmmcHostDeinit C.sdmmc_host_deinit
func SdmmcHostDeinit() EspErrT

/**
 * @brief Provides a real frequency used for an SD card installed on specific slot
 * of SD/MMC host controller
 *
 * This function calculates real working frequency given by current SD/MMC host
 * controller setup for required slot: it reads associated host and card dividers
 * from corresponding SDMMC registers, calculates respective frequency and stores
 * the value into the 'real_freq_khz' parameter
 *
 * @param slot slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param[out] real_freq_khz output parameter for the result frequency (in kHz)
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG on real_freq_khz == NULL or invalid slot number used
 */
//go:linkname SdmmcHostGetRealFreq C.sdmmc_host_get_real_freq
func SdmmcHostGetRealFreq(slot c.Int, real_freq_khz *c.Int) EspErrT

/**
 * @brief set input delay
 *
 * @note ESP32 doesn't support this feature, you will get an `ESP_ERR_NOT_SUPPORTED`
 *
 * - This API sets delay when the SDMMC Host samples the signal from the SD Slave.
 * - This API will check if the given `delay_phase` is valid or not.
 * - This API will print out the delay time, in picosecond (ps)
 *
 * @param slot         slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param delay_phase  delay phase, this API will convert the phase into picoseconds and print it out
 *
 * @return
 *        - ESP_OK:                ON success.
 *        - ESP_ERR_INVALID_ARG:   Invalid argument.
 *        - ESP_ERR_NOT_SUPPORTED: ESP32 doesn't support this feature.
 */
//go:linkname SdmmcHostSetInputDelay C.sdmmc_host_set_input_delay
func SdmmcHostSetInputDelay(slot c.Int, delay_phase SdmmcDelayPhaseT) EspErrT

/**
 * @brief Get the DMA memory information for the host driver
 *
 * @param[in]  slot slot number (SDMMC_HOST_SLOT_0 or SDMMC_HOST_SLOT_1)
 * @param[out] dma_mem_info  DMA memory information structure
 * @return
 *        - ESP_OK:                ON success.
 *        - ESP_ERR_INVALID_ARG:   Invalid argument.
 */
//go:linkname SdmmcHostGetDmaInfo C.sdmmc_host_get_dma_info
func SdmmcHostGetDmaInfo(slot c.Int, dma_mem_info *EspDmaMemInfoT) EspErrT

/**
 * @brief Check if the slot is set to uhs1 or not
 *
 * @param[in]  slot     Slot id
 * @param[out] is_uhs1  Is uhs1 or not
 *
 * @return
 *        - ESP_OK:                on success
 *        - ESP_ERR_INVALID_STATE: driver not in correct state
 */
//go:linkname SdmmcHostIsSlotSetToUhs1 C.sdmmc_host_is_slot_set_to_uhs1
func SdmmcHostIsSlotSetToUhs1(slot c.Int, is_uhs1 *bool) EspErrT

/**
 * @brief Get the state of SDMMC host
 *
 * @param[out] state output parameter for SDMMC host state structure
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_ERR_INVALID_ARG on invalid argument
 */
// llgo:link (*SdmmcHostStateT).SdmmcHostGetState C.sdmmc_host_get_state
func (recv_ *SdmmcHostStateT) SdmmcHostGetState() EspErrT {
	return 0
}
