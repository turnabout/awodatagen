package packerTree

// Find the next unused node in which the given dimensions fit
func FindNode(root *PackerNode, width, height int) (*PackerNode, bool) {
	// Given root node already used, return one of its adjacent nodes
	if root.Used {
		if rightNode, ok := FindNode(root.Right, width, height); ok {
			return rightNode, true
		} else {
			return FindNode(root.Down, width, height)
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
func SplitNode(node *PackerNode, width, height int) *PackerNode {
	node.Used = true

	node.Down = &PackerNode{
		X:      node.X,
		Y:      node.Y + height,
		Width:  node.Width,
		Height: node.Height - height,
	}

	node.Right = &PackerNode{
		X:      node.X + width,
		Y:      node.Y,
		Width:  node.Width - width,
		Height: height,
	}

	return node
}

// Grow the root node to fit a Frame Image of the given dimensions
func GrowNode(root *PackerNode, width, height int) (*PackerNode, bool) {
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

func growRight(root *PackerNode, width, height int) (*PackerNode, bool) {
	rootCopy := *root

	*root = PackerNode{
		Used:   true,
		X:      0,
		Y:      0,
		Width:  rootCopy.Width + width,
		Height: rootCopy.Height,
		Down:   &rootCopy,
		Right: &PackerNode{
			X:      rootCopy.Width,
			Y:      0,
			Width:  width,
			Height: rootCopy.Height,
		},
	}

	if node, ok := FindNode(root, width, height); ok {
		return SplitNode(node, width, height), true
	} else {
		return nil, false
	}
}

func growDown(root *PackerNode, width, height int) (*PackerNode, bool) {
	rootCopy := *root

	*root = PackerNode{
		Used:   true,
		X:      0,
		Y:      0,
		Width:  rootCopy.Width,
		Height: rootCopy.Height + height,
		Down: &PackerNode{
			X:      0,
			Y:      rootCopy.Height,
			Width:  rootCopy.Width,
			Height: height,
		},
		Right: &rootCopy,
	}

	if node, ok := FindNode(root, width, height); ok {
		return SplitNode(node, width, height), true
	} else {
		return nil, false
	}
}
