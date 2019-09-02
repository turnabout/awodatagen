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

    // Origin visual data (raw sprite sheet)
    // Unit Type -> Variation -> Animation -> Animation Frames
    unitsOriginVisualData := make([][][][]Frame, UnitTypeAmount)

    // Destination visual data (final, in-game sprite sheet generated for each army)
    // Unit Type -> Animation -> Animation Frame
    unitsDestVisualData := make([][][]Frame, UnitTypeAmount)

    // Generate origin data (sprite sheet & visual data)
    frameImgs := gatherUnitsFrameImages()
    packedFrameImgs, originWidth, originHeight := pack(frameImgs)


    // TODO: Generate visual data

    spriteSheet := drawPackedFrames(packedFrameImgs, originWidth, originHeight)

    // TODO: Generate destination data

    return spriteSheet, UnitsData{
        Origin: unitsOriginVisualData,
        Dest: unitsDestVisualData,
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
