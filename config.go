//
// Configuration for the project
//
package awodatagen


//
// Environment variables
//

// Environment variable holding the path where the sprite sheet should be output
const SSOutputEnvVar = "AWO_SPRITESHEET"

// Environment variable holding the path where the data file should be output
const JSONOutputEnvVar = "AWO_JSON"


//
// Output paths
//

// Default name for the sprite sheet output
const SSOutputDefaultName = "spritesheet.png"

// Default name for the data file output
const JSONOutputDefaultName = "visuals.json"


//
// Input paths
//

// Base inputs directory, containing all images & data files
const assetsDirName = "assets"

// Inputs subdirectories
const UnitsDir       = "units"
const TilesDir       = "tiles"
const PropertiesDir  = "properties"
const UIDir = "ui"
const AdditionalDir  = "additional"

// Names of extra data files
const PalettesFileName       = "palettes.json"
const BasePalettesFileName   = "basePalettes.json"
const StagesFileName         = "stages.json"
const AnimClocksFileName     = "animationClocks.json"
const TilesClockDataFileName = "tilesClockData.json"
const TilesAutoVarFileName   = "tilesAutoVarData.json"


//
// Other configuration
//

// Size of a regular tile
const RegularTileDimension int = 16
