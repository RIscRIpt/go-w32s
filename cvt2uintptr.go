package w32s

import (
	"fmt"
	"syscall"
	"unsafe"
)

func (d *dll) cvt2uintptr(iface ...interface{}) []uintptr {
	uiptr := make([]uintptr, len(iface))
	for i, x := range iface {
		switch v := x.(type) {
		case bool:
			if v {
				uiptr[i] = 1
			} else {
				uiptr[i] = 0
			}
		case []bool:
			bools := make([]int, len(v)) //typedef int BOOL
			for j, w := range v {
				if w {
					bools[j] = 1
				} else {
					bools[j] = 0
				}
			}
			uiptr[i] = uintptr(unsafe.Pointer(&bools[0]))
		case int8:
			uiptr[i] = uintptr(v)
		case uint8:
			uiptr[i] = uintptr(v)
		case int16:
			uiptr[i] = uintptr(v)
		case uint16:
			uiptr[i] = uintptr(v)
		case int32:
			uiptr[i] = uintptr(v)
		case uint32:
			uiptr[i] = uintptr(v)
		case int64:
			uiptr[i] = uintptr(v)
		case uint64:
			uiptr[i] = uintptr(v)
		case int:
			uiptr[i] = uintptr(v)
		case uint:
			uiptr[i] = uintptr(v)
		case uintptr:
			uiptr[i] = v
		case []int8:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []uint8:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []int16:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []uint16:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []int32:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []uint32:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []int64:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []uint64:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []int:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []uint:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case []uintptr:
			uiptr[i] = uintptr(unsafe.Pointer(&v[0]))
		case *int8:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *uint8:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *int16:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *uint16:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *int32:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *uint32:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *int64:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *uint64:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *int:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *uint:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case *uintptr:
			uiptr[i] = uintptr(unsafe.Pointer(v))
		case string:
			if d.encoding == Wide {
				uiptr[i] = uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(v)))
			} else {
				uiptr[i] = uintptr(unsafe.Pointer(&([]byte(v))[0]))
			}
		default:
			panic(fmt.Errorf("w32s: unsupported argument type %T", v))
		}
	}
	return uiptr
}
