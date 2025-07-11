package xtensa

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief  Perform an ERI read
 * @param  addr : ERI register to read from
 *
 * @return Value read
 */
//go:linkname EriRead C.eri_read
func EriRead(addr c.Int) c.Uint32T

/**
 * @brief  Perform an ERI write
 * @param  addr : ERI register to write to
 * @param  data : Value to write
 *
 * @return Value read
 */
//go:linkname EriWrite C.eri_write
func EriWrite(addr c.Int, data c.Uint32T)
