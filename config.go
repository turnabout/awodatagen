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

// Size of a regular Tile
const regularTileDimension int = 16

// Directory containing property images
const propertiesDirName string = "/properties"

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
