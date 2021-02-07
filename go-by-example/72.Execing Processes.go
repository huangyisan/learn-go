package main

import (
    "os"
    "os/exec"
    "syscall"
)

func main() {
    //  syscall.Exec用系统本身的执行进程替代当前go的进程

    // 获取ls的绝对路径
    binary, err := exec.LookPath("ls")

    if err != nil {
        panic(err)
    }

    args := []string{"ls", "-a", "-l", "-h"}

    // 获取env
    env := os.Environ()

    // 使用系统原生命令
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
