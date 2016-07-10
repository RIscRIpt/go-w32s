package w32s

import (
	"fmt"
	"testing"
)

func TestStringBufferMB(t *testing.T) {
	w32s := New(Multibyte)
	buf := w32s.StrBuf(8)
	buf.Store("Golang")
	if buf.String() != "Golang" {
		t.Fatalf("buf.String() / Multibyte: expected 'Golang', got '%s'", buf.String())
	}
}

func TestStringBufferWC(t *testing.T) {
	w32s := New(Wide)
	buf := w32s.StrBuf(8)
	buf.Store("Golang")
	if buf.String() != "Golang" {
		t.Fatalf("buf.String() / Wide: expected 'Golang', got '%s'", buf.String())
	}
}

func TestLoadRelease(t *testing.T) {
	w32s := New(Multibyte)
	if err := w32s.LoadDLL("user32.dll"); err != nil {
		t.Fatalf("LoadDLL returned: %+v", err)
	}
	if err := w32s.ReleaseDLL("user32.dll"); err != nil {
		t.Fatalf("ReleaseDLL returned: %+v", err)
	}
}

func TestLoadNonexistent(t *testing.T) {
	w32s := New(Wide)
	if err := w32s.LoadDLL("____THIS_LIBRARY_MUST_NOT_EXIST____.dll"); err == nil {
		t.Fatalf("LoadDLL didn't return a error for nonexistent library")
	}
}

func TestCallNotLoadedLib(t *testing.T) {
	w32s := New(Multibyte)
	if _, err := w32s.Call("user32.dll", "MessageBoxA", 0, 0, 0, 0); err != ErrDllNotLoaded {
		t.Fatalf("Call of not loaded DLL returned unexpected error: %+v", err)
	}
}

func TestGetModuleHandleW(t *testing.T) {
	w32s := New(Wide)
	if err := w32s.LoadDLL("kernel32.dll"); err != nil {
		t.Fatalf("LoadDLL returned: %+v", err)
	}
	hKernel32, err := w32s.Call(
		"kernel32.dll",
		"GetModuleHandleW",
		"kernel32.dll",
	)
	if err != Success {
		t.Fatalf("Call kernel32.GetModuleHandleW returned an error: %+v", err)
	}
	if uintptr(w32s.dlls["kernel32.dll"].Handle) != hKernel32 {
		t.Fatalf("Golang's kernel32.dll handle != GetModuleHandleW(L\"kernel32.dll\")")
	}
	if err := w32s.ReleaseDLL("kernel32.dll"); err != nil {
		t.Fatalf("ReleaseDLL returned: %+v", err)
	}
}

func ExampleW32s_Call_printf() {
	w32s := New(Multibyte)
	if err := w32s.LoadDLL("msvcrt.dll"); err != nil {
		// Handle an error here
		panic(err)
	}

	n, err := w32s.Call(
		"msvcrt.dll",
		"printf",
		"Hello, %s!\n",
		"world",
	)
	if err != Success {
		// Handle an error here
		panic(err)
	}
	fmt.Printf("printf just printed a string of length %d into your console!\n", n)

	if err := w32s.ReleaseDLL("msvcrt.dll"); err != nil {
		// Handle an error here
		panic(err)
	}

	// Output: printf just printed a string of length 14 into your console!
}

func ExampleW32s_Call_GetComputerNameW() {
	w32s := New(Wide)
	if err := w32s.LoadDLL("kernel32.dll"); err != nil {
		// Handle an error here
		panic(err)
	}

	buf := w32s.StrBuf(16)
	n, err := w32s.Call(
		"kernel32.dll",
		"GetComputerNameW",
		buf.Pointer(),
		buf.SizePtr(),
	)
	if n == 0 && err != Success {
		// Handle an error here
		panic(err)
	}
	//fmt.Printf("Your computer name: %s", buf.String())

	if err := w32s.ReleaseDLL("msvcrt.dll"); err != nil {
		// Handle an error here
		panic(err)
	}

	// Output:
}

// func ExampleW32s_Call_MessageBoxW() {
// 	w32s := New(Wide)
// 	if err := w32s.LoadDLL("user32.dll"); err != nil {
// 		// Handle an error here
// 		panic(err)
// 	}

// 	n, err := w32s.Call(
// 		"user32.dll",
// 		"MessageBoxW",
// 		0,
// 		"Hello, World!",
// 		"Greeting",
// 		0,
// 	)
// 	if err != Success {
// 		// Handle an error here
// 		panic(err)
// 	}
// 	fmt.Printf("user has pressed OK button [#%d] (he had no other choice)\n", n)

// 	if err := w32s.ReleaseDLL("user32.dll"); err != nil {
// 		// Handle an error here
// 		panic(err)
// 	}

// 	// Output: user has pressed OK button [#1] (he had no other choice)
// }
