package esp_driver_i2s

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type LpVadT c.Uint32T

type VadUnitCtxT struct {
	Unused [8]uint8
}
type VadUnitHandleT *VadUnitCtxT

/**
 * @brief LP VAD configurations
 */

type LpVadConfigT struct {
	InitFrameNum           c.Int
	MinEnergyThresh        c.Int
	SkipBandEnergyThresh   bool
	SpeakActivityThresh    c.Int
	NonSpeakActivityThresh c.Int
	MinSpeakActivityThresh c.Int
	MaxSpeakActivityThresh c.Int
}

/**
 * @brief LP VAD Init Configurations
 */

type LpVadInitConfigT struct {
	LpI2sChan LpI2sChanHandleT
	VadConfig LpVadConfigT
}

/**
 * @brief New LP VAD unit
 * @param[in]  vad_id        VAD id
 * @param[in]  init_config   Initial configurations
 * @param[out] ret_unit      Unit handle
 *
 * @return
 *        - ESP_OK:                 On success
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Driver state is invalid, you shouldn't call this API at this moment
 */
// llgo:link LpVadT.LpI2sVadNewUnit C.lp_i2s_vad_new_unit
func (recv_ LpVadT) LpI2sVadNewUnit(init_config *LpVadInitConfigT, ret_unit *VadUnitHandleT) EspErrT {
	return 0
}

/**
 * @brief Enable LP VAD
 *
 * @param[in] unit          VAD handle
 *
 * @return
 *        - ESP_OK:                 On success
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname LpI2sVadEnable C.lp_i2s_vad_enable
func LpI2sVadEnable(unit VadUnitHandleT) EspErrT

/**
 * @brief Disable LP VAD
 *
 * @param[in] unit          VAD handle
 *
 * @return
 *        - ESP_OK:                 On success
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname LpI2sVadDisable C.lp_i2s_vad_disable
func LpI2sVadDisable(unit VadUnitHandleT) EspErrT

/**
 * @brief Delete LP VAD unit
 * @param[in] unit          VAD handle
 *
 * @return
 *        - ESP_OK:                 On success
 *        - ESP_ERR_INVALID_ARG:    Invalid argument
 *        - ESP_ERR_INVALID_STATE:  Driver state is invalid, you shouldn't call this API at this moment
 */
//go:linkname LpI2sVadDelUnit C.lp_i2s_vad_del_unit
func LpI2sVadDelUnit(unit VadUnitHandleT) EspErrT
