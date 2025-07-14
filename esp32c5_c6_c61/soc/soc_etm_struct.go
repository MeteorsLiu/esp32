package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration Register */
/** Type of ch_ena_ad0 register
 *  channel enable register
 */

type SocEtmChEnaAd0RegT struct {
	Val c.Uint32T
}

/** Type of ch_ena_ad0_set register
 *  channel enable set register
 */

type SocEtmChEnaAd0SetRegT struct {
	Val c.Uint32T
}

/** Type of ch_ena_ad0_clr register
 *  channel enable clear register
 */

type SocEtmChEnaAd0ClrRegT struct {
	Val c.Uint32T
}

/** Type of ch_ena_ad1 register
 *  channel enable register
 */

type SocEtmChEnaAd1RegT struct {
	Val c.Uint32T
}

/** Type of ch_ena_ad1_set register
 *  channel enable set register
 */

type SocEtmChEnaAd1SetRegT struct {
	Val c.Uint32T
}

/** Type of ch_ena_ad1_clr register
 *  channel enable clear register
 */

type SocEtmChEnaAd1ClrRegT struct {
	Val c.Uint32T
}

/** Type of chn_evt_id register
 *  channeln event id register
 */

type SocEtmChnEvtIdRegT struct {
	Val c.Uint32T
}

/** Type of chn_task_id register
 *  channeln task id register
 */

type SocEtmChnTaskIdRegT struct {
	Val c.Uint32T
}

/** Type of clk_en register
 *  etm clock enable register
 */

type SocEtmClkEnRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  etm date register
 */

type SocEtmDateRegT struct {
	Val c.Uint32T
}

type SocEtmDevT struct {
	ChEnaAd0    SocEtmChEnaAd0RegT
	ChEnaAd0Set SocEtmChEnaAd0SetRegT
	ChEnaAd0Clr SocEtmChEnaAd0ClrRegT
	ChEnaAd1    SocEtmChEnaAd1RegT
	ChEnaAd1Set SocEtmChEnaAd1SetRegT
	ChEnaAd1Clr SocEtmChEnaAd1ClrRegT
	Channel     [50]struct {
		Eid SocEtmChnEvtIdRegT
		Tid SocEtmChnTaskIdRegT
	}
	ClkEn SocEtmClkEnRegT
	Date  SocEtmDateRegT
}
