package cogen

import (
    "fmt"
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
    "github.com/turnabout/awodatagen/pkg/packer"
    "io/ioutil"
    "path"
    "path/filepath"
    "strings"
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

    var CO awodatagen.CO
    var ok bool

    if CO, ok = awodatagen.COReverseStrings[CODirName]; !ok {
        awodatagen.LogFatal(
            []string{
                fmt.Sprintf(
                    "Found CO at '%s', doesn't match any CO set in CO enumeration\n",
                    COTypePath,
                ),
            },
        )
    }

    imgs, err := ioutil.ReadDir(COTypePath)
    awodatagen.LogFatalIfErr(err)

    // Add all images of this CO to the frame images
    for _, img := range imgs {
        var frameType awodatagen.COFrameType
        var ok bool

        // Ensure the looped image is a possible CO frame type image
        cleanFileName := strings.TrimSuffix(img.Name(), filepath.Ext(img.Name()))

        if frameType, ok = awodatagen.COFrameTypeReverseStrings[cleanFileName]; !ok {
            awodatagen.LogFatal(
                []string{
                    fmt.Sprintf(
                        "Found CO image at '%s', doesn't match any valid CO frame type\n",
                        path.Join(COTypePath, img.Name()),
                    ),
                },
            )
        }

        imageObj := genio.GetImage(path.Join(COTypePath, img.Name()))

        *frameImgs = append(*frameImgs, packer.FrameImage{
            Image: imageObj,
            Width: imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: packer.FrameImageMetaData{
                Type: uint8(CO),
                Variation: uint8(COArmyType),
                Index: int(frameType),
                FrameImageDataType: uint8(awodatagen.CODataType),
            },
        })
    }
}
