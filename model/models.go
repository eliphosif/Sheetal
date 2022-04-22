package model

type Trail struct {
	Question []int
}

type Item struct {
	Trails []Trail
}

type Forward struct {
	Items []Item
	score int
}

func init() {
	f := Forward{}

	f.Items[0].Trails[0].Question = []int{1, 7}
	f.Items[0].Trails[1].Question = []int{6, 3}
	f.Items[1].Trails[0].Question = []int{5, 8, 2}
	f.Items[1].Trails[1].Question = []int{6, 9, 4}
	f.Items[2].Trails[0].Question = []int{6, 4, 3, 9}
	f.Items[2].Trails[1].Question = []int{7, 2, 8, 6}
	f.Items[3].Trails[0].Question = []int{4, 2, 7, 3, 1}
	f.Items[3].Trails[1].Question = []int{7, 5, 8, 3, 6}
	f.Items[4].Trails[0].Question = []int{6, 1, 9, 4, 7, 3}
	f.Items[4].Trails[1].Question = []int{3, 9, 2, 4, 8, 7}
	f.Items[5].Trails[0].Question = []int{5, 9, 1, 7, 4, 2, 8}
	f.Items[5].Trails[1].Question = []int{4, 1, 7, 9, 3, 8, 6}
	f.Items[6].Trails[0].Question = []int{5, 8, 1, 9, 2, 6, 4, 7}
	f.Items[6].Trails[1].Question = []int{3, 8, 2, 9, 5, 1, 7, 4}
	f.Items[7].Trails[0].Question = []int{2, 7, 5, 8, 6, 2, 5, 8, 4}
	f.Items[7].Trails[1].Question = []int{7, 1, 3, 9, 4, 2, 5, 6, 8}
}
