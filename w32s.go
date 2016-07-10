package w32s

import (
	"errors"
	"syscall"
)

const (
	Success        = syscall.Errno(0)
	ResultW32Error = ^uintptr(1337)
)

// String encoding (multibyte / wide) character
type StringEncoding int

const (
	_ StringEncoding = iota
	Multibyte
	Wide
)

// Errors
var (
	ErrDllNotLoaded = errors.New("w32s: requested dll is not loaded")
)

type W32s struct {
	encoding StringEncoding
	dlls     map[string]*dll
}

func New(encoding StringEncoding) *W32s {
	return &W32s{
		encoding: encoding,
		dlls:     make(map[string]*dll),
	}
}

func (w *W32s) LoadDLL(name string) error {
	if _, ok := w.dlls[name]; !ok {
		dll, err := syscall.LoadDLL(name)
		if err != nil {
			return err
		}
		w.dlls[name] = newDll(dll, w.encoding)
	}
	return nil
}

func (w *W32s) ReleaseDLL(name string) error {
	if dll, ok := w.dlls[name]; ok {
		delete(w.dlls, name)
		if err := dll.Release(); err != nil {
			return err
		}
	}
	return nil
}

func (w *W32s) Call(dllname, procname string, args ...interface{}) (uintptr, error) {
	dll, ok := w.dlls[dllname]
	if !ok {
		return ResultW32Error, ErrDllNotLoaded
	}
	return dll.call(procname, args...)
}

func (w *W32s) StrBuf(length int) *StringBuffer {
	return newStrBuf(w.encoding, length)
}
