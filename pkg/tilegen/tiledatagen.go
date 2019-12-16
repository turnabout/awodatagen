package tilegen

import (
    "github.com/turnabout/awodatagen"
    "github.com/turnabout/awodatagen/pkg/packer"
    "github.com/turnabout/awodatagen/pkg/tilegen/autovargen"
)

// Generate Src visual data JSON & sprite sheet
func GetTileData(packedFrameImgs *[]packer.FrameImage) *awodatagen.TileData {

    // Get the base tiles data object containing frame source data
    var tilesData *awodatagen.TileData = getBaseTileData(packedFrameImgs)

    // Attach additional data to the tiles data: auto-var data, clock data...
    autovargen.AttachTilesAutoVarData(tilesData)
    attachTilesClockData(tilesData)

    return tilesData
}

// Generate visual data for Src
func getBaseTileData(packedFrameImgs *[]packer.FrameImage) *awodatagen.TileData {

    // Tile Type -> Tile Variation -> Tile Variation Frames
    tilesVData := make(awodatagen.TileData, awodatagen.NeutralTileTypeCount)

    // Initialize Variations on every TileTypeData
    for tileType := range tilesVData {
        tilesVData[tileType].Variations = make(map[string][]awodatagen.Frame)
    }

    // Process Frame Images
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-tile frame images
        if frameImg.MetaData.FrameImageDataType != uint8(awodatagen.TileDataType) {
            continue
        }


        tileType := awodatagen.TileType(frameImg.MetaData.Type)
        tileVar := awodatagen.TileVariation(frameImg.MetaData.Variation)
        tileFrame := frameImg.MetaData.Index

        // Get the accumulated Animation slice of this Tile Type's given Variation, or initialize it
        var animSlice []awodatagen.Frame
        var ok bool

        if animSlice, ok = tilesVData[tileType].Variations[tileVar.String()]; !ok {
            animSlice = make([]awodatagen.Frame, 0)
        }

        // Add any Frames missing up until the one we're adding
        if missingFrames := (tileFrame + 1) - len(animSlice); missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                animSlice = append(animSlice, awodatagen.Frame{})
            }
        }

        // Add the Frame data to the animation slice, and record it to the visual data
        frame := awodatagen.Frame{
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

        animSlice[tileFrame] = frame
        tilesVData[tileType].Variations[tileVar.String()] = animSlice
        tilesVData[tileType].ClockData = nil
    }

    return &tilesVData
}
