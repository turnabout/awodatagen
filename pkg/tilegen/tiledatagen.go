package tilegen

import (
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/packer"
    "github.com/turnabout/awossgen/pkg/tilegen/autovargen"
)

// Generate Src visual data JSON & sprite sheet
func GetTileData(packedFrameImgs *[]packer.FrameImage) *awossgen.TileData {

    // Get the base tiles data object containing frame source data
    var tilesData *awossgen.TileData = getBaseTileData(packedFrameImgs)

    // Attach additional data to the tiles data: auto-var data, clock data...
    autovargen.AttachTilesAutoVarData(tilesData)
    attachTilesClockData(tilesData)

    return tilesData
}

// Generate visual data for Src
func getBaseTileData(packedFrameImgs *[]packer.FrameImage) *awossgen.TileData {

    // Tile Type -> Tile Variation -> Tile Variation Frames
    tilesVData := make(awossgen.TileData, awossgen.NeutralTileTypeCount)

    // Initialize Variations on every TileTypeData
    for tileType := range tilesVData {
        tilesVData[tileType].Variations = make(map[string][]awossgen.Frame)
    }

    // Process Frame Images
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-tile frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awossgen.TileDataType) {
            continue
        }


        tileType := awossgen.TileType(frameImg.MetaData.Type)
        tileVar := awossgen.TileVariation(frameImg.MetaData.Variation)
        tileFrame := frameImg.MetaData.Index

        // Get the accumulated Animation slice of this Tile Type's given Variation, or initialize it
        var animSlice []awossgen.Frame
        var ok bool

        if animSlice, ok = tilesVData[tileType].Variations[tileVar.String()]; !ok {
            animSlice = make([]awossgen.Frame, 0)
        }

        // Add any Frames missing up until the one we're adding
        if missingFrames := (tileFrame + 1) - len(animSlice); missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                animSlice = append(animSlice, awossgen.Frame{})
            }
        }

        // Add the Frame data to the animation slice, and record it to the visual data
        frame := awossgen.Frame{
            X: frameImg.X,
            Y: frameImg.Y,
        }

        // To save space in the JSON file, omit adding SrcWidth/SrcHeight if they're the regular Tile sizes
        if frameImg.Width != awossgen.RegularTileDimension {
            frame.Width = frameImg.Width
        }

        if frameImg.Height != awossgen.RegularTileDimension {
            frame.Height = frameImg.Height
        }

        animSlice[tileFrame] = frame
        tilesVData[tileType].Variations[tileVar.String()] = animSlice
        tilesVData[tileType].ClockData = nil
    }

    return &tilesVData
}
