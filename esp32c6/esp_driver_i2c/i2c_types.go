package esp_driver_i2c

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type I2cPortNumT c.Int
type I2cMasterStatusT c.Int

const (
	I2C_STATUS_READ      I2cMasterStatusT = 0
	I2C_STATUS_WRITE     I2cMasterStatusT = 1
	I2C_STATUS_START     I2cMasterStatusT = 2
	I2C_STATUS_STOP      I2cMasterStatusT = 3
	I2C_STATUS_IDLE      I2cMasterStatusT = 4
	I2C_STATUS_ACK_ERROR I2cMasterStatusT = 5
	I2C_STATUS_DONE      I2cMasterStatusT = 6
	I2C_STATUS_TIMEOUT   I2cMasterStatusT = 7
)

type I2cMasterEventT c.Int

const (
	I2C_EVENT_ALIVE   I2cMasterEventT = 0
	I2C_EVENT_DONE    I2cMasterEventT = 1
	I2C_EVENT_NACK    I2cMasterEventT = 2
	I2C_EVENT_TIMEOUT I2cMasterEventT = 3
)

type I2cMasterCommandT c.Int

const (
	I2C_MASTER_CMD_START I2cMasterCommandT = 0
	I2C_MASTER_CMD_WRITE I2cMasterCommandT = 1
	I2C_MASTER_CMD_READ  I2cMasterCommandT = 2
	I2C_MASTER_CMD_STOP  I2cMasterCommandT = 3
)

type I2cAckValueT c.Int

const (
	I2C_ACK_VAL  I2cAckValueT = 0
	I2C_NACK_VAL I2cAckValueT = 1
)

type I2cMasterBusT struct {
	Unused [8]uint8
}
type I2cMasterBusHandleT *I2cMasterBusT

type I2cMasterDevT struct {
	Unused [8]uint8
}
type I2cMasterDevHandleT *I2cMasterDevT

type I2cSlaveDevT struct {
	Unused [8]uint8
}
type I2cSlaveDevHandleT *I2cSlaveDevT

/**
 * @brief Data type used in I2C event callback
 */

type I2cMasterEventDataT struct {
	Event I2cMasterEventT
}

// llgo:type C
type I2cMasterCallbackT func(I2cMasterDevHandleT, *I2cMasterEventDataT, c.Pointer) bool

/**
 * @brief Event structure used in I2C slave
 */

type I2cSlaveRxDoneEventDataT struct {
	Buffer *c.Uint8T
}

// llgo:type C
type I2cSlaveReceivedCallbackT func(I2cSlaveDevHandleT, *I2cSlaveRxDoneEventDataT, c.Pointer) bool

/**
 * @brief Stretch cause event structure used in I2C slave
 */

type I2cSlaveStretchEventDataT struct {
	StretchCause I2cSlaveStretchCauseT
}

// llgo:type C
type I2cSlaveStretchCallbackT func(I2cSlaveDevHandleT, *I2cSlaveStretchEventDataT, c.Pointer) bool

/**
 * @brief Event structure used in I2C slave request.
 */

type I2cSlaveRequestEventDataT struct {
	Unused [8]uint8
}

// llgo:type C
type I2cSlaveRequestCallbackT func(I2cSlaveDevHandleT, *I2cSlaveRequestEventDataT, c.Pointer) bool
