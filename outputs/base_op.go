package outputs

type URLItem struct {
	Pid    string `json:"pid"`
	ID     int    `json:"id"`
	ItemID string `json:"item_id"`
	Text   string `json:"text"`
	Src    string `json:"src"`
}

type SrcList []URLItem

type PicItem struct {
	ID      string  `json:"id"`
	Text    string  `json:"text"`
	AddTime string  `json:"add_time"`
	SrcList SrcList `json:"src_list"`
}

type PicItemList []PicItem

type PhotoOpt struct {
	Page  int       `json:"page"`
	Total int       `json:"total"`
	List  []URLItem `json:"list"`
}

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
