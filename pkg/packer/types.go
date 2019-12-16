package packer

import "image"

// A frame's image data (image/width/height)
type FrameImage struct {
    Image  image.Image
    Width  int
    Height int
    X int
    Y int

    // MetaData FrameImageMetaData
}

