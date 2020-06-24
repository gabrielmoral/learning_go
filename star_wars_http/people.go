package main

type peopleResult struct {
	Results []people
}

type people struct {
	Name   string
	Height string
}

func (p people) Print() string {
	return "Name:" + p.Name + ", Height:" + p.Height
}
