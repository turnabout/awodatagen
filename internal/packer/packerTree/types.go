package packerTree

// A packer node element, making up a packer tree along with other nodes.
// Represents a slot inside the sprite sheet that can be taken up.
// Can be split up into and linked to further nodes, one underneath it and another to its right.
type PackerNode struct {
	X, Y, Width, Height int
	Used                bool
	Down                *PackerNode
	Right               *PackerNode
}
