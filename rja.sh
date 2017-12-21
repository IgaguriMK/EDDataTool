#! /bin/bash

set -eu

OUT_DIR=out_jcheck

./rj.sh

ERR=`ls -1 ${OUT_DIR} | head -n 1`

if [ -n "$ERR" ]; then
	./rjn.sh $OUT_DIR/$ERR
fi
