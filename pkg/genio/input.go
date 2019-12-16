// Functions used to process inputs
package genio

import (
    "fmt"
    "github.com/turnabout/awossgen"
    "image"
    "os"
    "path/filepath"
)

// Gets image stored at the given path
func GetImage(path string) image.Image {

    // Ensure image is a png
    if extension := filepath.Ext(path); extension != ".png" {
        awossgen.LogFatal([]string{
            fmt.Sprintf(
                "Tried to get an image with extension '%s', only '.png' is valid",
                filepath.Ext(path),
            ),
        })
    }

    // Load image file
    imgFile, err := os.Open(path)
    awossgen.LogFatalIfErr(err)

    // Decode image and return
    img, _, err := image.Decode(imgFile)
    awossgen.LogFatalIfErr(err)

    return img
}
