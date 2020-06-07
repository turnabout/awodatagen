package config

const (

	//
	// Environment variables
	//

	// Env variable holding the path where the sprite sheet should be output
	SSOutputEnvVar = "AWO_SPRITESHEET"

	// Env variable holding the path where the data file should be output
	JSONOutputEnvVar = "AWO_JSON"

	// Env variable holding absolute path to directory containing the raw assets
	AssetsDirPath = "AWO_ASSETS_PATH"

	// Output paths
	// Default name for the sprite sheet output
	SSOutputDefaultName = "spritesheet.png"

	// Default name for the data file output
	JSONOutputDefaultName = "gamedata.json"

	//
	// Input paths
	//

	// Base inputs directory, containing all images & data files
	AssetsDirName = "assets"

	// Inputs subdirectories
	UnitsDir      = "units"
	TilesDir      = "tiles"
	PropertiesDir = "properties"
	UIDir         = "ui"
	CODir         = "co"
	OtherDir      = "other"
	FramesDir     = "frames"

	// Input data files
	PalettesFileName            = "palettes.json"
	BasePalettesFileName        = "basePalettes.json"
	StagesFileName              = "stages.json"
	ClocksFileName              = "clocks.json"
	TilesClockDataFileName      = "tilesClockData.json"
	TilesAutoVarFileName        = "tilesAutoVarData.json"
	TilesPlacementRulesFileName = "tilesPlacementRules.json"
	UnitDataFileName            = "data.json"
	WeaponTypesFileName         = "weaponTypes.json"

	//
	// Other configuration
	//

	// Size of a regular tile
	RegularTileDimension int = 16
)
