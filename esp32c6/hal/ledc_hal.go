package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Context that should be maintained by both the driver and the HAL
 */

type LedcHalContextT struct {
	Dev       *LedcDevT
	SpeedMode LedcModeT
}

/**
 * @brief Init the LEDC hal. This function should be called first before other hal layer function is called
 *
 * @param hal Context of the HAL layer
 * @param speed_mode speed_mode Select the LEDC speed_mode, high-speed mode and low-speed mod
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalInit C.ledc_hal_init
func (recv_ *LedcHalContextT) LedcHalInit(speed_mode LedcModeT) {
}

/**
 * @brief Update channel configure when select low speed mode
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index (0-7), select from ledc_channel_t
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalLsChannelUpdate C.ledc_hal_ls_channel_update
func (recv_ *LedcHalContextT) LedcHalLsChannelUpdate(channel_num LedcChannelT) {
}

/**
 * @brief Set the duty start
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index (0-7), select from ledc_channel_t
 * @param duty_start The duty start
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalSetDutyStart C.ledc_hal_set_duty_start
func (recv_ *LedcHalContextT) LedcHalSetDutyStart(channel_num LedcChannelT, duty_start bool) {
}

/**
 * @brief Set LEDC the integer part of duty value
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index (0-7), select from ledc_channel_t
 * @param duty_val LEDC duty value, the range of duty setting is [0, (2**duty_resolution)]
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalSetDutyIntPart C.ledc_hal_set_duty_int_part
func (recv_ *LedcHalContextT) LedcHalSetDutyIntPart(channel_num LedcChannelT, duty_val c.Uint32T) {
}

/**
 * @brief Set LEDC hpoint value
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index (0-7), select from ledc_channel_t
 * @param hpoint_val LEDC hpoint value(max: 0xfffff)
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalSetHpoint C.ledc_hal_set_hpoint
func (recv_ *LedcHalContextT) LedcHalSetHpoint(channel_num LedcChannelT, hpoint_val c.Uint32T) {
}

/**
 * @brief Get LEDC duty value
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index (0-7), select from ledc_channel_t
 * @param duty_val Pointer to accept the LEDC duty value
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalGetDuty C.ledc_hal_get_duty
func (recv_ *LedcHalContextT) LedcHalGetDuty(channel_num LedcChannelT, duty_val *c.Uint32T) {
}

/**
 * @brief Function to set fade parameters all-in-one
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index, select from ledc_channel_t
 * @param range Range index
 * @param dir LEDC duty change direction, increase or decrease
 * @param cycle The duty cycles
 * @param scale The step scale
 * @param step The number of increased or decreased times
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalSetFadeParam C.ledc_hal_set_fade_param
func (recv_ *LedcHalContextT) LedcHalSetFadeParam(channel_num LedcChannelT, range_ c.Uint32T, dir c.Uint32T, cycle c.Uint32T, scale c.Uint32T, step c.Uint32T) {
}

/**
 * @brief Set the total number of ranges in one fading
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index, select from ledc_channel_t
 * @param range_num Total number of ranges (1-16) of the fading configured
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalSetRangeNumber C.ledc_hal_set_range_number
func (recv_ *LedcHalContextT) LedcHalSetRangeNumber(channel_num LedcChannelT, range_num c.Uint32T) {
}

/**
 * @brief Get the total number of ranges in one fading
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index, select from ledc_channel_t
 * @param range_num Pointer to accept fade range number
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalGetRangeNumber C.ledc_hal_get_range_number
func (recv_ *LedcHalContextT) LedcHalGetRangeNumber(channel_num LedcChannelT, range_num *c.Uint32T) {
}

/**
 * @brief Read the fade parameters that are stored in gamma ram for a certain fade range
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index, select from ledc_channel_t
 * @param range Range index (0 - (SOC_LEDC_GAMMA_CURVE_FADE_RANGE_MAX-1)), it specifies to which range in gamma ram to read
 * @param dir Pointer to accept fade direction value
 * @param cycle Pointer to accept fade cycle value
 * @param scale Pointer to accept fade scale value
 * @param step Pointer to accept fade step value
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalGetFadeParam C.ledc_hal_get_fade_param
func (recv_ *LedcHalContextT) LedcHalGetFadeParam(channel_num LedcChannelT, range_ c.Uint32T, dir *c.Uint32T, cycle *c.Uint32T, scale *c.Uint32T, step *c.Uint32T) {
}

/**
 * @brief Clear left-off range fade parameters in LEDC gamma ram
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index, select from ledc_channel_t
 * @param start_range Start of the range to clear
 */
// llgo:link (*LedcHalContextT).LedcHalClearLeftOffFadeParam C.ledc_hal_clear_left_off_fade_param
func (recv_ *LedcHalContextT) LedcHalClearLeftOffFadeParam(channel_num LedcChannelT, start_range c.Uint32T) {
}

/**
 * @brief Get interrupt status of the specified channel
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index (0-7), select from ledc_channel_t
 * @param intr_status Pointer to accept the interrupt status
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalGetFadeEndIntrStatus C.ledc_hal_get_fade_end_intr_status
func (recv_ *LedcHalContextT) LedcHalGetFadeEndIntrStatus(intr_status *c.Uint32T) {
}

/**
 * @brief Clear interrupt status of the specified channel
 *
 * @param hal Context of the HAL layer
 * @param channel_num LEDC channel index (0-7), select from ledc_channel_t
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalClearFadeEndIntrStatus C.ledc_hal_clear_fade_end_intr_status
func (recv_ *LedcHalContextT) LedcHalClearFadeEndIntrStatus(channel_num LedcChannelT) {
}

/**
 * @brief Get clock config of LEDC timer
 *
 * @param hal Context of the HAL layer
 * @param timer_sel LEDC timer index (0-3), select from ledc_timer_t
 * @param clk_cfg Pointer to accept clock config
 *
 * @return None
 */
// llgo:link (*LedcHalContextT).LedcHalGetClkCfg C.ledc_hal_get_clk_cfg
func (recv_ *LedcHalContextT) LedcHalGetClkCfg(timer_sel LedcTimerT, clk_cfg *LedcClkCfgT) {
}
