package tilegen

import (
	"github.com/turnabout/awodatagen"
	"github.com/turnabout/awodatagen/pkg/framedata"
	"github.com/turnabout/awodatagen/pkg/genio"
	"github.com/turnabout/awodatagen/pkg/packer"
	"github.com/turnabout/awodatagen/pkg/tilegen/tiledata"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

// Gathers data on every single tile image
func GetTileFrameImgs(frameImgs *[]packer.FrameImage) {

	// Loop basic (non-property) tile types
	for tile := tiledata.NeutralTileTypeFirst; tile < tiledata.NeutralTileTypeCount; tile++ {
		tileDir := awodatagen.GetInputPath(awodatagen.TilesDir, tile.String())

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
	tile tiledata.TileType,
	tileDir string,
	files []os.FileInfo,
) {
	for _, file := range files {
		// Get the Tile Variation corresponding to this image file
		tileVar := tiledata.TileVarsReverseStrings[strings.TrimSuffix(file.Name(), path.Ext(file.Name()))]

		// Add this file's image data to its corresponding tile variation
		imageObj := genio.GetImage(path.Join(tileDir, file.Name()))

		*frameImgs = append(*frameImgs, packer.FrameImage{
			Image:  imageObj,
			Width:  imageObj.Bounds().Max.X,
			Height: imageObj.Bounds().Max.Y,
			MetaData: packer.FrameImageMetaData{
				Type:               uint8(tile),
				Variation:          uint8(tileVar),
				Index:              0,
				FrameImageDataType: uint8(framedata.TileDataType),
			},
		})
	}
}

// Gather frame images from a double level tile (variations are directories of images) and attach to given Frame Images
func gatherDoubleLvlTileFrameImgs(
	frameImgs *[]packer.FrameImage,
	tile tiledata.TileType,
	tileDir string, dirs []os.FileInfo,
) {
	// Loop every variation directory
	for _, dir := range dirs {
		// Get the Tile Variation corresponding to this image file
		tileVar := tiledata.TileVarsReverseStrings[dir.Name()]

		varDir := path.Join(tileDir, dir.Name())
		varFiles, err := ioutil.ReadDir(varDir)

		if err != nil {
			log.Fatal(err)
		}

		// Loop every frame image of this variation
		for index, varFile := range varFiles {
			// Add this file's image data to its corresponding tile variation
			imageObj := genio.GetImage(path.Join(varDir, varFile.Name()))

			*frameImgs = append(*frameImgs, packer.FrameImage{
				Image:  imageObj,
				Width:  imageObj.Bounds().Max.X,
				Height: imageObj.Bounds().Max.Y,
				MetaData: packer.FrameImageMetaData{
					Type:               uint8(tile),
					Variation:          uint8(tileVar),
					Index:              index,
					FrameImageDataType: uint8(framedata.TileDataType),
				},
			})
		}
	}
}
