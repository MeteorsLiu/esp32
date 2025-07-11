package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2sSlotModeT c.Int

const (
	I2S_SLOT_MODE_MONO   I2sSlotModeT = 1
	I2S_SLOT_MODE_STEREO I2sSlotModeT = 2
)

type I2sDirT c.Int

const (
	I2S_DIR_RX I2sDirT = 1
	I2S_DIR_TX I2sDirT = 2
)

type I2sRoleT c.Int

const (
	I2S_ROLE_MASTER I2sRoleT = 0
	I2S_ROLE_SLAVE  I2sRoleT = 1
)

type I2sDataBitWidthT c.Int

const (
	I2S_DATA_BIT_WIDTH_8BIT  I2sDataBitWidthT = 8
	I2S_DATA_BIT_WIDTH_16BIT I2sDataBitWidthT = 16
	I2S_DATA_BIT_WIDTH_24BIT I2sDataBitWidthT = 24
	I2S_DATA_BIT_WIDTH_32BIT I2sDataBitWidthT = 32
)

type I2sSlotBitWidthT c.Int

const (
	I2S_SLOT_BIT_WIDTH_AUTO  I2sSlotBitWidthT = 0
	I2S_SLOT_BIT_WIDTH_8BIT  I2sSlotBitWidthT = 8
	I2S_SLOT_BIT_WIDTH_16BIT I2sSlotBitWidthT = 16
	I2S_SLOT_BIT_WIDTH_24BIT I2sSlotBitWidthT = 24
	I2S_SLOT_BIT_WIDTH_32BIT I2sSlotBitWidthT = 32
)

type I2sClockSrcT SocPeriphI2sClkSrcT
type I2sPdmDsrT c.Int

const (
	I2S_PDM_DSR_8S  I2sPdmDsrT = 0
	I2S_PDM_DSR_16S I2sPdmDsrT = 1
	I2S_PDM_DSR_MAX I2sPdmDsrT = 2
)

type I2sPdmSigScaleT c.Int

const (
	I2S_PDM_SIG_SCALING_DIV_2 I2sPdmSigScaleT = 0
	I2S_PDM_SIG_SCALING_MUL_1 I2sPdmSigScaleT = 1
	I2S_PDM_SIG_SCALING_MUL_2 I2sPdmSigScaleT = 2
	I2S_PDM_SIG_SCALING_MUL_4 I2sPdmSigScaleT = 3
)

type I2sStdSlotMaskT c.Int

const (
	I2S_STD_SLOT_LEFT  I2sStdSlotMaskT = 1
	I2S_STD_SLOT_RIGHT I2sStdSlotMaskT = 2
	I2S_STD_SLOT_BOTH  I2sStdSlotMaskT = 3
)

type I2sPdmSlotMaskT c.Int

const (
	I2S_PDM_SLOT_RIGHT I2sPdmSlotMaskT = 1
	I2S_PDM_SLOT_LEFT  I2sPdmSlotMaskT = 2
	I2S_PDM_SLOT_BOTH  I2sPdmSlotMaskT = 3
)

type I2sEtmEventTypeT c.Int

const (
	I2S_ETM_EVENT_DONE         I2sEtmEventTypeT = 0
	I2S_ETM_EVENT_REACH_THRESH I2sEtmEventTypeT = 1
	I2S_ETM_EVENT_MAX          I2sEtmEventTypeT = 2
)

type I2sEtmTaskTypeT c.Int

const (
	I2S_ETM_TASK_START I2sEtmTaskTypeT = 0
	I2S_ETM_TASK_STOP  I2sEtmTaskTypeT = 1
	I2S_ETM_TASK_MAX   I2sEtmTaskTypeT = 2
)
