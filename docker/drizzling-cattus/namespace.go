package main

import (
  "fmt"
  "os"
  "os/exec"
  "syscall"
)

func main() {
  cmd := exec.Command("/bin/sh");

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  cmd.Env = append(cmd.Env, `PS1=$`)

  cmd.SysProcAttr = &syscall.SysProcAttr{
    Cloneflags: syscall.CLONE_NEWUTS,
  }

  syscall.Sethostname([]byte("hello world"))

  err := cmd.Run()
  
  fmt.Println("error is ", err)
}
