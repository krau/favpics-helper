#!/bin/sh
VERSION=$(git describe --tags)
_build() {
    local osarch=$1
    IFS=/ read -r -a arr <<<"$osarch"
    os="${arr[0]}"
    arch="${arr[1]}"

    # Go build to build the binary.
    export GOOS=$os
    export GOARCH=$arch

    out="release/favpics-helper_${VERSION}_${os}_${arch}"

    GOOS=$os GOARCH=$arch go build -o "${out}" main.go

    if [ "$os" = "windows" ]; then
        mv $out release/favpics-helper.exe
        cp config.toml release/
        zip -j -q "${out}.zip" release/favpics-helper.exe release/config.toml
        rm -f "release/favpics-helper.exe"
        rm -f "release/config.toml"
    else
        mv $out release/favpics-helper
        cp config.toml release/
        tar -zcvf "${out}.tar.gz" -C release favpics-helper config.toml
        rm -f "release/favpics-helper"
        rm -f "release/config.toml"
    fi
}

## List of architectures and OS to test coss compilation.
SUPPORTED_OSARCH="linux/amd64 linux/arm windows/amd64 linux/arm64"

echo "Release builds for OS/Arch: ${SUPPORTED_OSARCH}"
for each_osarch in ${SUPPORTED_OSARCH}; do
    echo "Building for ${each_osarch}"
    _build "${each_osarch}"
done
