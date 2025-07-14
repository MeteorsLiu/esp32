package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

/** Group: Configuration registers */
/** Type of slcconf0 register
 *  ******* Description ***********
 */

type SdioSlcconf0RegT struct {
	Val c.Uint32T
}

/** Type of slc0rxfifo_push register
 *  ******* Description ***********
 */

type SdioSlc0rxfifoPushRegT struct {
	Val c.Uint32T
}

/** Type of slc1rxfifo_push register
 *  reserved
 */

type SdioSlc1rxfifoPushRegT struct {
	Val c.Uint32T
}

/** Type of slc0rx_link register
 *  reserved
 */

type SdioSlc0rxLinkRegT struct {
	Val c.Uint32T
}

/** Type of slc0rx_link_addr register
 *  reserved
 */

type SdioSlc0rxLinkAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc0tx_link register
 *  reserved
 */

type SdioSlc0txLinkRegT struct {
	Val c.Uint32T
}

/** Type of slc0tx_link_addr register
 *  reserved
 */

type SdioSlc0txLinkAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc1rx_link register
 *  reserved
 */

type SdioSlc1rxLinkRegT struct {
	Val c.Uint32T
}

/** Type of slc1rx_link_addr register
 *  reserved
 */

type SdioSlc1rxLinkAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc1tx_link register
 *  reserved
 */

type SdioSlc1txLinkRegT struct {
	Val c.Uint32T
}

/** Type of slc1tx_link_addr register
 *  reserved
 */

type SdioSlc1txLinkAddrRegT struct {
	Val c.Uint32T
}

/** Type of slcintvec_tohost register
 *  reserved
 */

type SdioSlcintvecTohostRegT struct {
	Val c.Uint32T
}

/** Type of slc0token0 register
 *  reserved
 */

type SdioSlc0token0RegT struct {
	Val c.Uint32T
}

/** Type of slc0token1 register
 *  reserved
 */

type SdioSlc0token1RegT struct {
	Val c.Uint32T
}

/** Type of slc1token0 register
 *  ******* Description ***********
 */

type SdioSlc1token0RegT struct {
	Val c.Uint32T
}

/** Type of slc1token1 register
 *  reserved
 */

type SdioSlc1token1RegT struct {
	Val c.Uint32T
}

/** Type of slcconf1 register
 *  reserved
 */

type SdioSlcconf1RegT struct {
	Val c.Uint32T
}

/** Type of slcbridge_conf register
 *  ******* Description ***********
 */

type SdioSlcbridgeConfRegT struct {
	Val c.Uint32T
}

/** Type of slc0_to_eof_des_addr register
 *  reserved
 */

type SdioSlc0ToEofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_tx_eof_des_addr register
 *  reserved
 */

type SdioSlc0TxEofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_to_eof_bfr_des_addr register
 *  reserved
 */

type SdioSlc0ToEofBfrDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc1_to_eof_des_addr register
 *  reserved
 */

type SdioSlc1ToEofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc1_tx_eof_des_addr register
 *  reserved
 */

type SdioSlc1TxEofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc1_to_eof_bfr_des_addr register
 *  reserved
 */

type SdioSlc1ToEofBfrDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc_rx_dscr_conf register
 *  reserved
 */

type SdioSlcRxDscrConfRegT struct {
	Val c.Uint32T
}

/** Type of slc_tx_dscr_conf register
 *  reserved
 */

type SdioSlcTxDscrConfRegT struct {
	Val c.Uint32T
}

/** Type of slc0_len_conf register
 *  reserved
 */

type SdioSlc0LenConfRegT struct {
	Val c.Uint32T
}

/** Type of slc0_txpkt_h_dscr register
 *  reserved
 */

type SdioSlc0TxpktHDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_txpkt_e_dscr register
 *  reserved
 */

type SdioSlc0TxpktEDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_rxpkt_h_dscr register
 *  reserved
 */

type SdioSlc0RxpktHDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_rxpkt_e_dscr register
 *  reserved
 */

type SdioSlc0RxpktEDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_txpktu_h_dscr register
 *  reserved
 */

type SdioSlc0TxpktuHDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_txpktu_e_dscr register
 *  reserved
 */

type SdioSlc0TxpktuEDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_rxpktu_h_dscr register
 *  reserved
 */

type SdioSlc0RxpktuHDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_rxpktu_e_dscr register
 *  reserved
 */

type SdioSlc0RxpktuEDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc_seq_position register
 *  reserved
 */

type SdioSlcSeqPositionRegT struct {
	Val c.Uint32T
}

/** Type of slc0_dscr_rec_conf register
 *  reserved
 */

type SdioSlc0DscrRecConfRegT struct {
	Val c.Uint32T
}

/** Type of slc_sdio_crc_st1 register
 *  reserved
 */

type SdioSlcSdioCrcSt1RegT struct {
	Val c.Uint32T
}

/** Type of slc0_len_lim_conf register
 *  ******* Description ***********
 */

type SdioSlc0LenLimConfRegT struct {
	Val c.Uint32T
}

/** Type of slc0_tx_sharemem_start register
 *  reserved
 */

type SdioSlc0TxSharememStartRegT struct {
	Val c.Uint32T
}

/** Type of slc0_tx_sharemem_end register
 *  reserved
 */

type SdioSlc0TxSharememEndRegT struct {
	Val c.Uint32T
}

/** Type of slc0_rx_sharemem_start register
 *  reserved
 */

type SdioSlc0RxSharememStartRegT struct {
	Val c.Uint32T
}

/** Type of slc0_rx_sharemem_end register
 *  reserved
 */

type SdioSlc0RxSharememEndRegT struct {
	Val c.Uint32T
}

/** Type of slc1_tx_sharemem_start register
 *  reserved
 */

type SdioSlc1TxSharememStartRegT struct {
	Val c.Uint32T
}

/** Type of slc1_tx_sharemem_end register
 *  reserved
 */

type SdioSlc1TxSharememEndRegT struct {
	Val c.Uint32T
}

/** Type of slc1_rx_sharemem_start register
 *  reserved
 */

type SdioSlc1RxSharememStartRegT struct {
	Val c.Uint32T
}

/** Type of slc1_rx_sharemem_end register
 *  reserved
 */

type SdioSlc1RxSharememEndRegT struct {
	Val c.Uint32T
}

/** Type of hda_tx_sharemem_start register
 *  reserved
 */

type SdioHdaTxSharememStartRegT struct {
	Val c.Uint32T
}

/** Type of hda_rx_sharemem_start register
 *  reserved
 */

type SdioHdaRxSharememStartRegT struct {
	Val c.Uint32T
}

/** Type of slc_burst_len register
 *  reserved
 */

type SdioSlcBurstLenRegT struct {
	Val c.Uint32T
}

/** Type of slcid register
 *  ******* Description ***********
 */

type SdioSlcidRegT struct {
	Val c.Uint32T
}

/** Group: Interrupt registers */
/** Type of slc0int_raw register
 *  ******* Description ***********
 */

type SdioSlc0intRawRegT struct {
	Val c.Uint32T
}

/** Type of slc0int_st register
 *  ******* Description ***********
 */

type SdioSlc0intStRegT struct {
	Val c.Uint32T
}

/** Type of slc0int_ena register
 *  ******* Description ***********
 */

type SdioSlc0intEnaRegT struct {
	Val c.Uint32T
}

/** Type of slc0int_clr register
 *  ******* Description ***********
 */

type SdioSlc0intClrRegT struct {
	Val c.Uint32T
}

/** Type of slc1int_raw register
 *  reserved
 */

type SdioSlc1intRawRegT struct {
	Val c.Uint32T
}

/** Type of slc1int_st register
 *  reserved
 */

type SdioSlc1intStRegT struct {
	Val c.Uint32T
}

/** Type of slc1int_ena register
 *  reserved
 */

type SdioSlc1intEnaRegT struct {
	Val c.Uint32T
}

/** Type of slc1int_clr register
 *  reserved
 */

type SdioSlc1intClrRegT struct {
	Val c.Uint32T
}

/** Type of slc0int_st1 register
 *  reserved
 */

type SdioSlc0intSt1RegT struct {
	Val c.Uint32T
}

/** Type of slc0int_ena1 register
 *  reserved
 */

type SdioSlc0intEna1RegT struct {
	Val c.Uint32T
}

/** Type of slc1int_st1 register
 *  reserved
 */

type SdioSlc1intSt1RegT struct {
	Val c.Uint32T
}

/** Type of slc1int_ena1 register
 *  reserved
 */

type SdioSlc1intEna1RegT struct {
	Val c.Uint32T
}

/** Group: Status registers */
/** Type of slcrx_status register
 *  ******* Description ***********
 */

type SdioSlcrxStatusRegT struct {
	Val c.Uint32T
}

/** Type of slctx_status register
 *  ******* Description ***********
 */

type SdioSlctxStatusRegT struct {
	Val c.Uint32T
}

/** Type of slc0_state0 register
 *  reserved
 */

type SdioSlc0State0RegT struct {
	Val c.Uint32T
}

/** Type of slc0_state1 register
 *  ******* Description ***********
 */

type SdioSlc0State1RegT struct {
	Val c.Uint32T
}

/** Type of slc1_state0 register
 *  ******* Description ***********
 */

type SdioSlc1State0RegT struct {
	Val c.Uint32T
}

/** Type of slc1_state1 register
 *  ******* Description ***********
 */

type SdioSlc1State1RegT struct {
	Val c.Uint32T
}

/** Type of slc_sdio_st register
 *  reserved
 */

type SdioSlcSdioStRegT struct {
	Val c.Uint32T
}

/** Type of slc0_txlink_dscr register
 *  ******* Description ***********
 */

type SdioSlc0TxlinkDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_txlink_dscr_bf0 register
 *  ******* Description ***********
 */

type SdioSlc0TxlinkDscrBf0RegT struct {
	Val c.Uint32T
}

/** Type of slc0_txlink_dscr_bf1 register
 *  reserved
 */

type SdioSlc0TxlinkDscrBf1RegT struct {
	Val c.Uint32T
}

/** Type of slc0_rxlink_dscr register
 *  ******* Description ***********
 */

type SdioSlc0RxlinkDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_rxlink_dscr_bf0 register
 *  ******* Description ***********
 */

type SdioSlc0RxlinkDscrBf0RegT struct {
	Val c.Uint32T
}

/** Type of slc0_rxlink_dscr_bf1 register
 *  reserved
 */

type SdioSlc0RxlinkDscrBf1RegT struct {
	Val c.Uint32T
}

/** Type of slc1_txlink_dscr register
 *  reserved
 */

type SdioSlc1TxlinkDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc1_txlink_dscr_bf0 register
 *  reserved
 */

type SdioSlc1TxlinkDscrBf0RegT struct {
	Val c.Uint32T
}

/** Type of slc1_txlink_dscr_bf1 register
 *  reserved
 */

type SdioSlc1TxlinkDscrBf1RegT struct {
	Val c.Uint32T
}

/** Type of slc1_rxlink_dscr register
 *  ******* Description ***********
 */

type SdioSlc1RxlinkDscrRegT struct {
	Val c.Uint32T
}

/** Type of slc1_rxlink_dscr_bf0 register
 *  ******* Description ***********
 */

type SdioSlc1RxlinkDscrBf0RegT struct {
	Val c.Uint32T
}

/** Type of slc1_rxlink_dscr_bf1 register
 *  reserved
 */

type SdioSlc1RxlinkDscrBf1RegT struct {
	Val c.Uint32T
}

/** Type of slc0_tx_erreof_des_addr register
 *  reserved
 */

type SdioSlc0TxErreofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc1_tx_erreof_des_addr register
 *  reserved
 */

type SdioSlc1TxErreofDesAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc_token_lat register
 *  reserved
 */

type SdioSlcTokenLatRegT struct {
	Val c.Uint32T
}

/** Type of slc_cmd_infor0 register
 *  reserved
 */

type SdioSlcCmdInfor0RegT struct {
	Val c.Uint32T
}

/** Type of slc_cmd_infor1 register
 *  reserved
 */

type SdioSlcCmdInfor1RegT struct {
	Val c.Uint32T
}

/** Type of slc0_length register
 *  reserved
 */

type SdioSlc0LengthRegT struct {
	Val c.Uint32T
}

/** Type of slc_sdio_crc_st0 register
 *  reserved
 */

type SdioSlcSdioCrcSt0RegT struct {
	Val c.Uint32T
}

/** Type of slc0_eof_start_des register
 *  reserved
 */

type SdioSlc0EofStartDesRegT struct {
	Val c.Uint32T
}

/** Type of slc0_push_dscr_addr register
 *  ******* Description ***********
 */

type SdioSlc0PushDscrAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_done_dscr_addr register
 *  ******* Description ***********
 */

type SdioSlc0DoneDscrAddrRegT struct {
	Val c.Uint32T
}

/** Type of slc0_sub_start_des register
 *  ******* Description ***********
 */

type SdioSlc0SubStartDesRegT struct {
	Val c.Uint32T
}

/** Type of slc0_dscr_cnt register
 *  ******* Description ***********
 */

type SdioSlc0DscrCntRegT struct {
	Val c.Uint32T
}

/** Group: Debud registers */
/** Type of slc0txfifo_pop register
 *  reserved
 */

type SdioSlc0txfifoPopRegT struct {
	Val c.Uint32T
}

/** Type of slc1txfifo_pop register
 *  reserved
 */

type SdioSlc1txfifoPopRegT struct {
	Val c.Uint32T
}

/** Type of slc_ahb_test register
 *  reserved
 */

type SdioSlcAhbTestRegT struct {
	Val c.Uint32T
}

/** Group: Version registers */
/** Type of slcdate register
 *  ******* Description ***********
 */

type SdioSlcdateRegT struct {
	Val c.Uint32T
}

type SlcDevT struct {
	Slcconf0            SdioSlcconf0RegT
	Slc0intRaw          SdioSlc0intRawRegT
	Slc0intSt           SdioSlc0intStRegT
	Slc0intEna          SdioSlc0intEnaRegT
	Slc0intClr          SdioSlc0intClrRegT
	Slc1intRaw          SdioSlc1intRawRegT
	Slc1intSt           SdioSlc1intStRegT
	Slc1intEna          SdioSlc1intEnaRegT
	Slc1intClr          SdioSlc1intClrRegT
	SlcrxStatus         SdioSlcrxStatusRegT
	Slc0rxfifoPush      SdioSlc0rxfifoPushRegT
	Slc1rxfifoPush      SdioSlc1rxfifoPushRegT
	SlctxStatus         SdioSlctxStatusRegT
	Slc0txfifoPop       SdioSlc0txfifoPopRegT
	Slc1txfifoPop       SdioSlc1txfifoPopRegT
	Slc0rxLink          SdioSlc0rxLinkRegT
	Slc0rxLinkAddr      SdioSlc0rxLinkAddrRegT
	Slc0txLink          SdioSlc0txLinkRegT
	Slc0txLinkAddr      SdioSlc0txLinkAddrRegT
	Slc1rxLink          SdioSlc1rxLinkRegT
	Slc1rxLinkAddr      SdioSlc1rxLinkAddrRegT
	Slc1txLink          SdioSlc1txLinkRegT
	Slc1txLinkAddr      SdioSlc1txLinkAddrRegT
	SlcintvecTohost     SdioSlcintvecTohostRegT
	Slc0token0          SdioSlc0token0RegT
	Slc0token1          SdioSlc0token1RegT
	Slc1token0          SdioSlc1token0RegT
	Slc1token1          SdioSlc1token1RegT
	Slcconf1            SdioSlcconf1RegT
	Slc0State0          SdioSlc0State0RegT
	Slc0State1          SdioSlc0State1RegT
	Slc1State0          SdioSlc1State0RegT
	Slc1State1          SdioSlc1State1RegT
	SlcbridgeConf       SdioSlcbridgeConfRegT
	Slc0ToEofDesAddr    SdioSlc0ToEofDesAddrRegT
	Slc0TxEofDesAddr    SdioSlc0TxEofDesAddrRegT
	Slc0ToEofBfrDesAddr SdioSlc0ToEofBfrDesAddrRegT
	Slc1ToEofDesAddr    SdioSlc1ToEofDesAddrRegT
	Slc1TxEofDesAddr    SdioSlc1TxEofDesAddrRegT
	Slc1ToEofBfrDesAddr SdioSlc1ToEofBfrDesAddrRegT
	SlcAhbTest          SdioSlcAhbTestRegT
	SlcSdioSt           SdioSlcSdioStRegT
	SlcRxDscrConf       SdioSlcRxDscrConfRegT
	Slc0TxlinkDscr      SdioSlc0TxlinkDscrRegT
	Slc0TxlinkDscrBf0   SdioSlc0TxlinkDscrBf0RegT
	Slc0TxlinkDscrBf1   SdioSlc0TxlinkDscrBf1RegT
	Slc0RxlinkDscr      SdioSlc0RxlinkDscrRegT
	Slc0RxlinkDscrBf0   SdioSlc0RxlinkDscrBf0RegT
	Slc0RxlinkDscrBf1   SdioSlc0RxlinkDscrBf1RegT
	Slc1TxlinkDscr      SdioSlc1TxlinkDscrRegT
	Slc1TxlinkDscrBf0   SdioSlc1TxlinkDscrBf0RegT
	Slc1TxlinkDscrBf1   SdioSlc1TxlinkDscrBf1RegT
	Slc1RxlinkDscr      SdioSlc1RxlinkDscrRegT
	Slc1RxlinkDscrBf0   SdioSlc1RxlinkDscrBf0RegT
	Slc1RxlinkDscrBf1   SdioSlc1RxlinkDscrBf1RegT
	Slc0TxErreofDesAddr SdioSlc0TxErreofDesAddrRegT
	Slc1TxErreofDesAddr SdioSlc1TxErreofDesAddrRegT
	SlcTokenLat         SdioSlcTokenLatRegT
	SlcTxDscrConf       SdioSlcTxDscrConfRegT
	SlcCmdInfor0        SdioSlcCmdInfor0RegT
	SlcCmdInfor1        SdioSlcCmdInfor1RegT
	Slc0LenConf         SdioSlc0LenConfRegT
	Slc0Length          SdioSlc0LengthRegT
	Slc0TxpktHDscr      SdioSlc0TxpktHDscrRegT
	Slc0TxpktEDscr      SdioSlc0TxpktEDscrRegT
	Slc0RxpktHDscr      SdioSlc0RxpktHDscrRegT
	Slc0RxpktEDscr      SdioSlc0RxpktEDscrRegT
	Slc0TxpktuHDscr     SdioSlc0TxpktuHDscrRegT
	Slc0TxpktuEDscr     SdioSlc0TxpktuEDscrRegT
	Slc0RxpktuHDscr     SdioSlc0RxpktuHDscrRegT
	Slc0RxpktuEDscr     SdioSlc0RxpktuEDscrRegT
	SlcSeqPosition      SdioSlcSeqPositionRegT
	Slc0DscrRecConf     SdioSlc0DscrRecConfRegT
	SlcSdioCrcSt0       SdioSlcSdioCrcSt0RegT
	SlcSdioCrcSt1       SdioSlcSdioCrcSt1RegT
	Slc0EofStartDes     SdioSlc0EofStartDesRegT
	Slc0PushDscrAddr    SdioSlc0PushDscrAddrRegT
	Slc0DoneDscrAddr    SdioSlc0DoneDscrAddrRegT
	Slc0SubStartDes     SdioSlc0SubStartDesRegT
	Slc0DscrCnt         SdioSlc0DscrCntRegT
	Slc0LenLimConf      SdioSlc0LenLimConfRegT
	Slc0intSt1          SdioSlc0intSt1RegT
	Slc0intEna1         SdioSlc0intEna1RegT
	Slc1intSt1          SdioSlc1intSt1RegT
	Slc1intEna1         SdioSlc1intEna1RegT
	Slc0TxSharememStart SdioSlc0TxSharememStartRegT
	Slc0TxSharememEnd   SdioSlc0TxSharememEndRegT
	Slc0RxSharememStart SdioSlc0RxSharememStartRegT
	Slc0RxSharememEnd   SdioSlc0RxSharememEndRegT
	Slc1TxSharememStart SdioSlc1TxSharememStartRegT
	Slc1TxSharememEnd   SdioSlc1TxSharememEndRegT
	Slc1RxSharememStart SdioSlc1RxSharememStartRegT
	Slc1RxSharememEnd   SdioSlc1RxSharememEndRegT
	HdaTxSharememStart  SdioHdaTxSharememStartRegT
	HdaRxSharememStart  SdioHdaRxSharememStartRegT
	SlcBurstLen         SdioSlcBurstLenRegT
	Reserved180         [30]c.Uint32T
	Slcdate             SdioSlcdateRegT
	Slcid               SdioSlcidRegT
}
