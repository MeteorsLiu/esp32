package esp_phy

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspPhyAntT c.Int

const (
	ESP_PHY_ANT_ANT0 EspPhyAntT = 0
	ESP_PHY_ANT_ANT1 EspPhyAntT = 1
	ESP_PHY_ANT_MAX  EspPhyAntT = 2
)

type EspPhyAntModeT c.Int

const (
	ESP_PHY_ANT_MODE_ANT0 EspPhyAntModeT = 0
	ESP_PHY_ANT_MODE_ANT1 EspPhyAntModeT = 1
	ESP_PHY_ANT_MODE_AUTO EspPhyAntModeT = 2
	ESP_PHY_ANT_MODE_MAX  EspPhyAntModeT = 3
)

/**
 * @brief PHY GPIO configuration for antenna selection
 */

type EspPhyAntGpioT struct {
	GpioSelect c.Uint8T
	GpioNum    c.Uint8T
}

/**
 * @brief PHY GPIOs configuration for antenna selection
 */

type EspPhyAntGpioConfigT struct {
	GpioCfg [4]EspPhyAntGpioT
}

/**
 * @brief PHY antenna configuration
 */

type EspPhyAntConfigT struct {
	RxAntMode    EspPhyAntModeT
	RxAntDefault EspPhyAntT
	TxAntMode    EspPhyAntModeT
	EnabledAnt0  c.Uint8T
	EnabledAnt1  c.Uint8T
}

/**
 * @brief Set antenna GPIO configuration
 *
 * @param config : Antenna GPIO configuration.
 *
 * @return
 *                  - ESP_OK : success
 *                  - other  : failed
 */
// llgo:link (*EspPhyAntGpioConfigT).EspPhySetAntGpio C.esp_phy_set_ant_gpio
func (recv_ *EspPhyAntGpioConfigT) EspPhySetAntGpio() EspErrT {
	return 0
}

/**
 * @brief Get current antenna GPIO configuration
 *
 * @param config : Antenna GPIO configuration.
 *
 * @return
 *                  - ESP_OK : success
 *                  - other  : failed
 */
// llgo:link (*EspPhyAntGpioConfigT).EspPhyGetAntGpio C.esp_phy_get_ant_gpio
func (recv_ *EspPhyAntGpioConfigT) EspPhyGetAntGpio() EspErrT {
	return 0
}

/**
 * @brief Set antenna configuration
 *
 * @param config : Antenna configuration.
 *
 * @return
 *                  - ESP_OK : success
 *                  - other  : failed
 */
// llgo:link (*EspPhyAntConfigT).EspPhySetAnt C.esp_phy_set_ant
func (recv_ *EspPhyAntConfigT) EspPhySetAnt() EspErrT {
	return 0
}

/**
 * @brief Get current antenna configuration
 *
 * @param config : Antenna configuration.
 *
 * @return
 *                  - ESP_OK : success
 *                  - other  : failed
 */
// llgo:link (*EspPhyAntConfigT).EspPhyGetAnt C.esp_phy_get_ant
func (recv_ *EspPhyAntConfigT) EspPhyGetAnt() EspErrT {
	return 0
}
