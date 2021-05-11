#!/usr/bin/env bash
# This is a template script. Simply replace "dolphin" with your application of choice.
#
# This script's designed was aimed at having it assigned to a hotkey in order to prevent opening multiple windows of
# the same application in an X11 desktop. If a give command has its own X window, minimized or otherwise out of focus
# at runtime, this script will bring that window forward and focus it, instead of opening up a new application

wmctrl -xa dolphin || exec dolphin