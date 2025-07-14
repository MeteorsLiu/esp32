package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * Generate a linked list pointing to a (huge) buffer in an descriptor array.
 *
 * The caller should ensure there is enough size to hold the array, by calling
 * ``lldesc_get_required_num_constrained`` with the same max_desc_size argument.
 *
 * @param[out] out_desc_array Output of a descriptor array, the head should be fed to the DMA.
 * @param buffer Buffer for the descriptors to point to.
 * @param size Size (or length for TX) of the buffer
 * @param max_desc_size Maximum length of each descriptor
 * @param isrx The RX DMA may require the buffer to be word-aligned, set to true for a RX link, otherwise false.
 */
//go:linkname LldescSetupLinkConstrained C.lldesc_setup_link_constrained
func LldescSetupLinkConstrained(out_desc_array *LldescT, buffer c.Pointer, size c.Int, max_desc_size c.Int, isrx bool)

/**
 * @brief Get the received length of a linked list, until end of the link or eof.
 *
 * @param head      The head of the linked list.
 * @param[out] out_next Output of the next descriptor of the EOF descriptor. Return NULL if there's no
 *                 EOF. Can be set to NULL if next descriptor is not needed.
 * @return The accumulation of the `len` field of all descriptors until EOF or the end of the link.
 */
//go:linkname LldescGetReceivedLen C.lldesc_get_received_len
func LldescGetReceivedLen(head *LldescT, out_next **LldescT) c.Int
