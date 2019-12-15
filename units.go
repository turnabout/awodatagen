// Generates Units' visual data
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path"
)

// Generate units' sprite sheet & visuals data.
// srcX/srcY specifies the coordinates of the units' sprite sheet within the final raw sprite sheet
func getUnitsData(packedFrameImgs *[]FrameImage)  *UnitsData {
    var unitsData *UnitsData = getBaseUnitsData(packedFrameImgs)

    attachExtraUnitsVData(unitsData)

    return unitsData
}

// Gets Frame Images from every single Unit image
func getUnitsSrcFrameImgs(frameImgs *[]FrameImage) {

    // Loop units
    for unitType := FirstUnitType; unitType <= LastUnitType; unitType++ {

        // Loop Variations
        for unitVar := FirstUnitVariation; unitVar <= LastUnitVariation; unitVar++ {

            varDir := getFullProjectPath( unitsDir, unitType.String(), unitVar.String() )

            // Ignore this variation if it does not exist on this unit
            if _, err := os.Stat(varDir); os.IsNotExist(err) {
                break
            }

            // Loop Animations
            for anim := FirstUnitAnimation; anim <= LastUnitAnimation; anim++ {
                getAnimFrameImgs(
                    unitType,
                    unitVar,
                    anim,
                    path.Join( varDir, anim.String() ),
                    frameImgs,
                )
            }
        }
    }
}

// Get all Frame Images from the given Unit Animation
func getAnimFrameImgs(uType UnitType, uVar UnitVariation, uAnim UnitAnimation, animDir string, frameImgs *[]FrameImage) {
    imgFiles, err := ioutil.ReadDir(animDir)

    if err != nil {
        log.Fatal(err)
    }

    // Loop every image of this Animation
    for index, imgFile := range imgFiles {
        imageObj := getImage(path.Join(animDir, imgFile.Name()))

        *frameImgs = append(*frameImgs, FrameImage{
            Image: imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: FrameImageMetaData{
                Type:           uint8(uType),
                Variation:      uint8(uVar),
                Animation:      uint8(uAnim),
                Index:          index,
                FrameImageType: UnitFrameImage,
            },
        })
    }
}

// Generate the origin visual data (units' visual data on the raw sprite sheet) using packed Frame Images
func getBaseUnitsData(packedFrameImgs *[]FrameImage) *UnitsData {

    // Unit Type -> Variation -> Animation -> Animation Frames
    unitsData := make(UnitsData, UnitTypeAmount)

    for _, frameImg := range *packedFrameImgs {

        // Ignore non-unit frame images
        if frameImg.MetaData.FrameImageType != UnitFrameImage {
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
                unitsData[unitType] = append(unitsData[unitType], [][]Frame{})
            }
        }

        // Check if Animation is missing, add up to it if necessary
        missingAnims := int(unitAnim + 1) - len(unitsData[unitType][unitVar])

        if missingAnims > 0 {
            for i := 0; i < missingAnims; i++ {
                unitsData[unitType][unitVar] = append(unitsData[unitType][unitVar], []Frame{})
            }
        }

        // Check if Animation Frame is missing, add up to it if necessary
        missingFrames := int(unitFrame + 1) - len(unitsData[unitType][unitVar][unitAnim])

        if missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                unitsData[unitType][unitVar][unitAnim] = append(unitsData[unitType][unitVar][unitAnim], Frame{})
            }
        }

        // Store data
        if frameImg.X == 48 && frameImg.Y == 64 {
            fmt.Printf("%#v\n", frameImg)
        }

        unitsData[unitType][unitVar][unitAnim][unitFrame] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }
    }

    return &unitsData
}

// Attach extra data stored away in JSON files
func attachExtraUnitsVData(vData *UnitsData) {
    // unitsDir := baseDirPath + inputsDirName + unitsDir
    // attachJSONData(unitsDir + palettesFileName, &vData.Palettes)
    // attachJSONData(unitsDir + basePaletteFileName, &vData.BasePalette)
}
