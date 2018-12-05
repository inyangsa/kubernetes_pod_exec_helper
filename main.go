package main

import (
    "errors"
    "fmt"
    "os"
)

func main() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println(r)
        }
    }()

    serviceName := os.Args[1]
    fmt.Println("Let's see if I can get bash running on", serviceName, "...")

    podName, err := podName(serviceName)
    handleError(err)

    kubeCtlPath, err := kubectlLocation()
    handleError(err)

    err = runBashOnPod(kubeCtlPath, podName)
    handleError(err)

    fmt.Println("Goodbye :)")
}

func handleError(err error) {
    if err != nil {
        panic(err)
    }
}

func wrapError(snips ...interface{}) error {
    errMsg := ""

    for _, snip := range snips {
        errMsg = fmt.Sprintln(errMsg, "\n", snip)
    }

    return errors.New(errMsg)
}
