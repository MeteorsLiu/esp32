package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RtcCntlDevS struct {
	Options0 struct {
		Val c.Uint32T
	}
	SlpTimer0 c.Uint32T
	SlpTimer1 struct {
		Val c.Uint32T
	}
	TimeUpdate struct {
		Val c.Uint32T
	}
	Time0 c.Uint32T
	Time1 struct {
		Val c.Uint32T
	}
	State0 struct {
		Val c.Uint32T
	}
	Timer1 struct {
		Val c.Uint32T
	}
	Timer2 struct {
		Val c.Uint32T
	}
	Timer3 struct {
		Val c.Uint32T
	}
	Timer4 struct {
		Val c.Uint32T
	}
	Timer5 struct {
		Val c.Uint32T
	}
	AnaConf struct {
		Val c.Uint32T
	}
	ResetState struct {
		Val c.Uint32T
	}
	WakeupState struct {
		Val c.Uint32T
	}
	IntEna struct {
		Val c.Uint32T
	}
	IntRaw struct {
		Val c.Uint32T
	}
	IntSt struct {
		Val c.Uint32T
	}
	IntClr struct {
		Val c.Uint32T
	}
	RtcStore0  c.Uint32T
	RtcStore1  c.Uint32T
	RtcStore2  c.Uint32T
	RtcStore3  c.Uint32T
	ExtXtlConf struct {
		Val c.Uint32T
	}
	ExtWakeupConf struct {
		Val c.Uint32T
	}
	SlpRejectConf struct {
		Val c.Uint32T
	}
	CpuPeriodConf struct {
		Val c.Uint32T
	}
	SdioActConf struct {
		Val c.Uint32T
	}
	ClkConf struct {
		Val c.Uint32T
	}
	SdioConf struct {
		Val c.Uint32T
	}
	BiasConf struct {
		Val c.Uint32T
	}
	Rtc struct {
		Val c.Uint32T
	}
	RtcPwc struct {
		Val c.Uint32T
	}
	DigPwc struct {
		Val c.Uint32T
	}
	DigIso struct {
		Val c.Uint32T
	}
	WdtConfig0 struct {
		Val c.Uint32T
	}
	WdtConfig1 c.Uint32T
	WdtConfig2 c.Uint32T
	WdtConfig3 c.Uint32T
	WdtConfig4 c.Uint32T
	WdtFeed    struct {
		Val c.Uint32T
	}
	WdtWprotect c.Uint32T
	TestMux     struct {
		Val c.Uint32T
	}
	SwCpuStall struct {
		Val c.Uint32T
	}
	Store4    c.Uint32T
	Store5    c.Uint32T
	Store6    c.Uint32T
	Store7    c.Uint32T
	Diag0     c.Uint32T
	Diag1     c.Uint32T
	HoldForce struct {
		Val c.Uint32T
	}
	ExtWakeup1 struct {
		Val c.Uint32T
	}
	ExtWakeup1Status struct {
		Val c.Uint32T
	}
	BrownOut struct {
		Val c.Uint32T
	}
	Reserved39 c.Uint32T
	Reserved3d c.Uint32T
	Reserved41 c.Uint32T
	Reserved45 c.Uint32T
	Reserved49 c.Uint32T
	Reserved4d c.Uint32T
	Date       struct {
		Val c.Uint32T
	}
}
type RtcCntlDevT RtcCntlDevS
