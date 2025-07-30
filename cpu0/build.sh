llgo build -tags "nogc" -v  -o t
 python /Users/haolan/esp/esp-idf/components/esptool_py/esptool/esptool.py --chip esp32 elf2image --flash_mode dio --flash_freq 40m --flash_size 2MB --min-rev-full 0 --max-rev-full 399 -o t.bin t.elf 
 python -m esptool --chip esp32 -b 460800 --before default_reset --after hard_reset write_flash 0x10000 t.bin

