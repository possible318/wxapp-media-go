package outputs

type URLItem struct {
	Name string `json:"name"`
	Src  string `json:"src"`
}

type PicItem struct {
	ID      string    `json:"id"`
	Text    string    `json:"text"`
	AddTime string    `json:"add_time"`
	SrcList []URLItem `json:"src_list"`
}

type PicItemList []PicItem

func (a PicItemList) Len() int {
	return len(a)
}
func (a PicItemList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a PicItemList) Less(i, j int) bool {
	var flag bool
	if a[i].AddTime > a[j].AddTime {
		flag = true
	} else {
		flag = false
	}
	return flag
}
