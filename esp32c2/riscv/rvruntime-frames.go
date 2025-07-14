package riscv

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/*
-------------------------------------------------------------------------------

	INTERRUPT/EXCEPTION STACK FRAME FOR A EXCEPTION OR NESTED INTERRUPT

-------------------------------------------------------------------------------
*/
type RvExcFrame struct {
	Mepc    c.Long
	Ra      c.Long
	Sp      c.Long
	Gp      c.Long
	Tp      c.Long
	T0      c.Long
	T1      c.Long
	T2      c.Long
	S0      c.Long
	S1      c.Long
	A0      c.Long
	A1      c.Long
	A2      c.Long
	A3      c.Long
	A4      c.Long
	A5      c.Long
	A6      c.Long
	A7      c.Long
	S2      c.Long
	S3      c.Long
	S4      c.Long
	S5      c.Long
	S6      c.Long
	S7      c.Long
	S8      c.Long
	S9      c.Long
	S10     c.Long
	S11     c.Long
	T3      c.Long
	T4      c.Long
	T5      c.Long
	T6      c.Long
	Mstatus c.Long
	Mtvec   c.Long
	Mcause  c.Long
	Mtval   c.Long
	Mhartid c.Long
}
