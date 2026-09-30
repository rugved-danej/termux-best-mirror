#!/data/data/com.termux/files/usr/bin/bash

if [ -z "$TERMUX_PREFIX" ]; then
	TERMUX_PREFIX="/data/data/com.termux/files/usr"
fi

if [ ! -d "$TERMUX_PREFIX/bin" ]; then
	echo -e "\033[1;31m[!] Error: This script must be run inside Termux.\033[0m"
	exit 1
fi

SCRIPT_NAME="termux-best-mirror"
DESTINATION="$TERMUX_PREFIX/bin/$SCRIPT_NAME"

echo -e "\033[36m[*] Installing $SCRIPT_NAME to $DESTINATION...\033[0m"

ARCH=$(uname -m)
case "$ARCH" in
	aarch64) GOARCH="arm64" ;;
	armv7l|armv8l|arm) GOARCH="arm" ;;
	x86_64) GOARCH="amd64" ;;
	i686) GOARCH="386" ;;
	*)
		echo -e "\033[1;31m[!] Error: Unsupported architecture $ARCH\033[0m"
		exit 1
		;;
esac

if [ -f "cmd/termux-best-mirror/main.go" ] && command -v go >/dev/null 2>&1; then
	echo -e "\033[36m[*] Source files and Go detected. Building locally...\033[0m"
	rm -f "$DESTINATION"
	if ! go build -ldflags="-s -w" -o "$DESTINATION" ./cmd/termux-best-mirror/...; then
		echo -e "\033[1;31m[!] Error: Local build failed.\033[0m"
		exit 1
	fi
else
	echo -e "\033[36m[*] Downloading pre-compiled binary for $GOARCH...\033[0m"
	DOWNLOAD_URL="https://github.com/rugved-danej/termux-best-mirror/releases/latest/download/termux-best-mirror-android-$GOARCH"
	
	if ! curl -sL --fail "$DOWNLOAD_URL" -o "$DESTINATION"; then
		echo -e "\033[1;31m[!] Error: Failed to download the binary. Please ensure there is a release available for your architecture or install Go to build it locally.\033[0m"
		exit 1
	fi
	chmod +x "$DESTINATION"
fi

echo -e "\033[1;32m[+] Installation successful!\033[0m"
echo -e "\033[1;36mYou can now run the tool anywhere by simply typing: \033[1;32m$SCRIPT_NAME\033[0m"
