#!/bin/bash
gcc embed.c -o main \
	-std=gnu99 \
	-I'/usr/local/julia/include/julia' \
	-DJULIA_ENABLE_THREADING=1 \
	-fPIC \
	-lm \
	-L'/usr/local/julia/lib' \
	-Wl,--export-dynamic \
	-Wl,-rpath,'/usr/local/julia/lib' \
	-Wl,-rpath,'/usr/local/julia/lib/julia' \
	-ljulia
