package getsyscallid

import (
	"fmt"
	"syscall"
	"unsafe"
)

type (
	BOOL    uint32
	LPCSTR  uintptr
	HMODULE uintptr
	FARPROC uintptr
	HANDLE  uintptr
	SIZE_T  uintptr
	LPCVOID uintptr
	LPVOID  uintptr
)

func GetModuleHandle(moduleName string) (HMODULE, error) {
	lpModuleName_, err := syscall.BytePtrFromString(moduleName)
	if err != nil {
		return 0, err
	}
	lpModuleName := LPCSTR(unsafe.Pointer(lpModuleName_))
	hModule := _GetModuleHandleA(lpModuleName)
	if hModule == 0 {
		return 0, fmt.Errorf("call GetModuleHandle failed")
	}
	return hModule, nil
}

func GetProcAddress(hModule HMODULE, procName string) (FARPROC, error) {
	lpProcName_, err := syscall.BytePtrFromString(procName)
	if err != nil {
		return 0, err
	}
	lpProcName := LPCSTR(unsafe.Pointer(lpProcName_))
	pFunc := _GetProcAddress(hModule, lpProcName)
	if pFunc == 0 {
		return 0, fmt.Errorf("call GetProcAddress error")
	}
	return pFunc, nil
}

func GetCurrentProcess() (HANDLE, error) {
	handle := _GetCurrentProcess()
	if handle == 0 {
		return 0, fmt.Errorf("call GetCurrentProcess error")
	}
	return handle, nil
}

func ReadProcessMemory(hProcess HANDLE, lpBaseAddress LPCVOID, size int64) (buffer []byte, err error) {
	buffer = make([]byte, size)
	nSize := SIZE_T(size)
	lpNumberOfBytesRead := SIZE_T(0)
	ret := _ReadProcessMemory(
		hProcess,
		lpBaseAddress,
		LPVOID(unsafe.Pointer(&buffer[0])),
		nSize,
		&lpNumberOfBytesRead,
	)
	if ret == 0 {
		return buffer, fmt.Errorf("call ReadProcessMemory error")
	}

	return
}

//sys _GetModuleHandleA(lpModuleName LPCSTR) (hModule HMODULE) = kernel32.GetModuleHandleA
//sys _GetProcAddress(hModule HMODULE, lpProcName LPCSTR) (pFunc FARPROC) = kernel32.GetProcAddress
//sys _GetCurrentProcess() (handle HANDLE) = kernel32.GetCurrentProcess
//sys _ReadProcessMemory(hProcess HANDLE, lpBaseAddress LPCVOID, lpBuffer LPVOID, nSize SIZE_T, lpNumberOfBytesRead *SIZE_T) (ret BOOL) = kernel32.ReadProcessMemory
