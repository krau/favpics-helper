#!/bin/sh
VERSION=$(git describe --tags)
_build() {
    local osarch=$1
    IFS=/ read -r -a arr <<<"$osarch"
    os="${arr[0]}"
    arch="${arr[1]}"
    gcc="${arr[2]}"

    # Go build to build the binary.
    export GOOS=$os
    export GOARCH=$arch
    export CC=$gcc
    export CGO_ENABLED=1

    out="release/favpics-helper_${VERSION}_${os}_${arch}"

    go build -a -o "${out}"

    if [ "$os" = "windows" ]; then
        mv $out release/favpics-helper.exe
        cp config.toml release/
        zip -j -q "${out}.zip" release/favpics-helper.exe
        rm -f "release/favpics-helper.exe"
        rm -f "release/config.toml"
    else
        mv $out release/favpics-helper
        cp config.toml release/
        tar -zcvf "${out}.tar.gz" -C release favpics-helper
        rm -f "release/favpics-helper"
        rm -f "release/config.toml"
    fi
}

## List of architectures and OS to test coss compilation.
SUPPORTED_OSARCH="linux/amd64/gcc linux/arm/arm-linux-gnueabihf-gcc windows/amd64/x86_64-w64-mingw32-gcc linux/arm64/aarch64-linux-gnu-gcc"

echo "Release builds for OS/Arch/CC: ${SUPPORTED_OSARCH}"
for each_osarch in ${SUPPORTED_OSARCH}; do
    _build "${each_osarch}"
done
