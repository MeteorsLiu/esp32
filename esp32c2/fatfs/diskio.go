package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const STA_NOINIT = 0x01
const STA_NODISK = 0x02
const STA_PROTECT = 0x04
const CTRL_SYNC = 0
const GET_SECTOR_COUNT = 1
const GET_SECTOR_SIZE = 2
const GET_BLOCK_SIZE = 3
const CTRL_TRIM = 4
const CTRL_POWER = 5
const CTRL_LOCK = 6
const CTRL_EJECT = 7
const CTRL_FORMAT = 8
const MMC_GET_TYPE = 10
const MMC_GET_CSD = 11
const MMC_GET_CID = 12
const MMC_GET_OCR = 13
const MMC_GET_SDSTAT = 14
const ISDIO_READ = 55
const ISDIO_WRITE = 56
const ISDIO_MRITE = 57
const ATA_GET_REV = 20
const ATA_GET_MODEL = 21
const ATA_GET_SN = 22

type DSTATUS BYTE
type DRESULT c.Int

const (
	RES_OK     DRESULT = 0
	RES_ERROR  DRESULT = 1
	RES_WRPRT  DRESULT = 2
	RES_NOTRDY DRESULT = 3
	RES_PARERR DRESULT = 4
)

/*---------------------------------------*/
/* Prototypes for disk control functions */
// llgo:link BYTE.FfDiskInitialize C.ff_disk_initialize
func (recv_ BYTE) FfDiskInitialize() DSTATUS {
	return 0
}

// llgo:link BYTE.FfDiskStatus C.ff_disk_status
func (recv_ BYTE) FfDiskStatus() DSTATUS {
	return 0
}

// llgo:link BYTE.FfDiskRead C.ff_disk_read
func (recv_ BYTE) FfDiskRead(buff *BYTE, sector LBAT, count UINT) DRESULT {
	return 0
}

// llgo:link BYTE.FfDiskWrite C.ff_disk_write
func (recv_ BYTE) FfDiskWrite(buff *BYTE, sector LBAT, count UINT) DRESULT {
	return 0
}

// llgo:link BYTE.FfDiskIoctl C.ff_disk_ioctl
func (recv_ BYTE) FfDiskIoctl(cmd BYTE, buff c.Pointer) DRESULT {
	return 0
}
