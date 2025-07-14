package esp_eth

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief BMCR(Basic Mode Control Register)
 *
 */

type BmcrRegT struct {
	Val c.Uint32T
}

/**
 * @brief BMSR(Basic Mode Status Register)
 *
 */

type BmsrRegT struct {
	Val c.Uint32T
}

/**
 * @brief PHYIDR1(PHY Identifier Register 1)
 *
 */

type Phyidr1RegT struct {
	Val c.Uint32T
}

/**
 * @brief PHYIDR2(PHY Identifier Register 2)
 *
 */

type Phyidr2RegT struct {
	Val c.Uint32T
}

/**
 * @brief ANAR(Auto-Negotiation Advertisement Register)
 *
 */

type AnarRegT struct {
	Val c.Uint32T
}

/**
 * @brief ANLPAR(Auto-Negotiation Link Partner Ability Register)
 *
 */

type AnlparRegT struct {
	Val c.Uint32T
}

/**
 * @brief ANER(Auto-Negotiate Expansion Register)
 *
 */

type AnerRegT struct {
	Val c.Uint32T
}

/**
 * @brief MMD Access control register
 *
 */

type MmdctrlRegT struct {
	Val c.Uint32T
}

/**
 * @brief MMD Access address register
 *
 */

type MmdadRegT struct {
	Val c.Uint32T
}
