package awodatagen

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
    baseDirPath := path.Join(filepath.ToSlash(goPath), "src", "github.com", "turnabout", "awodatagen", assetsDirName)

    // Add up all given directories to make up the full path
    var result string = baseDirPath

    for _, loopedPath := range paths {
        result = path.Join(result, loopedPath)
    }

    return result
}

// Counts amount of bits in a number (hardcoded for 32-bit numbers)
func CountBits(n uint) uint {
    n = ((0xaaaaaaaa & n) >> 1) + (0x55555555 & n)
    n = ((0xcccccccc & n) >> 2) + (0x33333333 & n)
    n = ((0xf0f0f0f0 & n) >> 4) + (0x0f0f0f0f & n)
    n = ((0xff00ff00 & n) >> 8) + (0x00ff00ff & n)
    n = ((0xffff0000 & n) >> 16) + (0x0000ffff & n)
    return n
}
