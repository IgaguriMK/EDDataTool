#! /bin/bash

set -eu

OUT_DIR=out_jcheck

make
mv ${OUT_DIR} ${OUT_DIR}_old
rm -r ${OUT_DIR}_old &
mkdir -p ${OUT_DIR} || true
mkdir -p ${OUT_DIR}/types/ || true
./jcheck ~/Saved\ Games/Frontier\ Developments/Elite\ Dangerous/*.log
