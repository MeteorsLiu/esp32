package riscv

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Decode the offset value from a RISC-V JAL instruction
 * @note This API will abort if the instruction is not JAL formatted.
 *
 * @param inst_addr Address of JAL instruction
 * @return int offset value
 */
//go:linkname RiscvDecodeOffsetFromJalInstruction C.riscv_decode_offset_from_jal_instruction
func RiscvDecodeOffsetFromJalInstruction(inst_addr c.IntptrT) c.Int
