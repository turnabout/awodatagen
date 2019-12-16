// Generates Tiles' visual data
package tilegen

import (
    "io/ioutil"
    "log"
    "os"
    "path"
    "strings"
)

// Map for looking up a Tile Type using its corresponding full string
var tileReverseStrings = map[string]TileType {

    // Basic neutral tiles, represented visually
    "Plain": Plain,
    "Forest": Forest,
    "Mountain": Mountain,
    "Road": Road,
    "Bridge": Bridge,
    "River": River,
    "Sea": Sea,
    "Reef": Reef,
    "Shore": Shore,
    "Pipe": Pipe,
    "PipeFragile": PipeFragile,
    "Silo": Silo,

    // Additional neutral tiles, represented visually (standard size)
    "BaseSmoke": BaseSmoke,
    "Empty": Empty,

    // Additional neutral tiles, represented visually (non-standard size)
    "LandPiece": LandPiece,
}

// Map for looking up a Tile Variation using its corresponding full string
var tileVarsReverseStrings = map[string]TileVariation {
    "Default": Default,
    "Horizontal": Horizontal,
    "Vertical": Vertical,
    "VerticalEnd": VerticalEnd,
    "Top": Top,
    "Bottom": Bottom,
    "DirLeft": DirLeft,
    "DirRight": DirRight,
    "TopLeft": TopLeft,
    "TopRight": TopRight,
    "BottomLeft": BottomLeft,
    "BottomRight": BottomRight,
    "Middle": Middle,
    "ShadowedDefault": ShadowedDefault,
    "ShadowedTopLeft": ShadowedTopLeft,
    "ShadowedBottomLeft": ShadowedBottomLeft,
    "ShadowedLeft": ShadowedLeft,
    "ShadowedHorizontal": ShadowedHorizontal,
    "ShadowedVertical": ShadowedVertical,
    "ShadowedVerticalEnd": ShadowedVerticalEnd,
    "ShadowedTLeft": ShadowedTLeft,
    "TTop": TTop,
    "TBottom": TBottom,
    "TLeft": TLeft,
    "TRight": TRight,
    "Small": Small,
    "WaterfallUp": WaterfallUp,
    "WaterfallDown": WaterfallDown,
    "WaterfallLeft": WaterfallLeft,
    "WaterfallRight": WaterfallRight,
    "Hole": Hole,
    "HoleHorizontal": HoleHorizontal,
    "HoleVertical": HoleVertical,
    "HoleLeft": HoleLeft,
    "HoleRight": HoleRight,
    "HoleTop": HoleTop,
    "HoleBottom": HoleBottom,
    "TopConnectedLeft": TopConnectedLeft,
    "TopConnectedRight": TopConnectedRight,
    "TopConnectedFull": TopConnectedFull,
    "BottomConnectedLeft": BottomConnectedLeft,
    "BottomConnectedRight": BottomConnectedRight,
    "BottomConnectedFull": BottomConnectedFull,
    "LeftConnectedTop": LeftConnectedTop,
    "LeftConnectedBottom": LeftConnectedBottom,
    "LeftConnectedFull": LeftConnectedFull,
    "RightConnectedTop": RightConnectedTop,
    "RightConnectedBottom": RightConnectedBottom,
    "RightConnectedFull": RightConnectedFull,
    "TopLeftConnectedVertical": TopLeftConnectedVertical,
    "TopLeftConnectedHorizontal": TopLeftConnectedHorizontal,
    "TopLeftConnectedFull": TopLeftConnectedFull,
    "TopRightConnectedVertical": TopRightConnectedVertical,
    "TopRightConnectedHorizontal": TopRightConnectedHorizontal,
    "TopRightConnectedFull": TopRightConnectedFull,
    "BottomLeftConnectedVertical": BottomLeftConnectedVertical,
    "BottomLeftConnectedHorizontal": BottomLeftConnectedHorizontal,
    "BottomLeftConnectedFull": BottomLeftConnectedFull,
    "BottomRightConnectedVertical": BottomRightConnectedVertical,
    "BottomRightConnectedHorizontal": BottomRightConnectedHorizontal,
    "BottomRightConnectedFull": BottomRightConnectedFull,
    "HorizontalClosed": HorizontalClosed,
    "HorizontalOpen": HorizontalOpen,
    "VerticalClosed": VerticalClosed,
    "VerticalOpen": VerticalOpen,
    "Used": Used,
}

// Generate Src visual data JSON & sprite sheet
func getTilesData(packedFrameImgs *[]FrameImage) *TilesData {

    // Get the base tiles data object containing frame source data
    var tilesData *TilesData = getBaseTilesData(packedFrameImgs)

    // Attach additional data to the tiles data: auto-var data, clock data...
    attachTilesAutoVarData(tilesData)
    attachTilesClockData(tilesData)

    return tilesData
}

// Gathers data on every single tile image
func getTilesSrcFrameImgs(frameImgs *[]FrameImage) {

    // Loop basic (non-property) tile types
    for tile := FirstNeutralTileType; tile < NeutralTileTypeCount; tile++ {
        tileDir := getFullProjectPath( tilesDir, tile.String() )

        files, err := ioutil.ReadDir(tileDir)

        if err != nil {
            log.Fatal(err)
        }

        // Check if 1 or 2-level tile
        if files[0].IsDir() {
            gatherDoubleLvlTileFrameImgs(frameImgs, tile, tileDir, files)
        } else {
            gatherSingleLvlTileFrameImgs(frameImgs, tile, tileDir, files)
        }
    }
}

// Gather frame images from a single level tile (variations are single images) and attach to given Frame Images
func gatherSingleLvlTileFrameImgs(frameImgs *[]FrameImage, tile TileType, tileDir string, files []os.FileInfo) {
    for _, file := range files {
        // Get the Tile Variation corresponding to this image file
        tileVar := tileVarsReverseStrings[strings.TrimSuffix(file.Name(), path.Ext(file.Name()))]

        // Add this file's image data to its corresponding tile variation
        imageObj := getImage(path.Join(tileDir, file.Name()))

        *frameImgs = append(*frameImgs, FrameImage{
            Image: imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: FrameImageMetaData{
                Type:           uint8(tile),
                Variation:      uint8(tileVar),
                Index:          0,
                FrameImageType: TileFrameImage,
            },
        })
    }
}

// Gather frame images from a double level tile (variations are directories of images) and attach to given Frame Images
func gatherDoubleLvlTileFrameImgs(frameImgs *[]FrameImage, tile TileType, tileDir string, dirs []os.FileInfo) {
    // Loop every variation directory
    for _, dir := range dirs {
        // Get the Tile Variation corresponding to this image file
        tileVar := tileVarsReverseStrings[dir.Name()]

        varDir := path.Join(tileDir, dir.Name());
        varFiles, err := ioutil.ReadDir(varDir)

        if err != nil {
            log.Fatal(err)
        }

        // Loop every frame image of this variation
        for index, varFile := range varFiles {
            // Add this file's image data to its corresponding tile variation
            imageObj := getImage(path.Join(varDir, varFile.Name()))

            *frameImgs = append(*frameImgs, FrameImage{
                Image: imageObj,
                Width:  imageObj.Bounds().Max.X,
                Height: imageObj.Bounds().Max.Y,
                MetaData: FrameImageMetaData{
                    Type:           uint8(tile),
                    Variation:      uint8(tileVar),
                    Index:          index,
                    FrameImageType: TileFrameImage,
                },
            })
        }
    }
}

// Generate visual data for Src
func getBaseTilesData(packedFrameImgs *[]FrameImage) *TilesData {

    // Tile Type -> Tile Variation -> Tile Variation Frames
    tilesVData := make(TilesData, NeutralTileTypeCount)

    // Initialize Variations on every TileData
    for tileType := range tilesVData {
        tilesVData[tileType].Variations = make(map[string][]Frame)
    }

    // Process Frame Images
    for _, frameImg := range *packedFrameImgs {

        // Ignore non-tile frame images
        if frameImg.MetaData.FrameImageType != TileFrameImage {
            continue
        }

        tileType := TileType(frameImg.MetaData.Type)
        tileVar := TileVariation(frameImg.MetaData.Variation)
        tileFrame := frameImg.MetaData.Index

        // Get the accumulated Animation slice of this Tile Type's given Variation, or initialize it
        var animSlice []Frame
        var ok bool

        if animSlice, ok = tilesVData[tileType].Variations[tileVar.String()]; !ok {
            animSlice = make([]Frame, 0)
        }

        // Add any Frames missing up until the one we're adding
        if missingFrames := (tileFrame + 1) - len(animSlice); missingFrames > 0 {
            for i := 0; i < missingFrames; i++ {
                animSlice = append(animSlice, Frame{})
            }
        }

        // Add the Frame data to the animation slice, and record it to the visual data
        frame := Frame{
            X: frameImg.X,
            Y: frameImg.Y,
        }

        // To save space in the JSON file, omit adding SrcWidth/SrcHeight if they're the regular Tile sizes
        if frameImg.Width != regularTileDimension {
            frame.Width = frameImg.Width
        }

        if frameImg.Height != regularTileDimension {
            frame.Height = frameImg.Height
        }

        animSlice[tileFrame] = frame
        tilesVData[tileType].Variations[tileVar.String()] = animSlice
        tilesVData[tileType].ClockData = nil
    }

    return &tilesVData
}