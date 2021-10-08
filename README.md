# GetSyscallID

get windows system call number dynamically, it doesn't need the system call table.

## Try

### download

```shell
go get github.com/akkuman/getsyscallid/cmd/getsyscallid
```

### help

```shell
> ./getsyscallid.exe -help    
Usage of getsyscallid.exe:
  -proc string
        The name of the api name from ntdll.dll (default "NtCreateProcess")
```

### run

```shell
> ./get.exe -proc NtCreateProcess
the syscall Number is: b9
```

## As a package

```go
package main

import (
	"flag"
	"fmt"

	"github.com/akkuman/getsyscallid"
)

func main() {
	flag.Parse()
	sysID, err := getsyscallid.GetID("NtCreateProcess")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("the syscall Number is: %x\n", sysID)
}
```

## Reference

- [动态获取系统调用(syscall)号](https://idiotc4t.com/defense-evasion/dynamic-get-syscallid)