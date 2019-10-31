// Packer is used to tightly fit multiple images into a sprite sheet.
package main

import (
    "fmt"
    "image"
    "image/draw"
    "sort"
)

// A frame's image data (image/width/SrcHeight)
type FrameImage struct {
    Image  image.Image
    Width  int
    Height int
    MetaData FrameImageMetaData
    X int
    Y int
}

// Which Type/Var/Animation/Animation Index this Frame Image belongs to
// TODO: Rename/reorganize
type FrameImageMetaData struct {
    Type uint8
    Variation uint8
    Animation uint8
    Index int
}

func (f FrameImage) String() string {
    return fmt.Sprintf(
        "%s, %s, %s, [%d]: %dx%d (%d) (%d, %d)",
        PropertyType(f.MetaData.Type).String(),
        PropertyWeatherVariation(f.MetaData.Variation).String(),
        UnitVariation(f.MetaData.Animation).String(),
        f.MetaData.Index,
        f.Width,
        f.Height,
        f.Width * f.Height,
        f.X, f.Y,
    )
}

// Sorts Frame Images by image size (largest area to smallest area)
type SizeSorter []FrameImage

func (f SizeSorter) Len() int           { return len(f) }
func (f SizeSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f SizeSorter) Less(i, j int) bool { return (f[i].Width * f[i].Height) > (f[j].Width * f[j].Height) }

// Sorts Frame Images by Meta Data Type
type TypeSorter []FrameImage

func (f TypeSorter) Len() int           { return len(f) }
func (f TypeSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f TypeSorter) Less(i, j int) bool { return f[i].MetaData.Type < f[j].MetaData.Type }

// Represents a slot inside the sprite sheet that can be taken up.
// Can be split up into and linked to further nodes, one underneath it and another to its right
type Node struct {
    X, Y, Width, Height int
    Used bool
    Down *Node
    Right *Node
}

// Pack Frame Images into an expanding surface, attaching Nodes specifying coordinates to every FrameImage in the list,
// and returning it along with the packed surface's width and height
func pack(framesArg *[]FrameImage) (*[]FrameImage, int, int) {

    // Max encountered srcX/srcY (surface width/height)
    var xMax, yMax int

    if len(*framesArg) < 1 {
        return &[]FrameImage{}, 0, 0
    }

    frames := make([]FrameImage, len(*framesArg))
    copy(frames, *framesArg)

    // Sort the frames in descending order of size
    sort.Sort(SizeSorter(frames))

    root := Node{X: 0, Y: 0, Width: (frames)[0].Width, Height: (frames)[0].Height}

    for index, frame := range frames {
        node, ok := findNode(&root, frame.Width, frame.Height)

        if ok {
            // Fits within a found node, place it here and split the node
            node = splitNode(node, frame.Width, frame.Height)
        } else {
            // Doesn't fit within any found node, grow the node in either direction and return it
            node, ok = growNode(&root, frame.Width, frame.Height)

            if !ok {
                fmt.Printf("Failed to fit Frame #%d\n", index)
                continue
            }
        }

        // Assign Node's coordinates to this Image Frame
        frames[index].X = node.X
        frames[index].Y = node.Y

        // Update srcX/srcY max encountered values
        if xMax < node.X + frame.Width {
            xMax = node.X + frame.Width
        }

        if yMax < node.Y + frame.Height {
            yMax = node.Y + frame.Height
        }
    }

    return &frames, xMax, yMax
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

// Grow the root node to fit a Frame Image of the given dimensions
func growNode(root *Node, width, height int) (*Node, bool) {
    canGrowDown := width <= root.Width
    canGrowRight := height <= root.Height

    // Attempt to keep a square-like shape by growing in the appropriate direction
    shouldGrowRight := canGrowRight && (root.Height >= (root.Width + width))
    shouldGrowDown := canGrowDown && (root.Width >= (root.Height + height))

    if shouldGrowRight {
        return growRight(root, width, height)
    } else if shouldGrowDown {
        return growDown(root, width, height)
    } else if canGrowRight {
        return growRight(root, width, height)
    } else if canGrowDown {
        return growDown(root, width, height)
    }

    return nil, false // Shouldn't happen if all Frame Images are sorted from largest to smallest
}

func growRight(root *Node, width, height int) (*Node, bool) {
    rootCopy := *root

    *root = Node{
        Used: true,
        X: 0,
        Y: 0,
        Width: rootCopy.Width + width,
        Height: rootCopy.Height,
        Down: &rootCopy,
        Right: &Node{
            X: rootCopy.Width,
            Y: 0,
            Width: width,
            Height: rootCopy.Height,
        },
    }

    if node, ok := findNode(root, width, height); ok {
        return splitNode(node, width, height), true
    } else {
        return nil, false
    }
}

func growDown(root *Node, width, height int) (*Node, bool) {
    rootCopy := *root

    *root = Node{
        Used: true,
        X: 0,
        Y: 0,
        Width: rootCopy.Width,
        Height: rootCopy.Height + height,
        Down: &Node{
            X: 0,
            Y: rootCopy.Height,
            Width: rootCopy.Width,
            Height: height,
        },
        Right: &rootCopy,
    }

    if node, ok := findNode(root, width, height); ok {
        return splitNode(node, width, height), true
    } else {
        return nil, false
    }
}

// Draw previously packed frames onto a sprite sheet
func drawPackedFrames(frames *[]FrameImage, width, height int) *image.RGBA {
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