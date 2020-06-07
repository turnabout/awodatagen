// Functions used to process inputs
package genio

import (
	"github.com/turnabout/awodatagen/pkg/utilities"
	"image"
	"os"
	"path/filepath"
)

// Gets image stored at the given path
func GetImage(path string) image.Image {

	// Ensure image is a png
	if extension := filepath.Ext(path); extension != ".png" {
		utilities.LogFatalF(
			"Tried to get an image with extension '%s', only '.png' is valid",
			filepath.Ext(path),
		)
	}

	// Load image file
	imgFile, err := os.Open(path)
	utilities.LogFatalIfErr(err)

	// Decode image and return
	img, _, err := image.Decode(imgFile)
	utilities.LogFatalIfErr(err)

	return img
}
