#!/usr/bin/env bash
set -euo pipefail

dest="$1"
self="$(readlink -f  "$0")"
this="$(basename "$self")"
d="$(dirname "$self")"

set -x
while read f; do
  cp -r "$f" "$dest"
done < <(
find "$d" -mindepth 1 -maxdepth 1 \
    \! -name 'Makefile' \
    \! -name 'debian' \
    \! -name 'LICENSE' \
    \! -name "$this" \
    -print
)
