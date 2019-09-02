// Generates units' sprite sheet & visual data
package main

import (
    "image"
    "io/ioutil"
    "log"
    "os"
)

// Generate units' sprite sheet & visuals data
func generateUnitsData() (*image.RGBA, UnitsData) {

    // Destination visual data (final, in-game sprite sheet generated for each army)
    // Unit Type -> Animation -> Animation Frame
    destVData := make([][][]Frame, UnitTypeAmount)

    // Generate origin data (sprite sheet & visual data)
    frameImgs := gatherUnitsFrameImages()
    packedFrameImgs, originWidth, originHeight := pack(frameImgs)
    spriteSheet := drawPackedFrames(packedFrameImgs, originWidth, originHeight)

    // TODO: Generate destination data

    return spriteSheet, UnitsData{
        Origin: *generateOriginVData(packedFrameImgs),
        Dest: destVData,
        X: 0,
        Y: 0,
        Width: originWidth,
        Height: originHeight,
        FullWidth: 0,  // TODO
        FullHeight: 0, // TODO
    }
}

// Gathers Frame Images from every single Unit image
func gatherUnitsFrameImages() *[]FrameImage {
    var frameImgs []FrameImage

    unitsDir := baseDirPath + imageInputsDirName + unitsDirName + "/"

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
                gatherAnimFrameImages(unitType, unitVar, anim, varDir + anim.String() + "/", &frameImgs)
            }
        }
    }

    return &frameImgs
}

// Gather all Frame Images from the given Unit Animation
func gatherAnimFrameImages(uType UnitType, uVar UnitVariation, uAnim UnitAnimation, animDir string, frameImgs *[]FrameImage) {
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
func generateOriginVData(packedFrameImgs *[]FrameImage) *[][][][]Frame {
    // Unit Type -> Variation -> Animation -> Animation Frames
    originVData := make([][][][]Frame, UnitTypeAmount)

    // Generate visual data
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
