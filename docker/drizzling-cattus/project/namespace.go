package main

import (
  "fmt"
  "os"
  "os/exec"
  "syscall"
)

func main() {
  cmd := exec.Command("/bin/bash");

  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr

  cmd.Env = append(cmd.Env, `PS1=$`)

  //cmd.SysProcAttr = &syscall.SysProcAttr{
  //  Cloneflags: syscall.CLONE_NEWUTS,
  //}

  fmt.Printf("Running in prev namespace %d\n", os.Getpid())

  // https://medium.com/@teddyking/namespaces-in-go-network-fdcf63e76100
  syscall.Unshare(syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNET)

  syscall.Sethostname([]byte("hello world"))

  syscall.Chroot(os.Args[1]);
  syscall.Chdir("/");

  err := cmd.Run()

  fmt.Println("error is ", err)
}
