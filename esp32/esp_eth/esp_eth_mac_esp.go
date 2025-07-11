package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EmacRmiiClockModeT c.Int

const (
	EMAC_CLK_DEFAULT EmacRmiiClockModeT = 0
	EMAC_CLK_EXT_IN  EmacRmiiClockModeT = 1
	EMAC_CLK_OUT     EmacRmiiClockModeT = 2
)

type EmacRmiiClockGpioT c.Int

const (
	EMAC_CLK_IN_GPIO       EmacRmiiClockGpioT = 0
	EMAC_APPL_CLK_OUT_GPIO EmacRmiiClockGpioT = 0
	EMAC_CLK_OUT_GPIO      EmacRmiiClockGpioT = 16
	EMAC_CLK_OUT_180_GPIO  EmacRmiiClockGpioT = 17
)

/**
 * @brief Ethernet MAC Clock Configuration
 *
 */

type EthMacClockConfigT struct {
	Rmii struct {
		ClockMode EmacRmiiClockModeT
		ClockGpio EmacRmiiClockGpioT
	}
}

/**
 * @brief EMAC SMI GPIO configuration
 */

type EmacEspSmiGpioConfigT struct {
	MdcNum  c.Int
	MdioNum c.Int
}

/**
 * @brief EMAC MII data interface GPIO configuration
 */

type EthMacMiiGpioConfigT struct {
	TxClkNum c.Int
	TxEnNum  c.Int
	Txd0Num  c.Int
	Txd1Num  c.Int
	Txd2Num  c.Int
	Txd3Num  c.Int
	RxClkNum c.Int
	RxDvNum  c.Int
	Rxd0Num  c.Int
	Rxd1Num  c.Int
	Rxd2Num  c.Int
	Rxd3Num  c.Int
	ColInNum c.Int
	CrsInNum c.Int
	TxErNum  c.Int
	RxErNum  c.Int
}

/**
 * @brief EMAC RMII data interface GPIO configuration
 */

type EthMacRmiiGpioConfigT struct {
	TxEnNum  c.Int
	Txd0Num  c.Int
	Txd1Num  c.Int
	CrsDvNum c.Int
	Rxd0Num  c.Int
	Rxd1Num  c.Int
}

/**
* @brief EMAC specific configuration
*
 */
type EthEsp32EmacConfigT struct {
	Interface    EthDataInterfaceT
	ClockConfig  EthMacClockConfigT
	DmaBurstLen  EthMacDmaBurstLenT
	IntrPriority c.Int
}
type EthMacEspIoCmdT c.Int

const (
	ETH_MAC_ESP_CMD_SET_TDES0_CFG_BITS   EthMacEspIoCmdT = 4095
	ETH_MAC_ESP_CMD_CLEAR_TDES0_CFG_BITS EthMacEspIoCmdT = 4096
	ETH_MAC_ESP_CMD_PTP_ENABLE           EthMacEspIoCmdT = 4097
)

/**
* @brief Create ESP32 Ethernet MAC instance
*
* @param esp32_config: EMAC specific configuration
* @param config:       Ethernet MAC configuration
*
* @return
*      - instance: create MAC instance successfully
*      - NULL: create MAC instance failed because some error occurred
 */
// llgo:link (*EthEsp32EmacConfigT).EspEthMacNewEsp32 C.esp_eth_mac_new_esp32
func (recv_ *EthEsp32EmacConfigT) EspEthMacNewEsp32(config *EthMacConfigT) *EspEthMacT {
	return nil
}
