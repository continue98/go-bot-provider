package pawn_api

// #include <stdio.h>
// #include <stdlib.h>
// #cgo LDFLAGS: -lm -ldl -L ../../libs -lpawnrun -lpawncc
// int pc_compile(int argc, char **argv);
// int pawnrun(int argc, char **argv);
import "C"

import (
	"errors"
	"unsafe"
)

type SourceInfo struct {
	CurrentPath string
	Name        string
	IncludeDir  string
}

func CompileSrc(src *SourceInfo) error {
	argv := make([]*C.char, 4)
	argv[0] = C.CString(src.CurrentPath)                           // current path (argv[0])
	argv[1] = C.CString(src.CurrentPath + "/examples/" + src.Name) // filename for compile
	argv[2] = C.CString("-i" + src.IncludeDir)
	argv[3] = C.CString("-o" + src.CurrentPath + "/bin_pwn/" + src.Name)
	defer func() {
		C.free(unsafe.Pointer(argv[0]))
		C.free(unsafe.Pointer(argv[1]))
	}()

	ret := C.pc_compile(C.int(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))

	if ret != 0 {
		return errors.New("Source not compile")
	}
	return nil
}

func RunPwn(src *SourceInfo) error {
	argv := make([]*C.char, 2)
	argv[0] = C.CString(src.CurrentPath)                          // current path (argv[0])
	argv[1] = C.CString(src.CurrentPath + "/bin_pwn/" + src.Name) // filename for compile
	defer func() {
		C.free(unsafe.Pointer(argv[0]))
		C.free(unsafe.Pointer(argv[1]))
	}()

	ret := C.pawnrun(C.int(len(argv)), (**C.char)(unsafe.Pointer(&argv[0])))

	if ret != 0 {
		return errors.New("amx file not start")
	}
	return nil
}
