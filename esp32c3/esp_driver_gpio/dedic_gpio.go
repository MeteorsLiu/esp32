package esp_driver_gpio

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type DedicGpioBundleT struct {
	Unused [8]uint8
}
type DedicGpioBundleHandleT *DedicGpioBundleT

/**
 * @brief Type of Dedicated GPIO bundle configuration
 */

type DedicGpioBundleConfigT struct {
	GpioArray *c.Int
	ArraySize c.SizeT
	Flags     struct {
		InEn      c.Uint
		InInvert  c.Uint
		OutEn     c.Uint
		OutInvert c.Uint
	}
}

/**
 * @brief Create GPIO bundle and return the handle
 *
 * @param[in] config Configuration of GPIO bundle
 * @param[out] ret_bundle Returned handle of the new created GPIO bundle
 * @return
 *      - ESP_OK: Create GPIO bundle successfully
 *      - ESP_ERR_INVALID_ARG: Create GPIO bundle failed because of invalid argument
 *      - ESP_ERR_NO_MEM: Create GPIO bundle failed because of no capable memory
 *      - ESP_ERR_NOT_FOUND: Create GPIO bundle failed because of no enough continuous dedicated channels
 *      - ESP_FAIL: Create GPIO bundle failed because of other error
 *
 * @note One has to enable at least input or output mode in "config" parameter.
 */
// llgo:link (*DedicGpioBundleConfigT).DedicGpioNewBundle C.dedic_gpio_new_bundle
func (recv_ *DedicGpioBundleConfigT) DedicGpioNewBundle(ret_bundle *DedicGpioBundleHandleT) EspErrT {
	return 0
}

/**
 * @brief Destroy GPIO bundle
 *
 * @param[in] bundle Handle of GPIO bundle that returned from "dedic_gpio_new_bundle"
 * @return
 *      - ESP_OK: Destroy GPIO bundle successfully
 *      - ESP_ERR_INVALID_ARG: Destroy GPIO bundle failed because of invalid argument
 *      - ESP_FAIL: Destroy GPIO bundle failed because of other error
 */
//go:linkname DedicGpioDelBundle C.dedic_gpio_del_bundle
func DedicGpioDelBundle(bundle DedicGpioBundleHandleT) EspErrT

/**@{*/
/**
 * @brief Get allocated channel mask
 *
 * @param[in] bundle Handle of GPIO bundle that returned from "dedic_gpio_new_bundle"
 * @param[out] mask Returned mask value for on specific direction (in or out)
 * @return
 *      - ESP_OK: Get channel mask successfully
 *      - ESP_ERR_INVALID_ARG: Get channel mask failed because of invalid argument
 *      - ESP_FAIL: Get channel mask failed because of other error
 *
 * @note Each bundle should have at least one mask (in or/and out), based on bundle configuration.
 * @note With the returned mask, user can directly invoke LL function like "dedic_gpio_cpu_ll_write_mask"
 *       or write assembly code with dedicated GPIO instructions, to get better performance on GPIO manipulation.
 */
//go:linkname DedicGpioGetOutMask C.dedic_gpio_get_out_mask
func DedicGpioGetOutMask(bundle DedicGpioBundleHandleT, mask *c.Uint32T) EspErrT

//go:linkname DedicGpioGetInMask C.dedic_gpio_get_in_mask
func DedicGpioGetInMask(bundle DedicGpioBundleHandleT, mask *c.Uint32T) EspErrT

/**@{*/
/**
 * @brief Get the channel offset of the GPIO bundle
 *
 * A GPIO bundle maps the GPIOS of a particular direction to a consecutive set of channels within
 * a particular GPIO bank of a particular CPU. This function returns the offset to
 * the bundle's first channel of a particular direction within the bank.
 *
 * @param[in] bundle Handle of GPIO bundle that returned from "dedic_gpio_new_bundle"
 * @param[out] offset Offset value to the first channel of a specific direction (in or out)
 * @return
 *      - ESP_OK: Get channel offset successfully
 *      - ESP_ERR_INVALID_ARG: Get channel offset failed because of invalid argument
 *      - ESP_FAIL: Get channel offset failed because of other error
 */
//go:linkname DedicGpioGetOutOffset C.dedic_gpio_get_out_offset
func DedicGpioGetOutOffset(bundle DedicGpioBundleHandleT, offset *c.Uint32T) EspErrT

//go:linkname DedicGpioGetInOffset C.dedic_gpio_get_in_offset
func DedicGpioGetInOffset(bundle DedicGpioBundleHandleT, offset *c.Uint32T) EspErrT

/**
 * @brief Write value to GPIO bundle
 *
 * @param[in] bundle Handle of GPIO bundle that returned from "dedic_gpio_new_bundle"
 * @param[in] mask Mask of the GPIOs to be written in the given bundle
 * @param[in] value Value to write to given GPIO bundle, low bit represents low member in the bundle
 *
 * @note The mask is seen from the view of GPIO bundle.
 *       For example, bundleA contains [GPIO10, GPIO12, GPIO17], to set GPIO17 individually, the mask should be 0x04.
 * @note For performance reasons, this function doesn't check the validity of any parameters, and is placed in IRAM.
 */
//go:linkname DedicGpioBundleWrite C.dedic_gpio_bundle_write
func DedicGpioBundleWrite(bundle DedicGpioBundleHandleT, mask c.Uint32T, value c.Uint32T)

/**
 * @brief Read the value that output from the given GPIO bundle
 *
 * @param[in] bundle Handle of GPIO bundle that returned from "dedic_gpio_new_bundle"
 * @return Value that output from the GPIO bundle, low bit represents low member in the bundle
 *
 * @note For performance reasons, this function doesn't check the validity of any parameters, and is placed in IRAM.
 */
//go:linkname DedicGpioBundleReadOut C.dedic_gpio_bundle_read_out
func DedicGpioBundleReadOut(bundle DedicGpioBundleHandleT) c.Uint32T

/**
 * @brief Read the value that input to the given GPIO bundle
 *
 * @param[in] bundle Handle of GPIO bundle that returned from "dedic_gpio_new_bundle"
 * @return Value that input to the GPIO bundle, low bit represents low member in the bundle
 *
 * @note For performance reasons, this function doesn't check the validity of any parameters, and is placed in IRAM.
 */
//go:linkname DedicGpioBundleReadIn C.dedic_gpio_bundle_read_in
func DedicGpioBundleReadIn(bundle DedicGpioBundleHandleT) c.Uint32T
