package tilegen

import (
    "github.com/turnabout/awossgen"
    "github.com/turnabout/awossgen/pkg/genio"
    "github.com/turnabout/awossgen/pkg/packer"
    "io/ioutil"
    "log"
    "os"
    "path"
    "strings"
)

// Gathers data on every single tile image
func GetTileFrameImgs(frameImgs *[]packer.FrameImage) {

    // Loop basic (non-property) tile types
    for tile := awossgen.FirstNeutralTileType; tile < awossgen.NeutralTileTypeCount; tile++ {
        tileDir := awossgen.GetInputPath( awossgen.TilesDir, tile.String() )

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
func gatherSingleLvlTileFrameImgs(
    frameImgs *[]packer.FrameImage,
    tile awossgen.TileType,
    tileDir string,
    files []os.FileInfo,
) {
    for _, file := range files {
        // Get the Tile Variation corresponding to this image file
        tileVar := tileVarsReverseStrings[strings.TrimSuffix(file.Name(), path.Ext(file.Name()))]

        // Add this file's image data to its corresponding tile variation
        imageObj := genio.GetImage(path.Join(tileDir, file.Name()))

        *frameImgs = append(*frameImgs, packer.FrameImage{
            Image: imageObj,
            Width:  imageObj.Bounds().Max.X,
            Height: imageObj.Bounds().Max.Y,
            MetaData: packer.FrameImageMetaData{
                Type:           uint8(tile),
                Variation:      uint8(tileVar),
                Index:          0,
                FrameImageDataType: uint8(awossgen.TileDataType),
            },
        })
    }
}

// Gather frame images from a double level tile (variations are directories of images) and attach to given Frame Images
func gatherDoubleLvlTileFrameImgs(
    frameImgs *[]packer.FrameImage,
    tile awossgen.TileType,
    tileDir string, dirs []os.FileInfo,
) {
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
            imageObj := genio.GetImage(path.Join(varDir, varFile.Name()))

            *frameImgs = append(*frameImgs, packer.FrameImage{
                Image: imageObj,
                Width:  imageObj.Bounds().Max.X,
                Height: imageObj.Bounds().Max.Y,
                MetaData: packer.FrameImageMetaData{
                    Type:               uint8(tile),
                    Variation:          uint8(tileVar),
                    Index:              index,
                    FrameImageDataType: uint8(awossgen.TileDataType),
                },
            })
        }
    }
}
