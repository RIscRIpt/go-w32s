package w32s

import (
	"fmt"
	"syscall"
)

// StringBuffer represents a []byte or []uint16 buffer,
// if StringEncoding was set to Multibyte or Wide respectively.
// It is specially designed for using in Win32 API procedure calls, which
// require a string buffer and the buffer length as the parameters.
type StringBuffer struct {
	encoding StringEncoding
	size     int
	buffer   interface{}
}

func newStrBuf(encoding StringEncoding, size int) (sb *StringBuffer) {
	sb = &StringBuffer{
		encoding: encoding,
		size:     size,
	}
	if encoding == Wide {
		sb.buffer = make([]uint16, size)
		//sb.buffer = unsafe.Pointer(&(_buf.([]uint16))[0])
	} else {
		sb.buffer = make([]byte, size)
		//sb.buffer = unsafe.Pointer(&(_buf.([]byte))[0])
	}
	return
}

// Store copies specified string into the buffer.
func (sb *StringBuffer) Store(value string) {
	var bufLen int
	if sb.encoding == Wide {
		buf := syscall.StringToUTF16(value)
		bufLen = len(buf)
		if bufLen > sb.size {
			bufLen = sb.size
		}
		copy(sb.buffer.([]uint16)[0:bufLen], buf[0:bufLen])
	} else {
		bufLen = len(value)
		if bufLen > sb.size {
			bufLen = sb.size
		}
		copy(sb.buffer.([]byte)[0:bufLen], value[0:bufLen])
	}
	sb.size = bufLen
}

// Pointer returns a value which can be passed as a parameter to the Win32 API call.
func (sb *StringBuffer) Pointer() interface{} {
	return sb.buffer
}

// Size returns the buffer size.
func (sb *StringBuffer) Size() int {
	return sb.size
}

// SizePtr returns the pointer to the buffer size.
// This pointer can be passed as the parameter to the Win32 API call.
func (sb *StringBuffer) SizePtr() *int {
	return &sb.size
}

// String returns a string which is stored in the buffer.
func (sb *StringBuffer) String() string {
	if sb.encoding == Wide {
		bytes := sb.buffer.([]uint16)[0:sb.size]
		return syscall.UTF16ToString(bytes)
	} else if sb.encoding == Multibyte {
		bytes := sb.buffer.([]byte)[0:sb.size]
		for i, v := range bytes {
			if v == 0 {
				bytes = bytes[0:i]
				break
			}
		}
		return string(bytes)
	} else {
		panic(fmt.Errorf("w32s: unsupported encoding type: %v", sb.encoding))
	}
}
