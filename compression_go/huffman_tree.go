package main

type huffmanObject interface {
	isLeaf() bool
	getWeight() int
}

type huffmanNode struct {
	weight int
	value  string
}

func (n huffmanNode) isLeaf() bool {
	return true
}

func (n huffmanNode) getWeight() int {
	return n.weight
}

type huffmanTree struct {
	weight int
	left   huffmanObject
	right  huffmanObject
}

func (t huffmanTree) isLeaf() bool {
	return false
}

func (t huffmanTree) getWeight() int {
	return t.weight
}

func combineToTree(huffmanObject1 huffmanObject, huffmanObject2 huffmanObject) huffmanTree {
	var left, right huffmanObject
	if huffmanObject1.getWeight() < huffmanObject2.getWeight() {
		left = huffmanObject1
		right = huffmanObject2
	} else {
		left = huffmanObject2
		right = huffmanObject1
	}
	return huffmanTree{
		weight: left.getWeight() + right.getWeight(),
		left:   left,
		right:  right,
	}
}

func sortObjects(listOfObjects []huffmanObject) []huffmanObject {
	if len(listOfObjects) < 2 {
		return listOfObjects
	}
	partitionIndex := len(listOfObjects) / 2
	leftList := sortObjects(listOfObjects[:partitionIndex])
	rightList := sortObjects(listOfObjects[partitionIndex:])
	return mergeObjects(leftList, rightList)
}

func mergeObjects(leftList []huffmanObject, rightList []huffmanObject) []huffmanObject {
	combined := []huffmanObject{}
	var l_idx, r_idx int
	for {
		if l_idx == len(leftList) || r_idx == len(rightList) {
			break
		}
		if leftList[l_idx].getWeight() < rightList[r_idx].getWeight() {
			combined = append(combined, leftList[l_idx])
			l_idx++
		} else {
			combined = append(combined, rightList[r_idx])
			r_idx++
		}
	}
	for l_idx < len(leftList) {
		combined = append(combined, leftList[l_idx])
		l_idx++
	}

	for r_idx < len(rightList) {
		combined = append(combined, rightList[r_idx])
		r_idx++
	}
	return combined
}

func buildTree(listOfNodes []huffmanNode) huffmanTree {

	if len(listOfNodes) < 2 {
		return huffmanTree{}
	}

	remainingObjects := []huffmanObject{}
	for _, node := range listOfNodes {
		remainingObjects = append(remainingObjects, node)
	}

	leftoverObjects := []huffmanObject{}
	var treeOfFirstTwo huffmanTree
	for len(remainingObjects) > 2 {
		leftoverObjects = remainingObjects[2:]
		treeOfFirstTwo = combineToTree(remainingObjects[0], remainingObjects[1])
		remainingObjects = append(leftoverObjects, treeOfFirstTwo)
		remainingObjects = sortObjects(remainingObjects)
	}

	return combineToTree(remainingObjects[0], remainingObjects[1])
}
