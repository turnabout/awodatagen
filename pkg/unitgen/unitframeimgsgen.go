package unitgen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
    "github.com/turnabout/awodatagen/pkg/packer"
    "io/ioutil"
    "log"
    "os"
    "path"
)

// Gets frame images from all unit images
func GetUnitFrameImgs(frameImgs *[]packer.FrameImage) {

    // Loop units
    for unitType := awodatagen.UnitTypeFirst; unitType <= awodatagen.UnitTypeLast; unitType++ {

        // Loop Variations
        for unitVar := awodatagen.ArmyTypeFirst; unitVar <= awodatagen.ArmyTypeLast; unitVar++ {

            varDir := awodatagen.GetInputPath( awodatagen.UnitsDir, unitType.String(), unitVar.String() )

            // Ignore this variation if it does not exist on this unit
            if _, err := os.Stat(varDir); os.IsNotExist(err) {
                break
            }

            // Loop Animations
            for anim := awodatagen.UnitAnimFirst; anim <= awodatagen.UnitAnimLast; anim++ {
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

// Gets frame images from a unit animation
func getAnimFrameImgs(
    uType awodatagen.UnitType,
    uVar awodatagen.ArmyType,
    uAnim awodatagen.UnitAnimation,
    animDir string,
    frameImgs *[]packer.FrameImage,
) {
    imgFiles, err := ioutil.ReadDir(animDir)

    if err != nil {
        log.Fatal(err)
    }

    // Loop every image of this Animation
    for index, imgFile := range imgFiles {
        imageObj := genio.GetImage(path.Join(animDir, imgFile.Name()))

        *frameImgs = append(*frameImgs, packer.FrameImage{
            Image: imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: packer.FrameImageMetaData{
                Type:               uint8(uType),
                Variation:          uint8(uVar),
                Animation:          uint8(uAnim),
                Index:              index,
                FrameImageDataType: uint8(awodatagen.UnitDataType),
            },
        })
    }
}
