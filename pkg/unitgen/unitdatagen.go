package unitgen

import (
	"fmt"
	"github.com/turnabout/awodatagen/internal/config"
	"github.com/turnabout/awodatagen/internal/genio"
	"github.com/turnabout/awodatagen/internal/packer"
	"github.com/turnabout/awodatagen/internal/utilities"
	"github.com/turnabout/awodatagen/pkg/framedata"
	"os"
)

type rawUnitData struct {
	MovementType    string `json:"movementType"`
	Movement        uint8  `json:"movement"`
	Vision          uint8  `json:"vision"`
	Fuel            uint8  `json:"fuel"`
	WeaponPrimary   string `json:"weaponPrimary"`
	WeaponSecondary string `json:"weaponSecondary"`
}

// Generates units game data.
func GetUnitData(packedFrameImgs *[]packer.FrameImage) *UnitData {

	var unitData UnitData

	unitData.UnitTypesData = *getUnitTypesData(packedFrameImgs)
	unitData.WeaponTypesData = *getWeaponTypesData()

	return &unitData
}

// Generates all unit types data, including origin visual data (units' visual data on the raw
// sprite sheet) using packed Frame Images
func getUnitTypesData(packedFrameImgs *[]packer.FrameImage) *UnitTypesData {

	var unitsData UnitTypesData

	// Add frames
	for _, frameImg := range *packedFrameImgs {

		// Ignore non-unit frame images
		if frameImg.MetaData.FrameImageDataType != uint8(framedata.UnitDataType) {
			continue
		}

		unitType := frameImg.MetaData.Type
		unitVar := frameImg.MetaData.Variation
		unitAnim := frameImg.MetaData.Animation
		unitFrame := frameImg.MetaData.Index

		// Check if variation slice is missing, add up to it if necessary
		missingVars := int(unitVar+1) - len(unitsData[unitType].Variations)

		if missingVars > 0 {
			for i := 0; i < missingVars; i++ {
				unitsData[unitType].Variations = append(
					unitsData[unitType].Variations,
					[][]framedata.Frame{},
				)
			}
		}

		// Check if animation slice is missing, add up to it if necessary
		missingAnims := int(unitAnim+1) - len(unitsData[unitType].Variations[unitVar])

		if missingAnims > 0 {
			for i := 0; i < missingAnims; i++ {
				unitsData[unitType].Variations[unitVar] = append(
					unitsData[unitType].Variations[unitVar],
					[]framedata.Frame{},
				)
			}
		}

		// Check if animation frame is missing, add up to it if necessary
		missingFrames := int(unitFrame+1) - len(unitsData[unitType].Variations[unitVar][unitAnim])

		if missingFrames > 0 {
			for i := 0; i < missingFrames; i++ {
				unitsData[unitType].Variations[unitVar][unitAnim] = append(unitsData[unitType].Variations[unitVar][unitAnim], framedata.Frame{})
			}
		}

		// Store data
		if frameImg.X == 48 && frameImg.Y == 64 {
			fmt.Printf("%#v\n", frameImg)
		}

		unitsData[unitType].Variations[unitVar][unitAnim][unitFrame] = framedata.Frame{
			X:      frameImg.X,
			Y:      frameImg.Y,
			Width:  frameImg.Width,
			Height: frameImg.Height,
		}
	}

	// Load other data from the unit's source, raw JSON data
	for unitType := UnitTypeFirst; unitType <= UnitTypeLast; unitType++ {
		var rawData rawUnitData

		rawDataPath := utilities.GetInputPath(
			config.UnitsDir,
			unitType.String(),
			config.UnitDataFileName,
		)

		// Ensure raw data file exists
		if _, err := os.Stat(rawDataPath); os.IsNotExist(err) {
			utilities.LogFatalF(
				"Unit '%s' raw data file path '%s' is invalid",
				unitType.String(),
				rawDataPath,
			)
		}

		genio.AttachJSONData(rawDataPath, &rawData)
		unitsData[unitType].Fuel = rawData.Fuel
		unitsData[unitType].Movement = rawData.Movement
		unitsData[unitType].Vision = rawData.Vision

		var weaponPrimary WeaponType
		var weaponSecondary WeaponType
		var movementType MovementType
		var ok bool

		weaponPrimary, ok = WeaponTypeReverseStrings[rawData.WeaponPrimary]

		if !ok && rawData.WeaponPrimary != "" {
			utilities.LogFatalF(
				"Missing or invalid primary weapon type '%s' on unit '%s'",
				rawData.WeaponPrimary,
				unitType.String(),
			)
		}

		weaponSecondary, ok = WeaponTypeReverseStrings[rawData.WeaponSecondary]
		if !ok && rawData.WeaponSecondary != "" {
			utilities.LogFatalF(
				"Missing or invalid secondary weapon type '%s' on unit '%s'",
				rawData.WeaponSecondary,
				unitType.String(),
			)
		}

		if movementType, ok = MovementTypeReverseStrings[rawData.MovementType]; !ok {
			utilities.LogFatalF(
				"Missing or invalid movement type '%s' on unit '%s'",
				rawData.MovementType,
				unitType.String(),
			)
		}

		unitsData[unitType].WeaponPrimary = weaponPrimary
		unitsData[unitType].WeaponSecondary = weaponSecondary
		unitsData[unitType].MovementType = movementType
	}

	return &unitsData
}

// Generates all weapon types data
func getWeaponTypesData() *[WeaponTypeCount]WeaponTypeData {

	// Get raw weapon types data map
	var rawData map[string]WeaponTypeData

	dataPath := utilities.GetInputPath(
		config.UnitsDir,
		config.WeaponTypesFileName,
	)

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		utilities.LogFatalF("Weapon types file path '%s' invalid", dataPath)
	}

	genio.AttachJSONData(dataPath, &rawData)

	// Transform raw map into processed array
	var data [WeaponTypeCount]WeaponTypeData

	for wTypeStr, wData := range rawData {
		var wType WeaponType
		var ok bool

		if wType, ok = WeaponTypeReverseStrings[wTypeStr]; !ok {
			utilities.LogFatalF(
				"Unknown weapon type '%s' found in data",
				wTypeStr,
			)
		}

		data[wType] = wData
	}

	return &data
}
