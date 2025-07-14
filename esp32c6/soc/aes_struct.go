package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: key register */
/** Type of key_0 register
 *  Key material key_0 configure register
 */

type AesKey0RegT struct {
	Val c.Uint32T
}

/** Type of key_1 register
 *  Key material key_1 configure register
 */

type AesKey1RegT struct {
	Val c.Uint32T
}

/** Type of key_2 register
 *  Key material key_2 configure register
 */

type AesKey2RegT struct {
	Val c.Uint32T
}

/** Type of key_3 register
 *  Key material key_3 configure register
 */

type AesKey3RegT struct {
	Val c.Uint32T
}

/** Type of key_4 register
 *  Key material key_4 configure register
 */

type AesKey4RegT struct {
	Val c.Uint32T
}

/** Type of key_5 register
 *  Key material key_5 configure register
 */

type AesKey5RegT struct {
	Val c.Uint32T
}

/** Type of key_6 register
 *  Key material key_6 configure register
 */

type AesKey6RegT struct {
	Val c.Uint32T
}

/** Type of key_7 register
 *  Key material key_7 configure register
 */

type AesKey7RegT struct {
	Val c.Uint32T
}

/** Group: text in register */
/** Type of text_in_0 register
 *  source text material text_in_0 configure register
 */

type AesTextIn0RegT struct {
	Val c.Uint32T
}

/** Type of text_in_1 register
 *  source text material text_in_1 configure register
 */

type AesTextIn1RegT struct {
	Val c.Uint32T
}

/** Type of text_in_2 register
 *  source text material text_in_2 configure register
 */

type AesTextIn2RegT struct {
	Val c.Uint32T
}

/** Type of text_in_3 register
 *  source text material text_in_3 configure register
 */

type AesTextIn3RegT struct {
	Val c.Uint32T
}

/** Group: text out register */
/** Type of text_out_0 register
 *  result text material text_out_0 configure register
 */

type AesTextOut0RegT struct {
	Val c.Uint32T
}

/** Type of text_out_1 register
 *  result text material text_out_1 configure register
 */

type AesTextOut1RegT struct {
	Val c.Uint32T
}

/** Type of text_out_2 register
 *  result text material text_out_2 configure register
 */

type AesTextOut2RegT struct {
	Val c.Uint32T
}

/** Type of text_out_3 register
 *  result text material text_out_3 configure register
 */

type AesTextOut3RegT struct {
	Val c.Uint32T
}

/** Group: Configuration register */
/** Type of mode register
 *  AES Mode register
 */

type AesModeRegT struct {
	Val c.Uint32T
}

/** Type of endian register
 *  AES Endian configure register
 */

type AesEndianRegT struct {
	Val c.Uint32T
}

/** Type of block_mode register
 *  AES cipher block mode register
 */

type AesBlockModeRegT struct {
	Val c.Uint32T
}

/** Type of block_num register
 *  AES block number register
 */

type AesBlockNumRegT struct {
	Val c.Uint32T
}

/** Type of inc_sel register
 *  Standard incrementing function configure register
 */

type AesIncSelRegT struct {
	Val c.Uint32T
}

/** Type of aad_block_num register
 *  Additional Authential Data block number register
 */

type AesAadBlockNumRegT struct {
	Val c.Uint32T
}

/** Type of remainder_bit_num register
 *  AES remainder bit number register
 */

type AesRemainderBitNumRegT struct {
	Val c.Uint32T
}

/** Group: Control/Status register */
/** Type of trigger register
 *  AES trigger register
 */

type AesTriggerRegT struct {
	Val c.Uint32T
}

/** Type of state register
 *  AES state register
 */

type AesStateRegT struct {
	Val c.Uint32T
}

/** Type of dma_enable register
 *  DMA-AES working mode register
 */

type AesDmaEnableRegT struct {
	Val c.Uint32T
}

/** Type of continue register
 *  AES continue register
 */

type AesContinueRegT struct {
	Val c.Uint32T
}

/** Type of dma_exit register
 *  AES-DMA exit config
 */

type AesDmaExitRegT struct {
	Val c.Uint32T
}

/** Group: interrupt register */
/** Type of int_clear register
 *  AES Interrupt clear register
 */

type AesIntClearRegT struct {
	Val c.Uint32T
}

/** Type of int_ena register
 *  AES Interrupt enable register
 */

type AesIntEnaRegT struct {
	Val c.Uint32T
}

/** Group: Version control register */
/** Type of date register
 *  AES version control register
 */

type AesDateRegT struct {
	Val c.Uint32T
}

type AesDevT struct {
	Key0            AesKey0RegT
	Key1            AesKey1RegT
	Key2            AesKey2RegT
	Key3            AesKey3RegT
	Key4            AesKey4RegT
	Key5            AesKey5RegT
	Key6            AesKey6RegT
	Key7            AesKey7RegT
	TextIn0         AesTextIn0RegT
	TextIn1         AesTextIn1RegT
	TextIn2         AesTextIn2RegT
	TextIn3         AesTextIn3RegT
	TextOut0        AesTextOut0RegT
	TextOut1        AesTextOut1RegT
	TextOut2        AesTextOut2RegT
	TextOut3        AesTextOut3RegT
	Mode            AesModeRegT
	Endian          AesEndianRegT
	Trigger         AesTriggerRegT
	State           AesStateRegT
	Iv              [4]c.Uint32T
	H               [4]c.Uint32T
	J0              [4]c.Uint32T
	T0              [4]c.Uint32T
	DmaEnable       AesDmaEnableRegT
	BlockMode       AesBlockModeRegT
	BlockNum        AesBlockNumRegT
	IncSel          AesIncSelRegT
	AadBlockNum     AesAadBlockNumRegT
	RemainderBitNum AesRemainderBitNumRegT
	Conti           AesContinueRegT
	IntClear        AesIntClearRegT
	IntEna          AesIntEnaRegT
	Date            AesDateRegT
	DmaExit         AesDmaExitRegT
}
