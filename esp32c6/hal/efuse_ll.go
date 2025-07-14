package hal

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type EfuseControllerStateT c.Int

const (
	EFUSE_CONTROLLER_STATE_RESET          EfuseControllerStateT = 0
	EFUSE_CONTROLLER_STATE_IDLE           EfuseControllerStateT = 1
	EFUSE_CONTROLLER_STATE_READ_INIT      EfuseControllerStateT = 2
	EFUSE_CONTROLLER_STATE_READ_BLK0      EfuseControllerStateT = 3
	EFUSE_CONTROLLER_STATE_BLK0_CRC_CHECK EfuseControllerStateT = 4
	EFUSE_CONTROLLER_STATE_READ_RS_BLK    EfuseControllerStateT = 5
)
