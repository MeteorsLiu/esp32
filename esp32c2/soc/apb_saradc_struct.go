package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Registers */
/** Type of saradc_ctrl register
 *  register description
 */

type ApbSaradcCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_ctrl2 register
 *  register description
 */

type ApbSaradcCtrl2RegT struct {
	Val c.Uint32T
}

/** Type of saradc_filter_ctrl1 register
 *  register description
 */

type ApbSaradcFilterCtrl1RegT struct {
	Val c.Uint32T
}

/** Type of saradc_fsm_wait register
 *  register description
 */

type ApbSaradcFsmWaitRegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar1_status register
 *  register description
 */

type ApbSaradcSar1StatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar2_status register
 *  register description
 */

type ApbSaradcSar2StatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar_patt_tab1 register
 *  register description
 */

type ApbSaradcSarPattTab1RegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar_patt_tab2 register
 *  register description
 */

type ApbSaradcSarPattTab2RegT struct {
	Val c.Uint32T
}

/** Type of saradc_onetime_sample register
 *  register description
 */

type ApbSaradcOnetimeSampleRegT struct {
	Val c.Uint32T
}

/** Type of saradc_apb_adc_arb_ctrl register
 *  register description
 */

type ApbSaradcApbAdcArbCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_filter_ctrl0 register
 *  register description
 */

type ApbSaradcFilterCtrl0RegT struct {
	Val c.Uint32T
}

/** Type of saradc1_data_status register
 *  register description
 */

type ApbSaradc1DataStatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc2_data_status register
 *  register description
 */

type ApbSaradc2DataStatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc_thres0_ctrl register
 *  register description
 */

type ApbSaradcThres0CtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_thres1_ctrl register
 *  register description
 */

type ApbSaradcThres1CtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_thres_ctrl register
 *  register description
 */

type ApbSaradcThresCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_ena register
 *  register description
 */

type ApbSaradcIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_raw register
 *  register description
 */

type ApbSaradcIntRawRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_st register
 *  register description
 */

type ApbSaradcIntStRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_clr register
 *  register description
 */

type ApbSaradcIntClrRegT struct {
	Val c.Uint32T
}

/** Type of saradc_dma_conf register
 *  register description
 */

type ApbSaradcDmaConfRegT struct {
	Val c.Uint32T
}

/** Type of saradc_apb_adc_clkm_conf register
 *  register description
 */

type ApbSaradcApbAdcClkmConfRegT struct {
	Val c.Uint32T
}

/** Type of saradc_apb_tsens_ctrl register
 *  register description
 */

type ApbSaradcApbTsensCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_apb_tsens_ctrl2 register
 *  register description
 */

type ApbSaradcApbTsensCtrl2RegT struct {
	Val c.Uint32T
}

/** Type of saradc_cali register
 *  register description
 */

type ApbSaradcCaliRegT struct {
	Val c.Uint32T
}

/** Type of saradc_apb_ctrl_date register
 *  register description
 */

type ApbSaradcApbCtrlDateRegT struct {
	Val c.Uint32T
}

type ApbDevT struct {
	SaradcCtrl           ApbSaradcCtrlRegT
	SaradcCtrl2          ApbSaradcCtrl2RegT
	SaradcFilterCtrl1    ApbSaradcFilterCtrl1RegT
	SaradcFsmWait        ApbSaradcFsmWaitRegT
	SaradcSar1Status     ApbSaradcSar1StatusRegT
	SaradcSar2Status     ApbSaradcSar2StatusRegT
	SaradcSarPattTab1    ApbSaradcSarPattTab1RegT
	SaradcSarPattTab2    ApbSaradcSarPattTab2RegT
	SaradcOnetimeSample  ApbSaradcOnetimeSampleRegT
	SaradcApbAdcArbCtrl  ApbSaradcApbAdcArbCtrlRegT
	SaradcFilterCtrl0    ApbSaradcFilterCtrl0RegT
	Saradc1DataStatus    ApbSaradc1DataStatusRegT
	Saradc2DataStatus    ApbSaradc2DataStatusRegT
	SaradcThres0Ctrl     ApbSaradcThres0CtrlRegT
	SaradcThres1Ctrl     ApbSaradcThres1CtrlRegT
	SaradcThresCtrl      ApbSaradcThresCtrlRegT
	SaradcIntEna         ApbSaradcIntEnaRegT
	SaradcIntRaw         ApbSaradcIntRawRegT
	SaradcIntSt          ApbSaradcIntStRegT
	SaradcIntClr         ApbSaradcIntClrRegT
	SaradcDmaConf        ApbSaradcDmaConfRegT
	SaradcApbAdcClkmConf ApbSaradcApbAdcClkmConfRegT
	SaradcApbTsensCtrl   ApbSaradcApbTsensCtrlRegT
	SaradcApbTsensCtrl2  ApbSaradcApbTsensCtrl2RegT
	SaradcCali           ApbSaradcCaliRegT
	Reserved064          [230]c.Uint32T
	SaradcApbCtrlDate    ApbSaradcApbCtrlDateRegT
}
