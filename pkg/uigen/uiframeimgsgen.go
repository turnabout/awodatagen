package uigen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/genio"
	"github.com/turnabout/awodatagen/pkg/packer"
	"github.com/turnabout/awodatagen/pkg/utilities"
	"io/ioutil"
	"log"
	"path"
	"strconv"
	"strings"
)

// Gathers data on every single UI image
func GetUIFrameImgs(frameImgs *[]packer.FrameImage) {

	// Gather frame images from the elements found in the UI directory
	UIDirElements, err := ioutil.ReadDir(utilities.GetInputPath(awodatagen.UIDir))
	utilities.LogFatalIfErr(err)

	for _, UIDirElement := range UIDirElements {
		if UIDirElement.IsDir() {
			gatherUISubDirFrameImgs(
				frameImgs,
				UIDirElement.Name(),
				utilities.GetInputPath(awodatagen.UIDir, UIDirElement.Name()),
			)
		} else {
			appendUIFrameImgs(
				utilities.GetInputPath(awodatagen.UIDir),
				UIDirElement.Name(),
				0,
				UIElementNone,
				frameImgs,
			)
		}
	}
}

func gatherUISubDirFrameImgs(frameImgs *[]packer.FrameImage, dirName string, dirPath string) {

	// Get the UI Element corresponding to this directory
	uiElement := getUiElementByString(dirName)

	// Loop all frames for this UI element
	uiSubDirFiles, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range uiSubDirFiles {
		appendUIFrameImgs(dirPath, file.Name(), -1, uiElement, frameImgs)
	}
}

func appendUIFrameImgs(dirPath string, fileName string, frameIndex int, uiElement UIElement, frameImgs *[]packer.FrameImage) {

	// Create the frame image for this UI element
	imageObj := genio.GetImage(path.Join(dirPath, fileName))

	// If frame index not given, the frame index should be the file's name itself
	if frameIndex == -1 {
		var err error
		frameIndex, err = strconv.Atoi(strings.TrimSuffix(fileName, path.Ext(fileName)))
		if err != nil {
			log.Fatal(err)
		}
	}

	// If ui element not given, the ui element should be the file's name itself
	if int(uiElement) == UIElementNone {
		uiElement = getUiElementByString(strings.TrimSuffix(fileName, path.Ext(fileName)))
	}

	*frameImgs = append(*frameImgs, packer.FrameImage{
		Image:  imageObj,
		Width:  imageObj.Bounds().Max.X,
		Height: imageObj.Bounds().Max.Y,
		MetaData: packer.FrameImageMetaData{
			Type:               uint8(uiElement),
			Index:              frameIndex,
			FrameImageDataType: uint8(framedata.UIDataType),
		},
	})
}
