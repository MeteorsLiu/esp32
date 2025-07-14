package fatfs

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const FF_DEFINED = 80286
const FF_INTDEF = 2
const FA_READ = 0x01
const FA_WRITE = 0x02
const FA_OPEN_EXISTING = 0x00
const FA_CREATE_NEW = 0x04
const FA_CREATE_ALWAYS = 0x08
const FA_OPEN_ALWAYS = 0x10
const FA_OPEN_APPEND = 0x30
const FM_FAT = 0x01
const FM_FAT32 = 0x02
const FM_EXFAT = 0x04
const FM_ANY = 0x07
const FM_SFD = 0x08
const FS_FAT12 = 1
const FS_FAT16 = 2
const FS_FAT32 = 3
const FS_EXFAT = 4
const AM_RDO = 0x01
const AM_HID = 0x02
const AM_SYS = 0x04
const AM_DIR = 0x10
const AM_ARC = 0x20

type UINT c.Uint
type BYTE c.Char
type WORD c.Uint16T
type DWORD c.Uint32T
type QWORD c.Uint64T
type WCHAR WORD
type FSIZET DWORD
type LBAT DWORD
type TCHAR c.Char

/* Definitions of volume management */

type PARTITION struct {
	Pd BYTE
	Pt BYTE
}

/* Filesystem object structure (FATFS) */

type FATFS struct {
	FsType   BYTE
	Pdrv     BYTE
	Ldrv     BYTE
	NFats    BYTE
	Wflag    BYTE
	FsiFlag  BYTE
	Id       WORD
	NRootdir WORD
	Csize    WORD
	Ssize    WORD
	LastClst DWORD
	FreeClst DWORD
	NFatent  DWORD
	Fsize    DWORD
	Volbase  LBAT
	Fatbase  LBAT
	Dirbase  LBAT
	Database LBAT
	Winsect  LBAT
	Win      [4096]BYTE
}

/* Object ID and allocation information (FFOBJID) */

type FFOBJID struct {
	Fs      *FATFS
	Id      WORD
	Attr    BYTE
	Stat    BYTE
	Sclust  DWORD
	Objsize FSIZET
}

/* File object structure (FIL) */

type FIL struct {
	Obj     FFOBJID
	Flag    BYTE
	Err     BYTE
	Fptr    FSIZET
	Clust   DWORD
	Sect    LBAT
	DirSect LBAT
	DirPtr  *BYTE
	Buf     [4096]BYTE
}

/* Directory object structure (FF_DIR) */

type FFDIR struct {
	Obj   FFOBJID
	Dptr  DWORD
	Clust DWORD
	Sect  LBAT
	Dir   *BYTE
	Fn    [12]BYTE
}

/* File information structure (FILINFO) */

type FILINFO struct {
	Fsize   FSIZET
	Fdate   WORD
	Ftime   WORD
	Fattrib BYTE
	Fname   [13]TCHAR
}

/* Format parameter structure (MKFS_PARM) */

type MKFSPARM struct {
	Fmt    BYTE
	NFat   BYTE
	Align  UINT
	NRoot  UINT
	AuSize DWORD
}
type FRESULT c.Int

const (
	FR_OK                  FRESULT = 0
	FR_DISK_ERR            FRESULT = 1
	FR_INT_ERR             FRESULT = 2
	FR_NOT_READY           FRESULT = 3
	FR_NO_FILE             FRESULT = 4
	FR_NO_PATH             FRESULT = 5
	FR_INVALID_NAME        FRESULT = 6
	FR_DENIED              FRESULT = 7
	FR_EXIST               FRESULT = 8
	FR_INVALID_OBJECT      FRESULT = 9
	FR_WRITE_PROTECTED     FRESULT = 10
	FR_INVALID_DRIVE       FRESULT = 11
	FR_NOT_ENABLED         FRESULT = 12
	FR_NO_FILESYSTEM       FRESULT = 13
	FR_MKFS_ABORTED        FRESULT = 14
	FR_TIMEOUT             FRESULT = 15
	FR_LOCKED              FRESULT = 16
	FR_NOT_ENOUGH_CORE     FRESULT = 17
	FR_TOO_MANY_OPEN_FILES FRESULT = 18
	FR_INVALID_PARAMETER   FRESULT = 19
)

/*--------------------------------------------------------------*/
/* FatFs Module Application Interface                           */
/*--------------------------------------------------------------*/
// llgo:link (*FIL).FOpen C.f_open
func (recv_ *FIL) FOpen(path *TCHAR, mode BYTE) FRESULT {
	return 0
}

// llgo:link (*FIL).FClose C.f_close
func (recv_ *FIL) FClose() FRESULT {
	return 0
}

// llgo:link (*FIL).FRead C.f_read
func (recv_ *FIL) FRead(buff c.Pointer, btr UINT, br *UINT) FRESULT {
	return 0
}

// llgo:link (*FIL).FWrite C.f_write
func (recv_ *FIL) FWrite(buff c.Pointer, btw UINT, bw *UINT) FRESULT {
	return 0
}

// llgo:link (*FIL).FLseek C.f_lseek
func (recv_ *FIL) FLseek(ofs FSIZET) FRESULT {
	return 0
}

// llgo:link (*FIL).FTruncate C.f_truncate
func (recv_ *FIL) FTruncate() FRESULT {
	return 0
}

// llgo:link (*FIL).FSync C.f_sync
func (recv_ *FIL) FSync() FRESULT {
	return 0
}

// llgo:link (*FFDIR).FOpendir C.f_opendir
func (recv_ *FFDIR) FOpendir(path *TCHAR) FRESULT {
	return 0
}

// llgo:link (*FFDIR).FClosedir C.f_closedir
func (recv_ *FFDIR) FClosedir() FRESULT {
	return 0
}

// llgo:link (*FFDIR).FReaddir C.f_readdir
func (recv_ *FFDIR) FReaddir(fno *FILINFO) FRESULT {
	return 0
}

// llgo:link (*TCHAR).FMkdir C.f_mkdir
func (recv_ *TCHAR) FMkdir() FRESULT {
	return 0
}

// llgo:link (*TCHAR).FUnlink C.f_unlink
func (recv_ *TCHAR) FUnlink() FRESULT {
	return 0
}

// llgo:link (*TCHAR).FRename C.f_rename
func (recv_ *TCHAR) FRename(path_new *TCHAR) FRESULT {
	return 0
}

// llgo:link (*TCHAR).FStat C.f_stat
func (recv_ *TCHAR) FStat(fno *FILINFO) FRESULT {
	return 0
}

// llgo:link (*TCHAR).FChmod C.f_chmod
func (recv_ *TCHAR) FChmod(attr BYTE, mask BYTE) FRESULT {
	return 0
}

// llgo:link (*TCHAR).FUtime C.f_utime
func (recv_ *TCHAR) FUtime(fno *FILINFO) FRESULT {
	return 0
}

// llgo:link (*TCHAR).FGetfree C.f_getfree
func (recv_ *TCHAR) FGetfree(nclst *DWORD, fatfs **FATFS) FRESULT {
	return 0
}

// llgo:link (*FIL).FExpand C.f_expand
func (recv_ *FIL) FExpand(fsz FSIZET, opt BYTE) FRESULT {
	return 0
}

// llgo:link (*FATFS).FMount C.f_mount
func (recv_ *FATFS) FMount(path *TCHAR, opt BYTE) FRESULT {
	return 0
}

// llgo:link (*TCHAR).FMkfs C.f_mkfs
func (recv_ *TCHAR) FMkfs(opt *MKFSPARM, work c.Pointer, len UINT) FRESULT {
	return 0
}

// llgo:link BYTE.FFdisk C.f_fdisk
func (recv_ BYTE) FFdisk(ptbl *LBAT, work c.Pointer) FRESULT {
	return 0
}

/* RTC function (provided by user) */
//go:linkname GetFattime C.get_fattime
func GetFattime() DWORD

/* O/S dependent functions (samples available in ffsystem.c) */
//go:linkname FfMutexCreate C.ff_mutex_create
func FfMutexCreate(vol c.Int) c.Int

//go:linkname FfMutexDelete C.ff_mutex_delete
func FfMutexDelete(vol c.Int)

//go:linkname FfMutexTake C.ff_mutex_take
func FfMutexTake(vol c.Int) c.Int

//go:linkname FfMutexGive C.ff_mutex_give
func FfMutexGive(vol c.Int)
