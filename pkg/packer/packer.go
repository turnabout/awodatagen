// Packer is used to tightly fit multiple images into a sprite sheet.
package packer

import (
    "fmt"
    "github.com/turnabout/awossgen/pkg/packer/packerTree"
    "sort"
)

/*
type FrameImageType uint8

const(
    UnitFrameImage FrameImageType = iota
    TileFrameImage
    PropertyFrameImage
    UiElementFrameImage
    SpriteSheetSectionFrameImage
)
*/

// Which Type/Var/Animation/Animation Index this Frame Image belongs to
// TODO: Rename/reorganize
/*
type FrameImageMetaData struct {
    Type uint8
    Variation uint8
    Animation uint8
    Index int
    FrameImageType FrameImageType
}
*/

/*
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
*/

// Sorts Frame Images by Meta Data Type

/*
type TypeSorter []FrameImage

func (f TypeSorter) Len() int           { return len(f) }
func (f TypeSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f TypeSorter) Less(i, j int) bool { return f[i].MetaData.Type < f[j].MetaData.Type }
*/


// Sorts frame images by image size (largest to smallest area)
type SizeSorter []FrameImage

func (f SizeSorter) Len() int           { return len(f) }
func (f SizeSorter) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f SizeSorter) Less(i, j int) bool { return (f[i].Width * f[i].Height) > (f[j].Width * f[j].Height) }


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

    // Create the root node of the packer packerTree
    root := packerTree.PackerNode{X: 0, Y: 0, Width: (frames)[0].Width, Height: (frames)[0].Height}

    // Loop every frame image, filling out their x/y values
    for index, frame := range frames {
        node, ok := packerTree.FindNode(&root, frame.Width, frame.Height)


        if ok {
            // Fits within a found node, place it here and split the node
            node = packerTree.SplitNode(node, frame.Width, frame.Height)
        } else {
            // Doesn't fit within any found node, grow the node in either direction and return it
            node, ok = packerTree.GrowNode(&root, frame.Width, frame.Height)

            if !ok {
                fmt.Printf("Failed to fit Frame #%d\n", index)
                continue
            }
        }

        // Assign packer node's coordinates to this Image Frame
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

