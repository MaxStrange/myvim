# This script adds a monitor called DP1 to the left of eDP1, which
# should be your main display.
# To determine what these names should actually be, use the xrandr tool.

xrandr --output DP1 --mode 3840x2160 --left-of eDP1
# xrandr --output DP2 --mode 3840x2160 --right-of eDP1
