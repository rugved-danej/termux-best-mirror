#!/data/data/com.termux/files/usr/bin/bash

# Define standard Termux prefix fallback if not set
if [ -z "$TERMUX_PREFIX" ]; then
	TERMUX_PREFIX="/data/data/com.termux/files/usr"
fi

if [ ! -d "$TERMUX_PREFIX/bin" ]; then
	echo -e "\033[1;31m✖ Error: This script must be run inside Termux.\033[0m"
	exit 1
fi

SCRIPT_NAME="termux-best-mirror"
DESTINATION="$TERMUX_PREFIX/bin/$SCRIPT_NAME"
RAW_URL="https://raw.githubusercontent.com/rugved-danej/termux-best-mirror/main/termux-best-mirror"

echo -e "\033[36m⌛ Installing $SCRIPT_NAME to $DESTINATION...\033[0m"

if [ -f "$SCRIPT_NAME" ]; then
	# Copy local file if repository is cloned
	cp "$SCRIPT_NAME" "$DESTINATION"
else
	# Download directly from GitHub if run via curl memory pipe
	echo -e "\033[36m⌛ Downloading latest version from GitHub...\033[0m"
	if ! curl -sL "$RAW_URL" -o "$DESTINATION"; then
		echo -e "\033[1;31m✖ Error: Failed to download the script. Check your internet connection.\033[0m"
		exit 1
	fi
fi

# Make sure the installed version is executable
chmod +x "$DESTINATION"

echo -e "\033[1;32m✔ Installation successful!\033[0m"
echo -e "\033[1;36mYou can now run the tool anywhere by simply typing: \033[1;32m$SCRIPT_NAME\033[0m"
