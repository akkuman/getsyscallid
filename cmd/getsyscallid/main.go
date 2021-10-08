package main

import (
	"flag"
	"fmt"

	"github.com/akkuman/getsyscallid"
)

var (
	procName string
)

func init() {
	flag.StringVar(&procName, "proc", "NtCreateProcess", "The name of the api name from ntdll.dll")
}

func main() {
	flag.Parse()
	sysID, err := getsyscallid.GetID(procName)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("the syscall Number is: %x\n", sysID)
}
