package awossgen

import (
    "log"
    "path"
    "runtime"
)

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
