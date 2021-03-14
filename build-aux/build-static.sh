#!/bin/bash -xe

FILE_PV="5.39"

export CC=musl-gcc
export CGO_CFLAGS="-I${BUILDDIR}/file-${FILE_PV}/src"
export CGO_LDFLAGS="-static -L${BUILDDIR}/file-${FILE_PV}/src/.libs"

wget "ftp://ftp.astron.com/pub/file/file-${FILE_PV}.tar.gz"

rm -rf "file-${FILE_PV}"
tar -xvf "file-${FILE_PV}.tar.gz"

pushd "file-${FILE_PV}" > /dev/null

./configure \
    CC=musl-gcc \
    --enable-static \
    --disable-shared \
    --disable-zlib \
    --disable-bzlib \
    --disable-xzlib \
    --disable-libseccomp

${MAKE_CMD} -j1 \
    clean \
    all

popd > /dev/null

rm -rf "${BUILDDIR}/${P}"
mkdir -p "${BUILDDIR}/${P}"

pushd "${BUILDDIR}/${P}" > /dev/null
go build "${@}" "${SRCDIR}"
cp "${BUILDDIR}/file-${FILE_PV}/magic/magic.mgc" .
cp "${BUILDDIR}/file-${FILE_PV}/COPYING" LICENSE-libmagic
cp "${SRCDIR}/LICENSE" .
popd > /dev/null

rm -f "file-${FILE_PV}.tar.gz"

tar -cvJf "${P}.tar.xz" "${P}"
