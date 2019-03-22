package pawn_api

// #include <stdio.h>
// #include <stdlib.h>
// #cgo LDFLAGS: -lm -ldl -L ../libs -lpawnrun -lpawncc
// int pc_compile(int argc, char **argv);
// int pawnrun(int argc, char **argv);
import "C"

import (
	"unsafe"
)

type SourceInfo struct {
	current_path string
	name         string
}

func compile_src(src *SourceInfo) int {
	argv := make([]*C.char, 2)
	argv[0] = C.CString(src.current_path) // current path (argv[0])
	argv[1] = C.CString(src.name)         // filename for compile
	defer func() {
		C.free(unsafe.Pointer(argv[0]))
		C.free(unsafe.Pointer(argv[1]))
	}()

	ret := C.pc_compile(C.int(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))

	return int(ret)
}

func run_pwn(src *SourceInfo) int {
	argv := make([]*C.char, 2)
	argv[0] = C.CString(src.current_path) // current path (argv[0])
	argv[1] = C.CString(src.name)         // filename for compile
	defer func() {
		C.free(unsafe.Pointer(argv[0]))
		C.free(unsafe.Pointer(argv[1]))
	}()

	ret := C.pawnrun(C.int(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))

	return int(ret)
}
