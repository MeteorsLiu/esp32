package riscv

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Route the peripheral interrupt signal to the CPU
 * @param periph_intr_source  Peripheral interrupt number, one of ETS_XXX_SOURCE
 * @param rv_int_num  CPU interrupt number
 */
//go:linkname IntrMatrixRoute C.intr_matrix_route
func IntrMatrixRoute(periph_intr_source c.Int, rv_int_num c.Int)
