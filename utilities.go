package awodatagen

import (
    "fmt"
    "os"
    "path"
    "path/filepath"
    "runtime/debug"
)

// Full, absolute path to the project's directory containing the raw assets
var assetsFullPath string

func LogFatalIfErr(err error) {
    if err != nil {
        LogFatalF("Error: %s", err.Error())
    }
}

// Helper to log a fatal error with a stack trace, then exit the program
func LogFatalF(format string, a ...interface{}) {
    fmt.Println("Fatal error:")
    fmt.Printf(format, a)
    fmt.Println("Stack trace:")
    debug.PrintStack()
    os.Exit(1)
}

// Attempt to set the full "assets" path from an environment variable
// If it doesn't exist, use the CWD as the base path to assets
func init() {
    var envExists bool

    if assetsFullPath, envExists = os.LookupEnv(AssetsDirPath); !envExists {
        cwd, err := os.Getwd()

        if err != nil {
            LogFatalIfErr(err)
        }

        // Use the project's assets path as a base
        assetsFullPath = path.Join(filepath.ToSlash(cwd), assetsDirName)
    }
}

// Gets the full path to a directory in the project's inputs
func GetInputPath(paths ...string) string {

    // Add up all given directories to make up the full path
    result := assetsFullPath

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
