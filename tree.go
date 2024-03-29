package fti

type Tree struct {
	levels []*Level
	bufs   []*Level
	height int
}

func NewTree() *Tree {
	levels := make([]*Level, 0, 10)
	bufs := make([]*Level, 0, 10)
	levels = append(levels, NewLevel(1))
	bufs = append(bufs, NewLevel(1))
	return &Tree{levels: levels, bufs: bufs}
}

func (tree *Tree) Search(key int) *Node {
	var node *Node
	for i, l := range tree.levels {
		if l.Empty() {
			continue
		}
		node = l.Search(key)
		if node != nil {
			return node
		}
		node = tree.bufs[i].Search(key)
		if node != nil {
			return node
		}
	}
	return nil
}

func (tree *Tree) RangeSearch(key1, key2 int) []*Node {
	var temp [][]*Node
	var count int
	for i, l := range tree.levels {
		if l.Empty() {
			continue
		}
		r := l.RangeSearch(key1, key2)
		if r != nil {
			temp = append(temp, r)
			count += len(r)
		}
		r = tree.bufs[i].RangeSearch(key1, key2)
		if r != nil {
			temp = append(temp, r)
			count += len(r)
		}
	}
	if count == 0 {
		return nil
	} else {
		return mergeNodes(temp, count)
	}
}

func mergeNodes(temp [][]*Node, count int) []*Node {
	merge := make([]*Node, 0, count)
	var mNode *Node
	var mIdx  int
	for len(temp) > 0 {
		mNode = temp[0][0]
		mIdx = 0
		for i, level := range temp[1:] {
			if level[0].Key() < mNode.Key() {
				mNode = level[0]
				mIdx = i+1
			}
		}
		merge = append(merge, mNode)
		temp[mIdx] = temp[mIdx][1:]
		if len(temp[mIdx]) == 0 {
			temp[mIdx], temp[len(temp)-1] = temp[len(temp)-1], temp[mIdx]
			temp = temp[:len(temp)-1]
		}
	}
	return merge
}

func (tree *Tree) Insert(key int, value string) {
	node := tree.Search(key)
	if node != nil {
		node.SetValue(value)
	} else {
		tree.FastInsert(key, value)
	}
}

// Assume no duplication.
func (tree *Tree) FastInsert(key int, value string) {
	node := NewNode(key, value)
	if !tree.levels[0].Full() {
		tree.levels[0].Insert(node)
	} else {
		tree.bufs[0].Insert(node)
		h := 0
		for tree.levels[h].Full() && tree.bufs[h].Full() {
			if h+1 > tree.height {
				tree.levels = append(tree.levels, NewLevel(1<<uint(h+1)))
				tree.bufs = append(tree.bufs, NewLevel(1<<uint(h+1)))
				tree.height++
			}
			if tree.levels[h+1].Full() {
				tree.levels[h].Merge(tree.bufs[h], tree.bufs[h+1])
			} else {
				tree.levels[h].Merge(tree.bufs[h], tree.levels[h+1])
			}
			h++
		}
	}
}