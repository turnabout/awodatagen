// Stores configuration for the project
package main

import (
    "path"
    "runtime"
)

// Default name for the resulting spritesheet output
const ssOutputDefaultName string = "spritesheet.png"

// Environment variable holding the path where the sprite sheet should be output
const ssOutputEnvVar string = "AWO_SPRITESHEET"

// Default name for the resulting spritesheet output
const jsonOutputDefaultName string = "visuals.json"

// Environment variable holding the path where the sprite sheet should be output
const jsonOutputEnvVar string = "AWO_JSON"

// Base directory containing all spritesheet images
const imageInputsDirName string = "/raw_inputs"

// Directory containing unit images
const unitsDirName string = "/units"

// Directory containing tile images
const tilesDirName string = "/tiles"

// Directory containing property images
const propertiesDirName string = "/properties"

// Width (in pixels) taken up by units/tiles on sprite sheet
const unitsSSWidth int = 170
const tilesSSWidth int = 200

// The base path of this project
var baseDirPath string = getDirPath()

// Grab this directory's full path
func getDirPath() string {
    // Grab awssgen's directory path
    _, filename, _, ok := runtime.Caller(0)

    if !ok {
        panic("No caller information")
    }

    return path.Dir(filename)
}
