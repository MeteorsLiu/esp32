package esp_app_format

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/**
 * @brief Description about application.
 */

type EspAppDescT struct {
	MagicWord          c.Uint32T
	SecureVersion      c.Uint32T
	Reserv1            [2]c.Uint32T
	Version            [32]c.Char
	ProjectName        [32]c.Char
	Time               [16]c.Char
	Date               [16]c.Char
	IdfVer             [32]c.Char
	AppElfSha256       [32]c.Uint8T
	MinEfuseBlkRevFull c.Uint16T
	MaxEfuseBlkRevFull c.Uint16T
	MmuPageSize        c.Uint8T
	Reserv3            [3]c.Uint8T
	Reserv2            [18]c.Uint32T
}

/**
 * @brief   Return esp_app_desc structure. This structure includes app version.
 *
 * Return description for running app.
 * @return Pointer to esp_app_desc structure.
 */
//go:linkname EspAppGetDescription C.esp_app_get_description
func EspAppGetDescription() *EspAppDescT

/**
 * @brief   Fill the provided buffer with SHA256 of the ELF file, formatted as hexadecimal, null-terminated.
 * If the buffer size is not sufficient to fit the entire SHA256 in hex plus a null terminator,
 * the largest possible number of bytes will be written followed by a null.
 * @param dst   Destination buffer
 * @param size  Size of the buffer
 * @return      Number of bytes written to dst (including null terminator)
 */
//go:linkname EspAppGetElfSha256 C.esp_app_get_elf_sha256
func EspAppGetElfSha256(dst *c.Char, size c.SizeT) c.Int
