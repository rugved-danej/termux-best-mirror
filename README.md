# termux-best-mirror ⚡

![termux-best-mirror preview](preview.gif)

A modern, ultra-fast, and responsive tool to automatically find and set the fastest repository mirror for your Termux installation. 

Built exclusively for Termux, this tool features highly concurrent speed-tests, an interactive terminal UI with seamless resize responsiveness, and real-time color-coded latency tracking.

## Features ✨
- **Ultra-Fast Benchmarking**: Tests all Termux mirrors globally using asynchronous background requests to find the fastest response times.
- **Modern Interface**: A pure bash implementation that abandons the clunky `dialog` menus for a beautiful, colorful, responsive arrow-key navigated menu.
- **Terminal Resize Support**: Automatically adapts layout and redraws instantly if your screen size changes vertically or horizontally.
- **Zero Dependencies**: Requires no external TUI packages, relying entirely on Termux standards (`curl` and `awk`).

## Installation 🚀

You can easily install this tool globally onto your device with a single command:

```bash
bash <(curl -sL https://raw.githubusercontent.com/rugved-danej/termux-best-mirror/main/install.sh)
```

**Alternatively, to install manually from a local clone:**
```bash
chmod +x install.sh
./install.sh
```

## Usage 💡

Once installed, simply run the tool from anywhere in Termux by typing:

```bash
termux-best-mirror
```

1. Use the **Up/Down arrow keys** to navigate regions.
2. Press **Enter** to select a region (or `All mirrors` to scan everything).
3. The script will test latency to all mirrors concurrently, showing a beautiful live progress bar.
4. Select your preferred mirror from the resulting list (sorted fastest to slowest) to automatically link it and update `apt`.

## Uninstallation 🗑️

To remove the command from Termux completely, run the following Termux command:

```bash
rm $PREFIX/bin/termux-best-mirror
```
