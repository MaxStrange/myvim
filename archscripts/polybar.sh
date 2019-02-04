#!/usr/bin/env sh
# Pilfered from https://www.reddit.com/r/i3wm/comments/6lo0z0/how_to_use_polybar/
# This script runs my status bar (polybar). It is invoked from .config/i3/config.

# Terminate already running bar instances
killall -q polybar

# Wait until the processes have been shut down
while pgrep -x polybar >/dev/null; do sleep 1; done

# Launch polybar
polybar main &
