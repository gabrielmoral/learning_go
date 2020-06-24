package main

type vehiclesResult struct {
	Results []vehicle
}

func (r vehiclesResult) OneSeater() []vehicle {
	oneSeater := []vehicle{}
	for _, v := range r.Results {
		if v.Passengers == 1 {
			oneSeater = append(oneSeater, v)
		}
	}

	return oneSeater
}

type vehicle struct {
	Name         string
	Length       string
	Manufacturer string
	Passengers   int `json:"passengers,string"`
}

func (v vehicle) Print() string {
	return "Name:" + v.Name + ", Length:" + v.Length + ", Manufacturer:" + v.Manufacturer
}
