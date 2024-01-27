package cards


// sort.Interface implementation for handy sorting
type ByFace []Card

func (r ByFace) Len() int {
	return len(r)
}

func (r ByFace) Less(i, j int) bool {
	return r[i].Face() < r[j].Face()
}

func (r ByFace) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
