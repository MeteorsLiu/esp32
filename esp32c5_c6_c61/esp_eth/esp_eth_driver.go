package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EspEthHandleT c.Pointer

/**
* @brief Configuration of Ethernet driver
*
 */
type EspEthConfigT struct {
	Mac                  *EspEthMacT
	Phy                  *EspEthPhyT
	CheckLinkPeriodMs    c.Uint32T
	StackInput           c.Pointer
	OnLowlevelInitDone   c.Pointer
	OnLowlevelDeinitDone c.Pointer
	ReadPhyReg           c.Pointer
	WritePhyReg          c.Pointer
}

/**
 * @brief Data structure to Read/Write PHY register via ioctl API
 *
 */

type EspEthPhyRegRwDataT struct {
	RegAddr   c.Uint32T
	RegValueP *c.Uint32T
}
type EspEthIoCmdT c.Int

const (
	ETH_CMD_G_MAC_ADDR      EspEthIoCmdT = 0
	ETH_CMD_S_MAC_ADDR      EspEthIoCmdT = 1
	ETH_CMD_G_PHY_ADDR      EspEthIoCmdT = 2
	ETH_CMD_S_PHY_ADDR      EspEthIoCmdT = 3
	ETH_CMD_G_AUTONEGO      EspEthIoCmdT = 4
	ETH_CMD_S_AUTONEGO      EspEthIoCmdT = 5
	ETH_CMD_G_SPEED         EspEthIoCmdT = 6
	ETH_CMD_S_SPEED         EspEthIoCmdT = 7
	ETH_CMD_S_PROMISCUOUS   EspEthIoCmdT = 8
	ETH_CMD_S_FLOW_CTRL     EspEthIoCmdT = 9
	ETH_CMD_G_DUPLEX_MODE   EspEthIoCmdT = 10
	ETH_CMD_S_DUPLEX_MODE   EspEthIoCmdT = 11
	ETH_CMD_S_PHY_LOOPBACK  EspEthIoCmdT = 12
	ETH_CMD_READ_PHY_REG    EspEthIoCmdT = 13
	ETH_CMD_WRITE_PHY_REG   EspEthIoCmdT = 14
	ETH_CMD_CUSTOM_MAC_CMDS EspEthIoCmdT = 4095
	ETH_CMD_CUSTOM_PHY_CMDS EspEthIoCmdT = 8191
)

/**
* @brief Install Ethernet driver
*
* @param[in]  config: configuration of the Ethernet driver
* @param[out] out_hdl: handle of Ethernet driver
*
* @return
*       - ESP_OK: install esp_eth driver successfully
*       - ESP_ERR_INVALID_ARG: install esp_eth driver failed because of some invalid argument
*       - ESP_ERR_NO_MEM: install esp_eth driver failed because there's no memory for driver
*       - ESP_FAIL: install esp_eth driver failed because some other error occurred
 */
// llgo:link (*EspEthConfigT).EspEthDriverInstall C.esp_eth_driver_install
func (recv_ *EspEthConfigT) EspEthDriverInstall(out_hdl *EspEthHandleT) EspErrT {
	return 0
}

/**
* @brief Uninstall Ethernet driver
* @note It's not recommended to uninstall Ethernet driver unless it won't get used any more in application code.
*       To uninstall Ethernet driver, you have to make sure, all references to the driver are released.
*       Ethernet driver can only be uninstalled successfully when reference counter equals to one.
*
* @param[in] hdl: handle of Ethernet driver
*
* @return
*       - ESP_OK: uninstall esp_eth driver successfully
*       - ESP_ERR_INVALID_ARG: uninstall esp_eth driver failed because of some invalid argument
*       - ESP_ERR_INVALID_STATE: uninstall esp_eth driver failed because it has more than one reference
*       - ESP_FAIL: uninstall esp_eth driver failed because some other error occurred
 */
//go:linkname EspEthDriverUninstall C.esp_eth_driver_uninstall
func EspEthDriverUninstall(hdl EspEthHandleT) EspErrT

/**
* @brief Start Ethernet driver **ONLY** in standalone mode (i.e. without TCP/IP stack)
*
* @note This API will start driver state machine and internal software timer (for checking link status).
*
* @param[in] hdl handle of Ethernet driver
*
* @return
*       - ESP_OK: start esp_eth driver successfully
*       - ESP_ERR_INVALID_ARG: start esp_eth driver failed because of some invalid argument
*       - ESP_ERR_INVALID_STATE: start esp_eth driver failed because driver has started already
*       - ESP_FAIL: start esp_eth driver failed because some other error occurred
 */
//go:linkname EspEthStart C.esp_eth_start
func EspEthStart(hdl EspEthHandleT) EspErrT

/**
* @brief Stop Ethernet driver
*
* @note This function does the oppsite operation of `esp_eth_start`.
*
* @param[in] hdl handle of Ethernet driver
* @return
*       - ESP_OK: stop esp_eth driver successfully
*       - ESP_ERR_INVALID_ARG: stop esp_eth driver failed because of some invalid argument
*       - ESP_ERR_INVALID_STATE: stop esp_eth driver failed because driver has not started yet
*       - ESP_FAIL: stop esp_eth driver failed because some other error occurred
 */
//go:linkname EspEthStop C.esp_eth_stop
func EspEthStop(hdl EspEthHandleT) EspErrT

/**
* @brief Update Ethernet data input path (i.e. specify where to pass the input buffer)
*
* @note After install driver, Ethernet still don't know where to deliver the input buffer.
*       In fact, this API registers a callback function which get invoked when Ethernet received new packets.
*
* @param[in] hdl handle of Ethernet driver
* @param[in] stack_input function pointer, which does the actual process on incoming packets
* @param[in] priv private resource, which gets passed to `stack_input` callback without any modification
* @return
*       - ESP_OK: update input path successfully
*       - ESP_ERR_INVALID_ARG: update input path failed because of some invalid argument
*       - ESP_FAIL: update input path failed because some other error occurred
 */
//go:linkname EspEthUpdateInputPath C.esp_eth_update_input_path
func EspEthUpdateInputPath(hdl EspEthHandleT, stack_input func(EspEthHandleT, *c.Uint8T, c.Uint32T, c.Pointer) EspErrT, priv c.Pointer) EspErrT

/**
* @brief General Transmit
*
* @param[in] hdl: handle of Ethernet driver
* @param[in] buf: buffer of the packet to transfer
* @param[in] length: length of the buffer to transfer
*
* @return
*       - ESP_OK: transmit frame buffer successfully
*       - ESP_ERR_INVALID_ARG: transmit frame buffer failed because of some invalid argument
*       - ESP_ERR_INVALID_STATE: invalid driver state (e.i. driver is not started)
*       - ESP_ERR_TIMEOUT: transmit frame buffer failed because HW was not get available in predefined period
*       - ESP_FAIL: transmit frame buffer failed because some other error occurred
 */
//go:linkname EspEthTransmit C.esp_eth_transmit
func EspEthTransmit(hdl EspEthHandleT, buf c.Pointer, length c.SizeT) EspErrT

/**
* @brief Special Transmit with variable number of arguments
*
* @param[in] hdl handle of Ethernet driver
* @param[in] argc number variable arguments
* @param ... variable arguments
* @return
*       - ESP_OK: transmit successful
*       - ESP_ERR_INVALID_STATE: invalid driver state (e.i. driver is not started)
*       - ESP_ERR_TIMEOUT: transmit frame buffer failed because HW was not get available in predefined period
*       - ESP_FAIL: transmit frame buffer failed because some other error occurred
 */
//go:linkname EspEthTransmitVargs C.esp_eth_transmit_vargs
func EspEthTransmitVargs(hdl EspEthHandleT, argc c.Uint32T, __llgo_va_list ...interface{}) EspErrT

/**
* @brief Misc IO function of Ethernet driver
*
* @param[in] hdl: handle of Ethernet driver
* @param[in] cmd: IO control command
* @param[in, out] data: address of data for `set` command or address where to store the data when used with `get` command
*
* @return
*       - ESP_OK: process io command successfully
*       - ESP_ERR_INVALID_ARG: process io command failed because of some invalid argument
*       - ESP_FAIL: process io command failed because some other error occurred
*       - ESP_ERR_NOT_SUPPORTED: requested feature is not supported
*
* The following common IO control commands are supported:
* @li @c ETH_CMD_S_MAC_ADDR sets Ethernet interface MAC address. @c data argument is pointer to MAC address buffer with expected size of 6 bytes.
* @li @c ETH_CMD_G_MAC_ADDR gets Ethernet interface MAC address. @c data argument is pointer to a buffer to which MAC address is to be copied. The buffer size must be at least 6 bytes.
* @li @c ETH_CMD_S_PHY_ADDR sets PHY address in range of <0-31>. @c data argument is pointer to memory of uint32_t datatype from where the configuration option is read.
* @li @c ETH_CMD_G_PHY_ADDR gets PHY address. @c data argument is pointer to memory of uint32_t datatype to which the PHY address is to be stored.
* @li @c ETH_CMD_S_AUTONEGO enables or disables Ethernet link speed and duplex mode autonegotiation. @c data argument is pointer to memory of bool datatype from which the configuration option is read.
*                           Preconditions: Ethernet driver needs to be stopped.
* @li @c ETH_CMD_G_AUTONEGO gets current configuration of the Ethernet link speed and duplex mode autonegotiation. @c data argument is pointer to memory of bool datatype to which the current configuration is to be stored.
* @li @c ETH_CMD_S_SPEED sets the Ethernet link speed. @c data argument is pointer to memory of eth_speed_t datatype from which the configuration option is read.
*                           Preconditions: Ethernet driver needs to be stopped and auto-negotiation disabled.
* @li @c ETH_CMD_G_SPEED gets current Ethernet link speed. @c data argument is pointer to memory of eth_speed_t datatype to which the speed is to be stored.
* @li @c ETH_CMD_S_PROMISCUOUS sets/resets Ethernet interface promiscuous mode. @c data argument is pointer to memory of bool datatype from which the configuration option is read.
* @li @c ETH_CMD_S_FLOW_CTRL sets/resets Ethernet interface flow control. @c data argument is pointer to memory of bool datatype from which the configuration option is read.
* @li @c ETH_CMD_S_DUPLEX_MODE sets the Ethernet duplex mode. @c data argument is pointer to memory of eth_duplex_t datatype from which the configuration option is read.
*                            Preconditions: Ethernet driver needs to be stopped and auto-negotiation disabled.
* @li @c ETH_CMD_G_DUPLEX_MODE gets current Ethernet link duplex mode.  @c data argument is pointer to memory of eth_duplex_t datatype to which the duplex mode is to be stored.
* @li @c ETH_CMD_S_PHY_LOOPBACK sets/resets PHY to/from loopback mode. @c data argument is pointer to memory of bool datatype from which the configuration option is read.
*
* @li Note that additional control commands may be available for specific MAC or PHY chips. Please consult specific MAC or PHY documentation or driver code.
 */
//go:linkname EspEthIoctl C.esp_eth_ioctl
func EspEthIoctl(hdl EspEthHandleT, cmd EspEthIoCmdT, data c.Pointer) EspErrT

/**
* @brief Get PHY instance memory address
*
* @param[in] hdl handle of Ethernet driver
* @param[out] phy pointer to memory to store the instance
* @return esp_err_t
*       - ESP_OK: success
*       - ESP_ERR_INVALID_ARG: failed because of some invalid argument
 */
//go:linkname EspEthGetPhyInstance C.esp_eth_get_phy_instance
func EspEthGetPhyInstance(hdl EspEthHandleT, phy **EspEthPhyT) EspErrT

/**
* @brief Get MAC instance memory address
*
* @param[in] hdl handle of Ethernet driver
* @param[out] mac pointer to memory to store the instance
* @return esp_err_t
*       - ESP_OK: success
*       - ESP_ERR_INVALID_ARG: failed because of some invalid argument
 */
//go:linkname EspEthGetMacInstance C.esp_eth_get_mac_instance
func EspEthGetMacInstance(hdl EspEthHandleT, mac **EspEthMacT) EspErrT

/**
* @brief Increase Ethernet driver reference
* @note Ethernet driver handle can be obtained by os timer, netif, etc.
*       It's dangerous when thread A is using Ethernet but thread B uninstall the driver.
*       Using reference counter can prevent such risk, but care should be taken, when you obtain Ethernet driver,
*       this API must be invoked so that the driver won't be uninstalled during your using time.
*
*
* @param[in] hdl: handle of Ethernet driver
* @return
*       - ESP_OK: increase reference successfully
*       - ESP_ERR_INVALID_ARG: increase reference failed because of some invalid argument
 */
//go:linkname EspEthIncreaseReference C.esp_eth_increase_reference
func EspEthIncreaseReference(hdl EspEthHandleT) EspErrT

/**
* @brief Decrease Ethernet driver reference
*
* @param[in] hdl: handle of Ethernet driver
* @return
*       - ESP_OK: increase reference successfully
*       - ESP_ERR_INVALID_ARG: increase reference failed because of some invalid argument
 */
//go:linkname EspEthDecreaseReference C.esp_eth_decrease_reference
func EspEthDecreaseReference(hdl EspEthHandleT) EspErrT
