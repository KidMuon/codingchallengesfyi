package main

type stringEncoding map[string]string

type encodedNode struct {
	huffmanNode
	encoding string
}

func makeEncodingStrings(tree huffmanTree) stringEncoding {
	foundEncodedNodes := []encodedNode{}
	if tree.left.isLeaf() {
		foundEncodedNodes = append(foundEncodedNodes,
			encodedNode{huffmanNode: tree.left.(huffmanNode), encoding: "1"})
	} else {
		foundEncodedNodes = append(foundEncodedNodes,
			searchTree(tree.left.(huffmanTree), "1")...)
	}

	if tree.right.isLeaf() {
		foundEncodedNodes = append(foundEncodedNodes,
			encodedNode{huffmanNode: tree.right.(huffmanNode), encoding: "0"})
	} else {
		foundEncodedNodes = append(foundEncodedNodes,
			searchTree(tree.right.(huffmanTree), "0")...)
	}

	encodings := make(stringEncoding)
	for _, node := range foundEncodedNodes {
		encodings[node.value] = node.encoding
	}

	return encodings
}

func searchTree(tree huffmanTree, prefix string) []encodedNode {
	foundEncodedNodes := []encodedNode{}
	if tree.left.isLeaf() {
		foundEncodedNodes = append(foundEncodedNodes,
			encodedNode{huffmanNode: tree.left.(huffmanNode), encoding: prefix + "1"})
	} else {
		foundEncodedNodes = append(foundEncodedNodes,
			searchTree(tree.left.(huffmanTree), prefix+"1")...)
	}

	if tree.right.isLeaf() {
		foundEncodedNodes = append(foundEncodedNodes,
			encodedNode{huffmanNode: tree.right.(huffmanNode), encoding: prefix + "0"})
	} else {
		foundEncodedNodes = append(foundEncodedNodes,
			searchTree(tree.right.(huffmanTree), prefix+"0")...)
	}

	return foundEncodedNodes
}
