package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const RX_BUFF_SIZE = 0x400
const TX_BUFF_SIZE = 100
const UART_INT_FLAG_MASK = 0x0E
const FRAME_FLAG = 0x7E

type UartIntType c.Int

const (
	UART_LINE_STATUS_INT_FLAG  UartIntType = 6
	UART_RCV_FIFO_INT_FLAG     UartIntType = 4
	UART_RCV_TMOUT_INT_FLAG    UartIntType = 12
	UART_TXBUFF_EMPTY_INT_FLAG UartIntType = 2
)

type UartRcvFifoTrgLvl c.Int

const (
	RCV_ONE_BYTE      UartRcvFifoTrgLvl = 0
	RCV_FOUR_BYTE     UartRcvFifoTrgLvl = 1
	RCV_EIGHT_BYTE    UartRcvFifoTrgLvl = 2
	RCV_FOURTEEN_BYTE UartRcvFifoTrgLvl = 3
)

type UartBitsNum4Char c.Int

const (
	FIVE_BITS  UartBitsNum4Char = 0
	SIX_BITS   UartBitsNum4Char = 1
	SEVEN_BITS UartBitsNum4Char = 2
	EIGHT_BITS UartBitsNum4Char = 3
)

type UartStopBitsNum c.Int

const (
	ONE_STOP_BIT      UartStopBitsNum = 1
	ONE_HALF_STOP_BIT UartStopBitsNum = 2
	TWO_STOP_BIT      UartStopBitsNum = 3
)

type UartParityMode c.Int

const (
	NONE_BITS UartParityMode = 0
	ODD_BITS  UartParityMode = 2
	EVEN_BITS UartParityMode = 3
)

type UartExistParity c.Int

const (
	STICK_PARITY_DIS UartExistParity = 0
	STICK_PARITY_EN  UartExistParity = 2
)

type UartBautRate c.Int

const (
	BIT_RATE_9600   UartBautRate = 9600
	BIT_RATE_19200  UartBautRate = 19200
	BIT_RATE_38400  UartBautRate = 38400
	BIT_RATE_57600  UartBautRate = 57600
	BIT_RATE_115200 UartBautRate = 115200
	BIT_RATE_230400 UartBautRate = 230400
	BIT_RATE_460800 UartBautRate = 460800
	BIT_RATE_921600 UartBautRate = 921600
)

type UartFlowCtrl c.Int

const (
	NONE_CTRL     UartFlowCtrl = 0
	HARDWARE_CTRL UartFlowCtrl = 1
	XON_XOFF_CTRL UartFlowCtrl = 2
)

type RcvMsgBuffState c.Int

const (
	EMPTY       RcvMsgBuffState = 0
	UNDER_WRITE RcvMsgBuffState = 1
	WRITE_OVER  RcvMsgBuffState = 2
)

type RcvMsgBuff struct {
	PRcvMsgBuff *c.Uint8T
	PWritePos   *c.Uint8T
	PReadPos    *c.Uint8T
	TrigLvl     c.Uint8T
	BuffState   RcvMsgBuffState
}

type TrxMsgBuff struct {
	TrxBuffSize c.Uint32T
	PTrxBuff    *c.Uint8T
}
type RcvMsgState c.Int

const (
	BAUD_RATE_DET RcvMsgState = 0
	WAIT_SYNC_FRM RcvMsgState = 1
	SRCH_MSG_HEAD RcvMsgState = 2
	RCV_MSG_BODY  RcvMsgState = 3
	RCV_ESC_CHAR  RcvMsgState = 4
)

type UartDevice struct {
	BautRate    UartBautRate
	DataBits    UartBitsNum4Char
	ExistParity UartExistParity
	Parity      UartParityMode
	StopBits    UartStopBitsNum
	FlowCtrl    UartFlowCtrl
	BuffUartNo  c.Uint8T
	RcvBuff     RcvMsgBuff
	RcvState    RcvMsgState
	Received    c.Int
}

/**
 * @brief Get uart configuration struct.
 *        Please do not call this function in SDK.
 *
 * @param  None
 *
 * @return UartDevice * : uart configuration struct pointer.
 */
//go:linkname GetUartDevice C.GetUartDevice
func GetUartDevice() *UartDevice
