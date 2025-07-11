package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * This structure lists pin numbers
 */

type EmacIoInfoT struct {
	MdcIdx       c.Uint32T
	MdoIdx       c.Uint32T
	MdiIdx       c.Uint32T
	MiiTxClkIIdx c.Uint32T
	MiiTxEnOIdx  c.Uint32T
	MiiTxd0OIdx  c.Uint32T
	MiiTxd1OIdx  c.Uint32T
	MiiTxd2OIdx  c.Uint32T
	MiiTxd3OIdx  c.Uint32T
	MiiRxClkIIdx c.Uint32T
	MiiRxDvIIdx  c.Uint32T
	MiiRxd0IIdx  c.Uint32T
	MiiRxd1IIdx  c.Uint32T
	MiiRxd2IIdx  c.Uint32T
	MiiRxd3IIdx  c.Uint32T
	MiiColIIdx   c.Uint32T
	MiiCrsIIdx   c.Uint32T
	MiiRxErIIdx  c.Uint32T
	MiiTxErOIdx  c.Uint32T
}

type EmacIomuxInfoT struct {
	GpioNum GpioNumT
	Func    c.Uint32T
}

type EmacRmiiIomuxInfoT struct {
	Clki  *EmacIomuxInfoT
	Clko  *EmacIomuxInfoT
	TxEn  *EmacIomuxInfoT
	Txd0  *EmacIomuxInfoT
	Txd1  *EmacIomuxInfoT
	CrsDv *EmacIomuxInfoT
	Rxd0  *EmacIomuxInfoT
	Rxd1  *EmacIomuxInfoT
	TxEr  *EmacIomuxInfoT
	RxEr  *EmacIomuxInfoT
}

type EmacMiiIomuxInfoT struct {
	ClkTx *EmacIomuxInfoT
	TxEn  *EmacIomuxInfoT
	Txd0  *EmacIomuxInfoT
	Txd1  *EmacIomuxInfoT
	Txd2  *EmacIomuxInfoT
	Txd3  *EmacIomuxInfoT
	ClkRx *EmacIomuxInfoT
	RxDv  *EmacIomuxInfoT
	Rxd0  *EmacIomuxInfoT
	Rxd1  *EmacIomuxInfoT
	Rxd2  *EmacIomuxInfoT
	Rxd3  *EmacIomuxInfoT
	TxEr  *EmacIomuxInfoT
	RxEr  *EmacIomuxInfoT
	ColIn *EmacIomuxInfoT
	CrsIn *EmacIomuxInfoT
}
