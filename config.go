//
// Configuration for the project
//
package awodatagen


//
// Environment variables
//

// Env variable holding the path where the sprite sheet should be output
const SSOutputEnvVar = "AWO_SPRITESHEET"

// Env variable holding the path where the data file should be output
const JSONOutputEnvVar = "AWO_JSON"

// Env variable holding absolute path to directory containing the raw assets
const AssetsDirPath = "AWO_ASSETS_PATH"


//
// Output paths
//

// Default name for the sprite sheet output
const SSOutputDefaultName = "spritesheet.png"

// Default name for the data file output
const JSONOutputDefaultName = "gamedata.json"


//
// Input paths
//

// Base inputs directory, containing all images & data files
const assetsDirName = "assets"

// Inputs subdirectories
const UnitsDir       = "units"
const TilesDir       = "tiles"
const PropertiesDir  = "properties"
const UIDir          = "ui"
const CODir          = "co"
const OtherDir       = "other"
const FramesDir      = "frames"

// Names of extra data files
const PalettesFileName              = "palettes.json"
const BasePalettesFileName          = "basePalettes.json"
const StagesFileName                = "stages.json"
const ClocksFileName                = "clocks.json"
const TilesClockDataFileName        = "tilesClockData.json"
const TilesAutoVarFileName          = "tilesAutoVarData.json"
const TilesPlacementRulesFileName   = "tilesPlacementRules.json"
const UnitDataFileName              = "data.json"
const WeaponTypesFileName           = "weaponTypes.json"


//
// Other configuration
//

// Size of a regular tile
const RegularTileDimension int = 16
