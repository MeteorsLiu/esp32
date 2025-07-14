package esp_phy

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Set btbb enable for BT/ieee802154
 * @param[in] print_version enable btbb version print.
 * @return NULL
 */
//go:linkname BtBbV2InitCmplx C.bt_bb_v2_init_cmplx
func BtBbV2InitCmplx(print_version c.Int)
