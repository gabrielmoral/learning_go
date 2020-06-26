package types

type VehiclesResult struct {
	Results []Vehicle
}

func (r VehiclesResult) OneSeater() []Vehicle {
	oneSeater := []Vehicle{}
	for _, v := range r.Results {
		if v.Passengers == 1 {
			oneSeater = append(oneSeater, v)
		}
	}

	return oneSeater
}

type Vehicle struct {
	Name         string
	Length       string
	Manufacturer string
	Passengers   int `json:"passengers,string"`
}

func (v Vehicle) Print() string {
	return "Name:" + v.Name + ", Length:" + v.Length + ", Manufacturer:" + v.Manufacturer
}
