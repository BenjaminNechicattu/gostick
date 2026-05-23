#!/bin/bash

set -e

APP_NAME="gostick"
VERSION="0.1.0"
ARCH="amd64"

BUILD_DIR="packaging/deb"
OUTPUT="${APP_NAME}_${VERSION}_${ARCH}.deb"

echo ""
echo "=== Building GoStick Debian Package ==="
echo ""

# CLEAN OLD BUILD
echo "[1/8] Cleaning old build..."
rm -rf "$BUILD_DIR"

# CREATE FOLDERS
echo "[2/8] Creating package structure..."

mkdir -p $BUILD_DIR/DEBIAN
mkdir -p $BUILD_DIR/usr/bin
mkdir -p $BUILD_DIR/usr/share/applications
mkdir -p $BUILD_DIR/usr/share/icons/hicolor/256x256/apps
mkdir -p $BUILD_DIR/etc/gostick

# BUILD GO BINARY
echo "[3/8] Building GoStick binary..."

go build -o gostick

# COPY BINARY
echo "[4/8] Copying binary..."

cp gostick $BUILD_DIR/usr/bin/

chmod +x $BUILD_DIR/usr/bin/gostick

# COPY PROFILE
echo "[5/8] Copying default profile..."

cp profiles/default.json $BUILD_DIR/etc/gostick/default.json

# CREATE CONTROL FILE
echo "[6/8] Creating DEBIAN/control..."

cat > $BUILD_DIR/DEBIAN/control <<EOF
Package: gostick
Version: ${VERSION}
Section: utils
Priority: optional
Architecture: ${ARCH}
Maintainer: BenG
Description: Native Linux desktop control using a game controller
Depends: libc6
EOF

chmod 644 $BUILD_DIR/DEBIAN/control

# CREATE DESKTOP ENTRY
echo "[7/8] Creating desktop launcher..."

cat > $BUILD_DIR/usr/share/applications/gostick.desktop <<EOF
[Desktop Entry]
Name=GoStick
Comment=Game controller desktop control
Exec=/usr/bin/gostick
Icon=gostick
Terminal=false
Type=Application
Categories=Utility;
EOF

# OPTIONAL ICON
if [ -f assets/gostick.png ]; then

    echo "[INFO] Adding icon..."

    cp assets/gostick.png \
       $BUILD_DIR/usr/share/icons/hicolor/256x256/apps/gostick.png

else

    echo "[WARN] No icon found at assets/gostick.png"
fi

# BUILD PACKAGE
echo "[8/8] Building .deb package..."

dpkg-deb --build $BUILD_DIR $OUTPUT

echo ""
echo "=== Package Created Successfully ==="
echo ""
echo "Output:"
echo "  $OUTPUT"
echo ""
echo "Install with:"
echo "  sudo dpkg -i $OUTPUT"
echo ""