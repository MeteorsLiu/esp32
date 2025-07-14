package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration registers */
/** Type of cfg_data0 register
 *  Configure sdio cis content
 */

type HinfCfgData0RegT struct {
	Val c.Uint32T
}

/** Type of cfg_data1 register
 *  SDIO configuration register
 */

type HinfCfgData1RegT struct {
	Val c.Uint32T
}

/** Type of cfg_timing register
 *  Timing configuration registers
 */

type HinfCfgTimingRegT struct {
	Val c.Uint32T
}

/** Type of cfg_update register
 *  update sdio configurations
 */

type HinfCfgUpdateRegT struct {
	Val c.Uint32T
}

/** Type of cfg_data7 register
 *  SDIO configuration register
 */

type HinfCfgData7RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w0 register
 *  SDIO cis configuration register
 */

type HinfCisConfW0RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w1 register
 *  SDIO cis configuration register
 */

type HinfCisConfW1RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w2 register
 *  SDIO cis configuration register
 */

type HinfCisConfW2RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w3 register
 *  SDIO cis configuration register
 */

type HinfCisConfW3RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w4 register
 *  SDIO cis configuration register
 */

type HinfCisConfW4RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w5 register
 *  SDIO cis configuration register
 */

type HinfCisConfW5RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w6 register
 *  SDIO cis configuration register
 */

type HinfCisConfW6RegT struct {
	Val c.Uint32T
}

/** Type of cis_conf_w7 register
 *  SDIO cis configuration register
 */

type HinfCisConfW7RegT struct {
	Val c.Uint32T
}

/** Type of cfg_data16 register
 *  SDIO cis configuration register
 */

type HinfCfgData16RegT struct {
	Val c.Uint32T
}

/** Type of cfg_uhs1_int_mode register
 *  configure int to start and end ahead of time in uhs1 mode
 */

type HinfCfgUhs1IntModeRegT struct {
	Val c.Uint32T
}

/** Type of sdio_slave_eco_low register
 *  sdio_slave redundant control registers
 */

type HinfSdioSlaveEcoLowRegT struct {
	Val c.Uint32T
}

/** Type of sdio_slave_eco_high register
 *  sdio_slave redundant control registers
 */

type HinfSdioSlaveEcoHighRegT struct {
	Val c.Uint32T
}

/** Type of sdio_slave_eco_conf register
 *  sdio_slave redundant control registers
 */

type HinfSdioSlaveEcoConfRegT struct {
	Val c.Uint32T
}

/** Type of sdio_slave_ldo_conf register
 *  sdio slave ldo control register
 */

type HinfSdioSlaveLdoConfRegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of conf_status register
 *  func0 config0 status
 */

type HinfConfStatusRegT struct {
	Val c.Uint32T
}

/** Group: Version register */
/** Type of sdio_date register
 *  ******* Description ***********
 */

type HinfSdioDateRegT struct {
	Val c.Uint32T
}

type HinfDevT struct {
	CfgData0         HinfCfgData0RegT
	CfgData1         HinfCfgData1RegT
	CfgTiming        HinfCfgTimingRegT
	CfgUpdate        HinfCfgUpdateRegT
	Reserved010      [3]c.Uint32T
	CfgData7         HinfCfgData7RegT
	CisConfW0        HinfCisConfW0RegT
	CisConfW1        HinfCisConfW1RegT
	CisConfW2        HinfCisConfW2RegT
	CisConfW3        HinfCisConfW3RegT
	CisConfW4        HinfCisConfW4RegT
	CisConfW5        HinfCisConfW5RegT
	CisConfW6        HinfCisConfW6RegT
	CisConfW7        HinfCisConfW7RegT
	CfgData16        HinfCfgData16RegT
	CfgUhs1IntMode   HinfCfgUhs1IntModeRegT
	Reserved048      [3]c.Uint32T
	ConfStatus       HinfConfStatusRegT
	Reserved058      [19]c.Uint32T
	SdioSlaveEcoLow  HinfSdioSlaveEcoLowRegT
	SdioSlaveEcoHigh HinfSdioSlaveEcoHighRegT
	SdioSlaveEcoConf HinfSdioSlaveEcoConfRegT
	SdioSlaveLdoConf HinfSdioSlaveLdoConfRegT
	Reserved0b4      [18]c.Uint32T
	SdioDate         HinfSdioDateRegT
}
