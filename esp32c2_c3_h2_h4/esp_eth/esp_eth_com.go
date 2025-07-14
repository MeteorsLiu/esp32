package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ETH_CMD_CUSTOM_MAC_CMDS_OFFSET = 0x0FFF
const ETH_CMD_CUSTOM_PHY_CMDS_OFFSET = 0x1FFF

type EspEthStateT c.Int

const (
	ETH_STATE_LLINIT EspEthStateT = 0
	ETH_STATE_DEINIT EspEthStateT = 1
	ETH_STATE_LINK   EspEthStateT = 2
	ETH_STATE_SPEED  EspEthStateT = 3
	ETH_STATE_DUPLEX EspEthStateT = 4
	ETH_STATE_PAUSE  EspEthStateT = 5
)

type EspEthMediatorS struct {
	Unused [8]uint8
}
type EspEthMediatorT EspEthMediatorS
type EthEventT c.Int

const (
	ETHERNET_EVENT_START        EthEventT = 0
	ETHERNET_EVENT_STOP         EthEventT = 1
	ETHERNET_EVENT_CONNECTED    EthEventT = 2
	ETHERNET_EVENT_DISCONNECTED EthEventT = 3
)
