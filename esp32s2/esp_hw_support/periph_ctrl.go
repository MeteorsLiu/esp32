package esp_hw_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** @cond */
// The following functions are not intended to be used directly by the developers
//go:linkname PeriphRccAcquireEnter C.periph_rcc_acquire_enter
func PeriphRccAcquireEnter(periph PeriphModuleT) c.Uint8T

//go:linkname PeriphRccAcquireExit C.periph_rcc_acquire_exit
func PeriphRccAcquireExit(periph PeriphModuleT, ref_count c.Uint8T)

//go:linkname PeriphRccReleaseEnter C.periph_rcc_release_enter
func PeriphRccReleaseEnter(periph PeriphModuleT) c.Uint8T

//go:linkname PeriphRccReleaseExit C.periph_rcc_release_exit
func PeriphRccReleaseExit(periph PeriphModuleT, ref_count c.Uint8T)

//go:linkname PeriphRccEnter C.periph_rcc_enter
func PeriphRccEnter()

//go:linkname PeriphRccExit C.periph_rcc_exit
func PeriphRccExit()

/**
 * @brief Enable peripheral module by un-gating the clock and de-asserting the reset signal.
 *
 * @param[in] periph Peripheral module
 *
 * @note If @c periph_module_enable() is called a number of times,
 *       @c periph_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
//go:linkname PeriphModuleEnable C.periph_module_enable
func PeriphModuleEnable(periph PeriphModuleT)

/**
 * @brief Disable peripheral module by gating the clock and asserting the reset signal.
 *
 * @param[in] periph Peripheral module
 *
 * @note If @c periph_module_enable() is called a number of times,
 *       @c periph_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
//go:linkname PeriphModuleDisable C.periph_module_disable
func PeriphModuleDisable(periph PeriphModuleT)

/**
 * @brief Reset peripheral module by asserting and de-asserting the reset signal.
 *
 * @param[in] periph Peripheral module
 *
 * @note Calling this function does not enable or disable the clock for the module.
 */
//go:linkname PeriphModuleReset C.periph_module_reset
func PeriphModuleReset(periph PeriphModuleT)

/**
 * @brief Enable Wi-Fi and BT common module
 *
 * @note If @c wifi_bt_common_module_enable() is called a number of times,
 *       @c wifi_bt_common_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
//go:linkname WifiBtCommonModuleEnable C.wifi_bt_common_module_enable
func WifiBtCommonModuleEnable()

/**
 * @brief Disable Wi-Fi and BT common module
 *
 * @note If @c wifi_bt_common_module_enable() is called a number of times,
 *       @c wifi_bt_common_module_disable() has to be called the same number of times,
 *       in order to put the peripheral into disabled state.
 */
//go:linkname WifiBtCommonModuleDisable C.wifi_bt_common_module_disable
func WifiBtCommonModuleDisable()

/**
 * @brief Enable Wi-Fi module
 *
 * @note Calling this function will only enable Wi-Fi module.
 */
//go:linkname WifiModuleEnable C.wifi_module_enable
func WifiModuleEnable()

/**
 * @brief Disable Wi-Fi module
 *
 * @note Calling this function will only disable Wi-Fi module.
 */
//go:linkname WifiModuleDisable C.wifi_module_disable
func WifiModuleDisable()
