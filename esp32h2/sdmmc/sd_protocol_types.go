package sdmmc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SCF_ITSDONE = 0x0001
const SCF_CMD_AC = 0x0000
const SCF_CMD_ADTC = 0x0010
const SCF_CMD_BC = 0x0020
const SCF_CMD_BCR = 0x0030
const SCF_CMD_READ = 0x0040
const SCF_RSP_BSY = 0x0100
const SCF_RSP_136 = 0x0200
const SCF_RSP_CRC = 0x0400
const SCF_RSP_IDX = 0x0800
const SCF_RSP_PRESENT = 0x1000
const SCF_RSP_R0 = 0
const SCF_WAIT_BUSY = 0x2000
const SDMMC_FREQ_DEFAULT = 20000
const SDMMC_FREQ_HIGHSPEED = 40000
const SDMMC_FREQ_PROBING = 400
const SDMMC_FREQ_52M = 52000
const SDMMC_FREQ_26M = 26000
const SDMMC_FREQ_DDR50 = 50000
const SDMMC_FREQ_SDR50 = 100000

/**
 * Decoded values from SD card Card Specific Data register
 */

type SdmmcCsdT struct {
	CsdVer           c.Int
	MmcVer           c.Int
	Capacity         c.Int
	SectorSize       c.Int
	ReadBlockLen     c.Int
	CardCommandClass c.Int
	TrSpeed          c.Int
}

/**
 * Decoded values from SD card Card IDentification register
 */

type SdmmcCidT struct {
	MfgId    c.Int
	OemId    c.Int
	Name     [8]c.Char
	Revision c.Int
	Serial   c.Int
	Date     c.Int
}

/**
 * Decoded values from SD Configuration Register
 * Note: When new member is added, update reserved bits accordingly
 */

type SdmmcScrT struct {
	SdSpec        c.Uint32T
	EraseMemState c.Uint32T
	BusWidth      c.Uint32T
	Reserved      c.Uint32T
	RsvdMnf       c.Uint32T
}

/**
 * Decoded values from SD Status Register
 * Note: When new member is added, update reserved bits accordingly
 */

type SdmmcSsrT struct {
	AllocUnitKb    c.Uint32T
	EraseSizeAu    c.Uint32T
	CurBusWidth    c.Uint32T
	DiscardSupport c.Uint32T
	FuleSupport    c.Uint32T
	EraseTimeout   c.Uint32T
	EraseOffset    c.Uint32T
	Reserved       c.Uint32T
}

/**
 * Decoded values of Extended Card Specific Data
 */

type SdmmcExtCsdT struct {
	Rev           c.Uint8T
	PowerClass    c.Uint8T
	EraseMemState c.Uint8T
	SecFeature    c.Uint8T
}
type SdmmcResponseT [4]c.Uint32T

/**
 * SD SWITCH_FUNC response buffer
 */

type SdmmcSwitchFuncRspT struct {
	Data [16]c.Uint32T
}

/**
 * SD/MMC command information
 */

type SdmmcCommandT struct {
	Opcode          c.Uint32T
	Arg             c.Uint32T
	Response        SdmmcResponseT
	Data            c.Pointer
	Datalen         c.SizeT
	Buflen          c.SizeT
	Blklen          c.SizeT
	Flags           c.Int
	Error           EspErrT
	TimeoutMs       c.Uint32T
	VoltSwitchCb    c.Pointer
	VoltSwitchCbArg c.Pointer
}
type SdmmcDelayPhaseT c.Int

const (
	SDMMC_DELAY_PHASE_0    SdmmcDelayPhaseT = 0
	SDMMC_DELAY_PHASE_1    SdmmcDelayPhaseT = 1
	SDMMC_DELAY_PHASE_2    SdmmcDelayPhaseT = 2
	SDMMC_DELAY_PHASE_3    SdmmcDelayPhaseT = 3
	SDMMC_DELAY_PHASE_AUTO SdmmcDelayPhaseT = 4
)

type SdmmcDriverStrengthT c.Int

const (
	SDMMC_DRIVER_STRENGTH_B SdmmcDriverStrengthT = 0
	SDMMC_DRIVER_STRENGTH_A SdmmcDriverStrengthT = 1
	SDMMC_DRIVER_STRENGTH_C SdmmcDriverStrengthT = 2
	SDMMC_DRIVER_STRENGTH_D SdmmcDriverStrengthT = 3
)

type SdmmcCurrentLimitT c.Int

const (
	SDMMC_CURRENT_LIMIT_200MA SdmmcCurrentLimitT = 0
	SDMMC_CURRENT_LIMIT_400MA SdmmcCurrentLimitT = 1
	SDMMC_CURRENT_LIMIT_600MA SdmmcCurrentLimitT = 2
	SDMMC_CURRENT_LIMIT_800MA SdmmcCurrentLimitT = 3
)

/**
 * SD/MMC Host description
 *
 * This structure defines properties of SD/MMC host and functions
 * of SD/MMC host which can be used by upper layers.
 */

type SdmmcHostT struct {
	Flags            c.Uint32T
	Slot             c.Int
	MaxFreqKhz       c.Int
	IoVoltage        c.Float
	DriverStrength   SdmmcDriverStrengthT
	CurrentLimit     SdmmcCurrentLimitT
	Init             c.Pointer
	SetBusWidth      c.Pointer
	GetBusWidth      c.Pointer
	SetBusDdrMode    c.Pointer
	SetCardClk       c.Pointer
	SetCclkAlwaysOn  c.Pointer
	DoTransaction    c.Pointer
	IoIntEnable      c.Pointer
	IoIntWait        c.Pointer
	CommandTimeoutMs c.Int
	GetRealFreq      c.Pointer
	InputDelayPhase  SdmmcDelayPhaseT
	SetInputDelay    c.Pointer
	DmaAlignedBuffer c.Pointer
	PwrCtrlHandle    SdPwrCtrlHandleT
	GetDmaInfo       c.Pointer
	IsSlotSetToUhs1  c.Pointer
}

/**
 * SD/MMC card information structure
 */

type SdmmcCardT struct {
	Host           SdmmcHostT
	Ocr            c.Uint32T
	Csd            SdmmcCsdT
	Scr            SdmmcScrT
	Ssr            SdmmcSsrT
	ExtCsd         SdmmcExtCsdT
	Rca            c.Uint16T
	MaxFreqKhz     c.Uint32T
	RealFreqKhz    c.Int
	IsMem          c.Uint32T
	IsSdio         c.Uint32T
	IsMmc          c.Uint32T
	NumIoFunctions c.Uint32T
	LogBusWidth    c.Uint32T
	IsDdr          c.Uint32T
	IsUhs1         c.Uint32T
	Reserved       c.Uint32T
}
type SdmmcEraseArgT c.Int

const (
	SDMMC_ERASE_ARG   SdmmcEraseArgT = 0
	SDMMC_DISCARD_ARG SdmmcEraseArgT = 1
)
