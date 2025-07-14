package pthread

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** pthread configuration structure that influences pthread creation */

type EspPthreadCfgT struct {
	StackSize      c.SizeT
	Prio           c.SizeT
	InheritCfg     bool
	ThreadName     *c.Char
	PinToCore      c.Int
	StackAllocCaps c.Uint32T
}

/**
 * @brief Creates a default pthread configuration based
 * on the values set via menuconfig.
 *
 * @return
 *      A default configuration structure.
 */
//go:linkname EspPthreadGetDefaultConfig C.esp_pthread_get_default_config
func EspPthreadGetDefaultConfig() EspPthreadCfgT

/**
 * @brief Configure parameters for creating pthread
 *
 * This API allows you to configure how the subsequent
 * pthread_create() call will behave. This call can be used to setup
 * configuration parameters like stack size, priority, configuration
 * inheritance etc.
 *
 * If the 'inherit' flag in the configuration structure is enabled,
 * then the same configuration is also inherited in the thread
 * subtree.
 *
 * @note If cfg->stack_alloc_caps is 0, it is automatically set to valid default stack memory
 *       capabilities. If cfg->stack_alloc_caps is non-zero, the developer is responsible for its correctenss.
 *       This function only checks that the capabilities are MALLOC_CAP_8BIT, the rest is unchecked.
 * @note Passing non-NULL attributes to pthread_create() will override
 *       the stack_size parameter set using this API
 *
 * @param cfg The pthread config parameters
 *
 * @return
 *      - ESP_OK if configuration was successfully set
 *      - ESP_ERR_NO_MEM if out of memory
 *      - ESP_ERR_INVALID_ARG if cfg is NULL
 *      - ESP_ERR_INVALID_ARG if stack_size is less than PTHREAD_STACK_MIN
 *      - ESP_ERR_INVALID_ARG if stack_alloc_caps does not include MALLOC_CAP_8BIT
 */
// llgo:link (*EspPthreadCfgT).EspPthreadSetCfg C.esp_pthread_set_cfg
func (recv_ *EspPthreadCfgT) EspPthreadSetCfg() EspErrT {
	return 0
}

/**
 * @brief Get current pthread creation configuration
 *
 * This will retrieve the current configuration that will be used for
 * creating threads.
 *
 * @param p Pointer to the pthread config structure that will be
 * updated with the currently configured parameters
 *
 * @return
 *      - ESP_OK if the configuration was available
 *      - ESP_ERR_INVALID_ARG if p is NULL
 *      - ESP_ERR_NOT_FOUND if a configuration wasn't previously set
 */
// llgo:link (*EspPthreadCfgT).EspPthreadGetCfg C.esp_pthread_get_cfg
func (recv_ *EspPthreadCfgT) EspPthreadGetCfg() EspErrT {
	return 0
}

/**
 * @brief Initialize pthread library
 */
//go:linkname EspPthreadInit C.esp_pthread_init
func EspPthreadInit() EspErrT
