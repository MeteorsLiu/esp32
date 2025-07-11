package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const TWAI_LL_BRP_DIV_THRESH = 128

/*
 * The following frame structure has an NEARLY identical bit field layout to
 * each byte of the TX buffer. This allows for formatting and parsing frames to
 * be done outside of time critical regions (i.e., ISRs). All the ISR needs to
 * do is to copy byte by byte to/from the TX/RX buffer. The two reserved bits in
 * TX buffer are used in the frame structure to store the self_reception and
 * single_shot flags which in turn indicate the type of transmission to execute.
 */

type TwaiLlFrameBufferT struct {
	Bytes [13]c.Uint8T
}

/**
 * Some errata workarounds will require a hardware reset of the peripheral. Thus
 * certain registers must be saved before the reset, and then restored after the
 * reset. This structure is used to hold some of those registers.
 */

type TwaiLlRegSaveT struct {
	ModeReg              c.Uint8T
	InterruptEnableReg   c.Uint8T
	BusTiming0Reg        c.Uint8T
	BusTiming1Reg        c.Uint8T
	ErrorWarningLimitReg c.Uint8T
	AcrReg               [4]c.Uint8T
	AmrReg               [4]c.Uint8T
	RxErrorCounterReg    c.Uint8T
	TxErrorCounterReg    c.Uint8T
	ClockDividerReg      c.Uint8T
}
type TwaiLlErrTypeT c.Int

const (
	TWAI_LL_ERR_BIT   TwaiLlErrTypeT = 0
	TWAI_LL_ERR_FORM  TwaiLlErrTypeT = 1
	TWAI_LL_ERR_STUFF TwaiLlErrTypeT = 2
	TWAI_LL_ERR_OTHER TwaiLlErrTypeT = 3
	TWAI_LL_ERR_MAX   TwaiLlErrTypeT = 4
)

type TwaiLlErrDirT c.Int

const (
	TWAI_LL_ERR_DIR_TX  TwaiLlErrDirT = 0
	TWAI_LL_ERR_DIR_RX  TwaiLlErrDirT = 1
	TWAI_LL_ERR_DIR_MAX TwaiLlErrDirT = 2
)

type TwaiLlErrSegT c.Int

const (
	TWAI_LL_ERR_SEG_SOF        TwaiLlErrSegT = 0
	TWAI_LL_ERR_SEG_ID_28_21   TwaiLlErrSegT = 2
	TWAI_LL_ERR_SEG_SRTR       TwaiLlErrSegT = 4
	TWAI_LL_ERR_SEG_IDE        TwaiLlErrSegT = 5
	TWAI_LL_ERR_SEG_ID_20_18   TwaiLlErrSegT = 6
	TWAI_LL_ERR_SEG_ID_17_13   TwaiLlErrSegT = 7
	TWAI_LL_ERR_SEG_CRC_SEQ    TwaiLlErrSegT = 8
	TWAI_LL_ERR_SEG_R0         TwaiLlErrSegT = 9
	TWAI_LL_ERR_SEG_DATA       TwaiLlErrSegT = 10
	TWAI_LL_ERR_SEG_DLC        TwaiLlErrSegT = 11
	TWAI_LL_ERR_SEG_RTR        TwaiLlErrSegT = 12
	TWAI_LL_ERR_SEG_R1         TwaiLlErrSegT = 13
	TWAI_LL_ERR_SEG_ID_4_0     TwaiLlErrSegT = 14
	TWAI_LL_ERR_SEG_ID_12_5    TwaiLlErrSegT = 15
	TWAI_LL_ERR_SEG_ACT_FLAG   TwaiLlErrSegT = 17
	TWAI_LL_ERR_SEG_INTER      TwaiLlErrSegT = 18
	TWAI_LL_ERR_SEG_SUPERPOS   TwaiLlErrSegT = 19
	TWAI_LL_ERR_SEG_PASS_FLAG  TwaiLlErrSegT = 22
	TWAI_LL_ERR_SEG_ERR_DELIM  TwaiLlErrSegT = 23
	TWAI_LL_ERR_SEG_CRC_DELIM  TwaiLlErrSegT = 24
	TWAI_LL_ERR_SEG_ACK_SLOT   TwaiLlErrSegT = 25
	TWAI_LL_ERR_SEG_EOF        TwaiLlErrSegT = 26
	TWAI_LL_ERR_SEG_ACK_DELIM  TwaiLlErrSegT = 27
	TWAI_LL_ERR_SEG_OVRLD_FLAG TwaiLlErrSegT = 28
	TWAI_LL_ERR_SEG_MAX        TwaiLlErrSegT = 29
)
