package hal

import _ "unsafe"

type DwGdmaDevT struct {
	Unused [8]uint8
}
type DwGdmaSocHandleT *DwGdmaDevT

/**
 * @brief DW_GDMA HAL driver context
 */

type DwGdmaHalContextT struct {
	Dev DwGdmaSocHandleT
}

/**
 * @brief DW_GDMA HAL driver configuration
 */

type DwGdmaHalConfigT struct {
	Unused [8]uint8
}
