package freertos

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
 * Definition of the only type of object that a list can contain.
 */

type XLIST struct {
	Unused [8]uint8
}

type XLISTITEM struct {
	XItemValue  c.Int
	PxNext      *XLISTITEM
	PxPrevious  *XLISTITEM
	PvOwner     c.Pointer
	PxContainer *XLIST
}
type ListItemT XLISTITEM
type MiniListItemT XLISTITEM
type ListT XLIST
