package packer

import "image"

// A frame's image data (image/width/height)
type FrameImage struct {
    Image  image.Image
    Width  int
    Height int
    X int
    Y int

    MetaData FrameImageMetaData
}

// Metadata attached to a frame image
type FrameImageMetaData struct {
    Type               uint8
    Variation          uint8
    Animation          uint8
    Index              int
    FrameImageDataType uint8
}
