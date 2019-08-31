package main

import (
    "fmt"
    "image"
    "io/ioutil"
    "log"
    "os"
    "path"
    "strings"
)

// Image data on every single tile frame
// Dimensions: Tile Type -> Tile Variation -> Tile Variation Frames
var tilesImgData []map[string][]FrameImage

// Tiles visual data to be exported in final JSON
var tilesVisualData []TileData

// Tiles' sprite sheet image
var tilesSSImg *image.RGBA

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

func init() {
    tilesImgData = make([]map[string][]FrameImage, LastBasicTileType + 1)
}

// Generate Tiles visual data JSON & sprite sheet
func generateTiles() {
    gatherTilesImgData()

    for tile, tileMap := range tilesImgData {
        fmt.Printf("%s\n", TileType(tile).String())
        fmt.Printf("%#v\n", tileMap)
    }
}

// Gathers data on every single image, filling out "tilesImgData"
func gatherTilesImgData() {
    // Base directory containing tile images
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
            gatherDoubleLvlTileImgData(tile, tileDir, files)
        } else {
            gatherSingleLvlTileImgData(tile, tileDir, files)
        }
    }
}

// Gather image data from a single level tile (variations are single images)
func gatherSingleLvlTileImgData(tile TileType, tileDir string, files []os.FileInfo) {
    // Image data for all variations of this tile
    tileVars := make(map[string][]FrameImage)

    for _, file := range files {
        // Get the Tile Variation corresponding to this image file
        tileVar := tileVarsReverseStrings[strings.TrimSuffix(file.Name(), path.Ext(file.Name()))]

        // Add this file's image data to its corresponding tile variation
        imageObj := getImage(tileDir + file.Name())

        tileVars[tileVar.String()] = append(tileVars[tileVar.String()], FrameImage{
            Image:  imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
        })
    }

    // Attach all gathered tile variations image data to tilesImgData
    tilesImgData[tile] = tileVars
}

// Gather image data from a double level tile (variations are directories of images)
func gatherDoubleLvlTileImgData(tile TileType, tileDir string, dirs []os.FileInfo) {
}
