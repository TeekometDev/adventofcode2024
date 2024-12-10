package main

type Place struct {
	row      int
	col      int
	height   byte
	children []*Place
}

func uniqueChild(place *Place, newChild *Place) bool {
	isUnique := true
	for _, child := range (*place).children {
		if child.row == (*newChild).row && child.col == (*newChild).col {
			return false
		}
		isUnique = isUnique && uniqueChild(child, newChild)
	}
	return isUnique
}
