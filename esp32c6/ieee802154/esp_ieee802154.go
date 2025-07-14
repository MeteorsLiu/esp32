package ieee802154

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief  Initialize the IEEE 802.15.4 subsystem.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 *
 */
//go:linkname EspIeee802154Enable C.esp_ieee802154_enable
func EspIeee802154Enable() EspErrT

/**
 * @brief  Deinitialize the IEEE 802.15.4 subsystem.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154Disable C.esp_ieee802154_disable
func EspIeee802154Disable() EspErrT

/**
 * @brief  Get the operational channel.
 *
 * @return The channel number (11~26).
 *
 */
//go:linkname EspIeee802154GetChannel C.esp_ieee802154_get_channel
func EspIeee802154GetChannel() c.Uint8T

/**
 * @brief  Set the operational channel.
 *
 * @param[in]  channel  The channel number (11-26).
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetChannel C.esp_ieee802154_set_channel
func EspIeee802154SetChannel(channel c.Uint8T) EspErrT

/**
 * @brief  Get the transmission power for the current channel.
 *
 * @return The transmission power in dBm.
 *
 */
//go:linkname EspIeee802154GetTxpower C.esp_ieee802154_get_txpower
func EspIeee802154GetTxpower() c.Int8T

/**
 * @brief  Set the transmission power for all channels.
 *
 * @param[in]  power  The transmission power in dBm.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetTxpower C.esp_ieee802154_set_txpower
func EspIeee802154SetTxpower(power c.Int8T) EspErrT

/**
 * @brief  Set the transmission power table.
 *
 * @param[in]  power_table  The power table.
 *
 * @return
 *        - ESP_OK   Set the transmission power table to successfully.
 */
// llgo:link EspIeee802154TxpowerTableT.EspIeee802154SetPowerTable C.esp_ieee802154_set_power_table
func (recv_ EspIeee802154TxpowerTableT) EspIeee802154SetPowerTable() EspErrT {
	return 0
}

/**
 * @brief  Get the transmission power table.
 *
 * @param[out]  out_power_table  The power table.
 *
 * @return
 *        - ESP_OK                  Get the transmission power table successfully.
 *        - ESP_ERR_INVALID_ARG     Invalid arguments.
 *
 */
// llgo:link (*EspIeee802154TxpowerTableT).EspIeee802154GetPowerTable C.esp_ieee802154_get_power_table
func (recv_ *EspIeee802154TxpowerTableT) EspIeee802154GetPowerTable() EspErrT {
	return 0
}

/**
 * @brief  Set the transmission power for a specific channel.
 *
 * @param[in]  channel  The channel.
 * @param[in]  power    The power.
 *
 * @return
 *        - ESP_OK                  Set the transmission power for a specific channel successfully.
 *        - ESP_ERR_INVALID_ARG     Invalid arguments.
 */
//go:linkname EspIeee802154SetPowerWithChannel C.esp_ieee802154_set_power_with_channel
func EspIeee802154SetPowerWithChannel(channel c.Uint8T, power c.Int8T) EspErrT

/**
 * @brief  Get the transmission power for a specific channel.
 *
 * @param[in]  channel    The channel.
 * @param[out] out_power  The power.
 *
 * @return
 *        - ESP_OK                  Get the transmission power for a specific channel successfully.
 *        - ESP_ERR_INVALID_ARG     Invalid arguments.
 */
//go:linkname EspIeee802154GetPowerWithChannel C.esp_ieee802154_get_power_with_channel
func EspIeee802154GetPowerWithChannel(channel c.Uint8T, out_power *c.Int8T) EspErrT

/**
 * @brief  Get the promiscuous mode.
 *
 * @return
 *      - True   The promiscuous mode is enabled.
 *      - False  The promiscuous mode is disabled.
 *
 */
//go:linkname EspIeee802154GetPromiscuous C.esp_ieee802154_get_promiscuous
func EspIeee802154GetPromiscuous() bool

/**
 * @brief  Set the promiscuous mode.
 *
 * @param[in]  enable  The promiscuous mode to be set.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetPromiscuous C.esp_ieee802154_set_promiscuous
func EspIeee802154SetPromiscuous(enable bool) EspErrT

/**
 * @brief  Get the IEEE 802.15.4 Radio state.
 *
 * @return  The IEEE 802.15.4 Radio state, refer to esp_ieee802154_state_t.
 *
 */
//go:linkname EspIeee802154GetState C.esp_ieee802154_get_state
func EspIeee802154GetState() EspIeee802154StateT

/**
 * @brief  Set the IEEE 802.15.4 Radio to sleep state.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure due to invalid state.
 *
 */
//go:linkname EspIeee802154Sleep C.esp_ieee802154_sleep
func EspIeee802154Sleep() EspErrT

/**
 * @brief  Set the IEEE 802.15.4 Radio to receive state.
 *
 * @note Radio will continue receiving until it receives a valid frame.
 *       Refer to `esp_ieee802154_receive_done()`.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_FAIL on failure due to invalid state.
 *
 */
//go:linkname EspIeee802154Receive C.esp_ieee802154_receive
func EspIeee802154Receive() EspErrT

/**
 * @brief  Transmit the given frame.
 *         The transmit result will be reported via `esp_ieee802154_transmit_done()`
 *         or `esp_ieee802154_transmit_failed()`.
 *
 * @param[in]  frame  The pointer to the frame, the frame format:
 *                    |-----------------------------------------------------------------------|
 *                    | Len | MHR |              MAC Payload                          |  FCS  |
 *                    |-----------------------------------------------------------------------|
 * @param[in]  cca    Perform CCA before transmission if it's true, otherwise transmit the frame directly.
 *
 * @note During transmission, the hardware calculates the FCS, and send it over the air right after the MAC payload,
 *       so you just need to prepare the length, mac header and mac payload content.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_ERR_INVALID_ARG on an invalid frame.
 *      - ESP_FAIL on failure due to invalid state.
 *
 */
//go:linkname EspIeee802154Transmit C.esp_ieee802154_transmit
func EspIeee802154Transmit(frame *c.Uint8T, cca bool) EspErrT

/**
 * @brief  Set the time to wait for the ack frame.
 *
 * @param[in]  timeout  The time to wait for the ack frame, in us.
 *                      It Should be a multiple of 16. The default value is 1728 us (108 * 16).
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetAckTimeout C.esp_ieee802154_set_ack_timeout
func EspIeee802154SetAckTimeout(timeout c.Uint32T) EspErrT

/**
 * @brief  Get the time to wait for the ack frame.
 *
 * @return  The time to wait for the ack frame, in us.
 */
//go:linkname EspIeee802154GetAckTimeout C.esp_ieee802154_get_ack_timeout
func EspIeee802154GetAckTimeout() c.Uint32T

/**
 * @brief  Get the device PAN ID.
 *
 * @return  The device PAN ID.
 *
 */
//go:linkname EspIeee802154GetPanid C.esp_ieee802154_get_panid
func EspIeee802154GetPanid() c.Uint16T

/**
 * @brief  Set the device PAN ID.
 *
 * @param[in]  panid  The device PAN ID.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetPanid C.esp_ieee802154_set_panid
func EspIeee802154SetPanid(panid c.Uint16T) EspErrT

/**
 * @brief  Get the device short address.
 *
 * @return  The device short address.
 *
 */
//go:linkname EspIeee802154GetShortAddress C.esp_ieee802154_get_short_address
func EspIeee802154GetShortAddress() c.Uint16T

/**
 * @brief  Set the device short address.
 *
 * @param[in]  short_address  The device short address.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetShortAddress C.esp_ieee802154_set_short_address
func EspIeee802154SetShortAddress(short_address c.Uint16T) EspErrT

/**
 * @brief  Get the device extended address.
 *
 * @param[out]  ext_addr  The pointer to the device extended address.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154GetExtendedAddress C.esp_ieee802154_get_extended_address
func EspIeee802154GetExtendedAddress(ext_addr *c.Uint8T) EspErrT

/**
 * @brief  Set the device extended address.
 *
 * @param[in]  ext_addr  The pointer to the device extended address.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetExtendedAddress C.esp_ieee802154_set_extended_address
func EspIeee802154SetExtendedAddress(ext_addr *c.Uint8T) EspErrT

/**
 * @brief  Get the device coordinator.
 *
 * @return
 *         - True   The coordinator is enabled.
 *         - False  The coordinator is disabled.
 *
 */
//go:linkname EspIeee802154GetCoordinator C.esp_ieee802154_get_coordinator
func EspIeee802154GetCoordinator() bool

/**
 * @brief  Set the device coordinator role.
 *
 * @param[in]  enable  The coordinator role to be set.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetCoordinator C.esp_ieee802154_set_coordinator
func EspIeee802154SetCoordinator(enable bool) EspErrT

/**
 * @brief  Get the auto frame pending mode.
 *
 * @return  The auto frame pending mode, refer to esp_ieee802154_pending_mode_t.
 *
 */
//go:linkname EspIeee802154GetPendingMode C.esp_ieee802154_get_pending_mode
func EspIeee802154GetPendingMode() EspIeee802154PendingModeT

/**
 * @brief  Set the auto frame pending mode.
 *
 * @param[in]  pending_mode  The auto frame pending mode, refer to esp_ieee802154_pending_mode_t.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
// llgo:link EspIeee802154PendingModeT.EspIeee802154SetPendingMode C.esp_ieee802154_set_pending_mode
func (recv_ EspIeee802154PendingModeT) EspIeee802154SetPendingMode() EspErrT {
	return 0
}

/**
 * @brief  Add address to the source matching table.
 *
 * @param[in]  addr      The pointer to the address.
 * @param[in]  is_short  Short address or Extended address.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_ERR_NO_MEM if the pending table is full.
 *
 */
//go:linkname EspIeee802154AddPendingAddr C.esp_ieee802154_add_pending_addr
func EspIeee802154AddPendingAddr(addr *c.Uint8T, is_short bool) EspErrT

/**
 * @brief  Remove address from the source matching table.
 *
 * @param[in]  addr      The pointer to the address.
 * @param[in]  is_short  Short address or Extended address.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_ERR_NOT_FOUND if the address was not found from the source matching table.
 *
 */
//go:linkname EspIeee802154ClearPendingAddr C.esp_ieee802154_clear_pending_addr
func EspIeee802154ClearPendingAddr(addr *c.Uint8T, is_short bool) EspErrT

/**
 * @brief  Clear the source matching table to empty.
 *
 * @param[in]  is_short  Clear Short address table or Extended address table.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154ResetPendingTable C.esp_ieee802154_reset_pending_table
func EspIeee802154ResetPendingTable(is_short bool) EspErrT

/**
 * @brief  Get the CCA threshold.
 *
 * @return  The CCA threshold in dBm.
 *
 */
//go:linkname EspIeee802154GetCcaThreshold C.esp_ieee802154_get_cca_threshold
func EspIeee802154GetCcaThreshold() c.Int8T

/**
 * @brief  Set the CCA threshold.
 *
 * @param[in]  cca_threshold  The CCA threshold in dBm.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetCcaThreshold C.esp_ieee802154_set_cca_threshold
func EspIeee802154SetCcaThreshold(cca_threshold c.Int8T) EspErrT

/**
 * @brief  Get the CCA mode.
 *
 * @return  The CCA mode, refer to esp_ieee802154_cca_mode_t.
 *
 */
//go:linkname EspIeee802154GetCcaMode C.esp_ieee802154_get_cca_mode
func EspIeee802154GetCcaMode() EspIeee802154CcaModeT

/**
 * @brief  Set the CCA mode.
 *
 * @param[in]  cca_mode  The CCA mode, refer to esp_ieee802154_cca_mode_t.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
// llgo:link EspIeee802154CcaModeT.EspIeee802154SetCcaMode C.esp_ieee802154_set_cca_mode
func (recv_ EspIeee802154CcaModeT) EspIeee802154SetCcaMode() EspErrT {
	return 0
}

/**
 * @brief  Enable rx_on_when_idle mode, radio will receive during idle.
 *
 * @param[in]  enable  Enable/Disable rx_on_when_idle mode.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetRxWhenIdle C.esp_ieee802154_set_rx_when_idle
func EspIeee802154SetRxWhenIdle(enable bool) EspErrT

/**
 * @brief  Get the rx_on_when_idle mode.
 *
 * @return  rx_on_when_idle mode.
 *
 */
//go:linkname EspIeee802154GetRxWhenIdle C.esp_ieee802154_get_rx_when_idle
func EspIeee802154GetRxWhenIdle() bool

/**
 * @brief  Perform energy detection.
 *
 * @param[in]  duration  The duration of energy detection, in symbol unit (16 us).
 *                       The result will be reported via esp_ieee802154_energy_detect_done().
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure due to invalid state.
 *
 */
//go:linkname EspIeee802154EnergyDetect C.esp_ieee802154_energy_detect
func EspIeee802154EnergyDetect(duration c.Uint32T) EspErrT

/**
 * @brief  Notify the IEEE 802.15.4 Radio that the frame is handled done by upper layer.
 *
 * @param[in]  frame  The pointer to the frame which was passed from the function `esp_ieee802154_receive_done()`
 *                    or ack frame from `esp_ieee802154_transmit_done()`.
 *
 * @return
 *      - ESP_OK on success
 *      - ESP_FAIL if frame is invalid.
 *
 */
//go:linkname EspIeee802154ReceiveHandleDone C.esp_ieee802154_receive_handle_done
func EspIeee802154ReceiveHandleDone(frame *c.Uint8T) EspErrT

/** Below are the events generated by IEEE 802.15.4 subsystem, which are in ISR context **/
/**
 * @brief  A Frame was received.
 *
 * @note   User must call the function `esp_ieee802154_receive_handle_done()` to notify 802.15.4 driver after the received frame is handled.
 *
 * @param[in]  frame  The point to the received frame, frame format:
 *                    |-----------------------------------------------------------------------|
 *                    | Len | MHR |              MAC Payload                       (no FCS)   |
 *                    |-----------------------------------------------------------------------|
 * @param[in]  frame_info  More information of the received frame, refer to esp_ieee802154_frame_info_t.
 *
 * @note  During receiving, the hardware calculates the FCS of the received frame, and may drop it if the FCS doesn't match, only the valid
 *        frames will be received and notified by esp_ieee802154_receive_done(). Please note that the FCS field is replaced by RSSI and LQI
 *        value of the received frame.
 *
 */
//go:linkname EspIeee802154ReceiveDone C.esp_ieee802154_receive_done
func EspIeee802154ReceiveDone(frame *c.Uint8T, frame_info *EspIeee802154FrameInfoT)

/**
 * @brief  The SFD field of the frame was received.
 *
 */
//go:linkname EspIeee802154ReceiveSfdDone C.esp_ieee802154_receive_sfd_done
func EspIeee802154ReceiveSfdDone()

/**
 * @brief  The Frame Transmission succeeded.
 *
 * @note   If the ack frame is not null, user must call the function `esp_ieee802154_receive_handle_done()` to notify 802.15.4 driver
 *         after the ack frame is handled.
 *
 * @param[in]  frame           The pointer to the transmitted frame.
 * @param[in]  ack             The received ACK frame, it could be NULL if the transmitted frame's AR bit is not set.
 * @param[in]  ack_frame_info  More information of the ACK frame, refer to esp_ieee802154_frame_info_t.
 *
 */
//go:linkname EspIeee802154TransmitDone C.esp_ieee802154_transmit_done
func EspIeee802154TransmitDone(frame *c.Uint8T, ack *c.Uint8T, ack_frame_info *EspIeee802154FrameInfoT)

/**
 * @brief  The Frame Transmission failed. Refer to `esp_ieee802154_transmit()`.
 *
 * @param[in]  frame  The pointer to the frame.
 * @param[in]  error  The transmission failure reason, refer to esp_ieee802154_tx_error_t.
 *
 */
//go:linkname EspIeee802154TransmitFailed C.esp_ieee802154_transmit_failed
func EspIeee802154TransmitFailed(frame *c.Uint8T, error EspIeee802154TxErrorT)

/**
 * @brief  The SFD field of the frame was transmitted.
 *
 */
//go:linkname EspIeee802154TransmitSfdDone C.esp_ieee802154_transmit_sfd_done
func EspIeee802154TransmitSfdDone(frame *c.Uint8T)

/**
 * @brief  The energy detection done. Refer to `esp_ieee802154_energy_detect()`.
 *
 * @param[in]  power  The detected power level, in dBm.
 *
 */
//go:linkname EspIeee802154EnergyDetectDone C.esp_ieee802154_energy_detect_done
func EspIeee802154EnergyDetectDone(power c.Int8T)

/**
 * @brief  The receive window for receive_at has finished.
 *
 */
//go:linkname EspIeee802154ReceiveAtDone C.esp_ieee802154_receive_at_done
func EspIeee802154ReceiveAtDone()

/**
 * @brief  Set the IEEE 802.15.4 Radio to receive state at a specific time, for a specific duration.
 *
 * @note   Radio will start receiving after the timestamp, and continue receiving for the specific duration.
 *
 * @param[in]  time      A specific timestamp for starting receiving.
 * @param[in]  duration  A specific duration after which to stop receiving. Set duration = 0 to rx indefinitely.
 * @return
 *      - ESP_OK on success
 *      - ESP_FAIL on failure due to invalid state.
 *
 */
//go:linkname EspIeee802154ReceiveAt C.esp_ieee802154_receive_at
func EspIeee802154ReceiveAt(time c.Uint32T, duration c.Uint32T) EspErrT

/**
 * @brief  Transmit the given frame at a specific time.
 *         The transmit result will be reported via `esp_ieee802154_transmit_done()`
 *         or `esp_ieee802154_transmit_failed()`.
 *
 * @param[in]  frame  The pointer to the frame. Refer to `esp_ieee802154_transmit()`.
 * @param[in]  cca    Perform CCA before transmission if it's true, otherwise transmit the frame directly.
 * @param[in]  time  A specific timestamp for starting transmission.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_ERR_INVALID_ARG on an invalid frame.
 *      - ESP_FAIL on failure due to invalid state.
 *
 */
//go:linkname EspIeee802154TransmitAt C.esp_ieee802154_transmit_at
func EspIeee802154TransmitAt(frame *c.Uint8T, cca bool, time c.Uint32T) EspErrT

/**
 * @brief  Get the RSSI of the most recent received frame.
 *
 * @return The value of RSSI.
 *
 */
//go:linkname EspIeee802154GetRecentRssi C.esp_ieee802154_get_recent_rssi
func EspIeee802154GetRecentRssi() c.Int8T

/**
 * @brief  Get the LQI of the most recent received frame.
 *
 * @return The value of LQI.
 *
 */
//go:linkname EspIeee802154GetRecentLqi C.esp_ieee802154_get_recent_lqi
func EspIeee802154GetRecentLqi() c.Uint8T

/**
 * @brief  Set the key and addr for a frame needs to be encrypted by HW.
 *
 * @param[in]  frame  A frame needs to be encrypted. Refer to `esp_ieee802154_transmit()`.
 * @param[in]  key    A 16-bytes key for encryption.
 * @param[in]  addr   An 8-bytes addr for HW to generate nonce, in general, is the device extended address.
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154SetTransmitSecurity C.esp_ieee802154_set_transmit_security
func EspIeee802154SetTransmitSecurity(frame *c.Uint8T, key *c.Uint8T, addr *c.Uint8T) EspErrT

/**
 * @brief  This function will be called when a received frame needs to be acked with Enh-Ack, the upper
 *         layer should generate the Enh-Ack frame in this callback function.
 *
 * @param[in]  frame          The received frame.
 * @param[in]  frame_info     The frame information. Refer to `esp_ieee802154_frame_info_t`.
 * @param[out] enhack_frame   The Enh-ack frame need to be generated via this function, HW will send it back after AIFS.
 *
 * @return
 *        - ESP_OK if Enh-Ack generates done.
 *        - ESP_FAIL if Enh-Ack generates failed.
 *
 */
//go:linkname EspIeee802154EnhAckGenerator C.esp_ieee802154_enh_ack_generator
func EspIeee802154EnhAckGenerator(frame *c.Uint8T, frame_info *EspIeee802154FrameInfoT, enhack_frame *c.Uint8T) EspErrT

/**
 * @brief  Set the IEEE802.15.4 coexist config.
 *
 * @param[in]  config     The config of IEEE802.15.4 coexist.
 *
 */
//go:linkname EspIeee802154SetCoexConfig C.esp_ieee802154_set_coex_config
func EspIeee802154SetCoexConfig(config EspIeee802154CoexConfigT)

/**
 * @brief  Get the IEEE802.15.4 coexist config.
 *
 * @return
 *        - The config of IEEE802.15.4 coexist.
 *
 */
//go:linkname EspIeee802154GetCoexConfig C.esp_ieee802154_get_coex_config
func EspIeee802154GetCoexConfig() EspIeee802154CoexConfigT

/**
 * @brief  Register process callbacks for events generated by the IEEE 802.15.4 subsystem.
 *
 * @param[in]  cb_list The event process callback list, please refer to `esp_ieee802154_event_cb_list_t`.
 *
 * @note  This API should be called only when IEEE 802.15.4 subsystem is not enabled
 *        or after IEEE 802.15.4 subsystem is disabled (refer to `esp_ieee802154_disable`).
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
// llgo:link EspIeee802154EventCbListT.EspIeee802154EventCallbackListRegister C.esp_ieee802154_event_callback_list_register
func (recv_ EspIeee802154EventCbListT) EspIeee802154EventCallbackListRegister() EspErrT {
	return 0
}

/**
 * @brief  Unregister process callbacks for events generated by the IEEE 802.15.4 subsystem.
 *
 * @note  This API should be called only when IEEE 802.15.4 subsystem is not enabled
 *        or after IEEE 802.15.4 subsystem is disabled (refer to `esp_ieee802154_disable`).
 *
 * @return
 *      - ESP_OK on success.
 *      - ESP_FAIL on failure.
 */
//go:linkname EspIeee802154EventCallbackListUnregister C.esp_ieee802154_event_callback_list_unregister
func EspIeee802154EventCallbackListUnregister() EspErrT
