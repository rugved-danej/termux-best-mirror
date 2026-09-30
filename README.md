# termux-best-mirror

[![GitHub Release](https://img.shields.io/github/v/release/rugved-danej/termux-best-mirror?style=flat-square&color=blueviolet)](https://github.com/rugved-danej/termux-best-mirror/releases)
[![License](https://img.shields.io/github/license/rugved-danej/termux-best-mirror?style=flat-square&color=success)](LICENSE)
[![Issues](https://img.shields.io/github/issues/rugved-danej/termux-best-mirror?style=flat-square&color=critical)](https://github.com/rugved-danej/termux-best-mirror/issues)
[![Stars](https://img.shields.io/github/stars/rugved-danej/termux-best-mirror?style=flat-square&color=yellow)](https://github.com/rugved-danej/termux-best-mirror/stargazers)

![termux-best-mirror preview](preview.gif)

A modern, ultra-fast, and responsive tool to automatically find and set the fastest repository mirror for your Termux installation. 

Built exclusively for Termux, this tool features highly concurrent speed-tests, an interactive terminal UI with seamless resize responsiveness, and real-time color-coded latency tracking.

## Features
- **Ultra-Fast Benchmarking**: Tests all Termux mirrors globally using highly concurrent requests to find the fastest response times.
- **Modern Interface**: A pure Go implementation that utilizes `pterm` for a beautiful, colorful, and responsive interactive terminal UI.
- **Cool Animations**: Features smooth spinners, loading indicators, and progress bars.
- **Universal Device Support**: Automatically detects your architecture (arm, aarch64, x86_64, i686) and downloads the optimized binary, ensuring it works on *all* devices running Termux.
- **Zero Dependencies**: Distributed as a single compiled binary so you don't need to install Go or any other dependencies to use it.

## Installation

You can easily install this tool globally onto your device with a single command. It will automatically detect your phone's CPU architecture and download the right file:

```bash
bash <(curl -sL https://raw.githubusercontent.com/rugved-danej/termux-best-mirror/main/install.sh)
```

## Usage

Once installed, simply run the tool from anywhere in Termux by typing:

```bash
termux-best-mirror
```

1. Use the **Up/Down arrow keys** to navigate regions.
2. Press **Enter** to select a region (or `All mirrors` to scan everything).
3. The script will test latency to all mirrors concurrently, showing a beautiful live progress bar.
4. Select your preferred mirror from the resulting list (sorted fastest to slowest) to automatically link it and update `apt`.

## Uninstallation

To remove the command from Termux completely, run the following Termux command:

```bash
rm $PREFIX/bin/termux-best-mirror
```
