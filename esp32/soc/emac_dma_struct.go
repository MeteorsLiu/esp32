package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EmacDmaDevS struct {
	Dmabusmode struct {
		Val c.Uint32T
	}
	Dmatxpolldemand c.Uint32T
	Dmarxpolldemand c.Uint32T
	Dmarxbaseaddr   c.Uint32T
	Dmatxbaseaddr   c.Uint32T
	Dmastatus       struct {
		Val c.Uint32T
	}
	DmaoperationMode struct {
		Val c.Uint32T
	}
	DmainEn struct {
		Val c.Uint32T
	}
	Dmamissedfr struct {
		Val c.Uint32T
	}
	Dmarintwdtimer struct {
		Val c.Uint32T
	}
	Reserved28       c.Uint32T
	Reserved2c       c.Uint32T
	Reserved30       c.Uint32T
	Reserved34       c.Uint32T
	Reserved38       c.Uint32T
	Reserved3c       c.Uint32T
	Reserved40       c.Uint32T
	Reserved44       c.Uint32T
	Dmatxcurrdesc    c.Uint32T
	Dmarxcurrdesc    c.Uint32T
	DmatxcurraddrBuf c.Uint32T
	DmarxcurraddrBuf c.Uint32T
}
type EmacDmaDevT EmacDmaDevS
