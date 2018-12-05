package main

import (
    "os/exec"
    "strings"
)

func kubectlLocation() (string, error) {
    b, err := exec.Command("which", "kubectl").Output()
    if err != nil {
        return "", wrapError("failed to get kubectl path", err)
    }

    return strings.TrimSpace(string(b)), nil
}
