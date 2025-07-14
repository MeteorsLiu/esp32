package bootloader_support

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

// llgo:type C
type BootloaderFlashReadStatusFnT func() c.Uint

// llgo:type C
type BootloaderFlashWriteStatusFnT func(c.Uint)

type BootloaderQioInfoT struct {
	Manufacturer  *c.Char
	MfgId         c.Uint8T
	FlashId       c.Uint16T
	IdMask        c.Uint16T
	ReadStatusFn  BootloaderFlashReadStatusFnT
	WriteStatusFn BootloaderFlashWriteStatusFnT
	StatusQioBit  c.Uint8T
}

/**
 * @brief Read 8 bit status using RDSR command
 *
 * @return Value of SR1.
 */
//go:linkname BootloaderReadStatus8bRdsr C.bootloader_read_status_8b_rdsr
func BootloaderReadStatus8bRdsr() c.Uint

/**
 * @brief Read 8 bit status (second byte) using RDSR2 command
 *
 * @return Value of SR2
 */
//go:linkname BootloaderReadStatus8bRdsr2 C.bootloader_read_status_8b_rdsr2
func BootloaderReadStatus8bRdsr2() c.Uint

/**
 * @brief Read 8 bit status (third byte) using RDSR3 command
 *
 * @return Value of SR3
 */
//go:linkname BootloaderReadStatus8bRdsr3 C.bootloader_read_status_8b_rdsr3
func BootloaderReadStatus8bRdsr3() c.Uint

/**
 * @brief Read 16 bit status using RDSR & RDSR2 (low and high bytes)
 *
 * @return Value of SR2#SR1.
 */
//go:linkname BootloaderReadStatus16bRdsrRdsr2 C.bootloader_read_status_16b_rdsr_rdsr2
func BootloaderReadStatus16bRdsrRdsr2() c.Uint

/**
 * @brief Write 8 bit status using WRSR
 */
//go:linkname BootloaderWriteStatus8bWrsr C.bootloader_write_status_8b_wrsr
func BootloaderWriteStatus8bWrsr(new_status c.Uint)

/**
 * @brief Write 8 bit status (second byte) using WRSR2.
 */
//go:linkname BootloaderWriteStatus8bWrsr2 C.bootloader_write_status_8b_wrsr2
func BootloaderWriteStatus8bWrsr2(new_status c.Uint)

/**
 * @brief Write 8 bit status (third byte) using WRSR3.
 */
//go:linkname BootloaderWriteStatus8bWrsr3 C.bootloader_write_status_8b_wrsr3
func BootloaderWriteStatus8bWrsr3(new_status c.Uint)

/**
 * @brief Write 16 bit status using WRSR, (both write SR1 and SR2)
 */
//go:linkname BootloaderWriteStatus16bWrsr C.bootloader_write_status_16b_wrsr
func BootloaderWriteStatus16bWrsr(new_status c.Uint)

/**
 * @brief Read 8 bit status of XM25QU64A.
 *
 * @return Value of 8 bit SR.
 */
//go:linkname BootloaderReadStatus8bXmc25qu64a C.bootloader_read_status_8b_xmc25qu64a
func BootloaderReadStatus8bXmc25qu64a() c.Uint

/**
 * @brief Write 8 bit status for XM25QU64A
 */
//go:linkname BootloaderWriteStatus8bXmc25qu64a C.bootloader_write_status_8b_xmc25qu64a
func BootloaderWriteStatus8bXmc25qu64a(new_status c.Uint)

/**
 * @brief Unlock Flash write protect.
 *        Please do not call this function in SDK.
 *
 * @note This can be overridden because it's attribute weak.
 */
//go:linkname BootloaderFlashUnlock C.bootloader_flash_unlock
func BootloaderFlashUnlock() EspErrT
