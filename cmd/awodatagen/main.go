package main

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/pkg/cogen"
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/gamedata"
	"github.com/turnabout/awodatagen/pkg/genio"
	"github.com/turnabout/awodatagen/pkg/packer"
	"github.com/turnabout/awodatagen/pkg/palettegen"
	"github.com/turnabout/awodatagen/pkg/propertygen"
	"github.com/turnabout/awodatagen/pkg/tilegen"
	"github.com/turnabout/awodatagen/pkg/uigen"
	"github.com/turnabout/awodatagen/pkg/unitgen"
	"github.com/turnabout/awodatagen/pkg/utilities"
	"image"
	"log"
)

// Callback used to get frame images from a section of the game data (tiles, units, etc)
type getFrameImagesCB func(*[]packer.FrameImage)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Gather all packed frame images & the sprite sheet image
	packedFrameImages, ssImg := gatherFrameImages()

	// Create game data object using the frame images
	var gameData = gamedata.GameData{
		Tiles:      *tilegen.GetTileData(packedFrameImages),
		Properties: *propertygen.GetPropertyData(packedFrameImages),
		Units:      *unitgen.GetUnitData(packedFrameImages),
		UI:         *uigen.GetUIData(packedFrameImages),
		COs:        *cogen.GetCOData(packedFrameImages),

		SpriteSheetDimensions: gamedata.SSDimensions{
			Width:  ssImg.Bounds().Max.X,
			Height: ssImg.Bounds().Max.Y,
		},
	}

	attachAdditionalData(&gameData)

	// Output results
	genio.OutputJSON(&gameData)
	genio.OutputSpriteSheet(ssImg)
}

// Gathers frame images from every category of entities making up the sprite sheet
func gatherFrameImages() (*[]packer.FrameImage, *image.RGBA) {

	var accumImg *image.RGBA = nil

	// 1. Gather frame images that need to be aligned with the top-left (tiles, properties, idle units)
	var alignedFrameImages *[]packer.FrameImage

	alignedFrameImages, accumImg = gatherStepFrameImages(
		accumImg,
		tilegen.GetTileFrameImgs,
		propertygen.GetPropertyFrameImgs,
		unitgen.GetUnitIdleFrameImgs,
	)

	// 2. Gather all other frame images together
	var otherFrameImages *[]packer.FrameImage

	otherFrameImages, accumImg = gatherStepFrameImages(
		accumImg,
		cogen.GetCOFrameImgs,
		uigen.GetUIFrameImgs,
		unitgen.GetUnitNonIdleFrameImgs,
	)

	// 3. Combine all frame images in one array
	var frameImagesOut []packer.FrameImage = append(*alignedFrameImages, *otherFrameImages...)

	return &frameImagesOut, accumImg
}

// Gathers the frame images for a single step.
// Can use one or many frame image callbacks to process frame images, pack them, use them to draw an image and return
// both the packed frame images and the drawn image.
func gatherStepFrameImages(
	accumImg *image.RGBA,
	frameImagesCBs ...getFrameImagesCB,
) (*[]packer.FrameImage, *image.RGBA) {

	var frameImages []packer.FrameImage

	// If an accumulated image is given, use it as a base for this step's frame images
	if accumImg != nil {
		frameImages = append(frameImages, packer.FrameImage{
			Image:  accumImg,
			Width:  accumImg.Bounds().Max.X,
			Height: accumImg.Bounds().Max.Y,
			MetaData: packer.FrameImageMetaData{
				FrameImageDataType: uint8(framedata.OtherDataType),
			},
		})
	}

	// Get the frame images for this step using the given callbacks
	for _, CB := range frameImagesCBs {
		CB(&frameImages)
	}

	// Pack the frame images
	packedFrameImages, sectionWidth, sectionHeight := packer.Pack(&frameImages)

	// Output packed frame images and new accumulated image
	return packedFrameImages, packer.DrawPackedFrames(packedFrameImages, sectionWidth, sectionHeight)
}

// Gather additional visual data and attach to the main visual data object
func attachAdditionalData(gameData *gamedata.GameData) {

	// Adds default stages data
	genio.AttachJSONData(
		utilities.GetInputPath(awodatagen.OtherDir, awodatagen.StagesFileName),
		&gameData.Stages,
	)

	// Adds animation clocks data
	genio.AttachJSONData(
		utilities.GetInputPath(awodatagen.OtherDir, awodatagen.ClocksFileName),
		&gameData.Clocks,
	)

	// Palette data
	gameData.Palettes = *palettegen.GetPaletteData()
}
