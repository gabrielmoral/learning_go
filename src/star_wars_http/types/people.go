package types

type PeopleResult struct {
	Results []People
}

type People struct {
	Name   string
	Height string
}

func (p People) Print() string {
	return "Name:" + p.Name + ", Height:" + p.Height
}
