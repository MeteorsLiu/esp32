package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// TODO: ZB-93, rewrite this file using regdesc tools when IEEE802154.csv is ready.
type EspIeee802154S struct {
	Cmd struct {
		Val c.Uint32T
	}
	Conf struct {
		Val c.Uint32T
	}
	Multipan [4]struct {
		ShortAddr struct {
			Val c.Uint32T
		}
		Panid struct {
			Val c.Uint32T
		}
		ExtAddr0 c.Uint32T
		ExtAddr1 c.Uint32T
	}
	Channel struct {
		Val c.Uint32T
	}
	Txpower struct {
		Val c.Uint32T
	}
	EdDuration struct {
		Val c.Uint32T
	}
	EdCfg struct {
		Val c.Uint32T
	}
	IfsCfg struct {
		Val c.Uint32T
	}
	AckTimeout struct {
		Val c.Uint32T
	}
	EventEn struct {
		Val c.Uint32T
	}
	EventStatus struct {
		Val c.Uint32T
	}
	RxAbortEventEn struct {
		Val c.Uint32T
	}
	PendingCfg struct {
		Val c.Uint32T
	}
	Pti struct {
		Val c.Uint32T
	}
	Reserved74     c.Uint32T
	TxAbortEventEn struct {
		Val c.Uint32T
	}
	EnhackGenerateDoneNotify c.Uint32T
	RxStatus                 struct {
		Val c.Uint32T
	}
	TxStatus struct {
		Val c.Uint32T
	}
	TxrxStatus struct {
		Val c.Uint32T
	}
	TxSecScheduleState c.Uint32T
	CoreGckCfg         struct {
		Val c.Uint32T
	}
	Reserved94          c.Uint32T
	Reserved98          c.Uint32T
	Reserved9c          c.Uint32T
	ReservedA0          c.Uint32T
	RxLength            c.Uint32T
	Timer0Threshold     c.Uint32T
	Timer0Value         c.Uint32T
	Timer1Threshold     c.Uint32T
	Timer1Value         c.Uint32T
	ClkCounterThreshold c.Uint32T
	ClkCounterValue     c.Uint32T
	IfsCounterCfg       struct {
		Val c.Uint32T
	}
	SfdWait struct {
		Val c.Uint32T
	}
	TxrxPathDelay struct {
		Val c.Uint32T
	}
	BbClk     c.Uint32T
	DmaTxAddr c.Uint32T
	DmaTxCfg  struct {
		Val c.Uint32T
	}
	DmaTxErr   c.Uint32T
	ReservedDc c.Uint32T
	DmaRxAddr  c.Uint32T
	DmaRxCfg   struct {
		Val c.Uint32T
	}
	DmaRxErr     c.Uint32T
	ReservedEc   c.Uint32T
	DmaGck       c.Uint32T
	DmaDummyData c.Uint32T
	ReservedF8   c.Uint32T
	ReservedFc   c.Uint32T
	PaOnDelay    struct {
		Val c.Uint32T
	}
	TxOnDelay struct {
		Val c.Uint32T
	}
	TxenStopDelay struct {
		Val c.Uint32T
	}
	TxOffDelay struct {
		Val c.Uint32T
	}
	RxOnDelay struct {
		Val c.Uint32T
	}
	TxrxSwitchDelay struct {
		Val c.Uint32T
	}
	ContRxDelay c.Uint32T
	DcdcCtrl    struct {
		Val c.Uint32T
	}
	DebugCtrl struct {
		Val c.Uint32T
	}
	TxDmaErrStsReg c.Uint32T
	SecurityCtrl   struct {
		Val c.Uint32T
	}
	SecurityAddr0           c.Uint32T
	SecurityAddr1           c.Uint32T
	SecurityKey0            c.Uint32T
	SecurityKey1            c.Uint32T
	SecurityKey2            c.Uint32T
	SecurityKey3            c.Uint32T
	DebugSfdTimeoutCnt      c.Uint32T
	DebugCrcErrorCnt        c.Uint32T
	DebugEdAbortCnt         c.Uint32T
	DebugCcaFailCnt         c.Uint32T
	DebugRxFilterFailCnt    c.Uint32T
	DebugNoRssDetectCnt     c.Uint32T
	DebugRxAbortCoexCnt     c.Uint32T
	DebugRxRestartCnt       c.Uint32T
	DebugTxAckAbortCoexCnt  c.Uint32T
	DebugEdScanBreakCoexCnt c.Uint32T
	DebugRxAckAbortCoexCnt  c.Uint32T
	DebugRxAckTimeoutCnt    c.Uint32T
	DebugTxBreakCoexCnt     c.Uint32T
	DebugTxSecurityErrorCnt c.Uint32T
	DebugCcaBusyCnt         c.Uint32T
	DebugCntClr             struct {
		Val c.Uint32T
	}
	I154Version c.Uint32T
}
type EspIeee802154T EspIeee802154S
