package awossgen

import (
    "fmt"
    "go/build"
    "os"
    "path"
    "path/filepath"
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

    // Get $GOPATH
    goPath := os.Getenv("GOPATH")

    if goPath == "" {
        goPath = build.Default.GOPATH
    }

    // Use the project's assets path as a base
    baseDirPath := path.Join(filepath.ToSlash(goPath), "src", "github.com", "turnabout", "awossgen", assetsDirName)

    // Add up all given directories to make up the full path
    var result string = baseDirPath

    for _, loopedPath := range paths {
        result = path.Join(result, loopedPath)
    }

    return result
}
