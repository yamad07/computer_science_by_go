package main

import (
    "fmt"
    "os"
    "syscall"
)

func main() {
    session_id, _ := syscall.Getsid(os.Getpid())
    fmt.Fprintf(os.Stderr, "Process Group ID:%d \nSession Group ID:%d\n", syscall.Getpgrp(), session_id)
}
