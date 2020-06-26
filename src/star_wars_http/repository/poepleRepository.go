package repository

import (
	"star_wars_http/types"
)

func GetPeople(url string) (res types.PeopleResult) {
	peopleRes := types.PeopleResult{}

	get(url, &peopleRes)

	return peopleRes
}
