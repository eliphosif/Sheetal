package model

type Trail struct {
	Question string
}

type Item struct {
	Trails [2]Trail
}

type Forward struct {
	Items [8]Item
	score int
}

var Fmodal Forward

func init() {
	Fmodal = Forward{}

	Fmodal.Items[0].Trails[0].Question = "17"
	Fmodal.Items[0].Trails[1].Question = "63"
	Fmodal.Items[1].Trails[0].Question = "582"
	Fmodal.Items[1].Trails[1].Question = "694"
	Fmodal.Items[2].Trails[0].Question = "6439"
	Fmodal.Items[2].Trails[1].Question = "7286"
	Fmodal.Items[3].Trails[0].Question = "42731"
	Fmodal.Items[3].Trails[1].Question = "75836"
	Fmodal.Items[4].Trails[0].Question = "619473"
	Fmodal.Items[4].Trails[1].Question = "392487"
	Fmodal.Items[5].Trails[0].Question = "5917428"
	Fmodal.Items[5].Trails[1].Question = "4179386"
	Fmodal.Items[6].Trails[0].Question = "58192647"
	Fmodal.Items[6].Trails[1].Question = "38295174"
	Fmodal.Items[7].Trails[0].Question = "275862584"
	Fmodal.Items[7].Trails[1].Question = "713942568"
}
