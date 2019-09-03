package main

import (
    "image"
    "io/ioutil"
    "log"
    "os"
    "path"
    "strings"
)

// Map for looking up a Tile Variation using its corresponding full string
var tileVarsReverseStrings = map[string]TileVariation{
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

// Generate Tiles visual data JSON & sprite sheet
func generateTilesData(x, y int) (*image.RGBA, *TilesData) {

    // Generate Frame Images data
    frameImgs := gatherTilesFrameImages()

    // Generate Tiles' sprite sheet & visual data
    // pack(frameImgs)

    return nil, &TilesData{
        Tiles: *generateTilesVData(frameImgs),
        ClockData: 0, // TODO
        X: x,
        Y: y,
        Width: 0, // TODO
        Height: 0, // TODO
    }
}

// Gathers data on every single image, filling out "tilesImgData"
func gatherTilesFrameImages() *[]FrameImage {
    var frameImgs []FrameImage

    tilesDir := baseDirPath + imageInputsDirName + tilesDirName + "/"

    // Loop basic (non-property) tile types
    for tile := FirstBasicTileType; tile <= LastBasicTileType; tile++ {
        tileDir := tilesDir + tile.String() + "/"
        files, err := ioutil.ReadDir(tileDir)

        if err != nil {
            log.Fatal(err)
        }

        // Check if 1 or 2-level tile
        if files[0].IsDir() {
            gatherDoubleLvlTileImgData(&frameImgs, tile, tileDir, files)
        } else {
            gatherSingleLvlTileImgData(&frameImgs, tile, tileDir, files)
        }
    }

    return &frameImgs
}

// Gather image data from a single level tile (variations are single images) and attach to given Frame Images
func gatherSingleLvlTileImgData(frameImgs *[]FrameImage, tile TileType, tileDir string, files []os.FileInfo) {
    for index, file := range files {
        // Get the Tile Variation corresponding to this image file
        tileVar := tileVarsReverseStrings[strings.TrimSuffix(file.Name(), path.Ext(file.Name()))]

        // Add this file's image data to its corresponding tile variation
        imageObj := getImage(tileDir + file.Name())

        *frameImgs = append(*frameImgs, FrameImage{
            Image: imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: FrameImageMetaData{
                Type: uint8(tile),
                Variation: uint8(tileVar),
                Index: index,
            },
        })
    }
}

// Gather image data from a double level tile (variations are directories of images) and attach to given Frame Images
func gatherDoubleLvlTileImgData(frameImgs *[]FrameImage, tile TileType, tileDir string, dirs []os.FileInfo) {
    // Loop every variation directory
    for _, dir := range dirs {
        // Get the Tile Variation corresponding to this image file
        tileVar := tileVarsReverseStrings[dir.Name()]

        varDir := tileDir + dir.Name() + "/"
        varFiles, err := ioutil.ReadDir(varDir)

        if err != nil {
            log.Fatal(err)
        }

        // Loop every frame image of this variation
        for index, varFile := range varFiles {
            // Add this file's image data to its corresponding tile variation
            imageObj := getImage(varDir + varFile.Name())

            *frameImgs = append(*frameImgs, FrameImage{
                Image: imageObj,
                Width:  imageObj.Bounds().Max.X,
                Height: imageObj.Bounds().Max.Y,
                MetaData: FrameImageMetaData{
                    Type: uint8(tile),
                    Variation: uint8(tileVar),
                    Index: index,
                },
            })
        }
    }
}

// Generate visual data for Tiles
func generateTilesVData(packedFrameImgs *[]FrameImage) *[]TileData {

    // Tile Type -> Tile Variation -> Tile Variation Frames
    tilesVData := make([]TileData, BasicTileAmount)

    return &tilesVData
}
