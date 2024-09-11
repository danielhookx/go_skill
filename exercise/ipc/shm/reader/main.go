package main

import (
	"flag"
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

const (
	// IpcCreate create if key is nonexistent
	IpcCreate = 00001000
)

var id = flag.Int("id", 0, "0:write 1:read")

func main() {
	flag.Parse()
	shmid := uintptr(*id)
	fmt.Printf("shmid: %v\n", shmid)

	shmaddr, _, err := syscall.Syscall(syscall.SYS_SHMAT, shmid, 0, 0)
	if err != 0 {
		fmt.Printf("syscall error, err: %v\n", err)
		os.Exit(-2)
	}
	fmt.Printf("shmaddr: %v\n", shmaddr)

	defer syscall.Syscall(syscall.SYS_SHMDT, shmaddr, 0, 0)

	data := (*string)(unsafe.Pointer(uintptr(shmaddr)))

	fmt.Println(*data)
}
