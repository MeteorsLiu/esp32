package riscv

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type IntrType c.Int

const (
	INTR_TYPE_LEVEL IntrType = 0
	INTR_TYPE_EDGE  IntrType = 1
)

// llgo:type C
type IntrHandlerT func(c.Pointer)

/** Set the interrupt handler function for the given CPU interrupt
 * @param rv_int_num  CPU interrupt number
 * @param fn  Handler function
 * @param arg  Handler argument
 */
//go:linkname IntrHandlerSet C.intr_handler_set
func IntrHandlerSet(rv_int_num c.Int, fn IntrHandlerT, arg c.Pointer)

/** Get the interrupt handler function for the given CPU interrupt
 *
 *@return interrupt handler registered for a particular interrupt number, or NULL otherwise
 */
//go:linkname IntrHandlerGet C.intr_handler_get
func IntrHandlerGet(rv_int_num c.Int) IntrHandlerT

/** Get the interrupt handler argument associated with the given CPU interrupt
 *
 *@return interrupt handler argument for a particular interrupt number, or NULL otherwise
 */
//go:linkname IntrHandlerGetArg C.intr_handler_get_arg
func IntrHandlerGetArg(rv_int_num c.Int) c.Pointer

/**
 * @brief Get the current type of an interrupt
 *
 * Get the current type of a particular interrupt (level or edge). An interrupt's
 * type can be set by calling esprv_int_set_type().
 *
 * @param intr_num Interrupt number
 * @return Interrupt type
 */
//go:linkname EsprvIntGetType C.esprv_int_get_type
func EsprvIntGetType(intr_num c.Int) IntrType

/**
 * @brief Get the current priority of an interrupt
 *
 * Get the current priority of an interrupt.
 *
 * @param rv_int_num CPU interrupt number
 * @return Interrupt priority level, 1 to 7
 */
//go:linkname EsprvIntGetPriority C.esprv_int_get_priority
func EsprvIntGetPriority(rv_int_num c.Int) c.Int

/**
 * @brief Get interrupt unmask
 * @param none
 * @return uint32_t interrupt unmask
 */
//go:linkname EsprvGetInterruptUnmask C.esprv_get_interrupt_unmask
func EsprvGetInterruptUnmask() c.Uint32T

/**
 * @brief Check if the given interrupt is hardware vectored
 *
 * @param rv_int_num Interrupt number
 *
 * @return true if the interrupt is vectored, false if it is not.
 */
//go:linkname EsprvIntIsVectored C.esprv_int_is_vectored
func EsprvIntIsVectored(rv_int_num c.Int) bool
