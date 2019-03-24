package pawn_api

// #include <stdio.h>
// #include <stdlib.h>
// #cgo LDFLAGS: -lm -ldl -L ../../libs -lpawnrun -lpawncc
// int pc_compile(int argc, char **argv);
// int pawnrun(int argc, char **argv);
import "C"

import (
	"bytes"
	"errors"
	"io"
	"os"
	"regexp"
	"syscall"
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

func GetStreamsOut(f func(src *SourceInfo) error, src *SourceInfo) (error, string) { // stdout, stderr capture
	stdout_copy, _ := syscall.Dup(1)
	r, w, _ := os.Pipe()
	fd := w.Fd()
	syscall.Dup2(int(fd), 1)

	stderr_copy, _ := syscall.Dup(2)
	r1, w1, _ := os.Pipe()
	fd1 := w1.Fd()
	syscall.Dup2(int(fd1), 2)

	err := f(src)

	w.Close()

	syscall.Dup2(stdout_copy, 1)
	//stdout_copy.Close()

	var buf bytes.Buffer
	io.Copy(&buf, r)

	w1.Close()

	syscall.Dup2(stderr_copy, 2)

	var buff bytes.Buffer
	io.Copy(&buff, r1)

	out := buf.String() + buff.String()

	re := regexp.MustCompile(`(?m)(\d+\s+(Error|Warning|Errors|Warnings)[.])`)

	res := re.FindAllStringSubmatch(out, -1)
	out = re.ReplaceAllString(out, "")

	if cap(res) != 0 {
		out += "\n" + res[0][0]
	}
	out = regexp.MustCompile(`(?m)^\s+\n$`).ReplaceAllString(out, "$1")
	return err, out
}
