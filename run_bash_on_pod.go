package main

import (
    "fmt"
    "os"
)

func runBashOnPod(kubectlPath, podName string) error {
    pa := os.ProcAttr {
        Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
    }

    fmt.Println(">> Starting a new interactive shell")

    cmd := []string{"kubectl", "exec", "-it", podName, "/bin/bash"}
    proc, err := os.StartProcess(kubectlPath, cmd, &pa)
    if err != nil {
        return wrapError("failed to start bash on pod", err)
    }

    state, err := proc.Wait()
    if err != nil {
        return wrapError("failed to wait on process execution")
    }

    fmt.Println("<< Exited shell:", state.String())

    return nil
}
