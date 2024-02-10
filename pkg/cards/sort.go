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

type ByFaceReversed []Card


func (r ByFaceReversed) Len() int {
	return len(r)
}

func (r ByFaceReversed) Less(i, j int) bool {
	return r[i].Face() > r[j].Face()
}

func (r ByFaceReversed) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}


type ByCombination []Combination

func (r ByCombination) Len() int {
	return len(r)
}

func (r ByCombination) Less(i, j int) bool {
	return r[i].Less(r[j])
}

func (r ByCombination) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

type ByCombinationReversed []Combination

func (r ByCombinationReversed) Len() int {
	return len(r)
}

func (r ByCombinationReversed) Less(i, j int) bool {
	return r[j].Less(r[i])
}

func (r ByCombinationReversed) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
