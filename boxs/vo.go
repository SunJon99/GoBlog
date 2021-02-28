package boxs

type BlogInfo struct {
	ID         uint64 `json:"id"`
	Title      string `json:"title"`
	Type       string `json:"type"`
	Recommend  string `json:"recommend"`
	UpdateTime string `json:"updateTime"`
}
