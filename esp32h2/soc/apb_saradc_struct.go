package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configure Register */
/** Type of saradc_ctrl register
 *  digital saradc configure register
 */

type ApbSaradcCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_ctrl2 register
 *  digital saradc configure register
 */

type ApbSaradcCtrl2RegT struct {
	Val c.Uint32T
}

/** Type of saradc_filter_ctrl1 register
 *  digital saradc configure register
 */

type ApbSaradcFilterCtrl1RegT struct {
	Val c.Uint32T
}

/** Type of saradc_fsm_wait register
 *  digital saradc configure register
 */

type ApbSaradcFsmWaitRegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar1_status register
 *  digital saradc configure register
 */

type ApbSaradcSar1StatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar2_status register
 *  digital saradc configure register
 */

type ApbSaradcSar2StatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar_patt_tab1 register
 *  digital saradc configure register
 */

type ApbSaradcSarPattTab1RegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar_patt_tab2 register
 *  digital saradc configure register
 */

type ApbSaradcSarPattTab2RegT struct {
	Val c.Uint32T
}

/** Type of saradc_onetime_sample register
 *  digital saradc configure register
 */

type ApbSaradcOnetimeSampleRegT struct {
	Val c.Uint32T
}

/** Type of saradc_arb_ctrl register
 *  digital saradc configure register
 */

type ApbSaradcArbCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_filter_ctrl0 register
 *  digital saradc configure register
 */

type ApbSaradcFilterCtrl0RegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar1data_status register
 *  digital saradc configure register
 */

type ApbSaradcSar1dataStatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc_sar2data_status register
 *  digital saradc configure register
 */

type ApbSaradcSar2dataStatusRegT struct {
	Val c.Uint32T
}

/** Type of saradc_thres0_ctrl register
 *  digital saradc configure register
 */

type ApbSaradcThres0CtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_thres1_ctrl register
 *  digital saradc configure register
 */

type ApbSaradcThres1CtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_thres_ctrl register
 *  digital saradc configure register
 */

type ApbSaradcThresCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_ena register
 *  digital saradc int register
 */

type ApbSaradcIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_raw register
 *  digital saradc int register
 */

type ApbSaradcIntRawRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_st register
 *  digital saradc int register
 */

type ApbSaradcIntStRegT struct {
	Val c.Uint32T
}

/** Type of saradc_int_clr register
 *  digital saradc int register
 */

type ApbSaradcIntClrRegT struct {
	Val c.Uint32T
}

/** Type of saradc_dma_conf register
 *  digital saradc configure register
 */

type ApbSaradcDmaConfRegT struct {
	Val c.Uint32T
}

/** Type of saradc_clkm_conf register
 *  digital saradc configure register
 */

type ApbSaradcClkmConfRegT struct {
	Val c.Uint32T
}

/** Type of saradc_apb_tsens_ctrl register
 *  digital tsens configure register
 */

type ApbSaradcApbTsensCtrlRegT struct {
	Val c.Uint32T
}

/** Type of saradc_tsens_ctrl2 register
 *  digital tsens configure register
 */

type ApbSaradcTsensCtrl2RegT struct {
	Val c.Uint32T
}

/** Type of saradc_cali register
 *  digital saradc configure register
 */

type ApbSaradcCaliRegT struct {
	Val c.Uint32T
}

/** Type of tsens_wake register
 *  digital tsens configure register
 */

type ApbTsensWakeRegT struct {
	Val c.Uint32T
}

/** Type of tsens_sample register
 *  digital tsens configure register
 */

type ApbTsensSampleRegT struct {
	Val c.Uint32T
}

/** Type of saradc_ctrl_date register
 *  version
 */

type ApbSaradcCtrlDateRegT struct {
	Val c.Uint32T
}

type ApbSaradcDevT struct {
	SaradcCtrl           ApbSaradcCtrlRegT
	SaradcCtrl2          ApbSaradcCtrl2RegT
	SaradcFilterCtrl1    ApbSaradcFilterCtrl1RegT
	SaradcFsmWait        ApbSaradcFsmWaitRegT
	SaradcSar1Status     ApbSaradcSar1StatusRegT
	SaradcSar2Status     ApbSaradcSar2StatusRegT
	SaradcSarPattTab1    ApbSaradcSarPattTab1RegT
	SaradcSarPattTab2    ApbSaradcSarPattTab2RegT
	SaradcOnetimeSample  ApbSaradcOnetimeSampleRegT
	SaradcArbCtrl        ApbSaradcArbCtrlRegT
	SaradcFilterCtrl0    ApbSaradcFilterCtrl0RegT
	SaradcSar1dataStatus ApbSaradcSar1dataStatusRegT
	SaradcSar2dataStatus ApbSaradcSar2dataStatusRegT
	SaradcThres0Ctrl     ApbSaradcThres0CtrlRegT
	SaradcThres1Ctrl     ApbSaradcThres1CtrlRegT
	SaradcThresCtrl      ApbSaradcThresCtrlRegT
	SaradcIntEna         ApbSaradcIntEnaRegT
	SaradcIntRaw         ApbSaradcIntRawRegT
	SaradcIntSt          ApbSaradcIntStRegT
	SaradcIntClr         ApbSaradcIntClrRegT
	SaradcDmaConf        ApbSaradcDmaConfRegT
	SaradcClkmConf       ApbSaradcClkmConfRegT
	SaradcApbTsensCtrl   ApbSaradcApbTsensCtrlRegT
	SaradcTsensCtrl2     ApbSaradcTsensCtrl2RegT
	SaradcCali           ApbSaradcCaliRegT
	TsensWake            ApbTsensWakeRegT
	TsensSample          ApbTsensSampleRegT
	Reserved06c          [228]c.Uint32T
	SaradcCtrlDate       ApbSaradcCtrlDateRegT
}
