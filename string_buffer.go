package w32s

import "syscall"

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

func (sb *StringBuffer) Pointer() interface{} {
	return sb.buffer
}

func (sb *StringBuffer) Size() int {
	return sb.size
}

func (sb *StringBuffer) SizePtr() *int {
	return &sb.size
}

func (sb *StringBuffer) String() string {
	if sb.encoding == Wide {
		return syscall.UTF16ToString(sb.buffer.([]uint16))
	} else {
		bytes := sb.buffer.([]byte)[0:sb.size]
		for i, v := range bytes {
			if v == 0 {
				bytes = bytes[0:i]
				break
			}
		}
		return string(bytes)
	}
}
