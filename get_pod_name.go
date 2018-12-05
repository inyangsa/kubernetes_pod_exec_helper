package main

import (
    "fmt"
    "os/exec"
    "regexp"
)

func podName(serviceName string) (string, error) {
    grepRegex := fmt.Sprint("(?m)^", serviceName, "-[A-z0-9]*-[A-z0-9]*")
    b, err := exec.Command("kubectl", "get", "pods").Output()
    if err != nil {
        return "", wrapError("failed to get pods from bash command", err)
    }

    podData := string(b)
    r, err := regexp.Compile(grepRegex)
    if err != nil {
        return "", wrapError("failed to create service name regex", err)
    }

    return r.FindString(podData), nil
}
