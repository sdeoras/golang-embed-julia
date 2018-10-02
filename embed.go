package main

/*
// https://docs.julialang.org/en/v1/manual/embedding/#High-Level-Embedding-1
// forked and modified from: https://github.com/glycerine/golang-embed-julia
// Usage shown for Ubuntu Linux
// Tested on following configuration:
//	$ gcc --version
//	gcc (Ubuntu 7.3.0-27ubuntu1~18.04) 7.3.0
//
//	$ go version
//	go version go1.11 linux/amd64
//
//	$ julia -v
//	julia version 1.0.1
//
//	$ uname -r
//	4.15.0-1021-gcp
//
// Create following folder structure next to binary built from this code
//	mkdir -p ./bin/../lib/julia
// Copy .so files to lib and julia folders
//	$ cp /usr/local/julia/lib/libjulia.so ./lib
//	$ cp /usr/local/julia/lib/julia/* ./lib/julia
// Build
//	$ go build -o main .
//	$ ldd ./main
//	linux-vdso.so.1 (0x00007ffdc995f000)
//	libjulia.so.1 => /usr/local/julia/lib/libjulia.so.1 (0x00007f38e6cfd000)
//	libpthread.so.0 => /lib/x86_64-linux-gnu/libpthread.so.0 (0x00007f38e6ade000)
//	libc.so.6 => /lib/x86_64-linux-gnu/libc.so.6 (0x00007f38e66ed000)
//	libLLVM-6.0.so => /usr/local/julia/lib/julia/libLLVM-6.0.so (0x00007f38e3a51000)
//	libdl.so.2 => /lib/x86_64-linux-gnu/libdl.so.2 (0x00007f38e384d000)
//	librt.so.1 => /lib/x86_64-linux-gnu/librt.so.1 (0x00007f38e3645000)
//	libstdc++.so.6 => /usr/local/julia/lib/julia/libstdc++.so.6 (0x00007f38e32c7000)
//	libm.so.6 => /lib/x86_64-linux-gnu/libm.so.6 (0x00007f38e2f29000)
//	libgcc_s.so.1 => /usr/local/julia/lib/julia/libgcc_s.so.1 (0x00007f38e2d11000)
//	/lib64/ld-linux-x86-64.so.2 (0x00007f38e746f000)
//
#cgo CFLAGS: -std=gnu99 -I'/usr/local/julia/include/julia' -DJULIA_ENABLE_THREADING=1 -fPIC
#cgo LDFLAGS: -lm -L'/usr/local/julia/lib' -Wl,-rpath,'/usr/local/julia/lib' -Wl,-rpath,'/usr/local/julia/lib/julia' -Wl,-rpath,'\$ORIGIN/lib' -ljulia
#include <julia.h>
JULIA_DEFINE_FAST_TLS() // only define this once, in an executable (not in a shared library) if you want fast code.
*/
import "C"

func main() {
	/* required: setup the Julia context */
	C.jl_init()

	/* run Julia commands */
	C.jl_eval_string(C.CString(`println(sqrt(2.0))`))
	C.jl_eval_string(C.CString(`@time x=randn(1024, 1024)`))
	C.jl_eval_string(C.CString(`@time x=randn(1024, 1024)`))

	/* strongly recommended: notify Julia that the
	   program is about to terminate. this allows
	   Julia time to cleanup pending write requests
	   and run all finalizers
	*/
	C.jl_atexit_hook(0)
}
