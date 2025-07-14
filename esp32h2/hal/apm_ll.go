package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const APM_LL_CTRL_EXCEPTION_ID_S = 18
const APM_LL_CTRL_EXCEPTION_MODE_S = 16
const APM_LL_CTRL_EXCEPTION_REGION_S = 0
const APM_LL_HP_MAX_REGION_NUM = 15
const APM_LL_LP_MAX_REGION_NUM = 3
const APM_LL_MASTER_MAX = 32
const HP_APM_MAX_ACCESS_PATH = 0x4
const LP_APM_MAX_ACCESS_PATH = 0x1

type ApmLlMasterIdT c.Int

const (
	APM_LL_MASTER_HPCORE      ApmLlMasterIdT = 0
	APM_LL_MASTER_LPCORE      ApmLlMasterIdT = 1
	APM_LL_MASTER_REGDMA      ApmLlMasterIdT = 2
	APM_LL_MASTER_SDIOSLV     ApmLlMasterIdT = 3
	APM_LL_MASTER_MODEM       ApmLlMasterIdT = 4
	APM_LL_MASTER_MEM_MONITOR ApmLlMasterIdT = 5
	APM_LL_MASTER_TRACE       ApmLlMasterIdT = 6
	APM_LL_MASTER_GDMA        ApmLlMasterIdT = 16
	APM_LL_MASTER_GDMA_SPI2   ApmLlMasterIdT = 16
	APM_LL_MASTER_GDMA_UHCI0  ApmLlMasterIdT = 18
	APM_LL_MASTER_GDMA_I2S0   ApmLlMasterIdT = 19
	APM_LL_MASTER_GDMA_AES    ApmLlMasterIdT = 22
	APM_LL_MASTER_GDMA_SHA    ApmLlMasterIdT = 23
	APM_LL_MASTER_GDMA_ADC    ApmLlMasterIdT = 24
	APM_LL_MASTER_GDMA_PARLIO ApmLlMasterIdT = 25
)

type ApmLlApmCtrlT c.Int

const (
	LP_APM_CTRL ApmLlApmCtrlT = 0
	HP_APM_CTRL ApmLlApmCtrlT = 1
)

type ApmLlSecureModeT c.Int

const (
	APM_LL_SECURE_MODE_TEE  ApmLlSecureModeT = 0
	APM_LL_SECURE_MODE_REE0 ApmLlSecureModeT = 1
	APM_LL_SECURE_MODE_REE1 ApmLlSecureModeT = 2
	APM_LL_SECURE_MODE_REE2 ApmLlSecureModeT = 3
)

type ApmLlCtrlAccessPathT c.Int

const (
	APM_CTRL_ACCESS_PATH_M0 ApmLlCtrlAccessPathT = 0
	APM_CTRL_ACCESS_PATH_M1 ApmLlCtrlAccessPathT = 1
	APM_CTRL_ACCESS_PATH_M2 ApmLlCtrlAccessPathT = 2
	APM_CTRL_ACCESS_PATH_M3 ApmLlCtrlAccessPathT = 3
)

/**
 * @brief APM Ctrl path.
 */

type ApmCtrlPathT struct {
	ApmCtrl  ApmLlApmCtrlT
	ApmMPath ApmLlCtrlAccessPathT
}

/**
 * @brief APM exception information
 */

type ApmCtrlExceptionInfoT struct {
	ApmPath  ApmCtrlPathT
	ExcpRegn c.Uint8T
	ExcpMode c.Uint8T
	ExcpId   c.Uint8T
	ExcpType c.Uint8T
	ExcpAddr c.Uint32T
}
