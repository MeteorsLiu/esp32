package sdmmc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

const SDMMC_GO_IDLE_DELAY_MS = 20
const SDMMC_IO_SEND_OP_COND_DELAY_MS = 10
const SDMMC_DEFAULT_CMD_TIMEOUT_MS = 1000
const SDMMC_WRITE_CMD_TIMEOUT_MS = 5000
const SDMMC_SD_DISCARD_TIMEOUT = 250
const SDMMC_SEND_OP_COND_MAX_RETRIES = 300
const SDMMC_SEND_OP_COND_MAX_ERRORS = 3
const SDMMC_SD_ERASE_ARG = 0
const SDMMC_SD_DISCARD_ARG = 1
const SDMMC_MMC_TRIM_ARG = 1
const SDMMC_MMC_DISCARD_ARG = 3
const SDMMC_FREQ_SDR104 = 208000
const SDMMC_IO_BLOCK_SIZE = 512

/* Functions to send individual commands */
// llgo:link (*SdmmcCardT).SdmmcSendCmd C.sdmmc_send_cmd
func (recv_ *SdmmcCardT) SdmmcSendCmd(cmd *SdmmcCommandT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendAppCmd C.sdmmc_send_app_cmd
func (recv_ *SdmmcCardT) SdmmcSendAppCmd(cmd *SdmmcCommandT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdGoIdleState C.sdmmc_send_cmd_go_idle_state
func (recv_ *SdmmcCardT) SdmmcSendCmdGoIdleState() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSendIfCond C.sdmmc_send_cmd_send_if_cond
func (recv_ *SdmmcCardT) SdmmcSendCmdSendIfCond(ocr c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSendOpCond C.sdmmc_send_cmd_send_op_cond
func (recv_ *SdmmcCardT) SdmmcSendCmdSendOpCond(ocr c.Uint32T, ocrp *c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdReadOcr C.sdmmc_send_cmd_read_ocr
func (recv_ *SdmmcCardT) SdmmcSendCmdReadOcr(ocrp *c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSendCid C.sdmmc_send_cmd_send_cid
func (recv_ *SdmmcCardT) SdmmcSendCmdSendCid(out_cid *SdmmcCidT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdAllSendCid C.sdmmc_send_cmd_all_send_cid
func (recv_ *SdmmcCardT) SdmmcSendCmdAllSendCid(out_raw_cid *SdmmcResponseT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSetRelativeAddr C.sdmmc_send_cmd_set_relative_addr
func (recv_ *SdmmcCardT) SdmmcSendCmdSetRelativeAddr(out_rca *c.Uint16T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSetBlocklen C.sdmmc_send_cmd_set_blocklen
func (recv_ *SdmmcCardT) SdmmcSendCmdSetBlocklen(csd *SdmmcCsdT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSwitchFunc C.sdmmc_send_cmd_switch_func
func (recv_ *SdmmcCardT) SdmmcSendCmdSwitchFunc(mode c.Uint32T, group c.Uint32T, function c.Uint32T, resp *SdmmcSwitchFuncRspT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSendCsd C.sdmmc_send_cmd_send_csd
func (recv_ *SdmmcCardT) SdmmcSendCmdSendCsd(out_csd *SdmmcCsdT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSelectCard C.sdmmc_send_cmd_select_card
func (recv_ *SdmmcCardT) SdmmcSendCmdSelectCard(rca c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSendScr C.sdmmc_send_cmd_send_scr
func (recv_ *SdmmcCardT) SdmmcSendCmdSendScr(out_scr *SdmmcScrT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSetBusWidth C.sdmmc_send_cmd_set_bus_width
func (recv_ *SdmmcCardT) SdmmcSendCmdSetBusWidth(width c.Int) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdSendStatus C.sdmmc_send_cmd_send_status
func (recv_ *SdmmcCardT) SdmmcSendCmdSendStatus(out_status *c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdCrcOnOff C.sdmmc_send_cmd_crc_on_off
func (recv_ *SdmmcCardT) SdmmcSendCmdCrcOnOff(crc_enable bool) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSendCmdNumOfWrittenBlocks C.sdmmc_send_cmd_num_of_written_blocks
func (recv_ *SdmmcCardT) SdmmcSendCmdNumOfWrittenBlocks(out_num_blocks *c.SizeT) EspErrT {
	return 0
}

/* Higher level functions */
// llgo:link (*SdmmcCardT).SdmmcEnterHigherSpeedMode C.sdmmc_enter_higher_speed_mode
func (recv_ *SdmmcCardT) SdmmcEnterHigherSpeedMode() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcEnableHsModeAndCheck C.sdmmc_enable_hs_mode_and_check
func (recv_ *SdmmcCardT) SdmmcEnableHsModeAndCheck() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcWriteSectorsDma C.sdmmc_write_sectors_dma
func (recv_ *SdmmcCardT) SdmmcWriteSectorsDma(src c.Pointer, start_block c.SizeT, block_count c.SizeT, buffer_len c.SizeT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcReadSectorsDma C.sdmmc_read_sectors_dma
func (recv_ *SdmmcCardT) SdmmcReadSectorsDma(dst c.Pointer, start_block c.SizeT, block_count c.SizeT, buffer_len c.SizeT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcGetEraseTimeoutMs C.sdmmc_get_erase_timeout_ms
func (recv_ *SdmmcCardT) SdmmcGetEraseTimeoutMs(arg c.Int, erase_size_kb c.SizeT) c.Uint32T {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSelectDriverStrength C.sdmmc_select_driver_strength
func (recv_ *SdmmcCardT) SdmmcSelectDriverStrength(driver_strength SdmmcDriverStrengthT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcSelectCurrentLimit C.sdmmc_select_current_limit
func (recv_ *SdmmcCardT) SdmmcSelectCurrentLimit(current_limit SdmmcCurrentLimitT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcDoTimingTuning C.sdmmc_do_timing_tuning
func (recv_ *SdmmcCardT) SdmmcDoTimingTuning() EspErrT {
	return 0
}

/* SD specific */
// llgo:link (*SdmmcCardT).SdmmcCheckScr C.sdmmc_check_scr
func (recv_ *SdmmcCardT) SdmmcCheckScr() EspErrT {
	return 0
}

// llgo:link SdmmcResponseT.SdmmcDecodeCid C.sdmmc_decode_cid
func (recv_ SdmmcResponseT) SdmmcDecodeCid(out_cid *SdmmcCidT) EspErrT {
	return 0
}

// llgo:link SdmmcResponseT.SdmmcDecodeCsd C.sdmmc_decode_csd
func (recv_ SdmmcResponseT) SdmmcDecodeCsd(out_csd *SdmmcCsdT) EspErrT {
	return 0
}

//go:linkname SdmmcDecodeScr C.sdmmc_decode_scr
func SdmmcDecodeScr(raw_scr *c.Uint32T, out_scr *SdmmcScrT) EspErrT

//go:linkname SdmmcDecodeSsr C.sdmmc_decode_ssr
func SdmmcDecodeSsr(raw_ssr *c.Uint32T, out_ssr *SdmmcSsrT) EspErrT

// llgo:link (*SdmmcCardT).SdmmcSdGetEraseTimeoutMs C.sdmmc_sd_get_erase_timeout_ms
func (recv_ *SdmmcCardT) SdmmcSdGetEraseTimeoutMs(arg c.Int, erase_size_kb c.SizeT) c.Uint32T {
	return 0
}

/* SDIO specific */
// llgo:link (*SdmmcCardT).SdmmcIoReset C.sdmmc_io_reset
func (recv_ *SdmmcCardT) SdmmcIoReset() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcIoEnableHsMode C.sdmmc_io_enable_hs_mode
func (recv_ *SdmmcCardT) SdmmcIoEnableHsMode() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcIoSendOpCond C.sdmmc_io_send_op_cond
func (recv_ *SdmmcCardT) SdmmcIoSendOpCond(ocr c.Uint32T, ocrp *c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcIoRwDirect C.sdmmc_io_rw_direct
func (recv_ *SdmmcCardT) SdmmcIoRwDirect(function c.Int, reg c.Uint32T, arg c.Uint32T, byte *c.Uint8T) EspErrT {
	return 0
}

// Requirement to `data` and `size` when using SDMMC host:
// Buffer pointer (`data`) needs to be aligned to 4 byte boundary, and also cache line size if the buffer is behind the
// cache, unless `SDMMC_HOST_FLAG_ALLOC_ALIGNED_BUF` flag is set when calling `sdmmc_card_init`. This flag is mandory
// when the buffer is behind the cache in byte mode.
// llgo:link (*SdmmcCardT).SdmmcIoRwExtended C.sdmmc_io_rw_extended
func (recv_ *SdmmcCardT) SdmmcIoRwExtended(function c.Int, reg c.Uint32T, arg c.Int, data c.Pointer, size c.SizeT) EspErrT {
	return 0
}

/* MMC specific */
// llgo:link (*SdmmcCardT).SdmmcMmcSendExtCsdData C.sdmmc_mmc_send_ext_csd_data
func (recv_ *SdmmcCardT) SdmmcMmcSendExtCsdData(out_data c.Pointer, datalen c.SizeT, buffer_len c.SizeT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcMmcSwitch C.sdmmc_mmc_switch
func (recv_ *SdmmcCardT) SdmmcMmcSwitch(set c.Uint8T, index c.Uint8T, value c.Uint8T) EspErrT {
	return 0
}

//go:linkname SdmmcMmcDecodeCid C.sdmmc_mmc_decode_cid
func SdmmcMmcDecodeCid(mmc_ver c.Int, resp SdmmcResponseT, out_cid *SdmmcCidT) EspErrT

// llgo:link SdmmcResponseT.SdmmcMmcDecodeCsd C.sdmmc_mmc_decode_csd
func (recv_ SdmmcResponseT) SdmmcMmcDecodeCsd(out_csd *SdmmcCsdT) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcMmcEnableHsMode C.sdmmc_mmc_enable_hs_mode
func (recv_ *SdmmcCardT) SdmmcMmcEnableHsMode() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcMmcGetEraseTimeoutMs C.sdmmc_mmc_get_erase_timeout_ms
func (recv_ *SdmmcCardT) SdmmcMmcGetEraseTimeoutMs(arg c.Int, erase_size_kb c.SizeT) c.Uint32T {
	return 0
}

/* Parts of card initialization flow */
// llgo:link (*SdmmcCardT).SdmmcInitSdIfCond C.sdmmc_init_sd_if_cond
func (recv_ *SdmmcCardT) SdmmcInitSdIfCond() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSelectCard C.sdmmc_init_select_card
func (recv_ *SdmmcCardT) SdmmcInitSelectCard() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitCsd C.sdmmc_init_csd
func (recv_ *SdmmcCardT) SdmmcInitCsd() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitCid C.sdmmc_init_cid
func (recv_ *SdmmcCardT) SdmmcInitCid() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitRca C.sdmmc_init_rca
func (recv_ *SdmmcCardT) SdmmcInitRca() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitMmcDecodeCid C.sdmmc_init_mmc_decode_cid
func (recv_ *SdmmcCardT) SdmmcInitMmcDecodeCid() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitOcr C.sdmmc_init_ocr
func (recv_ *SdmmcCardT) SdmmcInitOcr() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSpiCrc C.sdmmc_init_spi_crc
func (recv_ *SdmmcCardT) SdmmcInitSpiCrc() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitIo C.sdmmc_init_io
func (recv_ *SdmmcCardT) SdmmcInitIo() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcIoInitReadCardCap C.sdmmc_io_init_read_card_cap
func (recv_ *SdmmcCardT) SdmmcIoInitReadCardCap(card_cap *c.Uint8T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcIoInitCheckCardCap C.sdmmc_io_init_check_card_cap
func (recv_ *SdmmcCardT) SdmmcIoInitCheckCardCap(card_cap *c.Uint8T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdBlocklen C.sdmmc_init_sd_blocklen
func (recv_ *SdmmcCardT) SdmmcInitSdBlocklen() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdScr C.sdmmc_init_sd_scr
func (recv_ *SdmmcCardT) SdmmcInitSdScr() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdSsr C.sdmmc_init_sd_ssr
func (recv_ *SdmmcCardT) SdmmcInitSdSsr() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdWaitDataReady C.sdmmc_init_sd_wait_data_ready
func (recv_ *SdmmcCardT) SdmmcInitSdWaitDataReady() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitMmcReadExtCsd C.sdmmc_init_mmc_read_ext_csd
func (recv_ *SdmmcCardT) SdmmcInitMmcReadExtCsd() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitHostBusWidth C.sdmmc_init_host_bus_width
func (recv_ *SdmmcCardT) SdmmcInitHostBusWidth() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdBusWidth C.sdmmc_init_sd_bus_width
func (recv_ *SdmmcCardT) SdmmcInitSdBusWidth() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitIoBusWidth C.sdmmc_init_io_bus_width
func (recv_ *SdmmcCardT) SdmmcInitIoBusWidth() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitMmcBusWidth C.sdmmc_init_mmc_bus_width
func (recv_ *SdmmcCardT) SdmmcInitMmcBusWidth() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitCardHsMode C.sdmmc_init_card_hs_mode
func (recv_ *SdmmcCardT) SdmmcInitCardHsMode() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitHostFrequency C.sdmmc_init_host_frequency
func (recv_ *SdmmcCardT) SdmmcInitHostFrequency() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitMmcCheckExtCsd C.sdmmc_init_mmc_check_ext_csd
func (recv_ *SdmmcCardT) SdmmcInitMmcCheckExtCsd() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdUhs1 C.sdmmc_init_sd_uhs1
func (recv_ *SdmmcCardT) SdmmcInitSdUhs1() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdDriverStrength C.sdmmc_init_sd_driver_strength
func (recv_ *SdmmcCardT) SdmmcInitSdDriverStrength() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdCurrentLimit C.sdmmc_init_sd_current_limit
func (recv_ *SdmmcCardT) SdmmcInitSdCurrentLimit() EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcInitSdTimingTuning C.sdmmc_init_sd_timing_tuning
func (recv_ *SdmmcCardT) SdmmcInitSdTimingTuning() EspErrT {
	return 0
}

//go:linkname SdmmcFlipByteOrder C.sdmmc_flip_byte_order
func SdmmcFlipByteOrder(response *c.Uint32T, size c.SizeT)

// llgo:link (*SdmmcCardT).SdmmcFixHostFlags C.sdmmc_fix_host_flags
func (recv_ *SdmmcCardT) SdmmcFixHostFlags() EspErrT {
	return 0
}

// Use only with SDMMC mode (not SDSPI)
// llgo:link (*SdmmcCardT).SdmmcWaitForIdle C.sdmmc_wait_for_idle
func (recv_ *SdmmcCardT) SdmmcWaitForIdle(status c.Uint32T) EspErrT {
	return 0
}

// llgo:link (*SdmmcCardT).SdmmcAllocateAlignedBuf C.sdmmc_allocate_aligned_buf
func (recv_ *SdmmcCardT) SdmmcAllocateAlignedBuf() EspErrT {
	return 0
}
