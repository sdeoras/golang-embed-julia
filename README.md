# golang-embed-julia
=====

example taken from: https://docs.julialang.org/en/v1/manual/embedding/#High-Level-Embedding-1
originally forked and modified from: https://github.com/glycerine/golang-embed-julia

## system info
code was tested on Ubuntu-18 with following configuration.
```bash
$ gcc --version
gcc (Ubuntu 7.3.0-27ubuntu1~18.04) 7.3.0

$ go version
go version go1.11 linux/amd64

$ julia -v
julia version 1.0.1

$ uname -r
4.15.0-1021-gcp
```

## example use
```bash
$ go build -o main .
$ ./main
1.4142135623730951
  0.088042 seconds (183.60 k allocations: 17.906 MiB)
  0.012806 seconds (6 allocations: 8.000 MiB)
```

## protable app bundle contains:
```bash
$ tree ./
./
├── bin
├── lib
│   ├── julia
│   │   ├── libamd.so
│   │   ├── libcamd.so
│   │   ├── libccalltest.so
│   │   ├── libccalltest.so.debug
│   │   ├── libccolamd.so
│   │   ├── libcholmod.so
│   │   ├── libcolamd.so
│   │   ├── libcurl.so
│   │   ├── libcurl.so.4
│   │   ├── libcurl.so.4.5.0
│   │   ├── libdSFMT.so
│   │   ├── libgcc_s.so.1
│   │   ├── libgfortran.so.4
│   │   ├── libgit2.so
│   │   ├── libgit2.so.0.27.2
│   │   ├── libgit2.so.27
│   │   ├── libgmp.so
│   │   ├── libgmp.so.10
│   │   ├── libgmp.so.10.3.2
│   │   ├── libLLVM-6.0.0.so
│   │   ├── libLLVM-6.0.so
│   │   ├── libLLVM.so
│   │   ├── libmbedcrypto.so
│   │   ├── libmbedcrypto.so.0
│   │   ├── libmbedcrypto.so.2.6.0
│   │   ├── libmbedtls.so
│   │   ├── libmbedtls.so.10
│   │   ├── libmbedtls.so.2.6.0
│   │   ├── libmbedx509.so
│   │   ├── libmbedx509.so.0
│   │   ├── libmbedx509.so.2.6.0
│   │   ├── libmpfr.so
│   │   ├── libmpfr.so.6
│   │   ├── libmpfr.so.6.0.1
│   │   ├── libopenblas64_.so
│   │   ├── libopenblas64_.so.0
│   │   ├── libopenlibm.so
│   │   ├── libopenlibm.so.2
│   │   ├── libopenlibm.so.2.5
│   │   ├── libpcre2-8.so
│   │   ├── libpcre2-8.so.0
│   │   ├── libpcre2-8.so.0.6.0
│   │   ├── libquadmath.so.0
│   │   ├── libspqr.so
│   │   ├── libssh2.so
│   │   ├── libssh2.so.1
│   │   ├── libssh2.so.1.0.1
│   │   ├── libstdc++.so.6
│   │   ├── libsuitesparseconfig.so
│   │   ├── libsuitesparse_wrapper.so
│   │   ├── libumfpack.so
│   │   └── sys.so
│   ├── libjulia.so
│   ├── libjulia.so.1
│   ├── libjulia.so.1.0
│   ├── libLLVM-6.0.0.so
│   ├── libLLVM-6.0.so
│   └── libLLVM.so
└── main
```
