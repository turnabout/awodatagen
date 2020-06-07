package genio

import (
	"github.com/turnabout/awodatagen/internal/config"
	"github.com/turnabout/awodatagen/internal/utilities"
	"image"
	"os"
	"path"
	"path/filepath"
)

// Full, absolute path to the project's directory containing the raw assets
var assetsFullPath string

// Attempt to set the full "assets" path from an environment variable
// If it doesn't exist, use the CWD as the base path to assets
func init() {
	var envExists bool

	if assetsFullPath, envExists = os.LookupEnv(config.AssetsDirPath); !envExists {
		cwd, err := os.Getwd()

		if err != nil {
			utilities.LogFatalIfErr(err)
		}

		// Use the project's assets path as a base
		assetsFullPath = path.Join(filepath.ToSlash(cwd), config.AssetsDirName)
	}
}

// Gets the full path to a directory in the project's inputs
func GetInputPath(paths ...string) string {

	// Add up all given directories to make up the full path
	result := assetsFullPath

	for _, loopedPath := range paths {
		result = path.Join(result, loopedPath)
	}

	return result
}

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
