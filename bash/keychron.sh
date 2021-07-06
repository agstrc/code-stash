#!/bin/env sh
# This simple one liner enables the "normal" usage of the F keys on Apple keyboards. As in, you must hold the FN key 
# along the F{1..12} key in order to use the special key function.
#
# This is mainly aimed at keychrons, since it might be a little bit unexpected to have them behave as Apple keyboards
# even on Windows mode.

echo 'options hid_apple fnmode=2' | sudo tee /etc/modprobe.d/hid_apple.conf
echo 2 | sudo tee /sys/module/hid_apple/parameters/fnmode