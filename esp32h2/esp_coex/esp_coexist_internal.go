package esp_coex

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type CoexPreferT c.Int

const (
	COEX_PREFER_WIFI    CoexPreferT = 0
	COEX_PREFER_BT      CoexPreferT = 1
	COEX_PREFER_BALANCE CoexPreferT = 2
	COEX_PREFER_NUM     CoexPreferT = 3
)

type CoexSchmCallbackTypeT c.Int

const (
	COEX_SCHM_CALLBACK_TYPE_WIFI CoexSchmCallbackTypeT = 0
	COEX_SCHM_CALLBACK_TYPE_BT   CoexSchmCallbackTypeT = 1
	COEX_SCHM_CALLBACK_TYPE_I154 CoexSchmCallbackTypeT = 2
)

type CoexSchmStTypeT c.Int

const (
	COEX_SCHM_ST_TYPE_WIFI CoexSchmStTypeT = 0
	COEX_SCHM_ST_TYPE_BLE  CoexSchmStTypeT = 1
	COEX_SCHM_ST_TYPE_BT   CoexSchmStTypeT = 2
)

// llgo:type C
type CoexFuncCbT func(c.Uint32T, c.Int)

// llgo:type C
type CoexSetLpclkSourceCallbackT func() EspErrT

// llgo:type C
type CoexWifiChannelChangeCbT func(c.Uint8T, c.Uint8T)

/**
 * @brief Pre-Init software coexist
 *        extern function for internal use.
 *
 * @return Init ok or failed.
 */
//go:linkname CoexPreInit C.coex_pre_init
func CoexPreInit() EspErrT

/**
 * @brief Init software coexist
 *        extern function for internal use.
 *
 * @return Init ok or failed.
 */
//go:linkname CoexInit C.coex_init
func CoexInit() EspErrT

/**
 * @brief De-init software coexist
 *        extern function for internal use.
 */
//go:linkname CoexDeinit C.coex_deinit
func CoexDeinit()

/**
 * @brief Enable software coexist
 *        extern function for internal use.
 *
 * @return Enable ok or failed.
 */
//go:linkname CoexEnable C.coex_enable
func CoexEnable() EspErrT

/**
 * @brief Disable software coexist
 *        extern function for internal use.
 */
//go:linkname CoexDisable C.coex_disable
func CoexDisable()

/**
 * @brief Get software coexist version string
 *        extern function for internal use.
 * @return : version string
 */
//go:linkname CoexVersionGet C.coex_version_get
func CoexVersionGet() *c.Char

/**
 * @brief Get software coexist version value
 *        extern function for internal use.
 * @param ptr_version : points to version structure
 * @return : ESP_OK - success, other - failed
 */
// llgo:link (*CoexVersionT).CoexVersionGetValue C.coex_version_get_value
func (recv_ *CoexVersionT) CoexVersionGetValue() EspErrT {
	return 0
}

/**
 * @brief Coexist performance preference set from libbt.a
 *        extern function for internal use.
 *
 *  @param prefer : the prefer enumeration value
 *  @return : ESP_OK - success, other - failed
 */
// llgo:link CoexPreferT.CoexPreferenceSet C.coex_preference_set
func (recv_ CoexPreferT) CoexPreferenceSet() EspErrT {
	return 0
}

/**
 * @brief Get software coexist status.
 *
 * @param bitmap : bitmap of the module getting status.
 * @return : software coexist status
 */
//go:linkname CoexStatusGet C.coex_status_get
func CoexStatusGet(bitmap c.Uint8T) c.Uint32T

/**
 * @brief WiFi requests coexistence.
 *
 *  @param event : WiFi event
 *  @param latency : WiFi will request coexistence after latency
 *  @param duration : duration for WiFi to request coexistence
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexWifiRequest C.coex_wifi_request
func CoexWifiRequest(event c.Uint32T, latency c.Uint32T, duration c.Uint32T) c.Int

/**
 * @brief WiFi release coexistence.
 *
 *  @param event : WiFi event
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexWifiRelease C.coex_wifi_release
func CoexWifiRelease(event c.Uint32T) c.Int

/**
 * @brief Set WiFi channel to coexistence module.
 *
 *  @param primary : WiFi primary channel
 *  @param secondary : WiFi secondary channel
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexWifiChannelSet C.coex_wifi_channel_set
func CoexWifiChannelSet(primary c.Uint8T, secondary c.Uint8T) c.Int

/**
 * @brief Get WiFi channel from coexistence module.
 *
 *  @param primary : pointer to value of WiFi primary channel
 *  @param secondary : pointer to value of WiFi secondary channel
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexWifiChannelGet C.coex_wifi_channel_get
func CoexWifiChannelGet(primary *c.Uint8T, secondary *c.Uint8T) c.Int

/**
 * @brief Register application callback function to Wi-Fi update low power clock module.
 *
 * @param callback : Wi-Fi update low power clock callback function
 */
//go:linkname CoexWifiRegisterUpdateLpclkCallback C.coex_wifi_register_update_lpclk_callback
func CoexWifiRegisterUpdateLpclkCallback(callback CoexSetLpclkSourceCallbackT)

/**
 * @brief Bluetooth requests coexistence
 *
 *  @param event : Bluetooth event
 *  @param latency : Bluetooth will request coexistence after latency
 *  @param duration : duration for Bluetooth to request coexistence
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexBtRequest C.coex_bt_request
func CoexBtRequest(event c.Uint32T, latency c.Uint32T, duration c.Uint32T) c.Int

/**
 * @brief Bluetooth release coexistence.
 *
 *  @param event : Bluetooth event
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexBtRelease C.coex_bt_release
func CoexBtRelease(event c.Uint32T) c.Int

/**
 * @brief Bluetooth registers callback function to receive notification when Wi-Fi channel changes
 *
 *  @param callback: callback function registered to coexistence module
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexRegisterWifiChannelChangeCallback C.coex_register_wifi_channel_change_callback
func CoexRegisterWifiChannelChangeCallback(callback CoexWifiChannelChangeCbT) c.Int

/**
 * @brief Update low power clock interval
 */
//go:linkname CoexUpdateLpclkInterval C.coex_update_lpclk_interval
func CoexUpdateLpclkInterval()

/**
 * @brief Get coexistence event duration.
 *
 *  @param event : Coexistence event
 *  @param duration: Coexistence event duration
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexEventDurationGet C.coex_event_duration_get
func CoexEventDurationGet(event c.Uint32T, duration *c.Uint32T) c.Int

/**
 * @brief Get coexistence event priority.
 *
 *  @param event : Coexistence event
 *  @param pti: Coexistence event priority
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexPtiGet C.coex_pti_get
func CoexPtiGet(event c.Uint32T, pti *c.Uint8T) c.Int

/**
 * @brief Clear coexistence status.
 *
 *  @param type : Coexistence status type
 *  @param status: Coexistence status
 */
//go:linkname CoexSchmStatusBitClear C.coex_schm_status_bit_clear
func CoexSchmStatusBitClear(type_ c.Uint32T, status c.Uint32T)

/**
 * @brief Set coexistence status.
 *
 *  @param type : Coexistence status type
 *  @param status: Coexistence status
 */
//go:linkname CoexSchmStatusBitSet C.coex_schm_status_bit_set
func CoexSchmStatusBitSet(type_ c.Uint32T, status c.Uint32T)

/**
 * @brief Set coexistence scheme interval.
 *
 *  @param interval : Coexistence scheme interval
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexSchmIntervalSet C.coex_schm_interval_set
func CoexSchmIntervalSet(interval c.Uint32T) c.Int

/**
 * @brief Get coexistence scheme interval.
 *
 *  @return : Coexistence scheme interval
 */
//go:linkname CoexSchmIntervalGet C.coex_schm_interval_get
func CoexSchmIntervalGet() c.Uint32T

/**
 * @brief Get current coexistence scheme period.
 *
 *  @return : Coexistence scheme period
 */
//go:linkname CoexSchmCurrPeriodGet C.coex_schm_curr_period_get
func CoexSchmCurrPeriodGet() c.Uint8T

/**
 * @brief Get current coexistence scheme phase.
 *
 *  @return : Coexistence scheme phase
 */
//go:linkname CoexSchmCurrPhaseGet C.coex_schm_curr_phase_get
func CoexSchmCurrPhaseGet() c.Pointer

/**
 * @brief Set current coexistence scheme phase index.
 *
 *  @param idx : Coexistence scheme phase index
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexSchmCurrPhaseIdxSet C.coex_schm_curr_phase_idx_set
func CoexSchmCurrPhaseIdxSet(idx c.Int) c.Int

/**
 * @brief Get current coexistence scheme phase index.
 *
 *  @return : Coexistence scheme phase index
 */
//go:linkname CoexSchmCurrPhaseIdxGet C.coex_schm_curr_phase_idx_get
func CoexSchmCurrPhaseIdxGet() c.Int

/**
 * @brief Register WiFi callback for coexistence starts.
 *
 *  @param cb : WiFi callback
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexRegisterStartCb C.coex_register_start_cb
func CoexRegisterStartCb(cb func() c.Int) c.Int

/**
 * @brief Restart current coexistence scheme.
 *
 *  @return : 0 - success, other - failed
 */
//go:linkname CoexSchmProcessRestart C.coex_schm_process_restart
func CoexSchmProcessRestart() c.Int

/**
 * @brief Register callback for coexistence scheme.
 *
 *  @param type : callback type
 *  @param callback : callback
 *  @return : 0 - success, other - failed
 */
// llgo:link CoexSchmCallbackTypeT.CoexSchmRegisterCallback C.coex_schm_register_callback
func (recv_ CoexSchmCallbackTypeT) CoexSchmRegisterCallback(callback c.Pointer) c.Int {
	return 0
}

/**
 * @brief Register coexistence adapter functions.
 *
 *  @param funcs : coexistence adapter functions
 *  @return : ESP_OK - success, other - failed
 */
// llgo:link (*CoexAdapterFuncsT).EspCoexAdapterRegister C.esp_coex_adapter_register
func (recv_ *CoexAdapterFuncsT) EspCoexAdapterRegister() EspErrT {
	return 0
}

/**
 * @brief     Get coexistence scheme phase by phase index.
 *
 * @param     phase_idx    Coexistence phase index
 *
 * @return    Coexistence scheme phase
 */
//go:linkname CoexSchmGetPhaseByIdx C.coex_schm_get_phase_by_idx
func CoexSchmGetPhaseByIdx(phase_idx c.Int) c.Pointer

/**
 * @brief     Check the MD5 values of the coexistence adapter header files in IDF and WiFi library
 *
 * @attention 1. It is used for internal CI version check
 *
 * @return
 *     - ESP_OK : succeed
 *     - ESP_WIFI_INVALID_ARG : MD5 check fail
 */
//go:linkname EspCoexAdapterFuncsMd5Check C.esp_coex_adapter_funcs_md5_check
func EspCoexAdapterFuncsMd5Check(md5 *c.Char) EspErrT
