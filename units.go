// Generates Units' visual data
package main

import (
    "io/ioutil"
    "log"
    "os"
)

// Generate units' sprite sheet & visuals data.
// srcX/srcY specifies the coordinates of the units' sprite sheet within the final raw sprite sheet
func getUnitsData()  *UnitsData {

    // Get source frame images
    srcFrameImgs := getUnitsSrcFrameImgs()
    packedSrcFrameImgs, srcWidth, srcHeight := pack(srcFrameImgs)

    vData := UnitsData{
        Src:       *getUnitsSrcVData(packedSrcFrameImgs),
        frameImg: FrameImage{
            Image:    drawPackedFrames(packedSrcFrameImgs, srcWidth, srcHeight),
            Width:    srcWidth,
            Height:   srcHeight,
            MetaData: FrameImageMetaData{Type: uint8(VisualDataUnits)},
        },
    }

    attachExtraUnitsVData(&vData)
    return &vData
}

// Gets Frame Images from every single Unit image
func getUnitsSrcFrameImgs() *[]FrameImage {
    var frameImgs []FrameImage

    unitsDir := baseDirPath + inputsDirName + unitsDirName + "/"

    // Loop Units
    for unitType := FirstUnitType; unitType <= LastUnitType; unitType++ {
        unitDir := unitsDir + unitType.String() + "/"

        // Loop Variations
        for unitVar := FirstUnitVariation; unitVar <= LastUnitVariation; unitVar++ {
            varDir := unitDir + unitVar.String() + "/"

            // Ignore this variation if it does not exist on this unit
            if _, err := os.Stat(varDir); os.IsNotExist(err) {
                break
            }

            // Loop Animations
            for anim := FirstUnitAnimation; anim <= LastUnitAnimation; anim++ {
                getAnimFrameImgs(unitType, unitVar, anim, varDir + anim.String() + "/", &frameImgs)
            }
        }
    }

    return &frameImgs
}

// Get all Frame Images from the given Unit Animation
func getAnimFrameImgs(uType UnitType, uVar UnitVariation, uAnim UnitAnimation, animDir string, frameImgs *[]FrameImage) {
    imgFiles, err := ioutil.ReadDir(animDir)

    if err != nil {
        log.Fatal(err)
    }

    // Loop every image of this Animation
    for index, imgFile := range imgFiles {
        imageObj := getImage(animDir + imgFile.Name())

        *frameImgs = append(*frameImgs, FrameImage{
            Image: imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: FrameImageMetaData{
                Type: uint8(uType),
                Variation: uint8(uVar),
                Animation: uint8(uAnim),
                Index: index,
            },
        })
    }
}

// Generate the origin visual data (units' visual data on the raw sprite sheet) using packed Frame Images
func getUnitsSrcVData(packedFrameImgs *[]FrameImage) *[][][][]Frame {

    // Unit Type -> Variation -> Animation -> Animation Frames
    originVData := make([][][][]Frame, UnitTypeAmount)

    for _, frameImg := range *packedFrameImgs {
        unitType := frameImg.MetaData.Type
        unitVar := frameImg.MetaData.Variation
        unitAnim := frameImg.MetaData.Animation
        unitFrame := frameImg.MetaData.Index

        // Check if Variation is missing, add up to it if necessary
        missingVars := int(unitVar + 1) - len(originVData[unitType])

        if missingVars > 0 {
            for i := 0; i < missingVars; i++ {
                originVData[unitType] = append(originVData[unitType], [][]Frame{})
            }
        }

        // Check if Animation is missing, add up to it if necessary
        missingAnims := int(unitAnim + 1) - len(originVData[unitType][unitVar])

        if missingAnims > 0 {
            for i := 0; i < missingAnims; i++ {
                originVData[unitType][unitVar] = append(originVData[unitType][unitVar], []Frame{})
            }
        }

        // Check if Animation Frame is missing, add up to it if necessary
        missingFrames := int(unitFrame + 1) - len(originVData[unitType][unitVar][unitAnim])

        if missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                originVData[unitType][unitVar][unitAnim] = append(originVData[unitType][unitVar][unitAnim], Frame{})
            }
        }

        // Store data
        originVData[unitType][unitVar][unitAnim][unitFrame] = Frame{
            X: frameImg.X,
            Y: frameImg.Y,
            Width: frameImg.Width,
            Height: frameImg.Height,
        }
    }

    return &originVData
}

// Attach extra visual data stored away in JSON files
func attachExtraUnitsVData(vData *UnitsData) {
    unitsDir := baseDirPath + inputsDirName + unitsDirName

    attachJSONData(unitsDir + palettesFileName, &vData.Palettes)
    attachJSONData(unitsDir + basePaletteFileName, &vData.BasePalette)
}
