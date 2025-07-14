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
 * @brief Enable software coexist
 *        extern function for internal use.
 *
 * @return Enable ok or failed.
 */
//go:linkname CoexEnable C.coex_enable
func CoexEnable() EspErrT

/**
 * @brief Get software coexist version string
 *        extern function for internal use.
 * @return : version string
 */
//go:linkname CoexVersionGet C.coex_version_get
func CoexVersionGet() *c.Char

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
 * @brief Set coexistence status.
 *
 *  @param type : Coexistence status type
 *  @param status: Coexistence status
 */
//go:linkname CoexSchmStatusBitSet C.coex_schm_status_bit_set
func CoexSchmStatusBitSet(type_ c.Uint32T, status c.Uint32T)
