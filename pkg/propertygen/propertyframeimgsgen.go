package propertygen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/genio"
    "github.com/turnabout/awodatagen/pkg/packer"
    "os"
)

// Gather Frame Images for Properties' source
func GetPropertyFrameImgs(frameImgs *[]packer.FrameImage) {

    // Loop Weather Variations
    for weatherVar := awodatagen.PropWeatherVarFirst; weatherVar <= awodatagen.PropWeatherVarLast; weatherVar++ {

        // Loop Property Types
        for propType := awodatagen.PropTypeFirst; propType <= awodatagen.PropTypeLast; propType++ {
            // propDir := getFullProjectPath(propertiesDir) + weatherVar.String() + "/" + propType.String() + "/"

            // Loop army variations
            for unitVar := awodatagen.ArmyTypeFirst; unitVar <= awodatagen.ArmyTypeLast; unitVar++ {

                fullPath := awodatagen.GetInputPath(
                    awodatagen.PropertiesDir,
                    weatherVar.String(),
                    propType.String(),
                    unitVar.String(),
                ) + ".png"

                // Ignore this army variation if it does not exist on properties of this type
                if _, err := os.Stat(fullPath); os.IsNotExist(err) {
                    continue
                }

                imageObj := genio.GetImage(fullPath)

                *frameImgs = append(*frameImgs, packer.FrameImage{
                    Image: imageObj,
                    Width:  imageObj.Bounds().Max.X,
                    Height: imageObj.Bounds().Max.Y,
                    MetaData: packer.FrameImageMetaData{
                        Type:               uint8(propType),
                        Variation:          uint8(weatherVar),
                        Animation:          uint8(unitVar),
                        FrameImageDataType: uint8(awodatagen.PropertyDataType),
                    },
                })
            }
        }
    }
}
