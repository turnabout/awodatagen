// Stores configuration for the project
package main

import (
    "path"
    "runtime"
)

// Default name for the resulting spritesheet output
const ssOutputDefaultName = "spritesheet.png"

// Environment variable holding the path where the sprite sheet should be output
const ssOutputEnvVar = "AWO_SPRITESHEET"

// Default name for the resulting spritesheet output
const jsonOutputDefaultName = "visuals.json"

// Environment variable holding the path where the sprite sheet should be output
const jsonOutputEnvVar = "AWO_JSON"

// Base directory containing all sprite sheet images & visual data files
const inputsDirName     = "/raw_inputs"
const unitsDirName      = "/units"
const tilesDirName      = "/tiles"
const propertiesDirName = "/properties"

// Name of extra data files
const palettesFileName    = "/palettes.json"
const basePaletteFileName = "/basePalette.json"
const propsLightsOnColor  = "/lightsOnColor.json"

// Size of a regular Tile
const regularTileDimension int = 16

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
