package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: PGM Data Register */
/** Type of pgm_data0 register
 *  Register 0 that stores data to be programmed.
 */

type EfusePgmData0RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data1 register
 *  Register 1 that stores data to be programmed.
 */

type EfusePgmData1RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data2 register
 *  Register 2 that stores data to be programmed.
 */

type EfusePgmData2RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data3 register
 *  Register 3 that stores data to be programmed.
 */

type EfusePgmData3RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data4 register
 *  Register 4 that stores data to be programmed.
 */

type EfusePgmData4RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data5 register
 *  Register 5 that stores data to be programmed.
 */

type EfusePgmData5RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data6 register
 *  Register 6 that stores data to be programmed.
 */

type EfusePgmData6RegT struct {
	Val c.Uint32T
}

/** Type of pgm_data7 register
 *  Register 7 that stores data to be programmed.
 */

type EfusePgmData7RegT struct {
	Val c.Uint32T
}

/** Type of pgm_check_value0 register
 *  Register 0 that stores the RS code to be programmed.
 */

type EfusePgmCheckValue0RegT struct {
	Val c.Uint32T
}

/** Type of pgm_check_value1 register
 *  Register 1 that stores the RS code to be programmed.
 */

type EfusePgmCheckValue1RegT struct {
	Val c.Uint32T
}

/** Type of pgm_check_value2 register
 *  Register 2 that stores the RS code to be programmed.
 */

type EfusePgmCheckValue2RegT struct {
	Val c.Uint32T
}

/** Group: Read Data Register */
/** Type of rd_wr_dis register
 *  BLOCK0 data register 0.
 */

type EfuseRdWrDisRegT struct {
	Val c.Uint32T
}

/** Type of rd_repeat_data0 register
 *  BLOCK0 data register 1.
 */

type EfuseRdRepeatData0RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk1_data0 register
 *  BLOCK1 data register 0.
 */

type EfuseRdBlk1Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk1_data1 register
 *  BLOCK1 data register 1.
 */

type EfuseRdBlk1Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk1_data2 register
 *  BLOCK1 data register 2.
 */

type EfuseRdBlk1Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data0 register
 *  Register 0 of BLOCK2.
 */

type EfuseRdBlk2Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data1 register
 *  Register 1 of BLOCK2.
 */

type EfuseRdBlk2Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data2 register
 *  Register 2 of BLOCK2.
 */

type EfuseRdBlk2Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data3 register
 *  Register 3 of BLOCK2.
 */

type EfuseRdBlk2Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data4 register
 *  Register 4 of BLOCK2.
 */

type EfuseRdBlk2Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data5 register
 *  Register 5 of BLOCK2.
 */

type EfuseRdBlk2Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data6 register
 *  Register 6 of BLOCK2.
 */

type EfuseRdBlk2Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk2_data7 register
 *  Register 7 of BLOCK2.
 */

type EfuseRdBlk2Data7RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data0 register
 *  Register 0 of BLOCK3.
 */

type EfuseRdBlk3Data0RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data1 register
 *  Register 1 of BLOCK3.
 */

type EfuseRdBlk3Data1RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data2 register
 *  Register 2 of BLOCK3.
 */

type EfuseRdBlk3Data2RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data3 register
 *  Register 3 of BLOCK3.
 */

type EfuseRdBlk3Data3RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data4 register
 *  Register 4 of BLOCK3.
 */

type EfuseRdBlk3Data4RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data5 register
 *  Register 5 of BLOCK3.
 */

type EfuseRdBlk3Data5RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data6 register
 *  Register 6 of BLOCK3.
 */

type EfuseRdBlk3Data6RegT struct {
	Val c.Uint32T
}

/** Type of rd_blk3_data7 register
 *  Register 7 of BLOCK3.
 */

type EfuseRdBlk3Data7RegT struct {
	Val c.Uint32T
}

/** Group: Report Register */
/** Type of rd_repeat_err register
 *  Programming error record register 0 of BLOCK0.
 */

type EfuseRdRepeatErrRegT struct {
	Val c.Uint32T
}

/** Type of rd_rs_err register
 *  Programming error record register 0 of BLOCK1-10.
 */

type EfuseRdRsErrRegT struct {
	Val c.Uint32T
}

/** Group: Configuration Register */
/** Type of clk register
 *  eFuse clcok configuration register.
 */

type EfuseClkRegT struct {
	Val c.Uint32T
}

/** Type of conf register
 *  eFuse operation mode configuration register
 */

type EfuseConfRegT struct {
	Val c.Uint32T
}

/** Type of cmd register
 *  eFuse command register.
 */

type EfuseCmdRegT struct {
	Val c.Uint32T
}

/** Type of dac_conf register
 *  Controls the eFuse programming voltage.
 */

type EfuseDacConfRegT struct {
	Val c.Uint32T
}

/** Type of rd_tim_conf register
 *  Configures read timing parameters.
 */

type EfuseRdTimConfRegT struct {
	Val c.Uint32T
}

/** Type of wr_tim_conf0 register
 *  Configurarion register 0 of eFuse programming timing parameters.
 */

type EfuseWrTimConf0RegT struct {
	Val c.Uint32T
}

/** Type of wr_tim_conf1 register
 *  Configurarion register 1 of eFuse programming timing parameters.
 */

type EfuseWrTimConf1RegT struct {
	Val c.Uint32T
}

/** Type of wr_tim_conf2 register
 *  Configurarion register 2 of eFuse programming timing parameters.
 */

type EfuseWrTimConf2RegT struct {
	Val c.Uint32T
}

/** Group: Status Register */
/** Type of status register
 *  eFuse status register.
 */

type EfuseStatusRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt Register */
/** Type of int_raw register
 *  eFuse raw interrupt register.
 */

type EfuseIntRawRegT struct {
	Val c.Uint32T
}

/** Type of int_st register
 *  eFuse interrupt status register.
 */

type EfuseIntStRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  eFuse interrupt enable register.
 */

type EfuseIntEnaRegT struct {
	Val c.Uint32T
}

/** Type of int_clr register
 *  eFuse interrupt clear register.
 */

type EfuseIntClrRegT struct {
	Val c.Uint32T
}

/** Group: Version Register */
/** Type of date register
 *  eFuse version register.
 */

type EfuseDateRegT struct {
	Val c.Uint32T
}

type EfuseDevT struct {
	PgmData0       EfusePgmData0RegT
	PgmData1       EfusePgmData1RegT
	PgmData2       EfusePgmData2RegT
	PgmData3       EfusePgmData3RegT
	PgmData4       EfusePgmData4RegT
	PgmData5       EfusePgmData5RegT
	PgmData6       EfusePgmData6RegT
	PgmData7       EfusePgmData7RegT
	PgmCheckValue0 EfusePgmCheckValue0RegT
	PgmCheckValue1 EfusePgmCheckValue1RegT
	PgmCheckValue2 EfusePgmCheckValue2RegT
	RdWrDis        EfuseRdWrDisRegT
	RdRepeatData0  EfuseRdRepeatData0RegT
	RdBlk1Data0    EfuseRdBlk1Data0RegT
	RdBlk1Data1    EfuseRdBlk1Data1RegT
	RdBlk1Data2    EfuseRdBlk1Data2RegT
	RdBlk2Data0    EfuseRdBlk2Data0RegT
	RdBlk2Data1    EfuseRdBlk2Data1RegT
	RdBlk2Data2    EfuseRdBlk2Data2RegT
	RdBlk2Data3    EfuseRdBlk2Data3RegT
	RdBlk2Data4    EfuseRdBlk2Data4RegT
	RdBlk2Data5    EfuseRdBlk2Data5RegT
	RdBlk2Data6    EfuseRdBlk2Data6RegT
	RdBlk2Data7    EfuseRdBlk2Data7RegT
	RdBlk3Data0    EfuseRdBlk3Data0RegT
	RdBlk3Data1    EfuseRdBlk3Data1RegT
	RdBlk3Data2    EfuseRdBlk3Data2RegT
	RdBlk3Data3    EfuseRdBlk3Data3RegT
	RdBlk3Data4    EfuseRdBlk3Data4RegT
	RdBlk3Data5    EfuseRdBlk3Data5RegT
	RdBlk3Data6    EfuseRdBlk3Data6RegT
	RdBlk3Data7    EfuseRdBlk3Data7RegT
	RdRepeatErr    EfuseRdRepeatErrRegT
	RdRsErr        EfuseRdRsErrRegT
	Clk            EfuseClkRegT
	Conf           EfuseConfRegT
	Status         EfuseStatusRegT
	Cmd            EfuseCmdRegT
	IntRaw         EfuseIntRawRegT
	IntSt          EfuseIntStRegT
	Reserved0a0    [24]c.Uint32T
	IntEna         EfuseIntEnaRegT
	IntClr         EfuseIntClrRegT
	DacConf        EfuseDacConfRegT
	RdTimConf      EfuseRdTimConfRegT
	WrTimConf0     EfuseWrTimConf0RegT
	WrTimConf1     EfuseWrTimConf1RegT
	WrTimConf2     EfuseWrTimConf2RegT
	Reserved11c    [56]c.Uint32T
	Date           EfuseDateRegT
}
