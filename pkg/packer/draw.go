package packer

import (
    "image"
    "image/draw"
)

// Draw previously packed frames onto a sprite sheet
func DrawPackedFrames(frames *[]FrameImage, width, height int) *image.RGBA {
    spriteSheet := image.NewRGBA(image.Rectangle{
        Min: image.Point{X: 0, Y: 0},
        Max: image.Point{X: width, Y: height},
    })

    for _, frame := range *frames {
        drawFrame(&frame, spriteSheet)
    }

    return spriteSheet
}

// Draw a Frame Image in the given Node, onto the given sprite sheet
func drawFrame(frame *FrameImage, ss *image.RGBA) {
    // Move image to the node's srcX/srcY coordinates
    rect := frame.Image.Bounds()

    rect.Min.X += frame.X
    rect.Max.X += frame.X

    rect.Min.Y += frame.Y
    rect.Max.Y += frame.Y

    // Draw frame image onto the sprite sheet
    draw.Draw(ss, rect, frame.Image, image.Point{X: 0, Y: 0}, draw.Src)
}