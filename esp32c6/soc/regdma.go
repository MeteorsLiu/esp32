package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type RegdmaLinkPriorityT c.Int

const (
	REGDMA_LINK_PRI_0 RegdmaLinkPriorityT = 0
	REGDMA_LINK_PRI_1 RegdmaLinkPriorityT = 1
	REGDMA_LINK_PRI_2 RegdmaLinkPriorityT = 2
	REGDMA_LINK_PRI_3 RegdmaLinkPriorityT = 3
	REGDMA_LINK_PRI_4 RegdmaLinkPriorityT = 4
	REGDMA_LINK_PRI_5 RegdmaLinkPriorityT = 5
	REGDMA_LINK_PRI_6 RegdmaLinkPriorityT = 6
	REGDMA_LINK_PRI_7 RegdmaLinkPriorityT = 7
)

type RegdmaEntryBufT [4]c.Pointer
type RegdmaLinkMode c.Int

const (
	REGDMA_LINK_MODE_CONTINUOUS RegdmaLinkMode = 0
	REGDMA_LINK_MODE_ADDR_MAP   RegdmaLinkMode = 1
	REGDMA_LINK_MODE_WRITE      RegdmaLinkMode = 2
	REGDMA_LINK_MODE_WAIT       RegdmaLinkMode = 3
)

type RegdmaLinkModeT RegdmaLinkMode

type RegdmaLinkHead struct {
	Length   c.Uint32T
	Reserve0 c.Uint32T
	Mode     c.Uint32T
	Reserve1 c.Uint32T
	Branch   c.Uint32T
	SkipR    c.Uint32T
	SkipB    c.Uint32T
	Eof      c.Uint32T
}
type RegdmaLinkHeadT RegdmaLinkHead

/* Continuous type linked list node body type definition */

type RegdmaLinkContinuousBody struct {
	Next    c.Pointer
	Backup  c.Pointer
	Restore c.Pointer
	Mem     c.Pointer
}
type RegdmaLinkContinuousBodyT RegdmaLinkContinuousBody

/* Address Map type linked list node body type definition */

type RegdmaLinkAddrMapBody struct {
	Next    c.Pointer
	Backup  c.Pointer
	Restore c.Pointer
	Mem     c.Pointer
	Map     [4]c.Uint32T
}
type RegdmaLinkAddrMapBodyT RegdmaLinkAddrMapBody

/* Write/Wait type linked list node body type definition */

type RegdmaLinkWriteWaitBody struct {
	Next   c.Pointer
	Backup c.Pointer
	Value  c.Uint32T
	Mask   c.Uint32T
}
type RegdmaLinkWriteWaitBodyT RegdmaLinkWriteWaitBody

/* Branch Continuous type linked list node body type definition */

type RegdmaLinkBranchContinuousBody struct {
	Next    RegdmaEntryBufT
	Backup  c.Pointer
	Restore c.Pointer
	Mem     c.Pointer
}
type RegdmaLinkBranchContinuousBodyT RegdmaLinkBranchContinuousBody

/* Branch Address Map type linked list node body type definition */

type RegdmaLinkBranchAddrMapBody struct {
	Next    RegdmaEntryBufT
	Backup  c.Pointer
	Restore c.Pointer
	Mem     c.Pointer
	Map     [4]c.Uint32T
}
type RegdmaLinkBranchAddrMapBodyT RegdmaLinkBranchAddrMapBody

/* Branch Write/Wait type linked list node body type definition */

type RegdmaLinkBranchWriteWaitBody struct {
	Next   RegdmaEntryBufT
	Backup c.Pointer
	Value  c.Uint32T
	Mask   c.Uint32T
}
type RegdmaLinkBranchWriteWaitBodyT RegdmaLinkBranchWriteWaitBody

type RegdmaLinkStats struct {
	Ref     c.Uint32T
	Reserve c.Uint32T
	Id      c.Uint32T
	Module  c.Int
}
type RegdmaLinkStatsT RegdmaLinkStats

type RegdmaLinkContinuous struct {
	Stat RegdmaLinkStatsT
	Head RegdmaLinkHeadT
	Body RegdmaLinkContinuousBodyT
	Buff [0]c.Uint32T
}
type RegdmaLinkContinuousT RegdmaLinkContinuous

type RegdmaLinkAddrMap struct {
	Stat RegdmaLinkStatsT
	Head RegdmaLinkHeadT
	Body RegdmaLinkAddrMapBodyT
	Buff [0]c.Uint32T
}
type RegdmaLinkAddrMapT RegdmaLinkAddrMap

type RegdmaLinkWriteWait struct {
	Stat RegdmaLinkStatsT
	Head RegdmaLinkHeadT
	Body RegdmaLinkWriteWaitBodyT
}
type RegdmaLinkWriteWaitT RegdmaLinkWriteWait

type RegdmaLinkBranchContinuous struct {
	Stat RegdmaLinkStatsT
	Head RegdmaLinkHeadT
	Body RegdmaLinkBranchContinuousBodyT
	Buff [0]c.Uint32T
}
type RegdmaLinkBranchContinuousT RegdmaLinkBranchContinuous

type RegdmaLinkBranchAddrMap struct {
	Stat RegdmaLinkStatsT
	Head RegdmaLinkHeadT
	Body RegdmaLinkBranchAddrMapBodyT
	Buff [0]c.Uint32T
}
type RegdmaLinkBranchAddrMapT RegdmaLinkBranchAddrMap

type RegdmaLinkBranchWriteWait struct {
	Stat RegdmaLinkStatsT
	Head RegdmaLinkHeadT
	Body RegdmaLinkBranchWriteWaitBodyT
}
type RegdmaLinkBranchWriteWaitT RegdmaLinkBranchWriteWait

type RegdmaLinkConfig struct {
	Head RegdmaLinkHeadT
	Id   c.Int
}
type RegdmaLinkConfigT RegdmaLinkConfig

type RegdmaEntriesConfigT struct {
	Config RegdmaLinkConfigT
	Owner  c.Uint32T
}
