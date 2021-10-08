package getsyscallid

import "encoding/binary"

func GetID(procName string) (uint32, error) {
	hModule, err := GetModuleHandle("ntdll.dll")
	if err != nil {
		return 0, err
	}
	funcAddr, err := GetProcAddress(hModule, procName)
	if err != nil {
		return 0, err
	}
	sysNumberAddr := funcAddr + 4
	curProc, err := GetCurrentProcess()
	if err != nil {
		return 0, err
	}
	buf, err := ReadProcessMemory(curProc, LPCVOID(sysNumberAddr), 4)
	if err != nil {
		return 0, err
	}
	sysID := binary.LittleEndian.Uint32(buf)
	return sysID, nil
}
