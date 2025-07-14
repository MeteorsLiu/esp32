package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief IEEE 802.3 PHY object infostructure
 *
 */

type Phy8023T struct {
	Parent              EspEthPhyT
	Eth                 *EspEthMediatorT
	Addr                c.Int
	ResetTimeoutMs      c.Uint32T
	AutonegoTimeoutMs   c.Uint32T
	LinkStatus          EthLinkT
	ResetGpioNum        c.Int
	HwResetAssertTimeUs c.Int32T
	PostHwResetDelayMs  c.Int32T
}
type EspEthPhy8023MmdFuncT c.Int

const (
	MMD_FUNC_ADDRESS     EspEthPhy8023MmdFuncT = 0
	MMD_FUNC_DATA_NOINCR EspEthPhy8023MmdFuncT = 1
	MMD_FUNC_DATA_RWINCR EspEthPhy8023MmdFuncT = 2
	MMD_FUNC_DATA_WINCR  EspEthPhy8023MmdFuncT = 3
)

/**
 * @brief Set Ethernet mediator
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param eth Ethernet mediator pointer
 * @return
 *      - ESP_OK: Ethermet mediator set successfully
 *      - ESP_ERR_INVALID_ARG: if @c eth is @c NULL
 */
// llgo:link (*Phy8023T).EspEthPhy8023SetMediator C.esp_eth_phy_802_3_set_mediator
func (recv_ *Phy8023T) EspEthPhy8023SetMediator(eth *EspEthMediatorT) EspErrT {
	return 0
}

/**
 * @brief Reset PHY
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @return
 *      - ESP_OK: Ethernet PHY reset successfully
 *      - ESP_FAIL: reset Ethernet PHY failed because some error occurred
 */
// llgo:link (*Phy8023T).EspEthPhy8023Reset C.esp_eth_phy_802_3_reset
func (recv_ *Phy8023T) EspEthPhy8023Reset() EspErrT {
	return 0
}

/**
 * @brief Control autonegotiation mode of Ethernet PHY
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param cmd autonegotiation command enumeration
 * @param[out] autonego_en_stat autonegotiation enabled flag
 * @return
 *      - ESP_OK: Ethernet PHY autonegotiation configured successfully
 *      - ESP_FAIL: Ethernet PHY autonegotiation configuration fail because some error occurred
 *      - ESP_ERR_INVALID_ARG: invalid value of @c cmd
 */
// llgo:link (*Phy8023T).EspEthPhy8023AutonegoCtrl C.esp_eth_phy_802_3_autonego_ctrl
func (recv_ *Phy8023T) EspEthPhy8023AutonegoCtrl(cmd EthPhyAutonegCmdT, autonego_en_stat *bool) EspErrT {
	return 0
}

/**
 * @brief Retrieve link status and propagate the status to higher layers if the status changed
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @return
 *      - ESP_OK: Ethernet PHY link status retrieved successfully
 *      - ESP_FAIL: Error occurred during reading registry
 */
// llgo:link (*Phy8023T).EspEthPhy8023UpdtLinkDupSpd C.esp_eth_phy_802_3_updt_link_dup_spd
func (recv_ *Phy8023T) EspEthPhy8023UpdtLinkDupSpd() EspErrT {
	return 0
}

/**
 * @brief Power control of Ethernet PHY
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param enable set true to power ON Ethernet PHY; set false to power OFF Ethernet PHY
 * @return
 *      - ESP_OK: Ethernet PHY power down mode set successfully
 *      - ESP_FAIL: Ethernet PHY power up or power down failed because some error occurred
 */
// llgo:link (*Phy8023T).EspEthPhy8023Pwrctl C.esp_eth_phy_802_3_pwrctl
func (recv_ *Phy8023T) EspEthPhy8023Pwrctl(enable bool) EspErrT {
	return 0
}

/**
 * @brief Set Ethernet PHY address
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param addr new PHY address
 * @return
 *      - ESP_OK: Ethernet PHY address set
 */
// llgo:link (*Phy8023T).EspEthPhy8023SetAddr C.esp_eth_phy_802_3_set_addr
func (recv_ *Phy8023T) EspEthPhy8023SetAddr(addr c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Get Ethernet PHY address
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param[out] addr Ethernet PHY address
 * @return
 *      - ESP_OK: Ethernet PHY address read successfully
 *      - ESP_ERR_INVALID_ARG: @c addr pointer is @c NULL
 */
// llgo:link (*Phy8023T).EspEthPhy8023GetAddr C.esp_eth_phy_802_3_get_addr
func (recv_ *Phy8023T) EspEthPhy8023GetAddr(addr *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Advertise pause function ability
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param ability enable or disable pause ability
 * @return
 *      - ESP_OK: pause ability set successfully
 *      - ESP_FAIL: Advertise pause function ability failed because some error occurred
 */
// llgo:link (*Phy8023T).EspEthPhy8023AdvertisePauseAbility C.esp_eth_phy_802_3_advertise_pause_ability
func (recv_ *Phy8023T) EspEthPhy8023AdvertisePauseAbility(ability c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Set Ethernet PHY loopback mode
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param enable set true to enable loopback; set false to disable loopback
 * @return
 *      - ESP_OK: Ethernet PHY loopback mode set successfully
 *      - ESP_FAIL: Ethernet PHY loopback configuration failed because some error occurred
 */
// llgo:link (*Phy8023T).EspEthPhy8023Loopback C.esp_eth_phy_802_3_loopback
func (recv_ *Phy8023T) EspEthPhy8023Loopback(enable bool) EspErrT {
	return 0
}

/**
 * @brief Set Ethernet PHY speed
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param speed new speed of Ethernet PHY link
 * @return
 *      - ESP_OK: Ethernet PHY speed set successfully
 *      - ESP_FAIL: Set Ethernet PHY speed failed because some error occurred
 */
// llgo:link (*Phy8023T).EspEthPhy8023SetSpeed C.esp_eth_phy_802_3_set_speed
func (recv_ *Phy8023T) EspEthPhy8023SetSpeed(speed EthSpeedT) EspErrT {
	return 0
}

/**
 * @brief Set Ethernet PHY duplex mode
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param duplex new duplex mode for Ethernet PHY link
 * @return
 *      - ESP_OK: Ethernet PHY duplex mode set successfully
 *      - ESP_ERR_INVALID_STATE: unable to set duplex mode to Half if loopback is enabled
 *      - ESP_FAIL: Set Ethernet PHY duplex mode failed because some error occurred
 */
// llgo:link (*Phy8023T).EspEthPhy8023SetDuplex C.esp_eth_phy_802_3_set_duplex
func (recv_ *Phy8023T) EspEthPhy8023SetDuplex(duplex EthDuplexT) EspErrT {
	return 0
}

/**
 * @brief Set Ethernet PHY link status
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param link new link status
 * @return
 *      - ESP_OK: Ethernet PHY link set successfully
 */
// llgo:link (*Phy8023T).EspEthPhy8023SetLink C.esp_eth_phy_802_3_set_link
func (recv_ *Phy8023T) EspEthPhy8023SetLink(link EthLinkT) EspErrT {
	return 0
}

/**
 * @brief Initialize Ethernet PHY
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @return
 *      - ESP_OK: Ethernet PHY initialized successfully
 */
// llgo:link (*Phy8023T).EspEthPhy8023Init C.esp_eth_phy_802_3_init
func (recv_ *Phy8023T) EspEthPhy8023Init() EspErrT {
	return 0
}

/**
 * @brief Power off Eternet PHY
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @return
 *      - ESP_OK: Ethernet PHY powered off successfully
 */
// llgo:link (*Phy8023T).EspEthPhy8023Deinit C.esp_eth_phy_802_3_deinit
func (recv_ *Phy8023T) EspEthPhy8023Deinit() EspErrT {
	return 0
}

/**
 * @brief Delete Ethernet PHY infostructure
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @return
 *      - ESP_OK: Ethernet PHY infostructure deleted
 */
// llgo:link (*Phy8023T).EspEthPhy8023Del C.esp_eth_phy_802_3_del
func (recv_ *Phy8023T) EspEthPhy8023Del() EspErrT {
	return 0
}

/**
 * @brief Performs hardware reset with specific reset pin assertion time
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param reset_assert_us Hardware reset pin assertion time
 * @return
 *      - ESP_OK: reset Ethernet PHY successfully
 *      - ESP_ERR_NOT_ALLOWED: reset GPIO not defined
 */
// llgo:link (*Phy8023T).EspEthPhy8023ResetHw C.esp_eth_phy_802_3_reset_hw
func (recv_ *Phy8023T) EspEthPhy8023ResetHw(reset_assert_us c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Detect PHY address
 *
 * @param eth Mediator of Ethernet driver
 * @param[out] detected_addr: a valid address after detection
 * @return
 *       - ESP_OK: detect phy address successfully
 *       - ESP_ERR_INVALID_ARG: invalid parameter
 *       - ESP_ERR_NOT_FOUND: can't detect any PHY device
 *       - ESP_FAIL: detect phy address failed because some error occurred
 */
// llgo:link (*EspEthMediatorT).EspEthPhy8023DetectPhyAddr C.esp_eth_phy_802_3_detect_phy_addr
func (recv_ *EspEthMediatorT) EspEthPhy8023DetectPhyAddr(detected_addr *c.Int) EspErrT {
	return 0
}

/**
 * @brief Performs basic PHY chip initialization
 *
 * @note It should be called as the first function in PHY specific driver instance
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @return
 *      - ESP_OK: initialized Ethernet PHY successfully
 *      - ESP_FAIL: initialization of Ethernet PHY failed because some error occurred
 *      - ESP_ERR_INVALID_ARG: invalid argument
 *      - ESP_ERR_NOT_FOUND: PHY device not detected
 *      - ESP_ERR_TIMEOUT: MII Management read/write operation timeout
 *      - ESP_ERR_INVALID_STATE: PHY is in invalid state to perform requested operation
 */
// llgo:link (*Phy8023T).EspEthPhy8023BasicPhyInit C.esp_eth_phy_802_3_basic_phy_init
func (recv_ *Phy8023T) EspEthPhy8023BasicPhyInit() EspErrT {
	return 0
}

/**
 * @brief Performs basic PHY chip de-initialization
 *
 * @note It should be called as the last function in PHY specific driver instance
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @return
 *      - ESP_OK: de-initialized Ethernet PHY successfully
 *      - ESP_FAIL: de-initialization of Ethernet PHY failed because some error occurred
 *      - ESP_ERR_TIMEOUT: MII Management read/write operation timeout
 *      - ESP_ERR_INVALID_STATE: PHY is in invalid state to perform requested operation
 */
// llgo:link (*Phy8023T).EspEthPhy8023BasicPhyDeinit C.esp_eth_phy_802_3_basic_phy_deinit
func (recv_ *Phy8023T) EspEthPhy8023BasicPhyDeinit() EspErrT {
	return 0
}

/**
 * @brief Reads raw content of OUI field
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param[out] oui OUI value
 * @return
 *      - ESP_OK: OUI field read successfully
 *      - ESP_FAIL: OUI field read failed because some error occurred
 *      - ESP_ERR_INVALID_ARG: invalid @c oui argument
 *      - ESP_ERR_TIMEOUT: MII Management read/write operation timeout
 *      - ESP_ERR_INVALID_STATE: PHY is in invalid state to perform requested operation
 */
// llgo:link (*Phy8023T).EspEthPhy8023ReadOui C.esp_eth_phy_802_3_read_oui
func (recv_ *Phy8023T) EspEthPhy8023ReadOui(oui *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Reads manufacturer’s model and revision number
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param[out] model Manufacturer’s model number (can be NULL when not required)
 * @param[out] rev Manufacturer’s revision number (can be NULL when not required)
 * @return
 *      - ESP_OK: Manufacturer’s info read successfully
 *      - ESP_FAIL: Manufacturer’s info read failed because some error occurred
 *      - ESP_ERR_TIMEOUT: MII Management read/write operation timeout
 *      - ESP_ERR_INVALID_STATE: PHY is in invalid state to perform requested operation
 */
// llgo:link (*Phy8023T).EspEthPhy8023ReadManufacInfo C.esp_eth_phy_802_3_read_manufac_info
func (recv_ *Phy8023T) EspEthPhy8023ReadManufacInfo(model *c.Uint8T, rev *c.Uint8T) EspErrT {
	return 0
}

/**
 * @brief Reads MDIO device's internal address register
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param devaddr Address of MDIO device
 * @param[out] mmd_addr Current address stored in device's register
 * @return
 *      - ESP_OK: Address register read successfully
 *      - ESP_FAIL: Address register read failed because of some error occurred
 *      - ESP_ERR_INVALID_ARG: Device address provided is out of range (hardware limits device address to 5 bits)
 */
// llgo:link (*Phy8023T).EspEthPhy8023GetMmdAddr C.esp_eth_phy_802_3_get_mmd_addr
func (recv_ *Phy8023T) EspEthPhy8023GetMmdAddr(devaddr c.Uint8T, mmd_addr *c.Uint16T) EspErrT {
	return 0
}

/**
 * @brief Write to DIO device's internal address register
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param devaddr Address of MDIO device
 * @param[out] mmd_addr New value of MDIO device's address register value
 * @return
 *      - ESP_OK: Address register written to successfully
 *      - ESP_FAIL: Address register write failed because of some error occurred
 *      - ESP_ERR_INVALID_ARG: Device address provided is out of range (hardware limits device address to 5 bits)
 */
// llgo:link (*Phy8023T).EspEthPhy8023SetMmdAddr C.esp_eth_phy_802_3_set_mmd_addr
func (recv_ *Phy8023T) EspEthPhy8023SetMmdAddr(devaddr c.Uint8T, mmd_addr c.Uint16T) EspErrT {
	return 0
}

/**
 * @brief Read data of MDIO device's memory at address register
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param devaddr Address of MDIO device
 * @param function MMD function
 * @param[out] data Data read from the device's memory
 * @return
 *      - ESP_OK: Memory read successfully
 *      - ESP_FAIL: Memory read failed because of some error occurred
 *      - ESP_ERR_INVALID_ARG: Device address provided is out of range (hardware limits device address to 5 bits) or MMD access function is invalid
 */
// llgo:link (*Phy8023T).EspEthPhy8023ReadMmdData C.esp_eth_phy_802_3_read_mmd_data
func (recv_ *Phy8023T) EspEthPhy8023ReadMmdData(devaddr c.Uint8T, function EspEthPhy8023MmdFuncT, data *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Write data to MDIO device's memory at address register
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param devaddr Address of MDIO device
 * @param function MMD function
 * @param[out] data Data to write to the device's memory
 * @return
 *      - ESP_OK: Memory written successfully
 *      - ESP_FAIL: Memory write failed because of some error occurred
 *      - ESP_ERR_INVALID_ARG: Device address provided is out of range (hardware limits device address to 5 bits) or MMD access function is invalid
 */
// llgo:link (*Phy8023T).EspEthPhy8023WriteMmdData C.esp_eth_phy_802_3_write_mmd_data
func (recv_ *Phy8023T) EspEthPhy8023WriteMmdData(devaddr c.Uint8T, function EspEthPhy8023MmdFuncT, data c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Set MMD address to mmd_addr with function MMD_FUNC_NOINCR and read contents to *data
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param devaddr Address of MDIO device
 * @param mmd_addr Address of MDIO device register
 * @param[out] data Data read from the device's memory
 * @return
 *      - ESP_OK: Memory read successfully
 *      - ESP_FAIL: Memory read failed because of some error occurred
 *      - ESP_ERR_INVALID_ARG: Device address provided is out of range (hardware limits device address to 5 bits)
 */
// llgo:link (*Phy8023T).EspEthPhy8023ReadMmdRegister C.esp_eth_phy_802_3_read_mmd_register
func (recv_ *Phy8023T) EspEthPhy8023ReadMmdRegister(devaddr c.Uint8T, mmd_addr c.Uint16T, data *c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Set MMD address to mmd_addr with function MMD_FUNC_NOINCR and write data
 *
 * @param phy_802_3 IEEE 802.3 PHY object infostructure
 * @param devaddr Address of MDIO device
 * @param mmd_addr Address of MDIO device register
 * @param[out] data Data to write to the device's memory
 * @return
 *      - ESP_OK: Memory written to successfully
 *      - ESP_FAIL: Memory write failed because of some error occurred
 *      - ESP_ERR_INVALID_ARG: Device address provided is out of range (hardware limits device address to 5 bits)
 */
// llgo:link (*Phy8023T).EspEthPhy8023WriteMmdRegister C.esp_eth_phy_802_3_write_mmd_register
func (recv_ *Phy8023T) EspEthPhy8023WriteMmdRegister(devaddr c.Uint8T, mmd_addr c.Uint16T, data c.Uint32T) EspErrT {
	return 0
}

/**
 * @brief Initializes configuration of parent IEEE 802.3 PHY object infostructure
 *
 * @param phy_802_3 Address to IEEE 802.3 PHY object infostructure
 * @param config Configuration of the IEEE 802.3 PHY object
 * @return
 *      - ESP_OK: configuration initialized successfully
 *      - ESP_ERR_INVALID_ARG: invalid @c config argument
 */
// llgo:link (*Phy8023T).EspEthPhy8023ObjConfigInit C.esp_eth_phy_802_3_obj_config_init
func (recv_ *Phy8023T) EspEthPhy8023ObjConfigInit(config *EthPhyConfigT) EspErrT {
	return 0
}
