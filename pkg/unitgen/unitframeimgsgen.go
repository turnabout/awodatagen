package unitgen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/genio"
	"github.com/turnabout/awodatagen/pkg/packer"
	"github.com/turnabout/awodatagen/pkg/utilities"
	"io/ioutil"
	"os"
	"path"
)

// Gets frame images from all non-idle unit images
func GetUnitNonIdleFrameImgs(frameImgs *[]packer.FrameImage) {
	getUnitFrameImgs(frameImgs, []UnitAnimation{Right, Up, Down})
}

// Gets frame images from all idle unit images
func GetUnitIdleFrameImgs(frameImgs *[]packer.FrameImage) {
	getUnitFrameImgs(frameImgs, []UnitAnimation{Idle})
}

// Gets frame images for all unit images of the given unit animations
func getUnitFrameImgs(frameImgs *[]packer.FrameImage, animations []UnitAnimation) {
	// Loop units
	for unitType := UnitTypeFirst; unitType <= UnitTypeLast; unitType++ {

		// Loop variations
		for unitVar := ArmyTypeFirst; unitVar <= ArmyTypeLast; unitVar++ {

			varDir := utilities.GetInputPath(
				awodatagen.UnitsDir,
				unitType.String(),
				awodatagen.FramesDir,
				unitVar.String(),
			)

			// Ignore this variation if it does not exist on this unit
			if _, err := os.Stat(varDir); os.IsNotExist(err) {
				break
			}

			// Loop animations
			for _, animation := range animations {
				getAnimFrameImgs(
					unitType,
					unitVar,
					animation,
					path.Join(varDir, animation.String()),
					frameImgs,
				)
			}
		}
	}
}

// Gets frame images from a unit animation
func getAnimFrameImgs(
	uType UnitType,
	uVar ArmyType,
	uAnim UnitAnimation,
	animDir string,
	frameImgs *[]packer.FrameImage,
) {
	imgFiles, err := ioutil.ReadDir(animDir)
	utilities.LogFatalIfErr(err)

	// Loop every image of this Animation
	for index, imgFile := range imgFiles {
		imageObj := genio.GetImage(path.Join(animDir, imgFile.Name()))

		*frameImgs = append(*frameImgs, packer.FrameImage{
			Image:  imageObj,
			Width:  imageObj.Bounds().Max.X,
			Height: imageObj.Bounds().Max.Y,
			MetaData: packer.FrameImageMetaData{
				Type:               uint8(uType),
				Variation:          uint8(uVar),
				Animation:          uint8(uAnim),
				Index:              index,
				FrameImageDataType: uint8(framedata.UnitDataType),
			},
		})
	}
}
