package unitgen

import (
    "fmt"
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/packer"
)

type rawUnitData struct {
    MovementType    string `json:"movementType"`
    Movement        uint8  `json:"movement"`
    Vision          uint8  `json:"vision"`
    Fuel            uint8  `json:"fuel"`
    Ammo            uint8  `json:"ammo"`
    WeaponPrimary   string `json:"weaponPrimary"`
    WeaponSecondary string `json:"weaponSecondary"`
}

// Generates units game data.
func GetUnitData(packedFrameImgs *[]packer.FrameImage)  *awodatagen.UnitData {

    var unitsData *awodatagen.UnitData = getBaseUnitData(packedFrameImgs)


    attachExtraUnitsVData(unitsData)

    return unitsData
}

// Generates the origin visual data (units' visual data on the raw sprite sheet) using packed Frame Images
func getBaseUnitData(packedFrameImgs *[]packer.FrameImage) *awodatagen.UnitData {

    var unitsData awodatagen.UnitData

    // Load data from the unit's source, raw JSON data
    // TODO

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

    return &unitsData
}

// Attach extra data stored away in JSON files
func attachExtraUnitsVData(vData *awodatagen.UnitData) {
    // unitsDir := baseDirPath + inputsDirName + unitsDir
    // attachJSONData(unitsDir + palettesFileName, &vData.Palettes)
    // attachJSONData(unitsDir + basePaletteFileName, &vData.BasePalette)
}
