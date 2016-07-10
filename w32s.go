package w32s

import (
	"errors"
	"syscall"
)

const (
	// Success is the error code of a successful Win32 API procedure call
	Success = syscall.Errno(0)
	// ResultW32Error signals than an error occurred before making Win32 API call
	ResultW32Error = ^uintptr(1337)
)

// StringEncoding (multibyte / wide) character
type StringEncoding int

// StringEncoding possible values
const (
	_ StringEncoding = iota
	Multibyte
	Wide
)

// Errors
var (
	ErrDllNotLoaded = errors.New("w32s: requested dll is not loaded")
)

// W32s represents a collection of DLLs, methods of which can be invoked
// using the Call method.
// W32s also keeps a StringEncoding value which helps translating golang strings
// to multibyte character set or wide character set strings for WinAPI procedures.
type W32s struct {
	encoding StringEncoding
	dlls     map[string]*dll
}

// New creates a new W32s struct and returns it.
func New(encoding StringEncoding) *W32s {
	return &W32s{
		encoding: encoding,
		dlls:     make(map[string]*dll),
	}
}

// Encoding returns current string encoding.
func (w *W32s) Encoding() StringEncoding {
	return w.encoding
}

// SetEncoding sets specified string encoding.
func (w *W32s) SetEncoding(encoding StringEncoding) {
	w.encoding = encoding
}

// LoadDLL loads specified DLL using syscall.LoadDLL (i.e. LoadLibraryW)
// and returns an error if any occurred during the syscall.LoadDLL
// Calling LoadDLL twice or on already loaded DLL has no effect.
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

// ReleaseDLL releases specified DLL using syscall.DLL.Release (i.e. FreeLibrary)
// and returns an error if any occurred.
// Calling ReleaseDLL twice or on a non-loaded DLL has no effect.
func (w *W32s) ReleaseDLL(name string) error {
	if dll, ok := w.dlls[name]; ok {
		delete(w.dlls, name)
		if err := dll.Release(); err != nil {
			return err
		}
	}
	return nil
}

// Call invokes the procedure by specified name from the specified DLL,
// passing specified parameters (`args`) to the procedure.
//
// The first return value `uintptr` is the return value of the procedure,
// or the ResultW32Error value, if if any error occurred before calling the procedure.
// The second return value `error` is either syscall.Errno([GetLastError()]), if the call succeeded,
// or it's the error which occurred before the call, if the first return value is ResultW32Error.
func (w *W32s) Call(dllname, procname string, args ...interface{}) (uintptr, error) {
	dll, ok := w.dlls[dllname]
	if !ok {
		return ResultW32Error, ErrDllNotLoaded
	}
	return dll.call(procname, args...)
}

// StrBuf returns a StringBuffer with specified length
// and encoding which is set int W32s instance.
func (w *W32s) StrBuf(length int) *StringBuffer {
	return newStrBuf(w.encoding, length)
}
