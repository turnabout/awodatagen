package cogen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/internal/genio"
	"github.com/turnabout/awodatagen/internal/packer"
	"github.com/turnabout/awodatagen/internal/utilities"
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/unitgen"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
)

func GetCOFrameImgs(frameImgs *[]packer.FrameImage) {

	// Loop army types
	for COArmy := unitgen.ArmyTypeFirst; COArmy < unitgen.ArmyTypeCount; COArmy++ {

		// Get directory for COs of this army type & loop contents
		COArmyDirPath := utilities.GetInputPath(awodatagen.CODir, COArmy.String())
		subDirs, err := ioutil.ReadDir(COArmyDirPath)
		utilities.LogFatalIfErr(err)

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
	COArmyType unitgen.ArmyType,
) {

	var CO CO
	var ok bool

	if CO, ok = COReverseStrings[CODirName]; !ok {
		utilities.LogFatalF(
			"Found CO at '%s', doesn't match any CO set in CO enumeration\n",
			COTypePath,
		)
	}

	imgs, err := ioutil.ReadDir(COTypePath)
	utilities.LogFatalIfErr(err)

	// Add all images of this CO to the frame images
	for _, img := range imgs {
		var frameType COFrameType
		var ok bool

		// Ensure the looped image is a possible CO frame type image
		cleanFileName := strings.TrimSuffix(img.Name(), filepath.Ext(img.Name()))

		if frameType, ok = COFrameTypeReverseStrings[cleanFileName]; !ok {
			utilities.LogFatalF(
				"Found CO image at '%s' doesn't match any valid CO frame type\n",
				path.Join(COTypePath, img.Name()),
			)
		}

		imageObj := genio.GetImage(path.Join(COTypePath, img.Name()))

		*frameImgs = append(*frameImgs, packer.FrameImage{
			Image:  imageObj,
			Width:  imageObj.Bounds().Max.X,
			Height: imageObj.Bounds().Max.Y,
			MetaData: packer.FrameImageMetaData{
				Type:               uint8(CO),
				Variation:          uint8(COArmyType),
				Index:              int(frameType),
				FrameImageDataType: uint8(framedata.CODataType),
			},
		})
	}
}
