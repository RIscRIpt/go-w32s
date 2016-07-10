# W32s
Win32 '**s**tring' API -- a Go package which can be used to call procedures of DLLs of Win32 using only their _string_ names.
## Examples
### printf call
```go
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
```
### GetComputerNameW call
```go
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
```
### MessageBoxW call
```go
func ExampleW32s_Call_MessageBoxW() {
	w32s := New(Wide)
	if err := w32s.LoadDLL("user32.dll"); err != nil {
		// Handle an error here
		panic(err)
	}

	n, err := w32s.Call(
		"user32.dll",
		"MessageBoxW",
		0,
		"Hello, World!",
		"Greeting",
		0,
	)
	if err != Success {
		// Handle an error here
		panic(err)
	}
	fmt.Printf("user has pressed OK button [#%d] (he had no other choice)\n", n)

	if err := w32s.ReleaseDLL("user32.dll"); err != nil {
		// Handle an error here
		panic(err)
	}

	// Output: user has pressed OK button [#1] (he had no other choice)
}
```
# Alternatives to this package
If you are looking for a package which has WinDefs, type names and real go-code methods for Win32 API, then you should consider using package [AllenDang/w32](https://github.com/AllenDang/w32) instead.

