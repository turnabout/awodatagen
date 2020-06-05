package unitgen

import (
    "fmt"
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
    "github.com/turnabout/awodatagen/pkg/packer"
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
func GetUnitData(packedFrameImgs *[]packer.FrameImage)  *awodatagen.UnitData {

    var unitData awodatagen.UnitData

    unitData.UnitTypesData = *getUnitTypesData(packedFrameImgs)
    unitData.WeaponTypesData = *getWeaponTypesData()


    return &unitData
}

// Generates all unit types data, including origin visual data (units' visual data on the raw
// sprite sheet) using packed Frame Images
func getUnitTypesData(packedFrameImgs *[]packer.FrameImage) *awodatagen.UnitTypesData {

    var unitsData awodatagen.UnitTypesData

    // Add frames
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-unit frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awodatagen.UnitDataType) {
            continue
        }

        unitType := frameImg.MetaData.Type
        unitVar := frameImg.MetaData.Variation
        unitAnim := frameImg.MetaData.Animation
        unitFrame := frameImg.MetaData.Index

        // Check if variation slice is missing, add up to it if necessary
        missingVars := int(unitVar + 1) - len(unitsData[unitType].Variations)

        if missingVars > 0 {
            for i := 0; i < missingVars; i++ {
                unitsData[unitType].Variations = append(
                    unitsData[unitType].Variations,
                    [][]awodatagen.Frame{},
                )
            }
        }

        // Check if animation slice is missing, add up to it if necessary
        missingAnims := int(unitAnim + 1) - len(unitsData[unitType].Variations[unitVar])

        if missingAnims > 0 {
            for i := 0; i < missingAnims; i++ {
                unitsData[unitType].Variations[unitVar] = append(
                    unitsData[unitType].Variations[unitVar],
                    []awodatagen.Frame{},
                )
            }
        }

        // Check if animation frame is missing, add up to it if necessary
        missingFrames := int(unitFrame + 1) - len(unitsData[unitType].Variations[unitVar][unitAnim])

        if missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                unitsData[unitType].Variations[unitVar][unitAnim] = append(unitsData[unitType].Variations[unitVar][unitAnim], awodatagen.Frame{})
            }
        }

        // Store data
        if frameImg.X == 48 && frameImg.Y == 64 {
            fmt.Printf("%#v\n", frameImg)
        }

        unitsData[unitType].Variations[unitVar][unitAnim][unitFrame] = awodatagen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }
    }

    // Load other data from the unit's source, raw JSON data
    for unitType := awodatagen.UnitTypeFirst; unitType <= awodatagen.UnitTypeLast; unitType++ {
        var rawData rawUnitData

        rawDataPath := awodatagen.GetInputPath(
            awodatagen.UnitsDir,
            unitType.String(),
            awodatagen.UnitDataFileName,
        )

        // Ensure raw data file exists
        if _, err := os.Stat(rawDataPath); os.IsNotExist(err) {
            awodatagen.LogFatalF(
                  "Unit '%s' raw data file path '%s' is invalid",
                unitType.String(),
                rawDataPath,
            )
        }

        genio.AttachJSONData(rawDataPath, &rawData)
        unitsData[unitType].Fuel = rawData.Fuel
        unitsData[unitType].Movement = rawData.Movement
        unitsData[unitType].Vision = rawData.Vision

        var weaponPrimary awodatagen.WeaponType
        var weaponSecondary awodatagen.WeaponType
        var movementType awodatagen.MovementType
        var ok bool

        weaponPrimary, ok = awodatagen.WeaponTypeReverseStrings[rawData.WeaponPrimary]

        if !ok && rawData.WeaponPrimary != "" {
            awodatagen.LogFatalF(
                  "Missing or invalid primary weapon type '%s' on unit '%s'",
                rawData.WeaponPrimary,
                unitType.String(),
            )
        }

        weaponSecondary, ok = awodatagen.WeaponTypeReverseStrings[rawData.WeaponSecondary]
        if !ok && rawData.WeaponSecondary != "" {
            awodatagen.LogFatalF(
                  "Missing or invalid secondary weapon type '%s' on unit '%s'",
                rawData.WeaponSecondary,
                unitType.String(),
            )
        }

        if movementType, ok = awodatagen.MovementTypeReverseStrings[rawData.MovementType]; !ok {
            awodatagen.LogFatalF(
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
func getWeaponTypesData() *[awodatagen.WeaponTypeCount]awodatagen.WeaponTypeData {

	// Get raw weapon types data map
    var rawData map[string]awodatagen.WeaponTypeData

    dataPath := awodatagen.GetInputPath(
        awodatagen.UnitsDir,
        awodatagen.WeaponTypesFileName,
    )

    if _, err := os.Stat(dataPath); os.IsNotExist(err) {
        awodatagen.LogFatalF("Weapon types file path '%s' invalid", dataPath)
    }

    genio.AttachJSONData(dataPath, &rawData)

    // Transform raw map into processed array
    var data [awodatagen.WeaponTypeCount]awodatagen.WeaponTypeData

    for wTypeStr, wData := range rawData {
        var wType awodatagen.WeaponType
        var ok bool

        if wType, ok = awodatagen.WeaponTypeReverseStrings[wTypeStr]; !ok {
            awodatagen.LogFatalF(
                "Unknown weapon type '%s' found in data",
                wTypeStr,
            )
        }

        data[wType] = wData
    }

    return &data
}
