package main

import (
    "fmt"
    "github.com/go-errors/errors"
    "os"
)

func main() {
    serviceName := os.Args[1]
    fmt.Println("Let's see if I can get bash running on", serviceName, "...")

    podName, err := podName(serviceName)
    handleError(err)

    kubeCtlPath, err := kubectlLocation()
    handleError(err)

    fmt.Println(kubeCtlPath)

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
