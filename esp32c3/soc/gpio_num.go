package soc

import (
	"github.com/goplus/lib/c"
	_ "unsafe"
)

type GpioNumT c.Int

const (
	GPIO_NUM_NC  GpioNumT = -1
	GPIO_NUM_0   GpioNumT = 0
	GPIO_NUM_1   GpioNumT = 1
	GPIO_NUM_2   GpioNumT = 2
	GPIO_NUM_3   GpioNumT = 3
	GPIO_NUM_4   GpioNumT = 4
	GPIO_NUM_5   GpioNumT = 5
	GPIO_NUM_6   GpioNumT = 6
	GPIO_NUM_7   GpioNumT = 7
	GPIO_NUM_8   GpioNumT = 8
	GPIO_NUM_9   GpioNumT = 9
	GPIO_NUM_10  GpioNumT = 10
	GPIO_NUM_11  GpioNumT = 11
	GPIO_NUM_12  GpioNumT = 12
	GPIO_NUM_13  GpioNumT = 13
	GPIO_NUM_14  GpioNumT = 14
	GPIO_NUM_15  GpioNumT = 15
	GPIO_NUM_16  GpioNumT = 16
	GPIO_NUM_17  GpioNumT = 17
	GPIO_NUM_18  GpioNumT = 18
	GPIO_NUM_19  GpioNumT = 19
	GPIO_NUM_20  GpioNumT = 20
	GPIO_NUM_21  GpioNumT = 21
	GPIO_NUM_MAX GpioNumT = 22
)
