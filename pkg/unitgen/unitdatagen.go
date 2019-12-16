package unitgen

import (
    "fmt"
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/packer"
)

// Generates units game data.
func GetUnitData(packedFrameImgs *[]packer.FrameImage)  *awossgen.UnitData {

    var unitsData *awossgen.UnitData = getBaseUnitData(packedFrameImgs)


    attachExtraUnitsVData(unitsData)

    return unitsData
}

// Generates the origin visual data (units' visual data on the raw sprite sheet) using packed Frame Images
func getBaseUnitData(packedFrameImgs *[]packer.FrameImage) *awossgen.UnitData {

    // Unit Type -> Variation -> Animation -> Animation Frames
    unitsData := make(awossgen.UnitData, awossgen.UnitTypeCount)

    for _, frameImg := range *packedFrameImgs {

        // Ignore non-unit frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awossgen.UnitDataType) {
            continue
        }

        unitType := frameImg.MetaData.Type
        unitVar := frameImg.MetaData.Variation
        unitAnim := frameImg.MetaData.Animation
        unitFrame := frameImg.MetaData.Index

        // Check if Variation is missing, add up to it if necessary
        missingVars := int(unitVar + 1) - len(unitsData[unitType])

        if missingVars > 0 {
            for i := 0; i < missingVars; i++ {
                unitsData[unitType] = append(unitsData[unitType], [][]awossgen.Frame{})
            }
        }

        // Check if Animation is missing, add up to it if necessary
        missingAnims := int(unitAnim + 1) - len(unitsData[unitType][unitVar])

        if missingAnims > 0 {
            for i := 0; i < missingAnims; i++ {
                unitsData[unitType][unitVar] = append(unitsData[unitType][unitVar], []awossgen.Frame{})
            }
        }

        // Check if Animation Frame is missing, add up to it if necessary
        missingFrames := int(unitFrame + 1) - len(unitsData[unitType][unitVar][unitAnim])

        if missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                unitsData[unitType][unitVar][unitAnim] = append(unitsData[unitType][unitVar][unitAnim], awossgen.Frame{})
            }
        }

        // Store data
        if frameImg.X == 48 && frameImg.Y == 64 {
            fmt.Printf("%#v\n", frameImg)
        }

        unitsData[unitType][unitVar][unitAnim][unitFrame] = awossgen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }
    }

    return &unitsData
}

// Attach extra data stored away in JSON files
func attachExtraUnitsVData(vData *awossgen.UnitData) {
    // unitsDir := baseDirPath + inputsDirName + unitsDir
    // attachJSONData(unitsDir + palettesFileName, &vData.Palettes)
    // attachJSONData(unitsDir + basePaletteFileName, &vData.BasePalette)
}
