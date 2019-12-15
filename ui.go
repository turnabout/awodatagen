package main

import (
    "fmt"
)

// Generate Src visual data JSON & sprite sheet
func getUIData(packedFrameImgs *[]FrameImage) *TilesData {
    vData := TilesData{
        // Src: *getTilesSrcVData(packedFrameImgs),
    }

    attachExtraTilesVData(&vData)
    return &vData
}

// Gathers data on every single UI image
func getUISrcFrameImgs(frameImgs *[]FrameImage) {
    uiDir := baseDirPath + inputsDirName + uiDirName + "/"

    fmt.Printf("%d\n", uiDir)

    // Loop basic (non-property) tile types
    /*
    for tile := FirstNeutralTileType; tile < NeutralTileTypeCount; tile++ {
        tileDir := tilesDir + tile.String() + "/"
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
    */
}
