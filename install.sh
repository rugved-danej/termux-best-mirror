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

echo -e "\033[36m⌛ Installing $SCRIPT_NAME to $DESTINATION...\033[0m"

if [ ! -f "$SCRIPT_NAME" ]; then
	echo -e "\033[1;31m✖ Error: Could not find the file '$SCRIPT_NAME' in the current directory.\033[0m"
	exit 1
fi

# Ensure the script is executable
chmod +x "$SCRIPT_NAME"

# Copy the script to the Termux bin directory
cp "$SCRIPT_NAME" "$DESTINATION"

# Make sure the installed version is executable
chmod +x "$DESTINATION"

echo -e "\033[1;32m✔ Installation successful!\033[0m"
echo -e "\033[1;36mYou can now run the tool anywhere by simply typing: \033[1;32m$SCRIPT_NAME\033[0m"
