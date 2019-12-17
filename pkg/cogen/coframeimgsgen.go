package cogen

import (
    "fmt"
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/packer"
    "io/ioutil"
    "path"
)

func GetCOFrameImgs(frameImgs *[]packer.FrameImage) {

    // Loop army types
    for COArmy := awodatagen.ArmyTypeFirst; COArmy < awodatagen.ArmyTypeCount; COArmy++ {

        // Get directory for COs of this army type & loop contents
        COArmyDirPath := awodatagen.GetInputPath(awodatagen.CODir, COArmy.String())
        subDirs, err := ioutil.ReadDir(COArmyDirPath)
        awodatagen.LogFatalIfErr(err)

        // Loop every CO directory in this army type directory
        for _, subDir := range subDirs {
            getCOTypeFrameImgs(
                frameImgs,
                path.Join(COArmyDirPath, subDir.Name()),
                subDir.Name(),
                COArmy,
            )
        }
    }
}

func getCOTypeFrameImgs(
    frameImgs *[]packer.FrameImage,
    COTypePath string,
    CODirName string,
    COArmyType awodatagen.ArmyType,
) {

    imgs, err := ioutil.ReadDir(COTypePath)
    awodatagen.LogFatalIfErr(err)

    fmt.Printf("%s\n", CODirName)

    for _, img := range imgs {
        fmt.Printf("%s\n", img.Name())
    }

    fmt.Printf("\n---------------------------\n")

    // Add the body image


    // Add the face images

    /*
    *frameImgs = append(*frameImgs, packer.FrameImage{
        Image: imageObj,
        Width: imageObj.Bounds().Max.X,
        Height: imageObj.Bounds().Max.Y,
        MetaData: packer.FrameImageMetaData{
            Type: uint8(uiElement),
            Index: frameIndex,
            FrameImageDataType: uint8(awodatagen.UIDataType),
        },
    })
    */


}
