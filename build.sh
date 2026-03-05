#!/bin/bash
set -e

NOTIFIER_APP="YouLostTheGame.app"
CONTENTS="$NOTIFIER_APP/Contents"
MACOS="$CONTENTS/MacOS"
RESOURCES="$CONTENTS/Resources"

# Clean
rm -rf "$NOTIFIER_APP"

# Build Go daemon
go build -o you-lost-the-game .

# Build Swift notifier app
swiftc notifier/main.swift -o notifier-bin

# Create .app bundle for the notifier (so it gets notification permissions)
mkdir -p "$MACOS" "$RESOURCES"

cp notifier-bin "$MACOS/YouLostTheGame"
cp you-lost-the-game "$MACOS/"
cp icon.icns "$RESOURCES/AppIcon.icns"

# Info.plist - identifies the notifier as the main app
cat > "$CONTENTS/Info.plist" << 'PLIST'
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>CFBundleExecutable</key>
    <string>YouLostTheGame</string>
    <key>CFBundleIdentifier</key>
    <string>dev.salmont.useless.youlostthegame</string>
    <key>CFBundleName</key>
    <string>You Lost The Game</string>
    <key>CFBundleIconFile</key>
    <string>AppIcon</string>
    <key>CFBundlePackageType</key>
    <string>APPL</string>
    <key>LSUIElement</key>
    <true/>
    <key>CFBundleVersion</key>
    <string>1.0</string>
</dict>
</plist>
PLIST

echo "Built $NOTIFIER_APP successfully!"
echo ""
echo "Test:  open YouLostTheGame.app --args 'You lost The Game.'"
echo "Run:   $MACOS/you-lost-the-game"
echo "Install: $MACOS/you-lost-the-game --install"
