package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: configuration registers */
/** Type of addr_lock register
 *  hardware lock register
 */

type AtomicAddrLockRegT struct {
	Val c.Uint32T
}

/** Type of lr_addr register
 *  gloable lr address register
 */

type AtomicLrAddrRegT struct {
	Val c.Uint32T
}

/** Type of lr_value register
 *  gloable lr value register
 */

type AtomicLrValueRegT struct {
	Val c.Uint32T
}

/** Type of lock_status register
 *  lock status register
 */

type AtomicLockStatusRegT struct {
	Val c.Uint32T
}

/** Type of counter register
 *  wait counter register
 */

type AtomicCounterRegT struct {
	Val c.Uint32T
}

type AtomicDevT struct {
	AddrLock   AtomicAddrLockRegT
	LrAddr     AtomicLrAddrRegT
	LrValue    AtomicLrValueRegT
	LockStatus AtomicLockStatusRegT
	Counter    AtomicCounterRegT
}
