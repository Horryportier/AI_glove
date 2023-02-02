#!/bin/bash

path="build/esp32.esp32.esp32doit-devkit-v1/WebUpdater.ino.bin"

ip="$1"

echo -e "$(curl -F "image=@$path" $ip/update)"
