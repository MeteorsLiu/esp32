package esp_psram

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Initialize PSRAM interface/hardware.
 *
 * @return
 *        - ESP_OK:                On success
 *        - ESP_FAIL:              PSRAM isn't initialized successfully, potential reason would be: wrong VDDSDIO, invalid chip ID, etc.
 *        - ESP_ERR_INVALID_STATE: PSRAM is initialized already
 */
//go:linkname EspPsramInit C.esp_psram_init
func EspPsramInit() EspErrT

/**
 * @brief If PSRAM has been initialized
 *
 * @return
 *          - true:  PSRAM has been initialized successfully
 *          - false: PSRAM hasn't been initialized or initialized failed
 */
//go:linkname EspPsramIsInitialized C.esp_psram_is_initialized
func EspPsramIsInitialized() bool

/**
 * @brief Get the available size of the attached PSRAM chip
 *
 * @return Size in bytes, or 0 if PSRAM isn't successfully initialized
 */
//go:linkname EspPsramGetSize C.esp_psram_get_size
func EspPsramGetSize() c.SizeT
