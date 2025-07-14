package bootloader_support

import _ "unsafe"

/*
 * SPDX-FileCopyrightText: 2020-2021 Espressif Systems (Shanghai) CO LTD
 *
 * SPDX-License-Identifier: Apache-2.0
 */
//go:linkname BootloaderInitMem C.bootloader_init_mem
func BootloaderInitMem()
