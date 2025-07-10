package protocomm

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type X_Status c.Int

const (
	STATUS__Success          X_Status = 0
	STATUS__InvalidSecScheme X_Status = 1
	STATUS__InvalidProto     X_Status = 2
	STATUS__TooManySessions  X_Status = 3
	STATUS__InvalidArgument  X_Status = 4
	STATUS__InternalError    X_Status = 5
	STATUS__CryptoError      X_Status = 6
	STATUS__InvalidSession   X_Status = 7
	X_STATUS_IS_INT_SIZE     X_Status = 2147483647
)

type Status X_Status
