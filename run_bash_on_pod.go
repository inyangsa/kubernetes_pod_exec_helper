package main

import (
    "fmt"
    "os"
)

func runBashOnPod(kubectlPath, podName string) error {
    cwd, err := os.Getwd()
    if err != nil {
        panic(err)
    }

    pa := os.ProcAttr {
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
        Dir: cwd,
    }

    fmt.Println(">> Starting a new interactive shell")

    cmd := []string{"kubectl", "exec", "-it", podName, "/bin/bash"}
    proc, err := os.StartProcess(kubectlPath, cmd, &pa)
    if err != nil {
        panic(err)
    }

    state, err := proc.Wait()
    if err != nil {
        panic(err)
    }

    fmt.Println("<< Exited shell:", state.String())

    return nil
}
