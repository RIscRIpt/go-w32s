package w32s

import "syscall"

type dll struct {
	encoding StringEncoding
	*syscall.DLL
	procs map[string]*syscall.Proc
}

func newDll(sysdll *syscall.DLL, encoding StringEncoding) *dll {
	return &dll{
		encoding: encoding,
		DLL:      sysdll,
		procs:    make(map[string]*syscall.Proc),
	}
}

func (d *dll) proc(procname string) (proc *syscall.Proc, err error) {
	var ok bool
	if proc, ok = d.procs[procname]; !ok {
		proc, err = d.FindProc(procname)
		if err != nil {
			return nil, err
		}
		d.procs[procname] = proc
	}
	return proc, nil
}

func (d *dll) call(procname string, args ...interface{}) (res uintptr, err error) {
	proc, err := d.proc(procname)
	if err != nil {
		return ResultW32Error, err
	}
	res, _, err = proc.Call(d.cvt2uintptr(args...)...)
	return
}
