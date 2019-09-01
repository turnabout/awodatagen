// Packer is used to tightly fit multiple images into a sprite sheet.
package main

import (
    "fmt"
    "image"
    "image/draw"
    "sort"
)

// A frame's image data (image/width/Height)
type FrameImage struct {
    Image  image.Image
    Width  int
    Height int
    MetaData FrameImageMetaData
    Fit Node
}

// Which Type/Var/Animation/Animation Index this Frame Image belongs to
type FrameImageMetaData struct {
    Type uint8
    Variation uint8
    Animation uint8
    Index int
}

func (f FrameImage) String() string {
    return fmt.Sprintf(
        "%s, %s, %s: %d x %d (%d)",
        UnitType(f.MetaData.Type).String(),
        UnitVariation(f.MetaData.Variation).String(),
        UnitAnimation(f.MetaData.Animation).String(),
        f.Width,
        f.Height,
        f.Width * f.Height,
    )
}

// Sorts Frame Images by image size (largest area to smallest area)
type SizeSorter []FrameImage

func (f SizeSorter) Len() int           { return len(f) }
func (f SizeSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f SizeSorter) Less(i, j int) bool { return (f[i].Width * f[i].Height) > (f[j].Width * f[j].Height) }

// Represents a slot inside the sprite sheet that can be taken up.
// Can be split up into and linked to further nodes, one underneath it and another to its right
type Node struct {
    X, Y, Width, Height int
    Used bool
    Down *Node
    Right *Node
}

// Pack the given Frame Images into the given dimensions, filling out a sprite sheet and returning it
func pack(frames *[]FrameImage, width int, height int) *image.RGBA {
    root := Node{X: 0, Y: 0, Width: width, Height: height}

    spriteSheet := image.NewRGBA(image.Rectangle{
        Min: image.Point{X: 0, Y: 0},
        Max: image.Point{X: width, Y: height},
    })

    // Sort the frames in descending order of size
    sort.Sort(SizeSorter(*frames))

    for _, frame := range *frames {
        if node, ok := findNode(&root, frame.Width, frame.Height); ok {
            node = splitNode(node, frame.Width, frame.Height)

            // Add image to sprite sheet at node's coordinates
            drawFrame(&frame, node, spriteSheet)
            fmt.Printf("%s\nGoing to: (%d, %d)\n\n", frame, node.X, node.Y)
        } else {
            fmt.Printf("Skipped %s\n", frame)
        }
    }

    return spriteSheet
}

// Find the next unused node in which the given dimensions fit
func findNode(root *Node, width, height int) (*Node, bool) {
    // Given root node already used, return one of its adjacent nodes
    if root.Used {
        if rightNode, ok := findNode(root.Right, width, height); ok {
            return rightNode, true
        } else {
            return findNode(root.Down, width, height)
        }

    // Return this node if the given dimensions fit inside it
    } else if (width <= root.Width) && (height <= root.Height) {
        return root, true

    // Dimensions don't fit, dead end
    } else {
        return nil, false
    }
}

// Occupy a node with a block of given dimensions, then split it, adding one to its right and under it
func splitNode(node *Node, width, height int) *Node {
    node.Used = true

    node.Down = &Node{
        X: node.X,
        Y: node.Y + height,
        Width: node.Width,
        Height: node.Height - height,
    }

    node.Right = &Node{
        X: node.X + width,
        Y: node.Y,
        Width: node.Width - width,
        Height: height,
    }

    return node
}

// Draw a Frame Image in the given Node, onto the given sprite sheet
func drawFrame(frame *FrameImage, node *Node, ss *image.RGBA) {

    // Move image to the node's X/Y coordinates
    rect := frame.Image.Bounds()

    rect.Min.X += node.X
    rect.Max.X += node.X

    rect.Min.Y += node.Y
    rect.Max.Y += node.Y

    // Draw frame image onto the sprite sheet
    draw.Draw(ss, rect, frame.Image, image.Point{X: 0, Y: 0}, draw.Src)
}