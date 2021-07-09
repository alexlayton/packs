package packs

type Packs []Pack

type Pack struct {
	Size  int `json:"size"`
	Count int `json:"count"`
}

func (p Packs) Len() int {
	return len(p)
}

func (p Packs) Less(i, j int) bool {
	return p[i].Size < p[j].Size
}

func (p Packs) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
