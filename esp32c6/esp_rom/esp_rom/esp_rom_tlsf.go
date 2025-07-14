package esp_rom

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type PoisonFillPfuncT func(c.Pointer, c.SizeT, bool)

// llgo:type C
type PoisonCheckPfuncT func(c.Pointer, c.SizeT, bool, bool) bool

/*!
 * @brief Set the function to call for checking memory region when
 * poisoning is configured.
 *
 * @param pfunc The callback function to trigger for checking
 * the content of a memory region.
 */
//go:linkname TlsfPoisonCheckPfuncSet C.tlsf_poison_check_pfunc_set
func TlsfPoisonCheckPfuncSet(pfunc PoisonCheckPfuncT)
