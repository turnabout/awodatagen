//
// Stores configuration for the project
//
package main

// Default name for the resulting spritesheet output
const ssOutputDefaultName = "spritesheet.png"

// Environment variable holding the path where the sprite sheet should be output
const ssOutputEnvVar = "AWO_SPRITESHEET"

// Default name for the resulting spritesheet output
const jsonOutputDefaultName = "visuals.json"

// Environment variable holding the path where the sprite sheet should be output
const jsonOutputEnvVar = "AWO_JSON"

// Base inputs directory, containing all images & data files
const inputsDirName = "inputs"

// Inputs subdirectories
const unitsDir       = "units"
const tilesDir       = "tiles"
const propertiesDir  = "properties"
const uiDir          = "ui"
const additionalDir  = "additional"

// Names of extra data files
const palettesFileName       = "palettes.json"
const basePalettesFileName   = "basePalettes.json"
const stagesFileName         = "stages.json"
const animClocksFileName     = "animationClocks.json"
const tilesClockDataFileName = "tilesClockData.json"
const tilesAutoVarFileName   = "tilesAutoVarData.json"

// Size of a regular tile
const regularTileDimension int = 16
