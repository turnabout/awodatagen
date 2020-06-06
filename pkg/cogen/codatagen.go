package cogen

import (
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/packer"
	"github.com/turnabout/awodatagen/pkg/unitgen"
)

// Generates the CO-related game data
func GetCOData(packedFrameImgs *[]packer.FrameImage) *COData {
	var COData *COData = getCOBaseData(packedFrameImgs)

	return COData
}

func getCOBaseData(packedFrameImgs *[]packer.FrameImage) *COData {

	// CO -> CO Type Data
	data := make(COData, COCount)

	// Process frame images
	for _, frameImg := range *packedFrameImgs {

		// Ignore non-CO frame images
		if frameImg.MetaData.FrameImageDataType != uint8(framedata.CODataType) {
			continue
		}

		// Get metadata on the CO this frame image represents
		CO := CO(frameImg.MetaData.Type)
		army := unitgen.ArmyType(frameImg.MetaData.Variation)
		frameType := COFrameType(frameImg.MetaData.Index)

		// Set data on this CO
		data[CO].Name = CO.String()
		data[CO].Army = army

		data[CO].Frames[frameType] = framedata.Frame{
			X:      frameImg.X,
			Y:      frameImg.Y,
			Width:  frameImg.Width,
			Height: frameImg.Height,
		}
	}

	return &data
}
