#!/bin/sh

# wait for display
if [ ! -e /sys/class/graphics/fb0 ]; then
    echo "Display not detected.  Will wait for it to show up ..."
fi
while [ ! -e /sys/class/graphics/fb0 ]; do
    sleep 1
done

# use /tmp for config and cache data in read-only mode
mkdir -p ${XDG_CACHE_HOME} ${XDG_CONFIG_HOME}

# Execute all the rest
echo "Starting Chromium: $@"
/usr/bin/chromium \
  --no-sandbox --start-maximized --no-first-run --disable --disable-translate \
  --disable-infobars --disable-suggestions-service --disable-save-password-bubble \
  --incognito --autoplay-policy=no-user-gesture-required --temp-profile \
  --disable-dev-shm-usage --kiosk "$@"
