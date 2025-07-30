dd=($(ls $(pwd)))


for i in "${dd[@]}"; do
     echo replace '"github.com/MeteorsLiu/esp32/esp32/'$i'"' '=>' ../$i
done
