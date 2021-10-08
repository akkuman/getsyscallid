// Code generated by 'go generate'; DO NOT EDIT.

package getsyscallid

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
	errERROR_EINVAL     error = syscall.EINVAL
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return errERROR_EINVAL
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")

	procGetCurrentProcess = modkernel32.NewProc("GetCurrentProcess")
	procGetModuleHandleA  = modkernel32.NewProc("GetModuleHandleA")
	procGetProcAddress    = modkernel32.NewProc("GetProcAddress")
	procReadProcessMemory = modkernel32.NewProc("ReadProcessMemory")
)

func _GetCurrentProcess() (handle HANDLE) {
	r0, _, _ := syscall.Syscall(procGetCurrentProcess.Addr(), 0, 0, 0, 0)
	handle = HANDLE(r0)
	return
}

func _GetModuleHandleA(lpModuleName LPCSTR) (hModule HMODULE) {
	r0, _, _ := syscall.Syscall(procGetModuleHandleA.Addr(), 1, uintptr(lpModuleName), 0, 0)
	hModule = HMODULE(r0)
	return
}

func _GetProcAddress(hModule HMODULE, lpProcName LPCSTR) (pFunc FARPROC) {
	r0, _, _ := syscall.Syscall(procGetProcAddress.Addr(), 2, uintptr(hModule), uintptr(lpProcName), 0)
	pFunc = FARPROC(r0)
	return
}

func _ReadProcessMemory(hProcess HANDLE, lpBaseAddress LPCVOID, lpBuffer LPVOID, nSize SIZE_T, lpNumberOfBytesRead *SIZE_T) (ret BOOL) {
	r0, _, _ := syscall.Syscall6(procReadProcessMemory.Addr(), 5, uintptr(hProcess), uintptr(lpBaseAddress), uintptr(lpBuffer), uintptr(nSize), uintptr(unsafe.Pointer(lpNumberOfBytesRead)), 0)
	ret = BOOL(r0)
	return
}