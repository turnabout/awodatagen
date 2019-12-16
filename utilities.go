package awossgen

import (
    "fmt"
    "log"
    "os"
    "path"
    "runtime"
    "runtime/debug"
)

func LogFatalIfErr(err error, msgs ...string) {
    if err != nil {
        data := append([]string{err.Error()}, msgs...)
        LogFatal(data)
    }
}

func LogFatal(msgs []string) {

    fmt.Println("Fatal error:")

    for _, loopedMsg := range msgs {
        fmt.Println(loopedMsg)
    }

    fmt.Println("Stack trace:")
    debug.PrintStack()

    os.Exit(1)

}

// Gets the full path to a directory in the project's inputs
func GetInputPath(paths ...string) string {

    // Get the base directory path of the project
    var baseDirPath string
    _, fileName, _, ok := runtime.Caller(0)

    if !ok {
        log.Fatal("getFullProjectPath: No caller information")
    }

    baseDirPath = path.Dir(fileName)

    // Add up all given directories to make up the full path
    var fullPath string = path.Join(baseDirPath, inputsDirName)

    for _, loopedPath := range paths {
        fullPath = path.Join(fullPath, loopedPath)
    }

    return fullPath
}
