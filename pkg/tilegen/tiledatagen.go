package tilegen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/internal/packer"
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/tilegen/autovargen"
	"github.com/turnabout/awodatagen/pkg/tilegen/placementrulesgen"
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
)

// Generate Src visual data JSON & sprite sheet
func GetTileData(packedFrameImgs *[]packer.FrameImage) *tiledata.TileData {

	// Get the base tiles data object containing frame source data
	var tilesData *tiledata.TileData = getBaseTileData(packedFrameImgs)

	// Attach additional data to the tiles data: auto-var data, clock data...
	autovargen.AttachTilesAutoVarData(tilesData)
	placementrulesgen.AttachTilesPlacementRulesData(tilesData)
	attachTilesClockData(tilesData)

	return tilesData
}

// Generate visual data for Src
func getBaseTileData(packedFrameImgs *[]packer.FrameImage) *tiledata.TileData {

	// Tile Type -> Tile Variation -> Tile Variation Frames
	tilesVData := make(tiledata.TileData, tiledata.NeutralTileTypeCount)

	// Initialize Variations on every TileTypeData
	for tileType := range tilesVData {
		tilesVData[tileType].Variations = make(map[string]tiledata.TileVarData)
	}

	// Process Frame Images
	for _, frameImg := range *packedFrameImgs {

		// Ignore non-tile frame images
		if frameImg.MetaData.FrameImageDataType != uint8(framedata.TileDataType) {
			continue
		}

		tileType := tiledata.TileType(frameImg.MetaData.Type)
		tileVar := tiledata.TileVariation(frameImg.MetaData.Variation)
		tileFrame := frameImg.MetaData.Index

		// Get the accumulated Animation slice of this Tile Type's given Variation, or initialize it
		var tileVarData tiledata.TileVarData
		var ok bool

		if tileVarData, ok = tilesVData[tileType].Variations[tileVar.String()]; !ok {
			tileVarData.Frames = make([]framedata.Frame, 0)
		}

		// Add any Frames missing up until the one we're adding
		if missingFrames := (tileFrame + 1) - len(tileVarData.Frames); missingFrames > 0 {
			for i := 0; i < missingFrames; i++ {
				tileVarData.Frames = append(tileVarData.Frames, framedata.Frame{})
			}
		}

		// Add the Frame data to the animation slice, and record it to the visual data
		frame := framedata.Frame{
			X: frameImg.X,
			Y: frameImg.Y,
		}

		// To save space in the JSON file, omit adding SrcWidth/SrcHeight if they're the regular Tile sizes
		if frameImg.Width != awodatagen.RegularTileDimension {
			frame.Width = frameImg.Width
		}

		if frameImg.Height != awodatagen.RegularTileDimension {
			frame.Height = frameImg.Height
		}

		tileVarData.Frames[tileFrame] = frame
		tileVarData.ClockIndex = nil

		tilesVData[tileType].Variations[tileVar.String()] = tileVarData
	}

	return &tilesVData
}
