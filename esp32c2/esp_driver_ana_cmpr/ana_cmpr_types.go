package esp_driver_ana_cmpr

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const ANA_CMPR_UNIT_0 = 0

type AnaCmprUnitT c.Int
type AnaCmprRefSourceT c.Int

const (
	ANA_CMPR_REF_SRC_INTERNAL AnaCmprRefSourceT = 0
	ANA_CMPR_REF_SRC_EXTERNAL AnaCmprRefSourceT = 1
)

type AnaCmprChannelTypeT c.Int

const (
	ANA_CMPR_SOURCE_CHAN  AnaCmprChannelTypeT = 0
	ANA_CMPR_EXT_REF_CHAN AnaCmprChannelTypeT = 1
)

type AnaCmprCrossTypeT c.Int

const (
	ANA_CMPR_CROSS_DISABLE AnaCmprCrossTypeT = 0
	ANA_CMPR_CROSS_POS     AnaCmprCrossTypeT = 1
	ANA_CMPR_CROSS_NEG     AnaCmprCrossTypeT = 2
	ANA_CMPR_CROSS_ANY     AnaCmprCrossTypeT = 3
)

type AnaCmprRefVoltageT c.Int

const (
	ANA_CMPR_REF_VOLT_0_PCT_VDD  AnaCmprRefVoltageT = 0
	ANA_CMPR_REF_VOLT_10_PCT_VDD AnaCmprRefVoltageT = 1
	ANA_CMPR_REF_VOLT_20_PCT_VDD AnaCmprRefVoltageT = 2
	ANA_CMPR_REF_VOLT_30_PCT_VDD AnaCmprRefVoltageT = 3
	ANA_CMPR_REF_VOLT_40_PCT_VDD AnaCmprRefVoltageT = 4
	ANA_CMPR_REF_VOLT_50_PCT_VDD AnaCmprRefVoltageT = 5
	ANA_CMPR_REF_VOLT_60_PCT_VDD AnaCmprRefVoltageT = 6
	ANA_CMPR_REF_VOLT_70_PCT_VDD AnaCmprRefVoltageT = 7
)

type AnaCmprT struct {
	Unused [8]uint8
}
type AnaCmprHandleT *AnaCmprT
type AnaCmprClkSrcT c.Int

/**
 * @brief Analog comparator cross event data
 *
 */

type AnaCmprCrossEventDataT struct {
	CrossType AnaCmprCrossTypeT
}

// llgo:type C
type AnaCmprCrossCbT func(AnaCmprHandleT, *AnaCmprCrossEventDataT, c.Pointer) bool
